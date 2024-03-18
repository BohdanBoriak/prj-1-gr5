package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"project-go/domain"
	"sort"
	"strconv"
	"time"
)

var id uint64 = 1

const (
	totalPoints      = 5
	pointPerQuestion = 5
)

func main() {
	fmt.Println("Вітаю у грі HARGCORE-MATH!")
	time.Sleep(2 * time.Second)

	var users []domain.User
	users = append(users, domain.User{
		Id:       1,
		NickName: "Katastrofa",
		Time:     5 * time.Second})
	users = append(users, domain.User{
		Id:       2,
		NickName: "Revenger",
		Time:     3 * time.Second})
	users = append(users, domain.User{
		Id:       3,
		NickName: "Bohdan",
		Time:     10 * time.Second})

	sortAndSave(users)

	// for {
	// 	menu()
	// 	punct := ""
	// 	fmt.Scan(&punct)

	// 	switch punct {
	// 	case "1":
	// 		u := play()
	// 		users = append(users, u)
	// 	case "2":
	// 		for _, user := range users {
	// 			fmt.Printf("Id: %v, Name: %s, Time: %v\n",
	// 				user.Id, user.NickName, user.Time)
	// 		}
	// 	case "3":
	// 		return
	// 	}
	// }
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

	points := totalPoints
	myPoints := 0
	start := time.Now()
	for myPoints < totalPoints {
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

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile(
		"users.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
