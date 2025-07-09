package test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/PedroPereiraN/go-hexagonal/adapter/input/controller"
	"github.com/PedroPereiraN/go-hexagonal/domain"
	"github.com/PedroPereiraN/go-hexagonal/tests/config"
	"github.com/PedroPereiraN/go-hexagonal/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserController_List(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()
	service := mocks.NewMockUserService(crtl)
	controller := controller.NewUserController(service)


	t.Run("id_is_invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{"id": {"TEST_ERROR"}}

		config.MakeRequest(context, params, url, "GET", nil)
		controller.List(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("user_not_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		invalidId := uuid.New()
		url := url.Values{"id": {invalidId.String()}}

		service.EXPECT().List(invalidId).Return(domain.UserDomain{}, errors.New("User not found"))

		config.MakeRequest(context, params, url, "GET", nil)
		controller.List(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		validId := uuid.New()
		url := url.Values{"id": {validId.String()}}

		service.EXPECT().List(validId).Return(domain.UserDomain{}, nil)

		config.MakeRequest(context, params, url, "GET", nil)
		controller.List(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

	t.Run("users_not_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{}

		service.EXPECT().ListAll().Return([]domain.UserDomain{}, errors.New("Error while fetching users"))

		config.MakeRequest(context, params, url, "GET", nil)
		controller.List(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("users_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := config.GetTestGinContext(recorder)

		params := []gin.Param{}

		url := url.Values{}

		service.EXPECT().ListAll().Return([]domain.UserDomain{}, nil)

		config.MakeRequest(context, params, url, "GET", nil)
		controller.List(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
