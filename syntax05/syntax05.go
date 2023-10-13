package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

// shadowing
func main(){
	// 자료타입을 변수명으로 사용
	var test1 float64 = 9.1
	var test2 float64 = 7.01
	var univ string = "inha"
	var f1 string = "functions"
	var f2 = append([]string{}, "함수")
	fmt.Println(test1 + test2)
	fmt.Println(univ)
	fmt.Println(f1)
	fmt.Println(f2)
	
	// 자료타입을 변수명으로 사용
	// var float64 float64 = 9.1
	// var test float64 = 7.01
	// fmt.Println(float64)
	// 패키지를 변수명으로 사용
	// var fmt string = "inha"
	// fmt.Println()
  // 함수를 변수명으로 사용
	// var append string = "functions"
	// var f = append([]string{}, "함수")

}
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	// "math/rand"
// 	// "time"
// )

// // 입력(0과 1처리))된 소수 판정 프로그램 v.0
// // 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
// func main(){
// 	var number int
	
// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}


// 	for number < 2{
// 		fmt.Println(number, "는(은) 소수가 이닙니다!")
// 		os.Exit(0)
// 		// fmt.Println(n)
// 	}
	
// 	isPrime := true

// 	for i:=2; i< number; i++ {
// 		if number % i == 0{
// 			isPrime = false
// 			break // 첫 번째 약수가 발경되면 반복문 즉시 종료
		
// 		}

// 	}
// 	if isPrime { // 비교 연산 제거
// 			fmt.Println(number, "는(은) 소수입니다!")
// 		} else {
// 			fmt.Println(number, "는(은) 소수가 이닙니다!")

// 	}
// }


// package main

// import (
// 	"fmt"
// 	"log"
// 	// "math/rand"
// 	// "time"
// )

// // 입력(fmt 패키지의 Scanln())된 소수 판정 프로그램 v.0
// // 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
// func main(){
// 	var number int
// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number)

// 	// fmt.Println(n)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	isPrime := true

// 	for i:=2; i< number; i++ {
// 		if number % i == 0{
// 			isPrime = false
// 			break // 첫 번째 약수가 발경되면 반복문 즉시 종료
		
// 		}

// 	}
// 	if isPrime { // 비교 연산 제거
// 			fmt.Println(number, "는(은) 소수입니다!")
// 		} else {
// 			fmt.Println(number, "는(은) 소수가 이닙니다!")

// 	}
// }
// package main

// import (
// 	"fmt"
// 	"strconv"
// 	// "math/rand"
// 	// "time"
// 	"bufio"
// 	"log"
// 	"os"
// 	"strings"
// )

// // 난수 추출 소수 판정 프로그램 v.06
// // 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
// func main(){
// 	fmt.Print("정수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err:= rd.ReadString('\n')
	
// 	if err != nil{ // 에러가 발생하면
// 		log.Fatal(err)
// 	} 

// 	in = strings.TrimSpace(in)
// 	// dan, err := strconv.ParseInt(in, 10, 32)
// 	number, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}


// 	isPrime := true

// 	for i:=2; i< number; i++ {
// 		if number % i == 0{
// 			isPrime = false
// 			break // 첫 번째 약수가 발경되면 반복문 즉시 종료
		
// 		}

// 	}
// 	if isPrime { // 비교 연산 제거
// 			fmt.Println(number, "는(은) 소수입니다!")
// 		} else {
// 			fmt.Println(number, "는(은) 소수가 이닙니다!")

// 	}
// }