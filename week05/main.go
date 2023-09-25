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
	if inputScore >= 90 { //mismatched types string and untyped int
		grade := "A grade!"
	} else {
		grade := "BCDE grade~"
	}
}
	