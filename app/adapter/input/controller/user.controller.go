package controller

import (
	"net/http"
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/model"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/port/input"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserController(
  service port.UserService,
) UserController {
  return &userController{
    service: service,
  }
}

type UserController interface {
  Create(c *gin.Context)
  List(c *gin.Context)
}

type userController struct {
  service port.UserService
}

func (controller *userController) Create(c *gin.Context) {

  var userRequest model.UserRequestModel

  if err := c.ShouldBindJSON(&userRequest); err != nil {

    c.JSON(http.StatusBadRequest, err.Error())

		return
	}

  uDomain := domain.UserDomain{
    Email: userRequest.Email,
		Password: userRequest.Password,
		Name: userRequest.Name,
  }

  result, err := controller.service.Create(uDomain)

  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  c.JSON(http.StatusOK, result)
}

func (controller *userController) List(c *gin.Context) {

  paramsId := c.Query("id")

  if paramsId != "" {
    userId, err := uuid.Parse(paramsId)

    if err != nil {
      c.JSON(http.StatusBadRequest, "Invalid id")
      return
    }

    result, err := controller.service.List(userId)

    if err != nil {
      c.JSON(http.StatusBadRequest, "User not found")
      return
    }

    c.JSON(http.StatusOK, result)

    return
  }

  result, err := controller.service.ListAll()

  if err != nil {
    c.JSON(http.StatusInternalServerError, "Error while fetching users")
    return
  }


  c.JSON(http.StatusOK, result)
}
