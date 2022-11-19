package service

import (
	"mime/multipart"
	"server/global"
	"server/model/system"
	"server/utils/upload"
	"strings"
)

type FileUploadAndDownloadService struct{}

func (e *FileUploadAndDownloadService) Upload(file system.FileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (file system.FileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := system.FileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, e.Upload(f)
	}
	return
}
