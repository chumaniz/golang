package main

import (
	"fmt"
	"os"
)

type account struct {
	cheque  int
	savings int
	travel  int
}

type pins struct {
	User0p string
	User1p string
	User2p string
}

type usernames struct {
	User0 string
	User1 string
	User2 string
}

func main() {
	password := pins{"123", "456", "789"}
	username := usernames{"Amahle", "Chumani", "Mihle"}
	var Username string
	var Pin string
	var Option string
	fmt.Println("Welcome to Fairbanx, South Africa's innovative digital bank!")
	fmt.Println("Main Menu")
	fmt.Println("+--------+")
	fmt.Println("1. Log In")
	fmt.Println("2. Quit")
	fmt.Print("Option: ")
	fmt.Scan(&Option)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	if Option == "1" {
		fmt.Print("Username: ")
		fmt.Scan(&Username)

		if Username == username.User0 {
			fmt.Print("Password: ")
			fmt.Scan(&Pin)
			if Pin == password.User0p {
				User0p()
			} else {
				fmt.Println("Access Denied: Incorrect Password")
			}
		} else if Username == username.User1 {
			fmt.Print("Password: ")
			fmt.Scan(&Pin)
			if Pin == password.User1p {
				User1p()
			} else {
				fmt.Println("Access Denied: Incorrect Password")
			}
		} else if Username == username.User2 {
			fmt.Print("Password: ")
			fmt.Scan(&Pin)
			if Pin == password.User2p {
				User2p()
			} else {
				fmt.Println("Access Denied: Incorrect Password")
			}
		} else {
			fmt.Println("Access Denied: Incorrect Username")
		}
	} else if Option == "2" {
		os.Exit(0)
	}

}

func User0p() {
	password := pins{"123", "456", "789"}
	username := usernames{"Amahle", "Chumani", "Mihle"}
	account := account{30000, 100000, 250000}
	if password.User0p == "123" {
		username.User0 = "Amahle"
		fmt.Println("Welcome to Fairbanx,", username.User0)
		fmt.Println("")
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Println("")
		fmt.Println(username.User0, "'s Account:")
		fmt.Println("|_Cheque: ", account.cheque)
		fmt.Println("|_Savings: ", account.savings)
		fmt.Println("|_Travel: ", account.travel)
		fmt.Println("")
		fmt.Println(username.User0, " would like to update their travel account.")
		fmt.Println("Travel account: +R", 100000)
		fmt.Println("")
		fmt.Println(username.User0, "'s Account:")
		fmt.Println("|_Cheque: ", account.cheque)
		fmt.Println("|_Savings: ", account.savings)
		fmt.Println("|_Travel: Previously=>", account.travel, "Now =>", account.travel+100000)
		fmt.Println("")
		fmt.Println("-------------------------------------------------------------------------------------")
	}
}

func User1p() {
	password := pins{"123", "456", "789"}
	username := usernames{"Amahle", "Chumani", "Mihle"}
	account := account{25000, 60000, 20000}
	if password.User1p == "456" {
		username.User0 = "Chumani"
		fmt.Println("Welcome to Fairbanx,", username.User1)
		fmt.Println("")
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Println("")
		fmt.Println(username.User1, "'s Account:")
		fmt.Println("|_Cheque: ", account.cheque)
		fmt.Println("|_Savings: ", account.savings)
		fmt.Println("|_Travel: ", account.travel)
		fmt.Println("")
		fmt.Println("-------------------------------------------------------------------------------------")
	}
}

func User2p() {
	password := pins{"123", "456", "789"}
	username := usernames{"Amahle", "Chumani", "Mihle"}
	account := account{50000, 20000, 0}
	if password.User2p == "789" {
		username.User2 = "Mihle"
		fmt.Println("Welcome to Fairbanx,", username.User2)
		fmt.Println("")
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Println("")
		fmt.Println(username.User2, "'s Account:")
		fmt.Println("|_Cheque: ", account.cheque)
		fmt.Println("|_Savings: ", account.savings)
		fmt.Println("|_Travel: ", account.travel)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("-------------------------------------------------------------------------------------")
	}
}
