package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-app/config"
	"todo-app/services"
)

var reader = bufio.NewReader(os.Stdin)

func readInput(prompt string) string {
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

func clearScreen() {
	print("\033[H\033[2J")
}

func welcome() {
	clearScreen()
	for {
		if services.CurrentUser == nil {
			startNotLoggedIn()
		} else {
			userMenu()
		}
	}
}

func startNotLoggedIn() {
	clearScreen()
	println("----------------Todo App----------------")
	println("1. Register")
	println("2. Login")
	println("3. Exit")

	input := readInput("Choose Option and Press enter to continue :-> ")

	switch input {
		case "1":
			services.RegisterUser()
			break
		case "2":
			result := services.LoginUser()
			if !result {
				welcome()
				return
			} else {
				userMenu()
			}
		case "3":
			println("Goodbye!")
			os.Exit(0)
	}
		
}

func userMenu() {
	clearScreen()
	println("----------------Todo App----------------")
	fmt.Printf("Logged in as: %s\n", services.CurrentUser.Username)
	fmt.Println("1. List todos")
	fmt.Println("2. Create todo")
	fmt.Println("3. Toggle todo")
	fmt.Println("4. Delete todo")
	fmt.Println("5. Logout")
	fmt.Println("6. Exit")

	input := readInput("Choose Option and Press enter to continue :-> ")

	switch input {
		case "6":
			println("Goodbye!")
			os.Exit(0)
	}
}

func main() {
	db := config.InitDB()

	if db == nil {
		println("Failed to connect to database")
	}

	welcome()
}