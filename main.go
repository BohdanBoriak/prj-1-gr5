package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var points int = 50

const pointPerQuestion = 5

func main() {
	fmt.Println("Вітаю у грі HARGCORE-MATH!")
	time.Sleep(2 * time.Second)

	for i := 5; i > 0; i-- {
		fmt.Printf("До початку: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	for myPoints < 50 {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y

		fmt.Println(x, "+", y, "=")

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if res == ansInt {

			}
		}

	}
}
