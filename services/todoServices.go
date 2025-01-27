package services

import (
	"fmt"
	"time"
	"todo-app/models"
)

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

func AddTodo() {
    fmt.Println(colorBlue("\n📝 Create New Todo"))
    printSeparator()
    
    title := readInput("Title: ")
    description := readInput("Description: ")

    printLoadingAnimation(1*time.Second, "Creating todo...")

    todo := models.Todo{
        Title: title,
        Description: description,
        UserID: CurrentUser.ID,
    }

    result := db.Create(&todo)
    if result.Error != nil {
        fmt.Println(colorRed("\n❌ Failed to create todo"))
        time.Sleep(2 * time.Second)
        return
    }

    fmt.Println(colorGreen("\n✅ Todo created successfully!"))
    time.Sleep(2 * time.Second)
}

func GetTodos() {
    fmt.Printf("\n%s\n", colorBlue("📋 Your Todos"))
    var todos []models.Todo

    db.Where("user_id = ?", CurrentUser.ID).Find(&todos)
    if len(todos) == 0 {
        fmt.Printf("\n%s\n", colorYellow("📝 No todos found - Create your first todo!"))
        return
    }

    for _, todo := range todos {
        status := colorRed("✘")
        if todo.Completed {
            status = colorGreen("✔")
        }
        
        title := todo.Title
        if todo.Completed {
            title = colorGreen(title)
        }
        
        fmt.Printf("%s [%d] %-25s %s\n", 
            status,
            todo.ID,
            title,
            colorBlue("• "+todo.Description))
    }
    fmt.Println()
}

func ToggleTodo() {
	id := readInput("Enter todo id: ")
	var todo models.Todo

	result := db.Where("id = ? AND user_id = ?", id, CurrentUser.ID).First(&todo)
	if result.Error != nil {
		fmt.Println("Todo not found❌")
		time.Sleep(3 * time.Second)
		return
	}

	todo.Completed = !todo.Completed
	db.Save(&todo)

	fmt.Println("Todo updated successfully✅")
	time.Sleep(3 * time.Second)
}

func DeleteTodo() {
	id := readInput("Enter todo id: ")

	var todo models.Todo

	result := db.Where(("id = ? AND user_id = ?"), id, CurrentUser.ID).First(&todo)
	if result.Error != nil {
		fmt.Println("Todo not found❌")
		time.Sleep(3 * time.Second)
		return
	}

	db.Delete(&todo)

	fmt.Println("Todo deleted successfully✅")
	time.Sleep(3 * time.Second)
}