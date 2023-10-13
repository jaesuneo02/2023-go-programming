// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log" //Fatal function
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main(){
// 	fmt.Print("단 입력: ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err:= rd.ReadString('\n')
// 		if err != nil{ // 에러가 발생하면
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)
// 	// dan, err := strconv.ParseInt(in, 10, 32)
// 	dan, err := strconv.Atoi(in)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		for i := 1; i<10; i++ {
//			// fmt.Println(dan , " * ", i, " = ", (int(dan)*i))
//			// fmt.Println(dan , " * ", i, " = ", (dan*i))
//			fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"log" //Fatal function
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Print("단 입력: ")
	rd := bufio.NewReader(os.Stdin)
	in, err:= rd.ReadString('\n')
		if err != nil{ // 에러가 발생하면
		log.Fatal(err)
	} 

	in = strings.TrimSpace(in)
	// dan, err := strconv.ParseInt(in, 10, 32)
	dan, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	//다른 언어의 while문 구현
	i := 1
	for i< 10 {

		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
		i++
		} 
}