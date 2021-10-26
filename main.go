package main

import (
	"fmt"
	"github.com/fabian4/go-gin-template/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitBaseRouter(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
