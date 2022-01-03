package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func InitDB() *dynamodb.DynamoDB {
	region := os.Getenv("REGION")
	endpoint := os.Getenv("ENDPOINT")

	//Credentials for local DynamoDB connection
	crd := credentials.NewStaticCredentials("123", "123", "")

	//Create a new session
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: crd,
		Region:      aws.String(region),
		Endpoint:    aws.String(endpoint),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)
	return db
}
