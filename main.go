package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	accountBalance, _ := readFromFile()
	fmt.Println("Welcome to the Bank Application!")
	for {
		fmt.Println("\nSelect the Service :- \n1.Check Balance\n2.Deposit Money\n3.Withdraw Money\n4.Exit")
		var userChoice int
		fmt.Print("Your Choice:")
		fmt.Scan(&userChoice)
		switch userChoice {
		case 1:
			fmt.Printf("Account Balance is $%.2f\n", accountBalance)
		case 2:
			accountBalance = takeDeposit(accountBalance)
			writeToFile(accountBalance)
			fmt.Printf("Updated Account Balance is $%.2f\n", accountBalance)
		case 3:
			accountBalance = processWithdrawal(accountBalance)
			writeToFile(accountBalance)
			fmt.Printf("Updated Account Balance is $%.2f\n", accountBalance)
		case 4:
			fmt.Println("Good Byee")
			return
		default:
			fmt.Println("Invalid Input")
		}
	}
}

func readFromFile() (float64, error) {
	_, err := os.Stat("balance.txt")
	if err != nil {
		os.WriteFile("balance.txt", []byte("0.0"), 0644)
	}
	value, _ := os.ReadFile("balance.txt")
	balanceText := string(value)
	return strconv.ParseFloat(balanceText, 64)
}

func writeToFile(accountBalance float64) {
	balanceText := fmt.Sprintf("%.2f", accountBalance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func takeDeposit(accountBalance float64) float64 {
	var deposit float64
	fmt.Print("Type the amount you want to Deposit: ")
	fmt.Scan(&deposit)
	if deposit <= 0 {
		fmt.Printf("\nInvalid Amount ! Deposit amount should be more than 0\n")
		return accountBalance
	}
	return (accountBalance + deposit)
}

func processWithdrawal(accountBalance float64) float64 {
	var withdrawal float64
	fmt.Print("Type the amount you want to Withdraw: ")
	fmt.Scan(&withdrawal)
	if withdrawal <= 0 || withdrawal > accountBalance {
		fmt.Printf("\nInvalid Amount ! Withdrawal amount should be more than 0 and should be less than Account Balance\n")
		return accountBalance
	}
	return (accountBalance - withdrawal)
}
