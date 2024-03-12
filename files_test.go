package moonshot

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {

	token := os.Getenv("MOONSHOT_API_KEY")
	if token == "" {
		t.Skip("MOONSHOT_API_KEY not set")
	}

	client := New(token)

	fileDetail, err := client.UploadFile(context.Background(), "./tests/hello.txt")
	assert.NoError(t, err)
	assert.NotNil(t, fileDetail)
	assert.NotEmpty(t, fileDetail.ID)
	t.Logf("upload file result: %v", fileDetail)

	fileDetailList, err := client.ListFiles(context.Background())
	assert.NoError(t, err)
	assert.True(t, fileInList(fileDetail.ID, fileDetailList))
	t.Logf("list files result: %v", fileDetailList)

	assert.NoError(t, client.DeleteFile(context.Background(), fileDetail.ID))

	fileDetailList, err = client.ListFiles(context.Background())
	assert.NoError(t, err)
	assert.False(t, fileInList(fileDetail.ID, fileDetailList))
}

func fileInList(fid string, list []FileDetail) bool {
	for _, f := range list {
		if f.ID == fid {
			return true
		}
	}
	return false
}
