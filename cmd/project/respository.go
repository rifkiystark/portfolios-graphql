package project

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository interface {
	CreateProject(project *Project) error
	UpdateProject(id string, project *Project) error
	GetProject(id string) (Project, error)
	GetProjects(search *string) ([]Project, error)
	DeleteProject(id string) error
}

type ProjectRepositoryImpl struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) ProjectRepository {
	return &ProjectRepositoryImpl{db: db}
}

func (p *ProjectRepositoryImpl) CreateProject(project *Project) error {
	projectCollections := p.db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := projectCollections.InsertOne(ctx, &project)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	project.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (p *ProjectRepositoryImpl) UpdateProject(id string, project *Project) error {
	projectCollections := p.db.Collection("projects")
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

func (p *ProjectRepositoryImpl) GetProject(id string) (Project, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return Project{}, err
	}

	movieColl := p.db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := movieColl.FindOne(ctx, bson.M{"_id": ObjectID})

	project := Project{}

	res.Decode(&project)

	return project, nil
}

func (p *ProjectRepositoryImpl) GetProjects(search *string) ([]Project, error) {
	movieColl := p.db.Collection("projects")
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

func (p *ProjectRepositoryImpl) DeleteProject(id string) error {
	movieColl := p.db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	res, err := movieColl.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		fmt.Printf("error DeleteOne: %v", err)
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("no document found with id %s", id)
	}

	return nil
}
