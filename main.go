package main

import (
	"fmt"
	"math/rand"
)

func getName() {
	name := ""
	fmt.Println("Enter your name")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Hello %s", name)
}

func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("\nEnter your bet or enter 0 to quit(balance: $%d): ", balance)
		fmt.Scan(&bet)
		if bet > balance {
			fmt.Println("Bet can't be greater than balance")
		} else {
			break
		}
	}
	return bet
}

func generateReel(symbols map[string]uint) []string {
	reel := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			reel = append(reel, symbol)
		}
	}
	return reel
}

func createBet(reel []string) []string {
	spinResult := []string{}
	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(30)
		spinResult = append(spinResult, reel[randomIndex])
	}

	return spinResult
}

func main() {
	balance := uint(200)
	symbols := map[string]uint{
		"A": 3,
		"K": 6,
		"Q": 9,
		"J": 12,
	}
	// multiplyer := map[string]uint{
	// 	"A" : 20,
	// 	"K" : 15,
	// 	"Q" : 10,
	// 	"J" : 5,
	// }

	reel := generateReel(symbols)

	spin := createBet(reel)
	fmt.Println(spin)
	getName()
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		fmt.Printf("Your Bet is: $%d.\n", bet)

		balance = balance - bet

	}
	fmt.Printf("Your remaining balance: $%d. \n", balance)
}
