package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestVideoRepositoryDb_Insert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "123"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repository := repositories.NewVideoRepository(db)
	repository.Insert(video)

	v, err := repository.Find(video.ID)
	require.Nil(t, err)
	require.NotEmpty(t, v.ID)
	require.Equal(t, v.ID, video.ID)
}
