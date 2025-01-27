package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-app/config"
	"todo-app/models"

	"golang.org/x/crypto/bcrypt"
)

var reader = bufio.NewReader(os.Stdin)
var db = config.InitDB()
var CurrentUser *models.User

func readInput(prompt string) string {	
	fmt.Print(prompt)

	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func RegisterUser() {
	username := readInput("Enter username: ")
	password := readInput("Enter password: ")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		Username: username,
		HashedPassword: string(hashedPassword),
	}

	result := db.Create(&user)

	if result.Error != nil {
		fmt.Println("Failed to register user")
		return
	}

	fmt.Println("User registered successfully")
}

func LoginUser() bool {
	username := readInput("Enter username: ")
	password := readInput("Enter password: ")

	var user models.User

	isUser := db.Where("username = ?", username).First(&user)
	if isUser.Error != nil {
		fmt.Println("User not found")
		return false
	}

	isPassCorrect := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if isPassCorrect != nil {	
		fmt.Println("Invalid password")
		return false
	}

	CurrentUser = &user

	fmt.Println("Login successful")
	return true	
}