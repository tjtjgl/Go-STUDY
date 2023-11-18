package main

import "fmt"

func main() {
	// go에만 있는 기능, 짧은 선언, 일회성으로 사용
	// 반드시 함수 안에서만 사용(전역 X)
	// 선언 후 재할당 시 예외 발생
	// 주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있다.
	// 프로시저나 프록시

	shortVar1 := 3
	shortVar2 := "TEST"
	shortVar3 := false

	// shortVar3 := true 예외 발생
	fmt.Println(shortVar1, shortVar2, shortVar3)

	// example
	if i := 10; i < 11 {
		fmt.Println("short variable test success!")
	}

}
