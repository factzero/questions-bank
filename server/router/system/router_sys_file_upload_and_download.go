package system

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload").Use(middleware.OperationRecord())
	fileUploadAndDownloadApi := v1.ApiGroupApp.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", fileUploadAndDownloadApi.UploadFile)       // 上传文件
		fileUploadAndDownloadRouter.POST("getFileList", fileUploadAndDownloadApi.GetFileList) // 获取上传文件列表
		fileUploadAndDownloadRouter.POST("deleteFile", fileUploadAndDownloadApi.DeleteFile)   // 删除指定文件
	}
}
