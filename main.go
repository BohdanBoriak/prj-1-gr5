package main

import (
	"fmt"
	"math/rand"
	"project-go/domain"
	"strconv"
	"time"
)

var (
	points int    = 50
	id     uint64 = 1
)

const pointPerQuestion = 5

func main() {
	fmt.Println("Вітаю у грі HARGCORE-MATH!")
	time.Sleep(2 * time.Second)

	for {
		menu()
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
		case "2":
			fmt.Println("Рейтинг в розробці...")
		case "3":
			return
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 5; i > 0; i-- {
		fmt.Printf("До початку: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	start := time.Now()
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
				myPoints += pointPerQuestion
				points -= pointPerQuestion
				fmt.Printf(
					"Чудово! Ви набрали %v очок, залишилось %v\n",
					myPoints, points,
				)
			} else {
				fmt.Println("Спробуй ще!")
			}
		}
	}

	finish := time.Now()
	timeSpent := finish.Sub(start)

	fmt.Printf("Вітаю, ти впорався за %v!", timeSpent.Seconds())

	fmt.Println("Введіть свій нікнейм:")
	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:       id,
		NickName: name,
		Time:     timeSpent,
	}
	id++

	return user
}
