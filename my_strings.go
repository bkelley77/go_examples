package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "bruce allen kelley, jr"

	fmt.Println("Strings:")
	fmt.Println("  name = ", name)
	fmt.Println("  name[0] = ", name[0])
	fmt.Println("  name[1:3] = ", name[1:3])
	fmt.Println("  name[:3] = ", name[:3])

	pos := strings.IndexAny(name, ",")
	if pos == -1 {
		fmt.Println("Comma was not found")
	} else {
		fmt.Println("Command found at position ", pos)
	}


}