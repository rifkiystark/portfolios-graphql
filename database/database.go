package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
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

func (db *DB) CreateProject(project *Project) error {
	projectCollections := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	newProject := Project{Title: project.Title, ImageURL: project.ImageURL, AdditionalInfo: project.AdditionalInfo, Description: project.Description}
	result, err := projectCollections.InsertOne(ctx, newProject)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	project.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (db *DB) UpdateProject(id string, project *Project) error {
	projectCollections := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)
	// updateProject := bson.D{{"$set", Project{Title: project.Title, ImageURL: project.ImageURL, AdditionalInfo: project.AdditionalInfo, Description: project.Description}}}
	updateProject := bson.D{{"$set", project}}
	_, err := projectCollections.UpdateByID(ctx, objectId, updateProject)

	if err != nil {
		log.Printf("error UpdateByID: %v", err)
		return err
	}

	project.ID = objectId

	return nil
}

func (db *DB) DeleteProject(id string) (bool, error) {

	movieColl := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error: %v", err)
		return false, err
	}
	res, err := movieColl.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		fmt.Printf("error DeleteOne: %v", err)
		return false, err
	}

	return res.DeletedCount == 1, nil
}

func (db *DB) FindProjectById(id string) (Project, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return Project{}, err
	}

	movieColl := db.client.Database(db.databaseName).Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := movieColl.FindOne(ctx, bson.M{"_id": ObjectID})

	project := Project{}

	res.Decode(&project)

	return project, nil
}

func (db *DB) AllProjects(search *string) ([]Project, error) {
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

	var projects []Project
	err = cur.All(ctx, &projects)
	if err != nil {
		logrus.Errorf("error decoding document: %+v", err)
		return nil, err
	}

	return projects, nil
}
