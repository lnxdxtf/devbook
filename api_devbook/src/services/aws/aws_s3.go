package aws_devbook_s3

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWS_S3_SERVICE struct {
	aws_s3_client *s3.Client
	bucket        string
}

type UploadInput struct {
	File_name string
	File_type string
	File_body []byte
}

func NewS3Service() (*AWS_S3_SERVICE, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_DEFAULT_REGION")))
	if err != nil {
		return nil, err
	}
	aws_s3_client := s3.NewFromConfig(cfg)
	return &AWS_S3_SERVICE{aws_s3_client, os.Getenv("AWS_S3_BUCKET")}, nil
}

func (s3Service AWS_S3_SERVICE) Upload(data UploadInput) error {

	input := s3.PutObjectInput{
		Bucket:      &s3Service.bucket,
		Key:         &data.File_name,
		Body:        bytes.NewReader(data.File_body),
		ContentType: aws.String(data.File_type),
	}

	_, err := s3Service.aws_s3_client.PutObject(context.TODO(), &input)
	if err != nil {
		return err
	}
	return nil
}
