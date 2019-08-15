package accumulate

import (
	"fmt"

	"github.com/DouwaIO/hairtail/src/task"
)

func main() {
    // mySetting := make(map[string]interface{})
	// mySetting["rename"] = []string{"b=bbbbb", "data__sales__name=sale_name"}

    str := "[{\"order_no\":\"001\",\"sale_name\":\"张三\",\"qty\":30},{\"order_no\":\"001\",\"sale_name\":\"李四\",\"qty\":50}]"
    params := task.Params{
        Settings: map[string]interface{
			"rename":  []string{"b=bbbbb", "data__sales__name=sale_name"},
    		"map":     []string{"sale_name=sale_full_name"},
			"source":  "qty",
			"target": "qty",
			"compute": "+",
			"ignore": false,
		}
        Data:     []byte(str),
    }

    tk = new(Plugin)
	result, err := tk.Run(&params)
	if err != nil {
		fmt.Printf("task run error: %s\n", err)
	}
    fmt.Printf("result is nothing")
}
