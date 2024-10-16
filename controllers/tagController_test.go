package controllers

import (
	"bytes"
	"encoding/json"
	"myapp/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTag(t *testing.T) {
	r := gin.Default()
	r.POST("/tags", CreateTag)

	// Test creating a valid tag
	tagData := `{"name": "Wembley Stadium"}`
	req := httptest.NewRequest("POST", "/tags", bytes.NewBuffer([]byte(tagData)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responseTag models.Tag
	json.Unmarshal(w.Body.Bytes(), &responseTag)

	assert.Equal(t, "Wembley Stadium", responseTag.Name)
	assert.NotEmpty(t, responseTag.ID)
}

func TestCreateTagInvalid(t *testing.T) {
	r := gin.Default()
	r.POST("/tags", CreateTag)

	// Test creating a tag without a name
	tagData := `{}`
	req := httptest.NewRequest("POST", "/tags", bytes.NewBuffer([]byte(tagData)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListTags(t *testing.T) {
	r := gin.Default()
	r.GET("/tags", ListTags)

	// Assuming we already have some tags created
	models.Tags = []models.Tag{
		{ID: "1", Name: "Wembley Stadium"},
		{ID: "2", Name: "Old Trafford"},
	}

	req := httptest.NewRequest("GET", "/tags", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var tags []models.Tag
	json.Unmarshal(w.Body.Bytes(), &tags)

	assert.Equal(t, 2, len(tags))
	assert.Equal(t, "Wembley Stadium", tags[0].Name)
	assert.Equal(t, "Old Trafford", tags[1].Name)
}
