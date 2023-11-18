package main

import "fmt"

func main() {
	//채널 Close
	//주의! 닫힌 채널에 값 전송 시 패닉(예외) 발생
	//Range : 채널 안에서 값을 꺼낸다.(순회), 채널 닫아야 반복문 종료
	// 		-> 채널이 열려있고 값을 전송하지 않으면 무한 대기

	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기

	}()

	// 채널에서 값을 꺼내온다. 채널이 클로즈 될 때까지!
	for i := range ch {
		fmt.Println("Ex1 : ", i)
	}
}
