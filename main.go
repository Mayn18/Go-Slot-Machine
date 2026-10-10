package main

import "fmt"

func getName() {
	name := ""
	fmt.Println("Enter your name")
	_, err :=  fmt.Scanln(&name)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("Hello %s", name)
}

func getBet(balance uint) uint{
	var bet uint
	for true{
		fmt.Printf("\nEnter your bet or enter 0 to quit(balance: $%d): ", balance)
		fmt.Scan(&bet)
		if bet > balance{
			fmt.Println("Bet can't be greater than balance")
		}else{
			break
		}
	}
	return bet
}

func main() {
	balance := uint(200)
	getName()
	for balance > 0{
		bet := getBet(balance)
		if bet == 0{
			break
		}
		fmt.Printf("Your Bet is: $%d.\n", bet)
		balance = balance - bet
	}
	fmt.Printf("Your remaining balance: $%d", balance)
}