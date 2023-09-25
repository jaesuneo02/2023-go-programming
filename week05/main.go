package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // get the current data and time as an integer
	answer := rand.Intn(100) + 1// random integer number between (1 ~ 100)
	fmt.Println("Guess Number Game~")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)
	for i :=0; i< 10; i++ {
		fmt.Println("You have", 10-i, "chances~")
		fmt.Print("Input guess nuber : ")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)
		if err != nil {
			log.Fatal(err)
		}
		if inputNumber == answer {
			fmt.Println("Great U got number. congratulations~") // Answer is higher~
			

		}
		else if inputNumber < answer{
			fmt.Println("Guess number is lower then answer") // Answer is higher~
		} else if inputNumber > answer{
			fmt.Println("Guess number is higher then answer") // Answer is lower
		}
}
	
}
	