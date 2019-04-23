package main

import (
	"time"
	"fmt"
)

func main() {
	ticker:=time.NewTicker(time.Second*5)
	    go func() {
	        for _=range ticker.C {
			fmt.Println("test")
	        }
	    }()
	
	time.Sleep(time.Minute)
}
