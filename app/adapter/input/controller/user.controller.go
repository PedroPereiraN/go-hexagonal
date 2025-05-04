package controller

import (
	"net/http"
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/model"
	"github.com/PedroPereiraN/go-hexagonal/domain"
  "github.com/PedroPereiraN/go-hexagonal/helpers"
	"github.com/PedroPereiraN/go-hexagonal/port/input"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
  "fmt"
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
  Update(c *gin.Context)
  Delete(c *gin.Context)
  UpdatePassword(c *gin.Context)
}

type userController struct {
  service port.UserService
}

  // @Summary create user
  // @Description create a new user
  // @Tags user
  // @Accept json
  // @Produce json
  // @Param user body model.UserRequestModel true "user"
  // @Success 200 "User created successfully"
  // @Failure 400 "invalid values"
  // @Failure 500 "Internal server error"
  // @Router /user [post]
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


  // @Summary list users
  // @Description list all users or specify one user using his id
  // @Tags user
  // @Accept json
  // @Produce json
  // @Param id query string false "user id"
  // @Success 200 {array} domain.UserDomain
  // @Failure 400 "User not found"
  // @Failure 500 "Internal server error"
  // @Router /user [get]
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

  // @Summary edit user
  // @Description edit an user
  // @Tags user
  // @Accept json
  // @Produce json
  // @Param user body model.UserUpdateModel true "user"
  // @Param id query string true "user id"
  // @Success 200 "User edited successfully"
  // @Failure 400 "invalid values"
  // @Failure 500 "Internal server error"
  // @Router /user [put]
func (controller *userController) Update(c *gin.Context) {

  paramsId := c.Query("id")

  if paramsId == "" {
    c.JSON(http.StatusBadRequest, "Unspecified user")
    return
  }

  userId, err := uuid.Parse(paramsId)

  if err != nil {
    c.JSON(http.StatusBadRequest, "Invalid id")
    return
  }

  var userRequest model.UserUpdateModel

  if err := c.ShouldBindJSON(&userRequest); err != nil {

    c.JSON(http.StatusBadRequest, err.Error())

		return
	}

  uDomain := domain.UserDomain{
    Email: userRequest.Email,
		Name: userRequest.Name,
  }

  result, err := controller.service.Update(userId, uDomain)

  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  c.JSON(http.StatusOK, result)
}


  // @Summary delete user
  // @Description delete an user
  // @Tags user
  // @Accept json
  // @Produce json
  // @Param id query string true "user id"
  // @Success 200 "User deleted successfully"
  // @Failure 400 "invalid id"
  // @Failure 500 "Internal server error"
  // @Router /user [delete]
func (controller *userController) Delete(c *gin.Context) {

  paramsId := c.Query("id")

  if paramsId == "" {
    c.JSON(http.StatusBadRequest, "Unspecified user")
    return
  }

  userId, err := uuid.Parse(paramsId)

  if err != nil {
    c.JSON(http.StatusBadRequest, "Invalid id")
    return
  }

  result, err := controller.service.Delete(userId)

  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  c.JSON(http.StatusOK, result)
}

  // @Summary update user password
  // @Description update an user password
  // @Tags user
  // @Accept json
  // @Produce json
  // @Param id query string true "user id"
  // @Param password body model.UserUpdatePasswordModel true "new password"
  // @Success 200 "User password edited successfully"
  // @Failure 400 "invalid values"
  // @Failure 500 "Internal server error"
  // @Router /user/change-password [patch]
func (controller *userController) UpdatePassword(c *gin.Context) {
  paramsId := c.Query("id")

  if paramsId == "" {
    c.JSON(http.StatusBadRequest, "Unspecified user")
    return
  }

  userId, err := uuid.Parse(paramsId)

  if err != nil {
    c.JSON(http.StatusBadRequest, "Invalid id")
    return
  }

  var passwordRequest model.UserUpdatePasswordModel

  if err := c.ShouldBindJSON(&passwordRequest); err != nil {

    c.JSON(http.StatusBadRequest, err.Error())

		return
	}

  uDomain := domain.UserDomain{
    Password: passwordRequest.Password,
  }

  userInfo, err := controller.service.List(userId)

  if err != nil {
    fmt.Println("estamos pegando o info do usuário")
    fmt.Println(err)
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  isTheSame, err := helpers.VerifyOldPassword(userInfo, uDomain.Password)

  if err != nil {
    fmt.Println("estamos usando a função de verificação")
    fmt.Println(err)
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  if isTheSame {
    c.JSON(http.StatusBadRequest, "New password cannot be the same as the previous one")
    return
  }

  result, err := controller.service.UpdatePassword(userId, uDomain.Password)

  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  c.JSON(http.StatusOK, result)
}
