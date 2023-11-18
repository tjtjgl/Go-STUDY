package main

import (
	"fmt"
	"time"
)

func main() {
	//동기 : 버퍼 미사용
	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			fmt.Println("GO : ", i)
			ch <- true // 보내고 다음 for 문이 실행 되는 것이 아니라 수신 받을 때까지 대기
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch // 수신
		fmt.Println("Main: ", i)
	}

}
