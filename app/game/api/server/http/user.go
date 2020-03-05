package http

import (
	"github.com/gin-gonic/gin"

	"solo/solo-go-common/ecode"
)

var ec ecode.Codes

func init()  {
	ec = ecode.New(1)
}


func User(c *gin.Context)  {


	c.Set("err", nil)


	users := srv.User(c)

	c.JSON(200, gin.H{
		"code": ec.Code(),
		"message": ec.Message(),
		"data": users,
	})
}
