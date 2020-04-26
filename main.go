package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var password = generateRandomPassword()

	fmt.Println(password)
}

func generateRandomPassword() [4]int {
	var password [4]int

	rand.Seed(time.Now().UTC().UnixNano())

	min := 1
	max := 8

	for i := 0; i < 4; i++ {
		password[i] = rand.Intn(max) + min
	}

	return password
}
