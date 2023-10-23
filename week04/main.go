package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	
	// var a int // declaration
	// a= 9 // assign value

	// var a int = 9 // declaration & assign value
	// var a = 9 // declaration & assign value
	
	a := 9
	// b := 2.71
	var b float32 = 2.71
	c:= 'z'
	d, e := 4.10, 3.99
	f := "문자열"
	var g int
	var h float32
  var i bool
	var j string
  K := "변수명이 대문자로 시작하면 다른 패키지에서도 이 변수를 사용할 수 있음"
  // koreanzombie := "정찬성"
  koreanZombie := "정찬성"
	fmt.Println(j, g, h, i, K, koreanZombie)
	fmt.Println(
		a, reflect.TypeOf(a),
	  b, reflect.TypeOf(b),
	  c, reflect.TypeOf(c),
		d, reflect.TypeOf(d),
		e, reflect.TypeOf(e),
	  f, reflect.TypeOf(f))
  fmt.Println(float32(a) * b)
  fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)


	fmt.Println(reflect.TypeOf('z'))			// rune, int32
	fmt.Println(reflect.TypeOf(99))			
	fmt.Println(reflect.TypeOf(2.7))			
	fmt.Println(reflect.TypeOf(false))			
	fmt.Println(reflect.TypeOf("go!"))	
	fmt.Println('A', 'a', '0', '김', '인', '하', '\n')// 65, 97, 48, 44608  , 10
	fmt.Println(math.Floor(3.17))
	fmt.Println(math.Ceil(3.17))
	fmt.Println(strings.Title("오픈소스 프로그래밍~"))
	fmt.Println(strings.Title("open source\t programming~\n\"go\"!"))
	// fmt.Println(math.Floor("go!"))
	// fmt.Println(strings.Title(3.14))
}