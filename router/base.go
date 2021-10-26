package router

import (
	"github.com/fabian4/go-gin-template/controller"
	"github.com/gin-gonic/gin"
)

var baseApi = controller.BaseApi{}

func InitBaseRouter(r *gin.Engine) {
	r.GET("/demo", baseApi.DoController)
	r.GET("/get", baseApi.GetController)
	r.POST("/form", baseApi.PostFormController)
	r.POST("/json", baseApi.PostJsonController)
	r.POST("/res", baseApi.PostJsonFromStrutController)
}
