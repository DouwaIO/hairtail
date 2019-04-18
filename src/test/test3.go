package main

import (
	"fmt"
	"github.com/DouwaIO/hairtail/src/schema"
)
func main() {

	test := `version: v1
kind: model
name: storage
columns:
  - name: id
    desc: a identity data
    type: int
    default: 0
  - name: name
    desc: the name of people
    type: int
    default: ''
  - name: sex
    desc: '0: male, 1: female'
    type: int
    default: 0
    `
	branches, err := schema.ParseString(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(branches)
}
