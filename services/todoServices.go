package services

import (
	"fmt"
	"time"
	"todo-app/models"
)

func AddTodo() {
	title := readInput("Enter title: ")
	description := readInput("Enter description:\n")

	todo := models.Todo{
		Title: title,
		Description: description,
		UserID: CurrentUser.ID,
	}

	result := db.Create(&todo)
	if result.Error != nil {
		fmt.Println("Failed to create todo❌")
		time.Sleep(3 * time.Second)
		return
	}

	fmt.Println("Todo created successfully!✅")
	time.Sleep(3 * time.Second)
}

func GetTodos() {
	fmt.Println("----------------Todos----------------")
	var todos []models.Todo

	db.Where("user_id = ?", CurrentUser.ID).Find(&todos)
	if len(todos) == 0 {
		fmt.Println("No todos found")
		return
	}


	for _, todo := range todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}

		fmt.Printf("[%d] %s %-25s - %s\n", todo.ID, status, todo.Title, todo.Description)
	}
}