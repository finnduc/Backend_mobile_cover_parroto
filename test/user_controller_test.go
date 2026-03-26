package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-familytree/internal/controller"
	"go-familytree/internal/repo"
	"go-familytree/internal/service"
	"go-familytree/pkg/response"
)

func TestRegisterEndpoint(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Dependencies
	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Setup Router
	r := gin.Default()
	r.POST("/user/register", userController.Register)

	// Create Request Body
	input := service.UserRegisterInput{
		Email:   "test@example.com",
		Purpose: "study",
	}
	body, _ := json.Marshal(input)

	// Create Request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Assert Status
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert Response
	var resp response.ResponseData
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.Code)
	
	data := resp.Data.(map[string]interface{})
	assert.Equal(t, "test@example.com", data["email"])
}
