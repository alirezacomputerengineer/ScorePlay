package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTag godoc
// @Summary Create a tag
// @Description Create a new tag
// @Tags tags
// @Accept json
// @Produce json
// @Param tag body models.Tag true "Tag"
// @Success 201 {object} models.Tag
// @Router /tags [post]
func CreateTag(c *gin.Context) {
	var newTag models.Tag
	if err := c.ShouldBindJSON(&newTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTag.ID = uuid.New().String()
	models.Tags = append(models.Tags, newTag)
	c.JSON(http.StatusCreated, newTag)
}

// ListTags godoc
// @Summary List all tags
// @Description Get all tags
// @Tags tags
// @Produce json
// @Success 200 {array} models.Tag
// @Router /tags [get]
func ListTags(c *gin.Context) {
	c.JSON(http.StatusOK, models.Tags)
}
