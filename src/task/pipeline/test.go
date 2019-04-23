package pipeline

import (
    "log"
)

func Test2(data []byte) ([]byte, error) {
	cc := string(data) + "qweqwe"
	log.Printf("Data :", string(data))
	return []byte(cc), nil
}

func Test(data []byte) error {
	log.Printf("Data :", string(data))
	return nil
}
