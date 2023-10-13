package main

import (
	"fmt"
	"math"
	"strings"
)

func main(){
	var c rune = '가' // 유닛코드 한글자
	// var a int16 = 7
	//var a = 7
	//a := 7
	a := 7
	var b float64 = 5.3
	a = int(b) // Type Conversion, casting
	d := false
	// 타입이 바뀌지는 않는다 / 강타입 언어
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c) // 유니코드(UTF-08) 값 출력
	fmt.Printf("%c\n", c) // 한 글자 출력
	fmt.Printf("%T\n", c) // rune은 결국 int32의 별명


	fmt.Println(math.Round(2.71))
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java"))

}