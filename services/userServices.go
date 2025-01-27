package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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
		fmt.Println("Failed to register user❌")
		time.Sleep(3 * time.Second)
		return
	}

	fmt.Println("User registered successfully✅")
	time.Sleep(3 * time.Second)
}

func LoginUser() bool {
	username := readInput("Enter username: ")
	password := readInput("Enter password: ")

	var user models.User

	isUser := db.Where("username = ?", username).First(&user)
	if isUser.Error != nil {
		fmt.Println("❌User not found❌")
		//pause for 3 seconds
		time.Sleep(3 * time.Second)
		return false
	}

	isPassCorrect := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if isPassCorrect != nil {	
		fmt.Println("❌Invalid password❌")
		time.Sleep(3 * time.Second)
		return false
	}

	CurrentUser = &user

	fmt.Println("✅Login successful✅")
	time.Sleep(3 * time.Second)
	return true	
}