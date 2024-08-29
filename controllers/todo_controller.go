package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/klnswamy1702/go-todo-app/models"
    "github.com/klnswamy1702/go-todo-app/services"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// TodoController handles incoming HTTP requests related to todos.
type TodoController struct {
    Service *services.TodoService
}

// NewTodoController creates a new TodoController with the specified service.
func NewTodoController(service *services.TodoService) *TodoController {
    return &TodoController{Service: service}
}

// CreateTodo handles the creation of a new todo item.
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)
    
    result, err := c.Service.CreateTodo(todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}

// GetTodos handles fetching all todo items.
func (c *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := c.Service.GetTodos()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(todos)
}

// GetTodoByID handles fetching a single todo item by its ID.
func (c *TodoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    todo, err := c.Service.GetTodoByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(todo)
}

// UpdateTodo handles updating an existing todo item.
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
    var updatedTodo models.Todo
    json.NewDecoder(r.Body).Decode(&updatedTodo)

    idStr := mux.Vars(r)["id"]
    id, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    err = c.Service.UpdateTodo(id, updatedTodo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// DeleteTodo handles deleting a todo item.
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteTodo(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
