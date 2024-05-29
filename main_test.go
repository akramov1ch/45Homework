package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	users = append(users, User{ID: 1, Name: "Shaxboz", Email: "shaxbozakramovic@gmail.com"})
	users = append(users, User{ID: 2, Name: "Abduazim", Email: "abduazimyusufov@gmail.com"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response []User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
}

func TestGetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	users = append(users, User{ID: 1, Name: "Shaxboz", Email: "shaxbozakramovic@gmail.com"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Shaxboz", response.Name)
	assert.Equal(t, "shaxbozakramovic@gmail.com", response.Email)
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	newUser := User{Name: "Shaxboz", Email: "shaxbozakramovic@gmail.com"}
	body, _ := json.Marshal(newUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var response User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Shaxboz", response.Name)
	assert.Equal(t, "shaxbozakramovic@gmail.com", response.Email)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	return router
}
