package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("玩一个猜数游戏,输入1玩,输入2不玩")
	var isPlay int
	fmt.Scanln(&isPlay)

	switch isPlay {
	case 1:
		game()
	case 2:
		fmt.Println("哦哦好的")
		time.Sleep(5 * time.Second)
	default:
		fmt.Println("你到底玩不玩")
		time.Sleep(5 * time.Second)
	}
}

func game() {
	rand.Seed(time.Now().UnixNano())
	targetNumber := rand.Intn(100)

	fmt.Println("输入一个0-100的数字,看看跟我想的一样吗")
	status := false
	var trial int
	for status == false {
		var input int
		fmt.Scanln(&input)

		if input == targetNumber {
			trial++
			fmt.Println("恭喜主人用", trial, "次就猜对啦")
			time.Sleep(10 * time.Second)
			status = true
		} else if input < targetNumber {
			trial++
			fmt.Println("小了喵")
		} else if input > targetNumber {
			trial++
			fmt.Println("大了喵")
		}

	}

}
