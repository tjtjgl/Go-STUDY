package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	s := "Goroutine Closure: "

	runtime.GOMAXPROCS(1)

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) //반복문 클로저는 일반적으로 즉시 실행
		// 하지만 고루틴 클로저는 가장 나중에 실행(반복문 종료 후)
		// 반복문 끝나길 기다렷다가 실행되어 어떤게 먼저 나올지 또 모름
	}

	time.Sleep(3 * time.Second)

}
