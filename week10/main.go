package main

import (
	"fmt"
	"week10/src/greeting"
	"week10/src/mymath"
)

func main() {
	fmt.Println(mymath.MyPower(2,9))
	fmt.Println(mymath.myAbs(-99))
	greeting.Hello()
	fmt.Println(mymath.myAbs(17))
	greeting.Hi()

}