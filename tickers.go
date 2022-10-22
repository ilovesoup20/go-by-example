package main

import (
	"fmt"
	"time"
)

func main() {
	tickExample()
	// select {}
}

func tickExample() {
	c := time.Tick(1 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}
func statusUpdate() string { return "foo" }

func tickExample2() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				go sayHi()
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func sayHi() {
	fmt.Println("Saying hi!")
}
