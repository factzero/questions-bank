package service

import (
	"errors"
	"mime/multipart"
	"server/global"
	"server/model/common/request"
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

func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.GVA_DB.Model(&system.FileUploadAndDownload{})
	var fileLists []system.FileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

func (e *FileUploadAndDownloadService) FindFile(id uint) (system.FileUploadAndDownload, error) {
	var file system.FileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

func (e *FileUploadAndDownloadService) DeleteFile(file system.FileUploadAndDownload) (err error) {
	var fileFromDb system.FileUploadAndDownload
	fileFromDb, err = e.FindFile(file.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("ÎÄ¼þÉ¾³ýÊ§°Ü")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}
