package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leslieleung/gin-application-template/pkg/log"
)

// PingPong godoc
// @Summary A gin example
// @Description
// @Tags autogenerated
// @Produce json
// @Router /ping [get]
func PingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Echo godoc
// @Summary A gin example
// @Description
// @Tags autogenerated
// @Produce json
// @Param msg query string true "message"
// @Router /echo [get]
func Echo(c *gin.Context) {
	msg := c.Query("msg")
	log.GetLogger(c).Infof("echo: %s", msg)
	c.JSON(200, gin.H{
		"message": msg,
	})
}
