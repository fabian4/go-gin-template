package router

import (
	"github.com/fabian4/go-gin-template/controller"
	"github.com/gin-gonic/gin"
)

var demoApi = controller.DemoApi{}

func InitDemoRouter(r *gin.Engine) {
	r.GET("/demo", demoApi.DoController)
	r.GET("/get", demoApi.GetController)
}
