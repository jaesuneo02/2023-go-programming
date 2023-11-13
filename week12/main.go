package main

import "fmt"

func main() {
	a := make([]string,4,5) // 타입, 개수, 수용량
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "z"

	fmt.Println(a, len(a),cap(a))
}

