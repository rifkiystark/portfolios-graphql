package certificate

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

type CertificateRepository interface {
	CreateCertificate(Certificate *Certificate) error
	UpdateCertificate(id string, Certificate *Certificate) error
	GetCertificate(id string) (Certificate, error)
	GetCertificates(search *string) ([]Certificate, error)
	DeleteCertificate(id string) error
}

type CertificateRepositoryImpl struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) CertificateRepository {
	return &CertificateRepositoryImpl{db: db}
}

func (p *CertificateRepositoryImpl) CreateCertificate(certificate *Certificate) error {
	certificateCollection := p.db.Collection("certificate")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := certificateCollection.InsertOne(ctx, &certificate)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	certificate.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (p *CertificateRepositoryImpl) UpdateCertificate(id string, certificate *Certificate) error {
	certificateCollection := p.db.Collection("certificate")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)
	updateCertificate := bson.D{{"$set", certificate}}
	_, err := certificateCollection.UpdateByID(ctx, objectId, updateCertificate)

	if err != nil {
		log.Printf("error UpdateByID: %v", err)
		return err
	}

	certificate.ID = objectId

	return nil
}

func (p *CertificateRepositoryImpl) GetCertificate(id string) (Certificate, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return Certificate{}, err
	}

	certificateCollection := p.db.Collection("certificate")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := certificateCollection.FindOne(ctx, bson.M{"_id": ObjectID})

	certificate := Certificate{}

	res.Decode(&certificate)

	return certificate, nil
}

func (p *CertificateRepositoryImpl) GetCertificates(search *string) ([]Certificate, error) {
	certificateCollection := p.db.Collection("certificate")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{}
	if search != nil {
		filter = bson.D{{"$text", bson.D{{"$search", *search}}}}
	}

	cur, err := certificateCollection.Find(ctx, filter)
	if err != nil {
		fmt.Printf("error finding document: %v", err)
		return nil, err
	}

	var certificates []Certificate
	err = cur.All(ctx, &certificates)
	if err != nil {
		logrus.Errorf("error decoding document: %+v", err)
		return nil, err
	}

	return certificates, nil
}

func (p *CertificateRepositoryImpl) DeleteCertificate(id string) error {
	certificateCollection := p.db.Collection("certificate")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	res, err := certificateCollection.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		fmt.Printf("error DeleteOne: %v", err)
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("no document found with id %s", id)
	}

	return nil
}
