package main

import (
	"time"
)

func channels() {
	ch := make(chan int, 4)

	ch <- 12
	ch <- 3
	ch <- 4

	_ = <-ch

	//close(ch)

	go func() {
		for {
			ch <- 55
		}
	}()

	time.Sleep(1 * time.Second)

	return
}
