package main

import (
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/controller"
	"github.com/PedroPereiraN/go-hexagonal/adapter/output/repository/user"
	"github.com/PedroPereiraN/go-hexagonal/service"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  uRepository := repository.NewUserRepository()

  uService := service.NewUserService(uRepository)

  uController := controller.NewUserController(uService)

  r.POST("/user", uController.Create)
  r.GET("/user", uController.List)
  r.PUT("/user", uController.Update)

  r.Run()
}
