package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)

	fmt.Println("exe1 func end", time.Now())

}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)

	fmt.Println("exe2 func end", time.Now())

}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)

	fmt.Println("exe3 func end", time.Now())

}

func main() {

	// 고루틴
	// 타 언어의 쓰레드와 비슷한 기능
	// 생성 방법도 간단하고 리소스도 매우 적게 사용
	// 즉, 수많은 고루틴 동시 생성 실행 가능
	// 비동기적 함수 루틴 실행 -> 채널을 통한 통신 가능
	// 공유메모리 사용 시에 정확한 동기화 코딩 필요
	// 싱글 루틴에 비해 항상 빠른 처리 결과는 아니다.

	// 멀티스레드의 장점과 단점
	// 장점 : 응답섭 향상, 자원공유를 효율적으로 활용 가능, 작업이 분리되어 코드 간결
	// 단점 : 구현하기 어려움, 테스트/디버킹 어려움, 전체 프로세스에 사이드이팩트, 성능 저하, 동기화 코딩, 데드락(교착 상태)

	// main 함수가 끝나면 고루틴도 같이 끝남 = 데몬스레드

	exe1() // 가장 먼저 실행(일반적인 흐름)

	fmt.Println("Main Routine Start", time.Now())

	// 순서는 각 루틴마다 별도로 실행되기 때문에 어떤게 먼저 실행될지 모름
	go exe2()
	go exe3()

	time.Sleep(4 * time.Second) // 데몬스레드 방지를 위해

	fmt.Println("Main Routine End", time.Now())

}
