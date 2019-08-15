package selec

import (
	"fmt"

	"github.com/DouwaIO/hairtail/src/task"
)

func main() {
    // mySetting := make(map[string]interface{})
	// mySetting["rename"] = []string{"b=bbbbb", "data__sales__name=sale_name"}

    str := "[{\"a\":1,\"b\":2,\"c\":3}]"
    params := task.Params{
        Settings: map[string]interface{
			"rename":  []string{"b=bbbbb", "data__sales__name=sale_name"}
		}
        Data:     []byte(str),
    }

    tk = new(Plugin)
	result, err := tk.Run(&params)
	if err != nil {
		fmt.Printf("task run error: %s\n", err)
	}
    fmt.Printf("\nresult: \n%s\n", string(result.Data))
}
