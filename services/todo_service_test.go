package services

import (
    "testing"

    "github.com/klnswamy1702/go-todo-app/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// TestCreateTodo tests the CreateTodo function of the TodoService.
func TestCreateTodo(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    defer mt.Close()

    mt.Run("successfully inserts a todo", func(mt *mtest.T) {
        mt.AddMockResponses(mtest.CreateSuccessResponse())
        service := NewTodoService(mt.Coll)

        todo := models.Todo{
            Title:       "Sample Title",
            Description: "Sample Description",
            Completed:   false,
        }

        _, err := service.CreateTodo(todo)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }
    })
}

// TestGetTodos tests the GetTodos function of the TodoService.
func TestGetTodos(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    defer mt.Close()

    mt.Run("successfully retrieves todos", func(mt *mtest.T) {
        firstTodo := models.Todo{Title: "First Todo", Description: "First Description", Completed: false}
        secondTodo := models.Todo{Title: "Second Todo", Description: "Second Description", Completed: true}

        // Create mock responses
        mt.AddMockResponses(mtest.CreateCursorResponse(1, "todoapp.todos", mtest.FirstBatch, bson.D{
            {"_id", primitive.NewObjectID()},
            {"title", firstTodo.Title},
            {"description", firstTodo.Description},
            {"completed", firstTodo.Completed},
        }, bson.D{
            {"_id", primitive.NewObjectID()},
            {"title", secondTodo.Title},
            {"description", secondTodo.Description},
            {"completed", secondTodo.Completed},
        }))

        service := NewTodoService(mt.Coll)

        todos, err := service.GetTodos()
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        if len(todos) != 2 {
            t.Errorf("expected 2 todos, got %d", len(todos))
        }
    })
}

// TestGetTodoByID tests the GetTodoByID function of the TodoService.
func TestGetTodoByID(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    defer mt.Close()

    mt.Run("successfully retrieves a todo by ID", func(mt *mtest.T) {
        todoID := primitive.NewObjectID()
        todo := models.Todo{ID: todoID, Title: "Sample Todo", Description: "Sample Description", Completed: false}

        mt.AddMockResponses(mtest.CreateCursorResponse(1, "todoapp.todos", mtest.FirstBatch, bson.D{
            {"_id", todoID},
            {"title", todo.Title},
            {"description", todo.Description},
            {"completed", todo.Completed},
        }))

        service := NewTodoService(mt.Coll)

        result, err := service.GetTodoByID(todoID)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }

        if result.ID != todoID {
            t.Errorf("expected ID %v, got %v", todoID, result.ID)
        }
    })
}

// TestUpdateTodo tests the UpdateTodo function of the TodoService.
func TestUpdateTodo(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    defer mt.Close()

    mt.Run("successfully updates a todo", func(mt *mtest.T) {
        mt.AddMockResponses(mtest.CreateSuccessResponse())
        service := NewTodoService(mt.Coll)

        todoID := primitive.NewObjectID()
        updatedTodo := models.Todo{
            Title:       "Updated Title",
            Description: "Updated Description",
            Completed:   true,
        }

        err := service.UpdateTodo(todoID, updatedTodo)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }
    })
}

// TestDeleteTodo tests the DeleteTodo function of the TodoService.
func TestDeleteTodo(t *testing.T) {
    mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
    defer mt.Close()

    mt.Run("successfully deletes a todo", func(mt *mtest.T) {
        mt.AddMockResponses(mtest.CreateSuccessResponse())
        service := NewTodoService(mt.Coll)

        todoID := primitive.NewObjectID()

        err := service.DeleteTodo(todoID)
        if err != nil {
            t.Errorf("expected no error, got %v", err)
        }
    })
}
