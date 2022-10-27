package system

import (
	"fmt"
	"net/http"

	"server/model/system/request"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var l request.Login
	if err := c.ShouldBind(&l); err != nil {
		fmt.Println(l)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
