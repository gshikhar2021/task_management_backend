package controllers

import (
	"backend/config"
	"backend/models"
	"backend/websockets"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTask - Creates a new task and assigns it to a specific user
func CreateTask(c *gin.Context) {
	var task models.Task

	
	if err := c.ShouldBindJSON(&task); err != nil {
		fmt.Println("JSON Binding Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received Task: %+v\n", task) // Debugging log

	// Extract username from middleware
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Ensure username is properly set
	task.CreatedBy, _ = username.(string)

	// Set default values
	task.Status = "Pending"
	task.CreatedAt = time.Now().Unix()

	// Check if AssignedTo is empty and set it to the current user if it is
	if task.AssignedTo == "" {
		task.AssignedTo = task.CreatedBy
		fmt.Printf("Setting AssignedTo to current user: %s\n", task.AssignedTo) // Debugging log
	}

	// Insert into MongoDB
	collection := config.GetCollection("tasks")

	// Debugging: Print task before insertion
	fmt.Printf("Task to be inserted: %+v\n", task)

	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		fmt.Printf("MongoDB insertion error: %v\n", err) // Debugging log
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Debugging: Check if task was inserted successfully
	fmt.Println("Inserted Task ID:", result.InsertedID)

	// Set the ID in the task struct
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		task.ID = oid
	}

	// Notify assigned user via WebSocket
	websockets.NotifyUser(task.AssignedTo, task.Title)

	c.JSON(http.StatusOK, gin.H{"message": "Task created", "id": result.InsertedID, "task": task})
}

// GetTasks - Fetch tasks assigned to the logged-in user
func GetTasks(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Ensure username is a string
	user, _ := username.(string)

	collection := config.GetCollection("tasks")
	cursor, err := collection.Find(context.TODO(), bson.M{"assigned_to": user})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// MarkTaskAsDone - Updates a task's status to "Done"
func MarkTaskAsDone(c *gin.Context) {
	taskID := c.Param("id")

	// Validate ObjectID format
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID format"})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user, _ := username.(string)

	collection := config.GetCollection("tasks")

	// Check if the user is authorized to mark this task as done
	var task models.Task
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID, "assigned_to": user}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your assigned tasks"})
		return
	}

	result, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"status": "Done"}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or already marked as done"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task marked as done"})
}
