package main

import "fmt"

func main() {
	//Close: 채널 닫기

	// 데이터 수신 여부 두번째 변수로 받아 볼 수 있음
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 7777

		}
	}()

	val1, ok1 := <-ch
	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex1 : ", val3, ok3)

	//close(ch)
	val4, ok4 := <-ch
	fmt.Println("ex1 : ", val4, ok4)

}
