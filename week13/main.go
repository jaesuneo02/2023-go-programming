package main

import "fmt"

func main() {

	var a []string
	a = make([]string, 4,5)
	fmt.Printf("%#v\n",a)
	fmt.Println(a,len(a), cap(a))
}