package filemanager

import (
	"fmt"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func UploadFileToS3(c *gin.Context) {
	// Get file
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    src, err := file.Open()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer src.Close()

	var bucketName = os.Getenv("AWS_S3_BUCKET")

	cfg, err := config.LoadDefaultConfig(c.Request.Context())
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(c.Request.Context(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)),
		Body:   src,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{"message": "File uploaded successfully", "location": result.Location})
	return
}