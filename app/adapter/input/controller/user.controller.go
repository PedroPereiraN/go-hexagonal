package controller

import (
	"fmt"
	"net/http"
	"time"
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/model"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/ports/input"
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
	Delete(c *gin.Context)
	Update(c *gin.Context)
	UpdatePassword(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	service port.UserService
}

// @Summary create user
// @Description create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.CreateUserModel true "user"
// @Success 200 "User created successfully"
// @Failure 400 "invalid values"
// @Failure 500 "Internal server error"
// @Router /user [post]
func (controller *userController) Create(c *gin.Context) {

	var userData model.CreateUserModel

	if err := c.ShouldBindJSON(&userData); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	uDomain, err := domain.CreateUser(
		uuid.Nil,
		userData.Name,
		userData.Email,
		userData.Phone,
		userData.Password,
		time.Time{},
		time.Time{},
		time.Time{},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	result, err := controller.service.Create(uDomain)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, "Usuário criado com sucesso: " + result.String())
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

		fmt.Println(err)
		fmt.Println(result)

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

	if err != nil && err.Error() == "sql: no rows in result set" {
    c.JSON(http.StatusBadRequest, "User not found")
    return
  }

	if err !=nil && err.Error() != "sql: no rows in result set" {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, "User deleted successfully: " + result.String())

}

  // @Summary update user
  // @Description update an user
  // @Tags user
  // @Accept json
  // @Produce json
  // @Param id query string true "user id"
  // @Param user body model.UpdateUserModel true "user"
  // @Success 200 "User updated successfully"
  // @Failure 400 "invalid values"
  // @Failure 500 "Internal server error"
  // @Router /user [put]
func (controller *userController) Update(c *gin.Context) {
	paramsId := c.Query("id")
	var userData model.UpdateUserModel

	if paramsId == "" {
		c.JSON(http.StatusBadRequest, "Inform an ID to update user data.")

		return
	}

  userId, err := uuid.Parse(paramsId)

  if err != nil {
    c.JSON(http.StatusBadRequest, "Invalid id")
    return
  }

	if err := c.ShouldBindJSON(&userData); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())

		return
	}


	uDomain, err := domain.CreateUser(
		uuid.Nil,
		userData.Name,
		userData.Email,
		userData.Phone,
		"",
		time.Time{},
		time.Time{},
		time.Time{},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	result, err := controller.service.Update(userId, uDomain)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, "User updated successfully: " + result.String())
}


// @Summary update user password
// @Description update an user password
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "user id"
// @Param password body model.UpdateUserPasswordModel true "new password"
// @Success 200 "User password edited successfully"
// @Failure 400 "invalid values"
// @Failure 500 "Internal server error"
// @Router /user/update-password [patch]
func (controller *userController) UpdatePassword(c *gin.Context) {
	paramsId := c.Query("id")
	var userData model.UpdateUserPasswordModel

	if paramsId == "" {
		c.JSON(http.StatusBadRequest, "Inform an ID to update user data.")

		return
	}

  userId, err := uuid.Parse(paramsId)

  if err != nil {
    c.JSON(http.StatusBadRequest, "Invalid id")
    return
  }

	if err := c.ShouldBindJSON(&userData); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	result, err := controller.service.UpdatePassword(userId, userData.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, "User password updated successfully: " + result.String())
}

// @Summary login
// @Description login with an user
// @Tags user
// @Accept json
// @Produce json
// @Param loginInfo body model.UserLoginModel true "user"
// @Success 200 "User updated successfully"
// @Failure 400 "invalid values"
// @Failure 500 "Internal server error"
// @Router /user/login [post]
func (controller *userController) Login(c *gin.Context) {

	var loginInfo model.UserLoginModel

	if err := c.ShouldBindJSON(&loginInfo); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	result, err := controller.service.Login(loginInfo.Email, loginInfo.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": result,
	})
}
