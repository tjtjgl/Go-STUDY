package main

import (
	"fmt"
	"time"
)

// int만 들어올 수 있는(int만 송/수신) 채널 v 생성
func work1(v chan int) {
	fmt.Println("work1 : Start --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work1 : Eend --->", time.Now())

	v <- 1 //1을 채널로 송신
}

func work2(v chan int) {
	fmt.Println("work2 : Start --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work2 : Eend --->", time.Now())

	v <- 2 //2을 채널로 송신
}

func main() {
	//채널
	//고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용 : 채널(동ㄹ기식, 레퍼런스 타입)
	// 실행 흐름 제어 가능(동기, 비동기) -> 채널 자체를 일반 변수로 선언 후 사용 가능

	// 데이터 전달 자료형 선언 후 사용(지정된 타입만 주고 받을 수 있음)
	//interface{} 모든 자료형을 받을 수 있는, 자료형 상관없이 전송 및 수신 가능
	//값을 전달(복사 후:bool, int 등), 포인터(슬라이스,맵)등을 전달시에는 주의 왜? 원본값이 변할 수 있기 떼문에, 동기화 사용(뮤텍스)
	//멀티프로세싱 처리에서 교착 상태(경쟁) 주의
	//<-, ->
	// 채널 <- 데이터 : 송신, 데이터를 채널로 보내겠다.
	// <- 채널 : 수신, 고루틴에서 채널로 보낸 데이터를 채널에서 받아온다.

	fmt.Println("Main Strart ----> ", time.Now())

	//var c chan int
	//c = make(chan int)

	v := make(chan int) // int형 채널 선언

	go work1(v)
	go work2(v)

	//수신 될 때 까지 대기. sleep이 필요없음
	<-v
	<-v
	fmt.Println("Main End ----> ", time.Now())

}
