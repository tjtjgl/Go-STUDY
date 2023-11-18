package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 운영체제 명령어에 접근할 수 있음

func exec(name int) {
	r := rand.Intn(100) //100 미만의 랜덤한 수를 가져온다.

	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", r, i+1)

	}

	fmt.Println(name, " end : ", time.Now())
}
func main() {

	//고루틴
	//멀티 코어(다중 CPU) 최대한 활용
	//복잡한 수치 계산에 필요(빅데이터, GPU 사용)

	//내 PC 코어 개수 반환 후 설정
	//runtime.NumCPU() : 내 PC의 CPU 개수를 가져옴
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // 설정 값

	fmt.Println("Main Routine Start: ", time.Now()) // 퍼포먼스 시간 측정을 위해
	for i := 0; i < 100; i++ {
		go exec(i) // 고루틴 100개 생성
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Main Routine End: ", time.Now()) // 퍼포먼스 시간 측정을 위해

}
