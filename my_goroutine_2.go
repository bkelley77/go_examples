package main

import (
	"fmt"
	"time"
)

func function() {

	// This go function runs forever...
	for {
		fmt.Println("func> go function...")

		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		fmt.Println("func> going to sleep in go function...")

		time.Sleep(5 * time.Second)
		fmt.Println("func> waking up in go function...")
	}
}

func main() {
	go function()

	// run in main forever...
	for {
		fmt.Println("main> Going to sleep in main...")
		time.Sleep(1 * time.Second)
		fmt.Println("main> Woke up in main...")
	}
}
