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
}