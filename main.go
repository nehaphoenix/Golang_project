//To list S3 buckets on AWS using golang APIs
package main

import (
	"fmt"
  //AWS packages needed for s3 
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3session *s3.S3
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})))
}

func listBuckets() (resp *s3.ListBucketsOutput) {
	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	return resp
}

func main() {
	fmt.Println(listBuckets())
}
