package controller

import (
	base2 "github.com/fabian4/go-gin-template/request/base"
	"github.com/fabian4/go-gin-template/response"
	"github.com/fabian4/go-gin-template/response/base"
	"github.com/fabian4/go-gin-template/service"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

var baseService = service.BaseService{}

func (BaseApi *BaseApi) DoController(c *gin.Context) {
	// call service method
	baseService.DoService()
}

func (BaseApi *BaseApi) GetController(c *gin.Context) {
	name := c.Query("name")
	sex := c.DefaultQuery("sex", "male")
	response.Success(c, map[string]string{
		"name": name,
		"sex":  sex,
	})
	return
}

func (BaseApi *BaseApi) GetPathParamController(c *gin.Context) {
	name := c.Param("name")
	sex := c.DefaultQuery("sex", "male")
	response.Success(c, map[string]string{
		"name": name,
		"sex":  sex,
	})
	return
}

func (BaseApi *BaseApi) PostFormController(c *gin.Context) {
	name := c.PostForm("name")
	sex := c.DefaultPostForm("sex", "male")
	response.Success(c, map[string]string{
		"name": name,
		"sex":  sex,
	})
	return
}

func (BaseApi *BaseApi) PostJsonController(c *gin.Context) {
	// request参数绑定map结构
	json := make(map[string]string)
	_ = c.Bind(&json)
	response.Success(c, map[string]string{
		"name": json["name"],
		"sex":  json["sex"],
	})
	return
}

func (BaseApi *BaseApi) PostJsonFromStrutController(c *gin.Context) {
	// request参数绑定自定义结构体
	json := base2.NameAndAge{}
	_ = c.Bind(&json)
	response.Success(c, base.NameAndAge{
		Name: json.Name,
		Age:  json.Age,
	})
	return
}
