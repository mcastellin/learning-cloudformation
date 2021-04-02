package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func AskConfirmation(msg string) bool {
	fmt.Println(msg)
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	return strings.HasPrefix(
		strings.ToLower(response), "y")
}

func askToCreateBucket(client *s3.Client, bucketName string) (string, error) {
	msg := fmt.Sprintf("Bucket with name %s does not exists. Would you like to create (Y/n)?", bucketName)
	if !AskConfirmation(msg) {
		return "", errors.New("Cannot proceed without a bucket. Aborting.")
	}

	createBucketInput := &s3.CreateBucketInput{
		Bucket: &bucketName,
	}
	_, err := client.CreateBucket(context.TODO(), createBucketInput)
	if err != nil {
		log.Fatal(err)
	}

	return bucketName, nil
}

func getOrCreateBucket(cfg aws.Config, bucketName string) string {
	client := s3.NewFromConfig(cfg)

	input := &s3.ListBucketsInput{}
	buckets, err := client.ListBuckets(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range buckets.Buckets {
		if *bucket.Name == bucketName {
			// found valid bucket. returning
			return *bucket.Name
		}
	}

	_, err = askToCreateBucket(client, bucketName)
	if err != nil {
		log.Fatal(err)
	}
	return bucketName
}

func AddFileToS3(client *s3.Client, bucketName string, path string) (*s3.PutObjectOutput, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	return client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(path),
		Body:                 bytes.NewReader(buffer),
		ACL:                  types.ObjectCannedACLPrivate,
		ContentLength:        size,
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: types.ServerSideEncryptionAes256,
	})
}

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	bucketName := getOrCreateBucket(cfg, "mycloudformationbucket1234")

	client := s3.NewFromConfig(cfg)

	fileslist := []string{
		"vpc.yaml",
		"jenkins.yaml",
	}

	fmt.Println("Uploading files to s3...")
	for _, path := range fileslist {
		fmt.Println(path)
		AddFileToS3(client, bucketName, path)
	}
	fmt.Println("Completed.")
}
