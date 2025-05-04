package main

import (
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/controller"
	"github.com/PedroPereiraN/go-hexagonal/adapter/output/repository/user"
	"github.com/PedroPereiraN/go-hexagonal/service"
	"github.com/gin-gonic/gin"
  swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
  _ "github.com/PedroPereiraN/go-hexagonal/docs"
)
//@title GO HEXAGONAL
//@description Simple api made with golang and hexagonal architecture (ports and adapters).
//@host localhost:8080
//@BasePath /

func main() {
  r := gin.Default()

  uRepository := repository.NewUserRepository()

  uService := service.NewUserService(uRepository)

  uController := controller.NewUserController(uService)

  r.POST("/user", uController.Create)

  r.GET("/user", uController.List)
  r.PUT("/user", uController.Update)
  r.PATCH("/user/change-password", uController.UpdatePassword)
  r.DELETE("/user", uController.Delete)

  r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

  r.Run()
}
