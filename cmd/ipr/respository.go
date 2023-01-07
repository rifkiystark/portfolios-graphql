package ipr

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

type IPRRepository interface {
	CreateIPR(ipr *IPR) error
	UpdateIPR(id string, ipr *IPR) error
	GetIPR(id string) (IPR, error)
	GetIPRs(search *string) ([]IPR, error)
	DeleteIPR(id string) error
}

type IPRRepositoryImpl struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) IPRRepository {
	return &IPRRepositoryImpl{db: db}
}

func (p *IPRRepositoryImpl) CreateIPR(ipr *IPR) error {
	projectiprCollection := p.db.Collection("ipr")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := projectiprCollection.InsertOne(ctx, &ipr)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	ipr.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (p *IPRRepositoryImpl) UpdateIPR(id string, ipr *IPR) error {
	projectiprCollection := p.db.Collection("ipr")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)
	updateIPR := bson.D{{"$set", ipr}}
	_, err := projectiprCollection.UpdateByID(ctx, objectId, updateIPR)

	if err != nil {
		log.Printf("error UpdateByID: %v", err)
		return err
	}

	ipr.ID = objectId

	return nil
}

func (p *IPRRepositoryImpl) GetIPR(id string) (IPR, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return IPR{}, err
	}

	iprCollection := p.db.Collection("ipr")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := iprCollection.FindOne(ctx, bson.M{"_id": ObjectID})

	ipr := IPR{}

	res.Decode(&ipr)

	return ipr, nil
}

func (p *IPRRepositoryImpl) GetIPRs(search *string) ([]IPR, error) {
	iprCollection := p.db.Collection("ipr")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{}
	if search != nil {
		filter = bson.D{{"$text", bson.D{{"$search", *search}}}}
	}

	cur, err := iprCollection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("error finding document: %v", err)
		return nil, err
	}

	var iprs []IPR
	err = cur.All(ctx, &iprs)
	if err != nil {
		logrus.Errorf("error decoding document: %+v", err)
		return nil, err
	}

	return iprs, nil
}

func (p *IPRRepositoryImpl) DeleteIPR(id string) error {
	iprCollection := p.db.Collection("ipr")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	res, err := iprCollection.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		fmt.Printf("error DeleteOne: %v", err)
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("no document found with id %s", id)
	}

	return nil
}
