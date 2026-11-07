package main

import (
	"fmt"
	"time"
)

func main() {
	username := "Nigora"
	goal := "Master Go and get into Vistula University"
	currentYear := time.Now().Year()

	fmt.Printf("Hello, %s!\n", username)
	fmt.Printf("Year: %d\n", currentYear)
	fmt.Printf("Current Goal: %s\n", goal)
	fmt.Println("Status: Code running succesfully!")
}
