package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"path/filepath"

	"github.com/google/uuid"
)

// CreateMedia godoc
// @Summary Create media
// @Description Upload and create a media with tags
// @Tags media
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Media Name"
// @Param tags formData []string true "Tags"
// @Param file formData file true "File"
// @Success 201 {object} models.Media
// @Router /media [post]
func CreateMedia(c *gin.Context) {
	name := c.PostForm("name")
	tags := c.PostFormArray("tags")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
		return
	}

	// Save the file
	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}

	newMedia := models.Media{
		ID:      uuid.New().String(),
		Name:    name,
		Tags:    tags,
		FileUrl: "/uploads/" + file.Filename,
	}

	models.Medias = append(models.Medias, newMedia)
	c.JSON(http.StatusCreated, newMedia)
}

// SearchMedia godoc
// @Summary Search media by tag
// @Description Search media items by tag
// @Tags media
// @Produce json
// @Param tag query string true "Tag ID"
// @Success 200 {array} models.Media
// @Router /media [get]
func SearchMedia(c *gin.Context) {
	tag := c.Query("tag")
	var result []models.Media

	for _, media := range models.Medias {
		for _, t := range media.Tags {
			if t == tag {
				result = append(result, media)
				break
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
