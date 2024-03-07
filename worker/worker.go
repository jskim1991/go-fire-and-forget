package worker

import (
	"fmt"
	"time"
)

func Work(ch chan string) {
	for {
		select {
		case input := <-ch:
			fmt.Println("Worker: Preparing...", input)
			time.Sleep(50 * time.Millisecond)
			fmt.Println("Worker: Processed", input)
			fmt.Println()
		}
	}

}
