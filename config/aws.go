package config

import (
	"os"
)

// S3 aws s3 config
type S3 struct {
	region string
	bucket string
}

// AccessKeyID aws s3 access key id
func (s *S3) AccessKeyID() string {
	if env := os.Getenv("AWS_S3_ACCESS_KEY_ID"); env != "" {
		return env
	}
	return "access-key-id"
}

// SecretAccessKey aws s3 secret access key
func (s *S3) SecretAccessKey() string {
	if env := os.Getenv("AWS_S3_SECRET_ACCESS_KEY"); env != "" {
		return env
	}
	return "secret-access-key"
}

// Endpoint aws s3 endpoint
func (s *S3) Endpoint() string {
	if env := os.Getenv("AWS_S3_ENDPOINT"); env != "" {
		return env
	}
	return "http://localhost:4566"
}

// Region aws s3 region
func (s *S3) Region() string {
	if env := os.Getenv("AWS_S3_REGION"); env != "" {
		return env
	}
	return "us-east-1"
}

// Bucket s3 bucket
func (s *S3) Bucket() string {
	if env := os.Getenv("AWS_S3_BUCKET"); env != "" {
		return env
	}
	return "bucket"
}
