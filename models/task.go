package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    AssignedTo  string             `bson:"assigned_to" json:"assigned_to"`
    CreatedBy   string             `bson:"created_by" json:"created_by"`
    Status      string             `bson:"status" json:"status"`
    CreatedAt   int64              `bson:"created_at" json:"created_at"`
}