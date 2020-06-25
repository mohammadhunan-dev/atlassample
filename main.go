package main

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
func main(){

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(
    "mongodb+srv://mongomoe:<password>@cluster0-xbkka.mongodb.net/<dbname>?retryWrites=true&w=majority",
    ))
    if err != nil { log.Fatal(err) }
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))


	//     fmt.Println("Connected to MongoDB!")

    // // Rest of the code will go here
    // collection := client.Database("test").Collection("stuff")

    // rec := Trainer{"A", 10, "Town"}
    // insertResult, err := collection.InsertOne(context.TODO(), rec)
    // if err != nil {
    //     log.Fatal(err)
    // }else{
    //     fmt.Println("something else occurred")
    // }

    // fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}




// import "go.mongodb.org/mongo-driver/mongo"

// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, options.Client().ApplyURI(
//   "mongodb+srv://mongomoe:<password>@cluster0-xbkka.mongodb.net/<dbname>?retryWrites=true&w=majority",
// ))
// if err != nil { log.Fatal(err) }