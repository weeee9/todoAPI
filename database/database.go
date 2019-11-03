package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Credential for mongoDB atlas credential
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

// TodoCollection ...
var TodoCollection *mongo.Collection

// Init will connect to mongoDB atlas will given atlas credential file path
// see your_mongoDB_atlas_info.json
func Init(credFile string) {
	var c Credential
	// open credentail file
	file, err := ioutil.ReadFile(credFile)
	if err != nil {
		log.Fatalf("File error: %v\n", err)
	}
	json.Unmarshal(file, &c)

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", c.Username, c.Password, c.Host)
	// set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// connect to MongoDB
	// client, err := mongo.NewClient(clinetOpt)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	TodoCollection = client.Database("todos").Collection("task")
}
