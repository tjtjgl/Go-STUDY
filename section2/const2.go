package main

import "fmt"

func main() {

	const (
		a, b    int = 10, 99
		c, d, e     = true, 0.84, "test"
	)

	fmt.Println("a,b: ", a, b, "c,d,e: ", c, d, e)

}
