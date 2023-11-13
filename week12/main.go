package main

import "fmt"

func main() {

	// var s []int
	// s = make([]int, 5) // make함수를 이용해 슬라이스 생성 후 메모리 할당, 제로 값 적용
	
	// s :=make([]int, 5)
	s := []int{0,0,0,0,0} //  슬라이스 리터럴을 이용해 슬라이스 생성 및 메모리 할당, 초기화 진행
	for _, value := range s {
		fmt.Println(value)
	}
}