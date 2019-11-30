package api

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
)

func TestFolders(t *testing.T) {
	checkClient(t)

	web := NewSP(spClient).Web()
	newFolderName := uuid.New().String()
	rootFolderURI := getRelativeURL(spClient.AuthCnfg.GetSiteURL()) + "/Shared%20Documents"

	t.Run("Add", func(t *testing.T) {
		if _, err := web.GetFolder(rootFolderURI).Folders().Add(newFolderName); err != nil {
			t.Error(err)
		}
	})

	t.Run("Get", func(t *testing.T) {
		data, err := web.GetFolder(rootFolderURI).Folders().Select("Id").Conf(headers.verbose).Get()
		if err != nil {
			t.Error(err)
		}

		res := &struct {
			D struct {
				Results []interface{} `json:"results"`
			} `json:"d"`
		}{}

		err = json.Unmarshal(data, &res)
		if err != nil {
			t.Error(err)
		}

		if len(res.D.Results) == 0 {
			t.Error("can't get folders")
		}
	})

	t.Run("GetByName", func(t *testing.T) {
		if _, err := web.GetFolder(rootFolderURI).Folders().GetByName(newFolderName).Get(); err != nil {
			t.Error(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		if _, err := web.GetFolder(rootFolderURI + "/" + newFolderName).Delete(); err != nil {
			t.Error(err)
		}
	})

}
