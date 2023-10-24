package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CloudName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't log env file")
	}
	cloudinaryName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	return cloudinaryName
}

func CloudApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't log env file")
	}
	cloudinaryApiKey := os.Getenv("CLOUDINARY_API_KEY")
	return cloudinaryApiKey
}

func CloudApiSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	cloudinaryApiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	return cloudinaryApiSecret
}

func CloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}
	cloudinaryUploadFolder := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	return cloudinaryUploadFolder
}