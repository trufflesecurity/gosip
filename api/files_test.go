package api

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestFiles(t *testing.T) {
	checkClient(t)

	web := NewSP(spClient).Web()
	newFolderName := uuid.New().String()
	rootFolderURI := getRelativeURL(spClient.AuthCnfg.GetSiteURL()) + "/Shared%20Documents"
	newFolderURI := rootFolderURI + "/" + newFolderName
	if _, err := web.GetFolder(rootFolderURI).Folders().Add(newFolderName); err != nil {
		t.Error(err)
	}

	t.Run("AddSeries", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			fileName := fmt.Sprintf("File_%d.txt", i)
			fileData := []byte(fmt.Sprintf("File %d data", i))
			if _, err := web.GetFolder(newFolderURI).Files().Add(fileName, fileData, true); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Get", func(t *testing.T) {
		if _, err := web.GetFolder(newFolderURI).Files().Get(); err != nil {
			t.Error(err)
		}
	})

	t.Run("GetByName", func(t *testing.T) {
		if _, err := web.GetFolder(newFolderURI).Files().GetByName("File_1.txt").Get(); err != nil {
			t.Error(err)
		}
	})

	t.Run("GetFile", func(t *testing.T) {
		if _, err := web.GetFile(newFolderURI + "/File_2.txt").Get(); err != nil {
			t.Error(err)
		}
	})

	if _, err := web.GetFolder(newFolderURI).Delete(); err != nil {
		t.Error(err)
	}
}
