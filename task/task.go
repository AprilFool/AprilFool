package task

import (
	"fmt"
	"time"
)

func Say(s string, i int) {
	for {
		fmt.Println(s)
		time.Sleep(10000 * time.Millisecond)
	}
}
