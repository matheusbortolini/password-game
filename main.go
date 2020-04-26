package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	allowRepetition := false
	maxGuesses := 1
	password := generateRandomPassword(allowRepetition)

	fmt.Println(password)

	var guess [4]int

	for guessCount := 0; guessCount < maxGuesses; guessCount++ {
		guess = getUserGuess(guessCount)
		fmt.Println(guess)
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

func getUserGuess(guessCount int) [4]int {
	var guess [4]int
	scanner := bufio.NewScanner(os.Stdin)

OUTER:
	for {
		fmt.Printf("Tentativa %d: ", guessCount)

		if scanner.Scan() {
			line := scanner.Text()
			stringGuess := strings.Fields(line)

			if len(stringGuess) != 4 {
				fmt.Println("Chute deve conter exatamente 4 números!")
				continue
			}

			for position, stringGuessElement := range stringGuess {
				intGuessElement, err := strconv.Atoi(stringGuessElement)

				if err != nil || intGuessElement < 1 || intGuessElement > 8 {
					fmt.Printf("Entrada '%s' inválida!\n", stringGuessElement)
					continue OUTER
				}

				guess[position] = intGuessElement
			}

			break
		}
	}

	return guess
}
