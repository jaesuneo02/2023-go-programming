package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main(){

	texts := "G@ M@ney~~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)
	
	// 변수명은 영문자로 시작해야된다.
	// 영문 대문자의 경우 다른 패키지에서 접근할 수 있다.
	//소문자로 시작하는 변수는 동일 패키지에서만 접근 가능
	// zero value 값을 활당하지 않았을 때 생기는 값
	var e string
	var d bool
	var c rune
	var a int
	var b float64

	// naming conventtion
	var studentId string 
	var i int32 
	// a := 7

	fmt.Println(studentId)
	fmt.Println(i)
	
	// zero value
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))

}