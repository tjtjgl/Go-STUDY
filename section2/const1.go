package main

import "fmt"

func main() {

	// 선언과 동시에 초기화 필요
	// 한번 선언 후 값 변경 금지
	const a = "TEST1"
	fmt.Println(a)

	// const d = getHeight() 변할 수 있는 함수 리턴 값을 받을 수 없음

}
