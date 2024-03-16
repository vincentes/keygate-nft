package model

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)


type Presigner struct {
	PresignClient *s3.PresignClient
}

func (presigner Presigner) PutObject(
	bucketName string, objectKey string, lifetimeSecs int64) (*v4.PresignedHTTPRequest, error) {
	request, err := presigner.PresignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return request, err
}

type S3Actions struct {
	S3Client *s3.Client
}

func (basics S3Actions) BucketExists(bucketName string) (bool, error) {
	_, err := basics.S3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				log.Printf("Bucket %v is available.\n", bucketName)
				exists = false
				err = nil
			default:
				log.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", bucketName, err)
				return false, err
			}
		}
	} else {
		log.Printf("Bucket %v exists and you already own it.", bucketName)
	}

	return exists, err
}

type ImageURLRequest struct {
}

type PresignedURLData struct {
	PresignedURL string `json:"presigned_url"`
	URL string `json:"url"`
}

func GetSignedURL(tx *sql.Tx, request *ImageURLRequest) (*PresignedURLData, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("Failed to load configuration, %v", err)
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg)
	bucketBasics := S3Actions{S3Client: s3Client}
	presignClient := s3.NewPresignClient(s3Client)
	presigner := Presigner{PresignClient: presignClient}

	bucketName := os.Getenv("NFT_IMAGE_BUCKET")
	bucketExists, err := bucketBasics.BucketExists(bucketName)
	if err != nil {
		log.Printf("Couldn't check if bucket exists: %v", err)
		return nil, err
	}

	if !bucketExists {
		log.Printf("Bucket %v doesn't exist", bucketName)
		bucketDoesNotExist := errors.New("Bucket " + bucketName + " doesn't exist")
		return nil, bucketDoesNotExist
	}

	uuid := uuid.New().String()

	presignedPutRequest, err := presigner.PutObject(bucketName, uuid, 60)
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			bucketName, uuid, err)
		return nil, err
	}

	return &PresignedURLData{
		PresignedURL: presignedPutRequest.URL,
		URL: "https://" + bucketName + ".s3.amazonaws.com/" + uuid,
	}, nil
}

