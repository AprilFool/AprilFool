package main

import (
	"fmt"
	"github.com/AprilFool/AprilFool/httpd"
	"github.com/AprilFool/AprilFool/task"
)

var messages = make(chan int, 10)

func main() {
	go httpd.Start(8080)
	go task.Say("world", 1000)
	for m := range messages {
		fmt.Println("%v", m)
	}
}
