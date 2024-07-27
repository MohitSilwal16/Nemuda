package db

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/MohitSilwal16/Nemuda/server/pb"
)

var mongoDBCollection *mongo.Collection

func Init_Mongo() error {
	err := godotenv.Load("main.env")

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

func GetBlogByTitle(title string) (*pb.Blog, error) {
	filter := bson.M{"title": title}

	result := mongoDBCollection.FindOne(context.Background(), filter)

	var blog *pb.Blog
	err := result.Decode(&blog)

	if err != nil {
		return blog, err
	}

	return blog, nil
}

func SearchBlogByTitle(title string) (bool, error) {
	filter := bson.M{"title": title}

	result := mongoDBCollection.FindOne(context.Background(), filter)

	var blog pb.Blog
	err := result.Decode(&blog)

	if err == nil {
		return true, nil
	}

	if err.Error() == mongo.ErrNoDocuments.Error() {
		return false, nil
	}
	return false, err
}

func AddBlog(blog *pb.Blog) error {
	blogTitleFound, err := SearchBlogByTitle(blog.Title)
	if err != nil {
		return err
	}
	if blogTitleFound {
		return errors.New("TITLE IS ALREADY USED")
	}

	result, err := mongoDBCollection.InsertOne(context.Background(), blog)

	if err != nil {
		return err
	}
	if result.InsertedID == "" {
		return errors.New("result.InsertedID is empty")
	}
	return nil
}

func isBlogLikedByUser(title string, username string) (bool, error) {
	filter := bson.M{
		"title":         title,
		"likedUsername": bson.M{"$in": []string{username}},
	}

	var blog pb.Blog
	err := mongoDBCollection.FindOne(context.Background(), filter).Decode(&blog)

	if err == nil {
		return true, nil
	}
	if err.Error() == mongo.ErrNoDocuments.Error() {
		return false, nil
	}
	return false, err
}

func LikeBlog(title string, username string) error {
	doesBlogExists, err := SearchBlogByTitle(title)

	if err != nil {
		return err
	}

	if !doesBlogExists {
		return errors.New("BLOG NOT FOUND")
	}

	isBlogAlreadyLiked, err := isBlogLikedByUser(title, username)

	if err != nil {
		return err
	}

	if isBlogAlreadyLiked {
		return errors.New("BLOG ALREADY LIKED")
	}

	filter := bson.M{"title": title}
	update := bson.M{
		"$inc":      bson.M{"likes": 1},
		"$addToSet": bson.M{"likedUsername": username},
	}

	result, err := mongoDBCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("BLOG NOT FOUND")
	}

	return nil
}

func DislikeBlog(title string, username string) error {
	doesBlogExists, err := SearchBlogByTitle(title)

	if err != nil {
		return err
	}

	if !doesBlogExists {
		return errors.New("BLOG NOT FOUND")
	}

	isBlogAlreadyLiked, err := isBlogLikedByUser(title, username)

	if err != nil {
		return err
	}

	if !isBlogAlreadyLiked {
		return errors.New("BLOG ALREADY DISLIKED")
	}

	filter := bson.M{"title": title}
	update := bson.M{
		"$inc":  bson.M{"likes": -1},
		"$pull": bson.M{"likedUsername": username},
	}

	result, err := mongoDBCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("BLOG NOT FOUND")
	}

	return nil
}

func AddComment(title string, comment *pb.Comment) error {
	filter := bson.M{
		"title": title,
	}

	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}

	result, err := mongoDBCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("BLOG NOT FOUND")
	}

	return nil
}

func UpdateBlog(oldTitle string, username string, newTitle string, newDescription string, newImagePath string, newTag string) error {
	doesBlogExists, err := SearchBlogByTitle(oldTitle)

	if err != nil {
		return err
	}

	if !doesBlogExists {
		return errors.New("BLOG NOT FOUND")
	}

	filter := bson.M{
		"username": username,
		"title":    oldTitle,
	}

	updated := bson.M{
		"$set": bson.M{
			"title":       newTitle,
			"description": newDescription,
			"imagepath":   newImagePath,
			"tag":         newTag,
		},
	}

	result, err := mongoDBCollection.UpdateOne(context.Background(), filter, updated)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("USER CANNOT UPDATE THIS BLOG")
	}

	return nil
}

func DeleteBlog(title string, username string) error {
	doesBlogExists, err := SearchBlogByTitle(title)

	if err != nil {
		return err
	}

	if !doesBlogExists {
		return errors.New("BLOG NOT FOUND")
	}

	filter := bson.M{
		"username": username,
		"title":    title,
	}

	result, err := mongoDBCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("USER CANNOT DELETE THIS BLOG")
	}

	return nil
}

func GetBlogsByTagWithOffset(tag string, offset int, limit int) ([]*pb.Blog, error) {
	var filter primitive.M
	if tag == "All" {
		filter = bson.M{}
	} else {
		filter = bson.M{"tag": tag}
	}

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))

	cursor, err := mongoDBCollection.Find(context.Background(), filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var blogs []*pb.Blog

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var blog pb.Blog

		if err = cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}
	return blogs, nil
}
