package test

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestUploadS3(t *testing.T) {

	region := "ap-southeast-1"

	svc, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("AKIA255X7MZ73JK7DD4S", "aTIjy30fS+dx6STViMG5L1/5IH8bnNsIqFZTuKi5", ""),
		Region:      &region,
	})

	assert.NoError(t, err)

	s3Client := s3.New(svc)

	upFile, err := os.Open("image001.png")

	assert.NoError(t, err)

	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("codestar-paul-homeword"),
		Key:    aws.String("image001-key"),
		// ACL:                  aws.String("private"),
		Body:               bytes.NewReader(fileBuffer),
		ContentLength:      aws.Int64(fileSize),
		ContentType:        aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition: aws.String("attachment"),
	})

}

func TestGetS3(t *testing.T) {

	region := "ap-southeast-1"

	svc, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("AKIA255X7MZ73JK7DD4S", "aTIjy30fS+dx6STViMG5L1/5IH8bnNsIqFZTuKi5", ""),
		Region:      &region,
	})

	assert.NoError(t, err)

	s3Client := s3.New(svc)

	upFile, err := os.Open("image001.png")

	assert.NoError(t, err)

	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	data, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("codestar-paul-homeword"),
		Key:    aws.String("image001-key"),
	})

	assert.NoError(t, err)

	fmt.Println(data)

}

func TestDeleteS3(t *testing.T) {

	region := "ap-southeast-1"

	svc, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("AKIA255X7MZ73JK7DD4S", "aTIjy30fS+dx6STViMG5L1/5IH8bnNsIqFZTuKi5", ""),
		Region:      &region,
	})

	assert.NoError(t, err)

	s3Client := s3.New(svc)

	upFile, err := os.Open("image001.png")

	assert.NoError(t, err)

	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	data, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String("codestar-paul-homeword"),
		Key:    aws.String("image001-key"),
	})

	assert.NoError(t, err)

	fmt.Println(data)

}
