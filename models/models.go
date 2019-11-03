package models

import (
	"context"
	"fmt"
	"log"

	"github.com/weeee9/go-todo/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TodoList represent todo task
type TodoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

// GetAllTasks will get all the todo tasks
func GetAllTasks() ([]*TodoList, error) {
	var results []*TodoList

	cur, err := database.TodoCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println("[MongoDB] Error:", err.Error())
		return nil, err
	}
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem TodoList
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Println("[MongoDB] Error:", err.Error())
		return nil, err
	}
	// Close the cursor once finished
	cur.Close(context.TODO())
	return results, nil
}

// InsertOneTask will insert a new tasl to database
func InsertOneTask(task TodoList) error {
	insertResult, err := database.TodoCollection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Println("[MongoDB] Error:", err.Error())
		return err
	}
	log.Println("[MongoDB] Insert a new task", insertResult)
	return nil
}

// CompleteTask will set task status to true
func CompleteTask(task string) error {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"status", true},
		}},
	}
	updateResult, err := database.TodoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("[MongoDB] Error:", err.Error())
		return err
	}
	log.Println("[MongoDB] Update:", updateResult)
	return nil
}

// UndoTask will set task status to false
func UndoTask(task string) error {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"status", false},
		}},
	}
	updateResult, err := database.TodoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("[MongoDB] Error:", err.Error())
		return err
	}
	log.Println("[MongoDB] Update:", updateResult)
	return nil
}

// DeleteOneTask will delete the task will provided task
func DeleteOneTask(tsakname string) error {
	fmt.Println(tsakname)
	id, _ := primitive.ObjectIDFromHex(tsakname)
	filter := bson.M{"_id": id}
	deleteResult, err := database.TodoCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("[MongoDB] Error:", err.Error())
		return err
	}
	log.Println("[MongoDB] Delete:", deleteResult)
	return nil
}

// DeleteAllTasks will all the tasks
func DeleteAllTasks() error {
	deleteResult, err := database.TodoCollection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println("[MongoDB] Error", err.Error())
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return nil
}
