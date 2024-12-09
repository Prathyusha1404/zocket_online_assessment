package services

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nfnt/resize"
)

// ProcessImage - Main function to process the image
func ProcessImage(message []byte) {
	log.Printf("Processing image with message: %s", message)

	// Extract the product ID and image URLs from the message
	// For simplicity, assume message is formatted as: "Product ID: 1, Image URLs: [\"http://example.com/image1.jpg\"]"
	msgStr := string(message)
	parts := strings.Split(msgStr, ", ")
	if len(parts) < 2 {
		log.Printf("Invalid message format: %s", msgStr)
		return
	}

	// Extract image URLs from message
	imageURLs := strings.Split(strings.Trim(parts[1], "[]\""), ", ")
	if len(imageURLs) == 0 {
		log.Printf("No image URLs found in message: %s", msgStr)
		return
	}

	// Process each image
	for _, imageURL := range imageURLs {
		err := downloadAndProcessImage(imageURL)
		if err != nil {
			log.Printf("Error processing image %s: %v", imageURL, err)
		}
	}
}

// downloadAndProcessImage - Downloads the image, compresses, and uploads it
func downloadAndProcessImage(imageURL string) error {
	// Step 1: Download the image from the URL
	img, err := downloadImage(imageURL)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}

	// Step 2: Compress the image
	compressedImg := compressImage(img)

	// Step 3: Upload the compressed image to S3 (or another storage solution)
	// You can replace this with your own storage logic (e.g., S3, local storage)
	err = uploadToS3(compressedImg)
	if err != nil {
		return fmt.Errorf("failed to upload image: %v", err)
	}

	log.Printf("Successfully processed and uploaded image from URL: %s", imageURL)
	return nil
}

// downloadImage - Downloads an image from the URL
func downloadImage(url string) (image.Image, error) {
	// Send HTTP GET request to download the image
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get image from URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Decode the image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	return img, nil
}

// compressImage - Compresses the image (resize and reduce quality)
func compressImage(img image.Image) []byte {
	// Resize the image to a smaller size (50% of the original size)
	resizedImg := resize.Resize(800, 0, img, resize.Lanczos3)

	// Create a buffer to hold the compressed image
	var buf bytes.Buffer
	// Compress image into the buffer with lower quality (e.g., 75)
	err := jpeg.Encode(&buf, resizedImg, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Printf("Error compressing image: %v", err)
		return nil
	}

	return buf.Bytes()
}

// uploadToS3 - Uploads the compressed image to AWS S3 (replace with your storage solution)
func uploadToS3(imageData []byte) error {
	// Initialize a session that will be used to create the S3 service
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS session: %v", err)
	}

	svc := s3.New(sess)

	// Define the S3 bucket and key (path)
	bucket := "your-s3-bucket-name"
	key := fmt.Sprintf("product_images/%s.jpg", generateFileName())

	// Create an S3 PutObject request
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(imageData),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return fmt.Errorf("failed to upload image to S3: %v", err)
	}

	log.Printf("Image successfully uploaded to S3: %s/%s", bucket, key)
	return nil
}

// generateFileName - Generates a unique filename based on the current timestamp
func generateFileName() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
