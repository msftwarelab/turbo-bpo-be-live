package awsS3

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func UploadFileToS3(content []byte, filenameandURl string) (string, error) {
	// get the file size and read
	// the file content into a buffer
	//size := fileHeader.Size
	//buffer := make([]byte, size)

	// create a unique file name for the file
	//tempFileName := "pictures/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s3Connection).PutObject(&s3.PutObjectInput{
		Bucket: aws.String("turbo-bpo"),
		Key:    aws.String(filenameandURl),
		ACL:    aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:   bytes.NewReader(content),
		//ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(content)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://turbo-bpo.s3.ap-southeast-1.amazonaws.com/%s", filenameandURl)
	return url, err
}
