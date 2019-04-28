package model


type Proc struct {
	ID       string            `json:"id"                  gorm:"primary_key;type:varchar(50);column:proc_id"`
	BuildID  string            `json:"build_id"            gorm:"type:varchar(50);column:proc_build_id"`
	PID      int               `json:"pid"                  gorm:"type:integer;column:proc_pid"`
	Name     string            `json:"name"                 gorm:"type:varchar(250);column:proc_name"`
	State    string            `json:"state"                gorm:"type:varchar(250);column:proc_state"`
	Error    string            `json:"error,omitempty"      gorm:"type:varchar(500);column:proc_error"`
	ExitCode int               `json:"exit_code"            gorm:"type:integer;column:proc_exit_code"`
	Started  int64             `json:"start_time,omitempty" gorm:"type:integer;column:proc_started"`
	Stopped  int64             `json:"end_time,omitempty"   gorm:"type:integer;column:proc_stopped"`
	Children []*Proc           `json:"children,omitempty"   gorm:"-"`
}

// Running returns true if the process state is pending or running.
func (p *Proc) Running() bool {
	return p.State == StatusPending || p.State == StatusRunning
}

// Failing returns true if the process state is failed, killed or error.
func (p *Proc) Failing() bool {
	return p.State == StatusError || p.State == StatusKilled || p.State == StatusFailure
}

// Tree creates a process tree from a flat process list.
func Tree(procs []*Proc) []*Proc {
	var (
		nodes  []*Proc
		parent *Proc
	)
	for _, proc := range procs {
		parent.Children = append(parent.Children, proc)
	}
	return nodes
}

