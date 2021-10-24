package controller

import (
	"github.com/fabian4/go-gin-template/response"
	"github.com/fabian4/go-gin-template/service"
	"github.com/gin-gonic/gin"
)

type DemoApi struct {
}

var demoService = service.DemoService{}

func (DemoApi DemoApi) DoController(c *gin.Context) {
	// call service method
	demoService.DoService()
}

func (DemoApi DemoApi) GetController(c *gin.Context) {
	name := c.Query("name")
	response.Success(c, name)
	return
}
