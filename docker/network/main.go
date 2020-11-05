package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := GetClient()
		err := c.Ping(context.Background(), readpref.Primary())
		if err != nil {
			log.Fatal("Couldn't connect to the database", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			log.Println("Connected!")
			w.Write([]byte("Connected"))
		}
	})

	fmt.Println("Server running at http://localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// GetClient is...
func GetClient() *mongo.Client {
	var mongoserver string = fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	// var mongoserver string = fmt.Sprintf("mongodb://%s:%s", "localhost", "27017")

	clientOptions := options.Client().ApplyURI(mongoserver)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
