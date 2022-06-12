package repositories

import (
	"certificate_api/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	reqCollection = "certificate_requests"
	certCollection = "certificates"
)

type CertificatesRepo struct {
	mongoClient *mongo.Client
	database    string
}

func NewRequestsRepo(mongoClient *mongo.Client, database string) *CertificatesRepo {
	return &CertificatesRepo{
		mongoClient: mongoClient,
		database:    database,
	}
}

func (certRepo *CertificatesRepo) StoreRequest(req models.CertificateRequestModel) error {
	certificatesDatabase := certRepo.mongoClient.Database(certRepo.database)
	requestsCollection := certificatesDatabase.Collection(reqCollection)
	_, err2 := requestsCollection.InsertOne(context.TODO(), req)
	if err2 != nil {
		fmt.Println("error storing request")
		if err2 == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err2)
	}
	return nil
}

func (certRepo *CertificatesRepo) StoreCertificate(cert *models.CertificateModel) error {
	certificatesDatabase := certRepo.mongoClient.Database(certRepo.database)
	certificatesCollection := certificatesDatabase.Collection(certCollection)
	_, err2 := certificatesCollection.InsertOne(context.TODO(), cert)
	if err2 != nil {
		fmt.Println("error storing certificate")
		if err2 == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err2)
	}
	return nil
}