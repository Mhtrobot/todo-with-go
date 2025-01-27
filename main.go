package main

import "todo-app/config"

func main() {
	db := config.InitDB()

	if db != nil {
		println("Database connected")
	}
}