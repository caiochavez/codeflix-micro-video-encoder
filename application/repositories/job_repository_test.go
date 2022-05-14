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

func TestJobRepositoryDb_Insert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "123"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repositoryVideo := repositories.NewVideoRepository(db)
	repositoryVideo.Insert(video)

	job, err := domain.NewJob("output_path", "pending", video)
	require.Nil(t, err)

	repositoryJob := repositories.NewJobRepository(db)
	repositoryJob.Insert(job)

	j, err := repositoryJob.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, j.ID)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDb_Update(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "123"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repositoryVideo := repositories.NewVideoRepository(db)
	repositoryVideo.Insert(video)

	job, err := domain.NewJob("output_path", "pending", video)
	require.Nil(t, err)

	repositoryJob := repositories.NewJobRepository(db)
	repositoryJob.Insert(job)

	job.Status = "completed"
	repositoryJob.Update(job)

	j, err := repositoryJob.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, j.ID)
	require.Equal(t, j.Status, job.Status)
}
