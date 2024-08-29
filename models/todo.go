package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo represents a task with a title, description, and completion status.
type Todo struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    Completed   bool               `bson:"completed" json:"completed"`
}
