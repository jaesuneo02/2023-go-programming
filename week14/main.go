package main

import "fmt"



func status(name string){
	balls := map[string]int{"성기훈" : 20, "오일남" : 0,}
	// ball := balls[name]
	// fmt.Println(ball)
	ball, exists := balls[name]
	if !exists{
		fmt.Println(name, "님은 게임 참옂가 아닙니다.")
	} else if ball < 1 {
		fmt.Println(name, "님은", balls[name], "개로 탈락~ 탕탕")
		} else {
			fmt.Println(name, "님은 게임에서 승리하셨습니다.")
	}
}

func main(){
	status("오일남")
	status("강철")
	status("성기훈")
	
}