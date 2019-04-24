package pipeline

import (
	//"log"
	"fmt"
)

func Test2(data []byte) ([]byte, error) {
	cc := string(data) + "qweqwe"
	fmt.Println("Data :", string(data))
	return []byte(cc), nil
}

func Test(data []byte) error {
	fmt.Println("Data :", string(data))
	return nil
}
