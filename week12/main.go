package main

import "fmt"

func main() {

	a := []string{"a", "b", "c", "d"}

	as := a[0:2]
	as[1] = "z"
	fmt.Println(a)
	fmt.Println(as)

	b := [4]int(4,3,2,1)
	bs := b[1:3]
	fmt.Println(bs) 
}

