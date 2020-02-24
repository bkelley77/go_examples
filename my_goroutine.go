package main

import (
	"fmt"
	"time"
)

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func main() {
	go function()

	go func() {

		for {
			fmt.Println("Going to sleep in main...")

			for i := 10; i < 20; i++ {
				fmt.Print(i, " ")
				time.Sleep(1 * time.Second)
			}
		}

	}()

	for {

		fmt.Println("Going to sleep in main...")
		time.Sleep(1 * time.Second)
		fmt.Println("Woke up in main...")
	}
}

