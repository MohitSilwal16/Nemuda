package db

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDBCollection *mongo.Collection

func Init_Mongo() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	dbName := os.Getenv("mongoDBName")
	collectionName := os.Getenv("mongoCollectionName")

	if dbName == "" || collectionName == "" {
		return errors.New("DATABASE NAME & COLLECTION NAME NOT SPECIFIED IN .ENV FILE")
	}

	uri := "mongodb://localhost:27017/"

	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOption)

	if err != nil {
		return err
	}
	mongoDBCollection = client.Database(dbName).Collection(collectionName)

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Connection with Mongo DB isn't established")
		return err
	}

	log.Println("Connection with Mongo DB is established")
	return nil
}

func GetBlogsByTags(tag string) ([]models.Blog, error) {
	var filter primitive.M
	if tag == "All" {
		filter = bson.M{}
	} else {
		filter = bson.M{"Tags": bson.M{"$in": []string{tag}}}
	}

	cursor, err := mongoDBCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var blogs []models.Blog

	for cursor.Next(context.Background()) {
		var blog models.Blog

		if err = cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func AddBlog(blog models.Blog) error {
	result, err := mongoDBCollection.InsertOne(context.Background(), blog)

	if err != nil {
		return err
	}
	if result.InsertedID == "" {
		return errors.New("result.InsertedID is empty")
	}
	return nil
}
