package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		allowRepetition := true
		fmt.Println(generateRandomPassword(allowRepetition))
		time.Sleep(1)
	}
}

func generateRandomPassword(allowRepetition bool) [4]int {
	var password [4]int

	rand.Seed(time.Now().UTC().UnixNano())

	min := 1
	max := 8

	if allowRepetition {
		for i := 0; i < 4; i++ {
			password[i] = rand.Intn(max-min+1) + min
		}
	} else {
		var permutation = rand.Perm(max - min + 1)

		for i := 0; i < 4; i++ {
			password[i] = permutation[i] + min
		}
	}

	return password
}
