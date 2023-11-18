package main

import (
	"fmt"
	"runtime"
)

func main() {
	//비동기 : 버퍼 사용
	//버퍼 : 발신, 가득 찰때까지 대기
	//		수신, 비어 있으면 대기, 가득 차 있으면 작동

	runtime.GOMAXPROCS(1)

	ch := make(chan bool, 2) // 2: 채널 버퍼 용량, 버퍼가 가득 찰때까지 대기했다가 버퍼가 빌 때까지 수신
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			fmt.Println("GO : ", i+1)
			ch <- true
			//time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main: ", i+1)
	}

}
