package system

import (
	"fmt"
	"server/model/common/request"
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

func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		fmt.Println("获取失败!", err.Error())
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file system.FileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		fmt.Println("删除失败!", err.Error())
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
