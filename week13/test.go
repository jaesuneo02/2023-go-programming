package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args) // Args == Arguments
  fmt.Println(os.Args[2])
}