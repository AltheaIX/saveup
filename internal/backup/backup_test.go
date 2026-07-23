package backup

import (
	"testing"
)

func TestLatestArchives(t *testing.T) {
	archiveFile, err := LatestArchive("D:\\Palworld\\Workspace")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(archiveFile)
}
