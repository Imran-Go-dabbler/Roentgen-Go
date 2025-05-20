package main

import (
	"fmt"
	"os"
	"strconv"
)

const filename = "balance.txt"

func read(filename string) float64 {
	data, _ := os.ReadFile(filename)
	balancer, _ := strconv.ParseFloat(string(data), 64)
	return balancer
}

func write(balanceLatest float64) {
	balanceUpdate := fmt.Sprint(balanceLatest)
	os.WriteFile(filename, []byte(balanceUpdate), 0644)
}
func main() {
	fmt.Println("Welcome to Go bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var balance float64 = read(filename)
		var input int
		fmt.Print("Waiting for your choice: ")
		fmt.Scan(&input)
		fmt.Println("Your choice: ", input)

		switch input {
		case 1:
			fmt.Println("Balance: ", balance)
		case 2:
			fmt.Print("How much are you going to deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Print("Deposit cannot be 0")
				continue
			}
			balance += deposit
			fmt.Println("Your balance: ", balance)
			write(balance)
		case 3:
			fmt.Print("How much are you going to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			if withdraw >= balance {
				fmt.Print("Withdrawal cannot be more than your balance")
				continue
			}

			balance -= withdraw
			fmt.Println("Your balance: ", balance)
			write(balance)
		default:
			fmt.Println("You have exited application")
			fmt.Print("Thanks for choosing Go Bank")
			return
		}
	}

}
