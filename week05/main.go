package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // get the current data and time as an integer
	answer := rand.Intn(100) + 1// random integer number between (1 ~ 100)
	fmt.Println("Guess Number Game~")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input guess nuber : ")
	inputNumber, err := reader.ReadString('\n')
	if err = nil {
		log.Fatal(err)
	}
	
}
	