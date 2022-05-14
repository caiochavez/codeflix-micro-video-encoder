package services

import (
	"cloud.google.com/go/storage"
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (videoService *VideoService) Download(bucketName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bucket := client.Bucket(bucketName)
	object := bucket.Object(videoService.Video.FilePath)

	reader, err := object.NewReader(ctx)
	if err != nil {
		return err
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	file, err := os.Create(os.Getenv("localStoragePath") + "/" + videoService.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = file.Write(body)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func (videoService *VideoService) Fragment() error {
	err := os.Mkdir(os.Getenv("localStoragePath")+"/"+videoService.Video.ID, os.ModePerm)
	if err != nil {
		return err
	}

	source := os.Getenv("localStoragePath") + "/" + videoService.Video.ID + ".mp4"
	target := os.Getenv("localStoragePath") + "/" + videoService.Video.ID + ".frag"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	PrintOutput(output)
	return nil
}

func PrintOutput(output []byte) {
	if len(output) > 0 {
		log.Printf("=======> Output: %s", string(output))
	}
}
