package services

import (
    "context"
    "github.com/klnswamy1702/go-todo-app/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "time"
)

// TodoService provides operations on the Todo model.
type TodoService struct {
    Collection *mongo.Collection
}

// NewTodoService creates a new TodoService with the specified MongoDB collection.
func NewTodoService(collection *mongo.Collection) *TodoService {
    return &TodoService{Collection: collection}
}

// CreateTodo adds a new todo item to the collection.
func (s *TodoService) CreateTodo(todo models.Todo) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return s.Collection.InsertOne(ctx, todo)
}

// GetTodos retrieves all todo items from the collection.
func (s *TodoService) GetTodos() ([]models.Todo, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := s.Collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var todos []models.Todo
    for cursor.Next(ctx) {
        var todo models.Todo
        if err := cursor.Decode(&todo); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

// GetTodoByID retrieves a single todo item by its ID.
func (s *TodoService) GetTodoByID(id primitive.ObjectID) (*models.Todo, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var todo models.Todo
    err := s.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo)
    if err != nil {
        return nil, err
    }
    return &todo, nil
}

// UpdateTodo updates an existing todo item in the collection.
func (s *TodoService) UpdateTodo(id primitive.ObjectID, updatedTodo models.Todo) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"_id": id}
    update := bson.M{"$set": updatedTodo}

    _, err := s.Collection.UpdateOne(ctx, filter, update)
    return err
}

// DeleteTodo removes a todo item from the collection.
func (s *TodoService) DeleteTodo(id primitive.ObjectID) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"_id": id}
    _, err := s.Collection.DeleteOne(ctx, filter)
    return err
}
