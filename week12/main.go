package main

import "fmt"

func main() {

	s := []int{0,0,0,0,0}
	
	s[4] = 99
	s[2] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]

	for _, value := range copyS {
		fmt.Println(value)
	}
	test := [3] string{"inha","go","student"} // 배열 리터럴을 이용해서 test 배열 생성
	// testS := test[0:4]                     //invalid argument: index 4 out of bounds [0:4]
	testS:= test[:2] //testS := test[0:2]
	testS2 := testS[1:]
	testS2[0] = "python"
	// fmt.Println(testS2[1])
	
	fmt.Println(testS2)
	fmt.Println(testS, len(testS))
	fmt.Println(test)

}