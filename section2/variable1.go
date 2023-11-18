package main

import "fmt"

func main() {
	var a int
	var b string

	var c, d, e int
	var f, g, h int = 1, 2, 3 // 선언 동시 초기화

	var i float32 = 11.4

	var j string = "hi golang"

	var k = 4.74 // 자료형 자동 인식

	var l = "hi"

	var m = true

	a = 4
	b = "wow"
	c, d, e = 5, 6, 7

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c,d,e: ", c, d, e)
	fmt.Println("f,g,h: ", f, g, h)

	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)
}
