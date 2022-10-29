package response

import (
	"server/model/system"
)

type SysLoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
