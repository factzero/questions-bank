package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	// fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload").Use(middleware.OperationRecord())
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	fileUploadAndDownloadApi := v1.ApiGroupApp.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", fileUploadAndDownloadApi.UploadFile) // 上传文件
	}
}
