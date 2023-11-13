package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}


func main() {
	numbers , err:= GetFloats("data.txt")
	sum := 0.0
	for  _, v := range numbers{
		sum = sum + v 
	}
	avg := sum/float64(len(numbers))
	
	fmt.Println(avg)
	if err != nil{
		log.Fatal(err)
	}
	

	// f, err := os.Open("data.txt")
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// scanner := bufio.NewScanner(f)

	// for scanner.Scan(){
	// 	fmt.Println(scanner.Text())
	// }
	// err = f.Close()
	// if err != nil {
	// 	log.Fatal(err)
		
	// }
}

