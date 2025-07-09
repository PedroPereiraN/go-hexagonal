package main

import (
	"database/sql"
	"fmt"
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/controller"
	"github.com/PedroPereiraN/go-hexagonal/adapter/output/repository"
	"github.com/PedroPereiraN/go-hexagonal/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
  _ "github.com/PedroPereiraN/go-hexagonal/docs"
)

//@title GO HEXAGONAL
//@description Simple api made with golang and hexagonal architecture (ports and adapters).
//@host localhost:8080
//@BasePath /
func main() {
	router := gin.Default()

	db, err := sql.Open("postgres", "postgres://postgres:postgres@postgres:5432/db?sslmode=disable")
	defer db.Close()


	if err != nil {
		fmt.Println(err)

		return
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)

		return
	}

	// db e route user
	uRepository := repository.NewUserRepository(db)

	err = uRepository.CreateTable()

	if err != nil {
		fmt.Println(err)

		return
	}

	uService := service.NewUserService(uRepository)

	uController := controller.NewUserController(uService)

	router.POST("/user", uController.Create)
	router.GET("/user", uController.List)
	router.DELETE("/user", uController.Delete)
	router.PUT("/user", uController.Update)
	router.PATCH("/user/update-password", uController.UpdatePassword)
	router.POST("/user/login", uController.Login)

  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
