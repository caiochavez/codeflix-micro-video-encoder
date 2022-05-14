package services_test

import (
	"encoder/application/repositories"
	"encoder/application/services"
	"encoder/domain"
	"encoder/framework/database"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "hello_world.mp4"
	video.CreatedAt = time.Now()

	repository := repositories.NewVideoRepository(db)

	return video, repository
}

func TestVideoService_Download(t *testing.T) {
	video, repository := prepare()

	service := services.NewVideoService()
	service.Video = video
	service.VideoRepository = repository

	err := service.Download("codeflix-test")
	require.Nil(t, err)

	err = service.Fragment()
	require.Nil(t, err)
}
