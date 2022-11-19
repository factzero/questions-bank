package system

import (
	"fmt"
	"server/model/common/response"
	"server/model/system"
	systemRes "server/model/system/response"

	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadApi struct{}

func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file system.FileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("接收文件失败!", err.Error())
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		fmt.Println("修改数据库链接失败!", err.Error())
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(systemRes.FileResponse{File: file}, "上传成功", c)
}
