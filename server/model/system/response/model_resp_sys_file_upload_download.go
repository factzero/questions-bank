package response

import "server/model/system"

type FileResponse struct {
	File system.FileUploadAndDownload `json:"file"`
}
