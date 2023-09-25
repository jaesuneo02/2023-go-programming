package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err:= reader.ReadString('\n') // option 2
	log.Fatal(err)
	fmt.Println(inputScore)
}