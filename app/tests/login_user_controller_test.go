package test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/controller"
	"github.com/PedroPereiraN/go-hexagonal/adapter/input/model"
	"github.com/PedroPereiraN/go-hexagonal/tests/config"
	"github.com/PedroPereiraN/go-hexagonal/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	gomock "go.uber.org/mock/gomock"
)

func TestUserController_Login(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()
	service := mocks.NewMockUserService(crtl)
	controller := controller.NewUserController(service)


	t.Run("invalid_fields", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{}

		model := model.UserLoginModel{
			Email:    "not-an-email",
			Password: "123",
		}


		body, _ := json.Marshal(model)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		config.MakeRequest(context, params, url, "POST", stringReader)

		controller.Login(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("service_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{}

		model := model.UserLoginModel{
			Email: "test@email.com",
			Password: "password@123",
		}

		body, _ := json.Marshal(model)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		service.EXPECT().Login(model.Email, model.Password).Return("", errors.New("Invalid user values"))

		config.MakeRequest(context, params, url, "POST", stringReader)

		controller.Login(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("login_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{}

		model := model.UserLoginModel{
			Email: "test@email.com",
			Password: "password@123",
		}

		body, _ := json.Marshal(model)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		service.EXPECT().Login(model.Email, gomock.Any()).Return("super-token",nil)

		config.MakeRequest(context, params, url, "POST", stringReader)

		controller.Login(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
