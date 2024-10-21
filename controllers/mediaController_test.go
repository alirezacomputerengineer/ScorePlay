package controllers

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"myapp/models"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateMedia(t *testing.T) {
	r := gin.Default()
	r.POST("/media", CreateMedia)

	// Mock file to upload
	file, _ := os.CreateTemp("", "testimage.jpg")
	defer os.Remove(file.Name())

	// Prepare form data
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "super nice picture")
	writer.WriteField("tags", "1234")

	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	part.Write([]byte("fake image content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/media", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responseMedia models.Media
	json.Unmarshal(w.Body.Bytes(), &responseMedia)

	assert.Equal(t, "super nice picture", responseMedia.Name)
	assert.NotEmpty(t, responseMedia.FileUrl)
}

func TestCreateMediaMissingFields(t *testing.T) {
	r := gin.Default()
	r.POST("/media", CreateMedia)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("tags", "1234")
	writer.Close()

	req := httptest.NewRequest("POST", "/media", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSearchMedia(t *testing.T) {
	r := gin.Default()
	r.GET("/media", SearchMedia)

	// Add mock data to Media array
	models.Medias = []models.Media{
		{
			ID:      "c906cbbf-1a25-4a99-b223-34bcf6e3b8a7",
			Name:    "super nice picture",
			Tags:    []string{"Zinedine Zidane", "Real Madrid", "Champions League"},
			FileUrl: "https://some_url.com/file.jpg",
		},
	}

	// Test search by tag
	req := httptest.NewRequest("GET", "/media?tag=Real%20Madrid", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseMedia []models.Media
	json.Unmarshal(w.Body.Bytes(), &responseMedia)

	assert.Equal(t, 1, len(responseMedia))
	assert.Equal(t, "super nice picture", responseMedia[0].Name)
}
