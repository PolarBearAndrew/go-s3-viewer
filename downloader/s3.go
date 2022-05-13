package downloader

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Reader struct {
	region string
	bucket string
}

// NewS3Reader
// must setup env AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY
func NewS3Reader(region, bucket string) *S3Reader {

	return &S3Reader{
		region: region,
		bucket: bucket,
	}
}

func (r *S3Reader) GetObject(key string) (*io.ReadCloser, error) {

	conf := aws.NewConfig().
		WithRegion(r.region).
		WithCredentials(credentials.NewEnvCredentials())

	sess, err := session.NewSession(conf)

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess, &aws.Config{
		DisableRestProtocolURICleaning: aws.Bool(true),
	})

	fmt.Println("bucket >>>>>>>", r.bucket)
	fmt.Println("key >>>>>>>", key)

	s3Obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	return &s3Obj.Body, nil
}
