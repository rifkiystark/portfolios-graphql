package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rifkiystark/portfolios-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client       *mongo.Client
	databaseName string
}

func Connect(dbHost string, dbName string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbHost))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client:       client,
		databaseName: dbName,
	}
}

func (db *DB) CreateProject(project model.Project) (*model.Project, error) {
	projectCollections := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newProject := Project{Title: project.Title, ImageURL: project.ImageURL, AdditionalInfo: project.AdditionalInfo, Description: project.Description}
	result, err := projectCollections.InsertOne(ctx, newProject)

	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	project.ID = insertedID

	return &project, nil
}

func (db *DB) UpdateProject(id string, project model.Project) (*model.Project, error) {
	projectCollections := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)
	updateProject := bson.D{{"$set", Project{Title: project.Title, ImageURL: project.ImageURL, AdditionalInfo: project.AdditionalInfo, Description: project.Description}}}
	_, err := projectCollections.UpdateByID(ctx, objectId, updateProject)

	if err != nil {
		log.Printf("error UpdateByID: %v", err)
		return nil, err
	}

	project.ID = id

	return &project, nil
}

func (db *DB) FindProjectById(id string) (*model.Project, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	movieColl := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := movieColl.FindOne(ctx, bson.M{"_id": ObjectID})

	movie := model.Project{ID: id}

	res.Decode(&movie)

	return &movie, nil
}

func (db *DB) AllProjects(search *string) ([]*model.Project, error) {
	movieColl := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{}
	if search != nil {
		filter = bson.D{{"$text", bson.D{{"$search", *search}}}}
	}

	cur, err := movieColl.Find(ctx, filter)
	if err != nil {
		fmt.Printf("error finding document: %v", err)
		return nil, err
	}

	var projects []*model.Project
	for cur.Next(ctx) {
		var result Project
		err := cur.Decode(&result)
		if err != nil {
			fmt.Printf("error decoding cursor: %v", err)
			return nil, err
		}

		project := model.Project{ID: result.ID.Hex(), Title: result.Title, ImageURL: result.ImageURL, AdditionalInfo: result.AdditionalInfo, Description: result.Description}

		projects = append(projects, &project)
	}

	return projects, nil
}
