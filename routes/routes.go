package routes

import (
    "github.com/gorilla/mux"
    "github.com/klnswamy1702/go-todo-app/controllers"
)

// SetupRoutes sets up the application routes.
func SetupRoutes(todoController *controllers.TodoController) *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/todos", todoController.GetTodos).Methods("GET")
    router.HandleFunc("/todos/{id}", todoController.GetTodoByID).Methods("GET")
    router.HandleFunc("/todos", todoController.CreateTodo).Methods("POST")
    router.HandleFunc("/todos/{id}", todoController.UpdateTodo).Methods("PUT")
    router.HandleFunc("/todos/{id}", todoController.DeleteTodo).Methods("DELETE")

    return router
}
