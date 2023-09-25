package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time 
	now = time.Now()
	year := now.Year()
	var month string = now.Month().String()

	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}