package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	//welcome message
	fmt.Println("Welcome to the guessing game!")
	fmt.Println("I have chosen the number beteween 1 and 100")
	fmt.Println("Can you guess what number is it?")

	var guess int
	for {
		fmt.Println("enter your guess: ")
		fmt.Scanln(&guess)

		//check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations! you guessed the correct numbesg]gq.bkmkgrmkkmkfffer!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}

	}

}
