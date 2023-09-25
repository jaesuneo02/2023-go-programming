package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input Score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err:= reader.ReadString('\n') // option 2
	if err != nil{
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace((inputScoreString)) // remove space bar
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) // casting
	if err != nil{
		log.Fatal(err)
	}
	var grade string
	if inputScore >= 90 { 
		grade = "A grade!"
	} else {
		grade = "under A grade~"
	}
	fmt.Println("you got " + grade)
}
	