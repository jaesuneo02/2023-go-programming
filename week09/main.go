package main

import "fmt"


func main(){
	a := 10
	// b := 20 // var b int = 20
	
	c := &a
	fmt.Printf("%T\n", c)
}
