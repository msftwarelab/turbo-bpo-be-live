package awsS3

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
)

func Uploader(uploadedFile graphql.Upload) (*string, error) {

	port := os.Getenv("PORT")
	if port == "" {
		port = config.AppConfig.GetString("port")
	}

	appUrl := config.AppConfig.GetString("AppDefaultUrl")
	if port == "" {
		port = config.AppConfig.GetString("AppDefaultUrl")
	}
	baseURL := fmt.Sprintf("%s:%s", appUrl, port)

	content, err := ioutil.ReadAll(uploadedFile.File)
	if err != nil {
		return nil, err
	}

	format := fmt.Sprintf("prof-doc-.%v", uploadedFile.Filename)
	tempFile, err := ioutil.TempFile(os.TempDir(), format)
	if err != nil {
		return nil, err
	}
	if _, err = tempFile.Write(content); err != nil {
		return nil, err
	}
	if err = tempFile.Close(); err != nil {
		return nil, err
	}
	cdt := fmt.Sprintf("%s", utils.Int64ToStr(millis.NowInMillis()))
	newPath := filepath.Join("..", "resources", cdt, uploadedFile.Filename)
	CreateDirIfNotExist(filepath.Join("..", "resources", cdt))
	fileUrl := fmt.Sprintf("%s/%s/%s/%s", baseURL, "static", cdt, uploadedFile.Filename)
	newFile, err := os.Create(newPath)
	if err != nil {
		return nil, err
	}
	if _, err := newFile.Write(content); err != nil {
		return nil, err
	}
	defer newFile.Close()

	return &fileUrl, nil
}

func S3Uploader(uploadedFile graphql.Upload) (*string, error) {
	content, err := ioutil.ReadAll(uploadedFile.File)
	if err != nil {
		return nil, err
	}
	cdt := fmt.Sprintf("%s", utils.Int64ToStr(millis.NowInMillis()))
	s3FilenameAndUrl := fmt.Sprintf("%s%s%s%s", "resources/", cdt, "/", uploadedFile.Filename)
	fileName, err := UploadFileToS3(content, s3FilenameAndUrl)
	if err != nil {
		fmt.Println("Could not upload file error :", err)
	}
	return &fileName, nil
}

func S3UploaderFromFile(fileDir string) (*string, error) {

	file, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	if err != nil {
		return nil, err
	}
	cdt := fmt.Sprintf("%s", utils.Int64ToStr(millis.NowInMillis()))
	s3FilenameAndUrl := fmt.Sprintf("resources/%s/%s", cdt, filepath.Base(file.Name()))
	fileName, err := UploadFileToS3(buffer, s3FilenameAndUrl)
	if err != nil {
		fmt.Println("Could not upload file error :", err)
	}
	return &fileName, nil
}

func S3UploaderFromFileExcel(fileDir string) (*string, error) {

	file, err := os.Open(fileDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	if err != nil {
		return nil, err
	}
	cdt := fmt.Sprintf("%s", utils.Int64ToStr(millis.NowInMillis()))
	s3FilenameAndUrl := fmt.Sprintf("resources/tmp/%s/%s", cdt, filepath.Base(file.Name()))
	fileName, err := UploadFileToS3(buffer, s3FilenameAndUrl)
	if err != nil {
		fmt.Println("Could not upload file error :", err)
	}
	return &fileName, nil
}
