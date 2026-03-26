package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-familytree/internal/controller"
	"go-familytree/pkg/response"
)

func TestGetUserEndpoint(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Router
	r := gin.Default()
	userController := controller.NewUserController()
	r.GET("/user/User", userController.GetUser)

	// Create Request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/User", nil)
	r.ServeHTTP(w, req)

	// Assert Status
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert Response
	var resp response.ResponseData
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "Finn", resp.Data)
}

func TestGetFamilyEndpoint(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Router
	r := gin.Default()
	userController := controller.NewUserController()
	r.GET("/user/Family", userController.GetFamilyController)

	// Create Request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/Family", nil)
	r.ServeHTTP(w, req)

	// Assert Status
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert Response
	var resp response.ResponseData
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.Code)
	
	family, ok := resp.Data.([]interface{})
	assert.True(t, ok)
	assert.Len(t, family, 3)
}
