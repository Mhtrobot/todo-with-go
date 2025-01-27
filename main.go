package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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

func colorCyan(text string) string {
    return "\033[36m" + text + "\033[0m"
}

func colorGreen(text string) string {
    return "\033[32m" + text + "\033[0m"
}

func colorRed(text string) string {
    return "\033[31m" + text + "\033[0m"
}

func colorYellow(text string) string {
    return "\033[33m" + text + "\033[0m"
}

func colorBlue(text string) string {
    return "\033[34m" + text + "\033[0m"
}

func printSeparator() {
    fmt.Println(colorBlue("════════════════════════════════════════"))
}

func printLoadingAnimation(duration time.Duration, message string) {
    frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
    end := time.Now().Add(duration)
    
    for time.Now().Before(end) {
        for _, frame := range frames {
            fmt.Printf("\r%s %s", frame, message)
            time.Sleep(100 * time.Millisecond)
        }
    }
    fmt.Println()
}

func printBanner() {
    fmt.Println(colorCyan(`
   ____          _            
  / ___|___   __| | ___  _ __ 
 | |   / _ \ / _  |/ _ \| '__|
 | |__| (_) | (_| | (_) | |   
  \____\___/ \__,_|\___/|_|   
                              
        `))
}

func welcome() {
	clearScreen()
	printBanner()
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
	printBanner()
	fmt.Println(colorGreen("--------- Main Menu ---------"))
	println("1. Register")
	println("2. Login")
	println("3. Exit")

	input := readInput(colorCyan("Choose an option -> "))

	switch input {
		case "1":
			services.RegisterUser()
			time.Sleep(time.Second)
		case "2":
			result := services.LoginUser()
			time.Sleep(time.Second)
			if result {
				userMenu()
				return
			}
		case "3":
			fmt.Println(colorGreen("Goodbye!"))
			os.Exit(0)
		default:
			fmt.Println(colorRed("Invalid choice!"))
			time.Sleep(time.Second)
	}
		
}

func userMenu() {
	clearScreen()
	printBanner()
	printSeparator()
    fmt.Printf("👤 Logged in as: %s\n", colorGreen(services.CurrentUser.Username))
	printSeparator()

    services.GetTodos()
	printSeparator()

	fmt.Printf("%s %s\n", colorYellow("➊"), "Create todo")
    fmt.Printf("%s %s\n", colorYellow("➋"), "Toggle todo")
    fmt.Printf("%s %s\n", colorYellow("➌"), "Delete todo")
    fmt.Printf("%s %s\n", colorYellow("➍"), "Logout")
    fmt.Printf("%s %s\n", colorYellow("➎"), "Exit")
    printSeparator()

    input := readInput(colorCyan("Choose an option ▶ "))
   
    switch input {
    case "1":
        services.AddTodo()
    case "2":
        services.ToggleTodo()
    case "3":
        services.DeleteTodo()
    case "4":
        services.CurrentUser = nil
    case "5":
        fmt.Println(colorGreen("Goodbye!"))
        os.Exit(0)
    default:
        fmt.Println(colorRed("Invalid choice!"))
        time.Sleep(time.Second)
    }
}

func main() {
	db := config.InitDB()

	if db == nil {
        fmt.Println(colorRed("Failed to connect to database"))
        return
    }

	welcome()
}