package main

import "fmt"

func main() {
	var x[5] int
	name := "bruce"

	fmt.Println("Strings:")
	fmt.Println("  name = ", name)

	x[0] = 100
	x[4] = 401

	fmt.Println("Iterate through and array named x...   for _, value := range x")
	for _, value := range x {
		fmt.Println("   value = ", value)
	}
	fmt.Println("   length of array x is ", len(x))

	fmt.Println("Iterate using the folling... for i := 0;i < 15; i++")
	for i := 0; i < 15; i++ {
		if i % 2 == 0 {
			fmt.Println("   i = ", i, " and is a even number")
		} else {
			fmt.Println("   i = ", i, " and it is and odd number")
		}
	}

	fmt.Println("x = ", x)
}
