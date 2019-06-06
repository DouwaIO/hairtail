package main

import (
	"fmt"
	"container/list"
	"context"
	"log"
	"runtime"
	"sync"
	"time"
//	"net/http"
//	"io/ioutil"
	"errors"
//	"encoding/json"
//	"os"
)

var (
	// ErrCancel indicates the task was cancelled.
	ErrCancel = errors.New("queue: task cancelled")

	// ErrNotFound indicates the task was not found in the queue.
	ErrNotFound = errors.New("queue: task not found")
)

type MapType map[string]string

// Task defines a unit of work in the queue.
type Task struct {
	// ID identifies this task.
	ID string `json:"id,omitempty"`

	// Data is the actual data in the entry.
	Data []byte `json:"data"`

	// Labels represents the key-value pairs the entry is lebeled with.
	Labels map[string]string `json:"labels,omitempty"`
	//Labels MapType `json:"labels,omitempty"`
}

// InfoT provides runtime information.
type InfoT struct {
	Pending []*Task `json:"pending"`
	Running []*Task `json:"running"`
	Agents []string `json:"agents"`
	Stats   struct {
		Workers  int `json:"worker_count"`
		Pending  int `json:"pending_count"`
		Running  int `json:"running_count"`
		Complete int `json:"completed_count"`
	} `json:"stats"`
}

//func createFilterFunc(filter rpc.Filter) (queue.Filter, error) {
//	return func(task *queue.Task) bool {
//		var st *expr.Selector
//		var err error
//		logrus.Debugf("task.Labels: %v", task.Labels)
//		logrus.Debugf("filter.Labels: %v", filter.Labels)
//		vv := map[string]string{}
//		if filter.Labels["public"] == task.Labels["public"] {
//			if filter.Labels["public"] == "true" {
//				return true
//			} else {
//
//				if task.Labels["filter"] != "" {
//					if filter.Labels["tags_str"] != ""{
//						task_tags := strings.Split(task.Labels["filter"], ",")
//						for i := 0;i<=len(task_tags)-1;i++{
//						//	fmt.Printf("zzz:%v", zzz[i])
//							expr_str := "'"+task_tags[i]+ "' IN " + filter.Labels["tags_str"]
//							logrus.Debugf("expr_str: %s", expr_str)
//							st, err = expr.ParseString(expr_str)
//							if err != nil {
//								logrus.Debugf("err1: %v", err)
//								return false
//							}
//							if st != nil {
//								//if task.Labels["tags"] != ""{
//								//	 err = json.Unmarshal(task.Labels["tags"], &zzz)
//								//	    if err != nil {
//								//		logrus.Debugf("err2: %v", err)
//								//		    false
//								//	    }
//								//}
//								match, _ := st.Eval(expr.NewRow(vv))
//								if match == true{
//									return true
//								} else if i == len(task_tags)-1 {
//									return false
//								}
//							}
//						}
//
//					} else {
//						logrus.Debugf("1111")
//						return false
//					}
//				} else {
//					logrus.Debugf("2222")
//					return false
//				}
//			}
//		} else {
//			logrus.Debugf("public != task public")
//			return false
//		}
//		logrus.Debugf("just")
//		return true
//
//		//for k, v := range filter.Labels {
//		//	if task.Labels[k] != v {
//		//		return false
//		//	}
//		//}
//		//return true
//	}, nil
//}



// Filter filters tasks in the queue. If the Filter returns false,
// the Task is skipped and not returned to the subscriber.
type Filter func(*Task) bool

// Queue defines a task queue for scheduling tasks among
// a pool of workers.
type Queue interface {
	// Push pushes an task to the tail of this queue.
	Push(c context.Context, task *Task) error

	// Poll retrieves and removes a task head of this queue.
	Poll(c context.Context, f Filter, agent_id string) (*Task, error)

	// Extend extends the deadline for a task.
	Extend(c context.Context, id string) error

	// Done signals the task is complete.
	Done(c context.Context, id string) error

	// Error signals the task is complete with errors.
	Error(c context.Context, id string, err error) error

	// Evict removes a pending task from the queue.
	Evict(c context.Context, id string) error

	// Wait waits until the task is complete.
	Wait(c context.Context, id string) error

	// Info returns internal queue information.
	Info(c context.Context) InfoT
	GetRuning(c context.Context, project string) int
	GetWorkers(c context.Context) []*Task
}

type Nums struct {
	BuildCount  int `json:"build_num"`
	AgentCount int   `json:"agent_num"`
	ConCount int   `json:"concurrency_num"`
}

type entry struct {
	item     *Task
	done     chan bool
	retry    int
	error    error
	deadline time.Time
}

type worker struct {
	filter  Filter
	channel chan *Task
	agent_id string
}

type fifo struct {
	sync.Mutex

	workers   map[*worker]struct{}
	running   map[string]*entry
	pending   *list.List
	extension time.Duration
}

// New returns a new fifo queue.
func New() Queue {
	return &fifo{
		workers:   map[*worker]struct{}{},
		running:   map[string]*entry{},
		pending:   list.New(),
		extension: time.Minute * 10,
	}
}

// Push pushes an item to the tail of this queue.
func (q *fifo) Push(c context.Context, task *Task) error {
	q.Lock()
	q.pending.PushBack(task)
	q.Unlock()
	go q.process()
	return nil
}

// Poll retrieves and removes the head of this queue.
func (q *fifo) Poll(c context.Context, f Filter, agent_id string) (*Task, error) {
	q.Lock()
	w := &worker{
		channel: make(chan *Task, 1),
		filter:  f,
		agent_id: agent_id,
	}
	q.workers[w] = struct{}{}
	q.Unlock()
	go q.process()

	for {
		select {
		case <-c.Done():
			q.Lock()
			delete(q.workers, w)
			q.Unlock()
			return nil, nil
		case t := <-w.channel:
			return t, nil
		}
	}
}

// Done signals that the item is done executing.
func (q *fifo) Done(c context.Context, id string) error {
	return q.Error(c, id, nil)
}

// Error signals that the item is done executing with error.
func (q *fifo) Error(c context.Context, id string, err error) error {
	q.Lock()
	state, ok := q.running[id]
	if ok {
		state.error = err
		close(state.done)
		delete(q.running, id)
	}
	q.Unlock()
	return nil
}

// Evict removes a pending task from the queue.
func (q *fifo) Evict(c context.Context, id string) error {
	q.Lock()
	defer q.Unlock()

	var next *list.Element
	for e := q.pending.Front(); e != nil; e = next {
		next = e.Next()
		task, ok := e.Value.(*Task)
		if ok && task.ID == id {
			q.pending.Remove(e)
			return nil
		}
	}
	return ErrNotFound
}

// Wait waits until the item is done executing.
func (q *fifo) Wait(c context.Context, id string) error {
	q.Lock()
	state := q.running[id]
	q.Unlock()
	if state != nil {
		select {
		case <-c.Done():
		case <-state.done:
			return state.error
		}
	}
	return nil
}

// Extend extends the task execution deadline.
func (q *fifo) Extend(c context.Context, id string) error {
	log.Printf("extend running: %s\n", q.running)
	q.Lock()
	defer q.Unlock()

	state, ok := q.running[id]
	if ok {
		state.deadline = time.Now().Add(q.extension)
		return nil
	}
	return ErrNotFound
}

// Info returns internal queue information.
func (q *fifo) GetWorkers(c context.Context) []*Task {
	q.Lock()
	data := []*Task{}

	for w := range q.workers {
		t := <-w.channel
		data = append(data, t)
	}

	q.Unlock()
	return data
}

// Info returns internal queue information.
func (q *fifo) Info(c context.Context) InfoT {
	q.Lock()
	stats := InfoT{}
	var agents []string
	stats.Stats.Workers = len(q.workers)
	stats.Stats.Pending = q.pending.Len()
	stats.Stats.Running = len(q.running)

	for e := q.pending.Front(); e != nil; e = e.Next() {
		stats.Pending = append(stats.Pending, e.Value.(*Task))
	}
	for _, entry := range q.running {
		stats.Running = append(stats.Running, entry.item)
	}

	for w := range q.workers {
		agents = append(agents, w.agent_id)
	}
	stats.Agents = agents

	q.Unlock()
	return stats
}

func (q *fifo) GetRuning(c context.Context, project string) int {
	q.Lock()
	running_count := 0
	for _, entry := range q.running {
		if entry.item.Labels["project"] == project && entry.item.Labels["public"] == "true" {
					running_count = running_count + 1
		}
	}


	q.Unlock()
	return running_count
}


// helper function that loops through the queue and attempts to
// match the item to a single subscriber.
func (q *fifo) process() {
	defer func() {
		// the risk of panic is low. This code can probably be removed
		// once the code has been used in real world installs without issue.
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("queue: unexpected panic: %v\n%s", err, buf)
		}
	}()

	q.Lock()
	defer q.Unlock()

	// TODO(bradrydzewski) move this to a helper function
	// push items to the front of the queue if the item expires.
	for id, state := range q.running {
		if time.Now().After(state.deadline) {
			q.pending.PushFront(state.item)
			delete(q.running, id)
			close(state.done)
		}
	}

	var next *list.Element
loop:
	for e := q.pending.Front(); e != nil; e = next {
	//	running_count := 0
		next = e.Next()
		item := e.Value.(*Task)
		//log.Printf("project: %s", item.Labels["project"])
		if item.Labels["public"] == "true"  {
			//for _, entry := range q.running {
			//	if item.Labels["project"] == entry.item.Labels["project"] && entry.item.Labels["public"] == "true" {
			//		running_count = running_count + 1
			//	}
			//}
			//log.Printf("count: %s", running_count)

			//numURL := os.Getenv("DOUSER_NUM_URL")
			//nums, err := GetNum(numURL + "/" + item.Labels["project"])
			//if err != nil {
			//	log.Printf("get num error")
			//}
			//if nums.ConCount == 0 {
			//		break loop
			//}
			//if running_count > nums.ConCount - 1  {
			//		break loop
			//}
		}
		//if item.Labels["agent"] != "" {
		//	if item.Labels["agent"] != item.Labels["label"] {
		//			break loop
		//	}
		//}
	        log.Printf("workers: %s\n", q.workers)
		for w := range q.workers {
			if w.filter(item) {
				delete(q.workers, w)
				q.pending.Remove(e)

				q.running[item.ID] = &entry{
					item:     item,
					done:     make(chan bool),
					deadline: time.Now().Add(q.extension),
				}

				w.channel <- item
				break loop
			}
		}
	}
}

func main() {
	s := New()

	ctx := context.Background()
	fmt.Println("s", s)
	//fmt.Println("s %v", s.Info(ctx))
	log.Printf("start: %s\n", s.Info(ctx))
	task := &Task{
		ID: "abc",
		Data: []byte("aaa"),
	}
	s.Push(ctx, task)
	//fmt.Println("s %v", s.Info(ctx))
	log.Printf("push: %s\n", s.Info(ctx))
        fn, err := createFilterFunc("filter")
        if err != nil {
	     //return nil, err
	     fmt.Println("s", err)
        }
	s.Poll(ctx, fn, "abc2")
	//fmt.Println("s %v", s.Info(ctx))
	// log.Printf("poll: %s\n", s.Info(ctx))
	res := s.Extend(ctx, "abc")
	log.Printf("extend err: %s\n", res)
	log.Printf("extend: %s\n", s.Info(ctx))
	s.Done(ctx, "abc")
	log.Printf("done: %s\n", s.Info(ctx))
	//fmt.Println("s", s.GetWorkers(ctx))

}
	//Filter struct {
	//	Labels map[string]string `json:"labels"`
	//	Expr   string            `json:"expr"`
	//}

func createFilterFunc(filter string) (Filter, error) {
	return func(task *Task) bool {
	        // log.Printf("task: %s\n", task)
		return true
	}, nil
}

