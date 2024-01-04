package awsS3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// create an AWS session which can be
// reused if we're uploading many files

const s3Id = "AKIAI2M3S4PZOW7H5ULQ"
const s3Secret = "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"
const s3Region = "ap-southeast-1"

var s3Connection *session.Session

func Init() {
	var err error
	s3Connection, err = session.NewSession(&aws.Config{
		Region: aws.String(s3Region),
		Credentials: credentials.NewStaticCredentials(
			s3Id,     // id
			s3Secret, // secret
			""),      // token can be left blank for now
	})
	if err != nil {
		fmt.Println("s3 session error :", err)
	}
}
