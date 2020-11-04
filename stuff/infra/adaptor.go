package infra

import (
	"bytes"
	"github/four-servings/dropit-backend/config"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type (
	// S3Adaptor aws s3 adaptor interface
	S3Adaptor interface {
		Upload(fileHeader *multipart.FileHeader) string
	}

	s3Adaptor struct {
		bucket  string
		session *session.Session
	}
)

// NewS3Adaptor create s3adaptor instance
func NewS3Adaptor() S3Adaptor {
	s3Config := config.S3{}

	accessKeyID := s3Config.AccessKeyID()
	secretAccessKey := s3Config.SecretAccessKey()
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Region:           aws.String(s3Config.Region()),
		Endpoint:         aws.String(s3Config.Endpoint()),
		S3ForcePathStyle: aws.Bool(true),
	}))

	return &s3Adaptor{
		bucket:  s3Config.Bucket(),
		session: sess,
	}
}

// Upload upload file to aws s3 bucket
func (s *s3Adaptor) Upload(fileHeader *multipart.FileHeader) string {
	if fileHeader == nil {
		return ""
	}

	buffer := make([]byte, fileHeader.Size)
	file, err := fileHeader.Open()
	if err != nil {
		return ""
	}

	file.Read(buffer)

	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(s.session)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(uuid.String()),
		Body:   bytes.NewReader(buffer),
	})
	if err != nil {
		return ""
	}

	return uuid.String()
}
