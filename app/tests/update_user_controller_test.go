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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	gomock "go.uber.org/mock/gomock"
)

func TestUserController_Update(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()
	service := mocks.NewMockUserService(crtl)
	controller := controller.NewUserController(service)

	t.Run("id_is_invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{"id": {"TEST_ERROR"}}

		config.MakeRequest(context, params, url, "PUT", nil)
		controller.Update(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("invalid_fields", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{"id": {uuid.New().String()}}

		model := model.UpdateUserModel{
			Email:    "not-an-email",
			Name:     "ab",
			Phone:    "12345",
		}

		body, _ := json.Marshal(model)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		config.MakeRequest(context, params, url, "PUT", stringReader)

		controller.Update(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("service_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}


		userId := uuid.New()

		url := url.Values{"id": {userId.String()}}

		model := model.UpdateUserModel{
			Email: "test@email.com",
			Name: "Test Name",
			Phone: "00000000000",
		}

		body, _ := json.Marshal(model)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		service.EXPECT().Update(userId, gomock.Any()).Return(uuid.Nil, errors.New("Invalid user values"))

		config.MakeRequest(context, params, url, "POST", stringReader)

		controller.Update(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("update_user_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		userId := uuid.New()

		url := url.Values{"id": {userId.String()}}

		model := model.UpdateUserModel{
			Email: "test@email.com",
			Name: "Test Name",
			Phone: "00000000000",
		}

		body, _ := json.Marshal(model)
		stringReader := io.NopCloser(strings.NewReader(string(body)))

		service.EXPECT().Update(userId, gomock.Any()).Return(userId, nil)

		config.MakeRequest(context, params, url, "PUT", stringReader)

		controller.Update(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
