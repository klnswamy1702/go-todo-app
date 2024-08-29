package main

import (
    "log"
    "net/http"

    "github.com/klnswamy1702/go-todo-app/config"
    "github.com/klnswamy1702/go-todo-app/controllers"
    "github.com/klnswamy1702/go-todo-app/routes"
    "github.com/klnswamy1702/go-todo-app/services"
)

func main() {
    // Connect to MongoDB
    config.ConnectDB()

    // Initialize the Todo service and controller
    todoService := services.NewTodoService(config.DB.Collection("todos"))
    todoController := controllers.NewTodoController(todoService)

    // Setup routes
    router := routes.SetupRoutes(todoController)

    // Start the server
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
