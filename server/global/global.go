package global

import (
	"server/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var GVA_CONFIG config.CfgServer
var GVA_VP *viper.Viper
var GVA_DB *gorm.DB
