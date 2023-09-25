package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _:= reader.ReadString('\n') // option 1 ,  Ignore the error return value with the blank identifier
	fmt.Println(inputScore)
}