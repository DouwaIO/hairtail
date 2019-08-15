package even

import (
	"fmt"

	"github.com/DouwaIO/hairtail/src/task"
)

func main() {
    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}]}}]"
    str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}]},\"data1\":{\"order_no\":\"001\",\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}]}}]"
    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}]}}]"
    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"},{\"name\":\"特步\"}]}}]"
    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"},{\"name\":\"特步\"}]}}]"

    params := task.Params{
        Data:     []byte(str),
    }
    tk = new(Plugin)
	result, err := tk.Run(&params)
	if err != nil {
		fmt.Printf("task run error: %s\n", err)
	}
    fmt.Printf("\nresult: \n%s\n", string(result.Data))
}
