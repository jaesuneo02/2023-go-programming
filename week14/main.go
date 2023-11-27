package main

import "fmt"



func main(){
	games := map[int]string{
		456 : "성기훈",
		218 : "박해수",
		67 : "강새벽",
		1 : "오일남",
		199 : "알리",
		101: "아이오아이",
	}
name, ok := games[100]
fmt.Println(name, ok)

	for _, v := range games{
		fmt.Println(v)
	}
	games[101] = "장덕수" // update
	delete(games,199) // delete
	for k, v := range games {
		fmt.Println(k,v)
	}
}