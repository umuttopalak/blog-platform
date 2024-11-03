package controllers

import (
	"blog-platform/database"
	"blog-platform/models"
	"blog-platform/requests"
	"blog-platform/responses"
	"blog-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not retrieve categories", nil)
		return
	}

	var responseCategories []responses.CategoryResponse
	for _, category := range categories {
		responseCategories = append(responseCategories, responses.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	utils.CreateResponse(c, http.StatusOK, "Categories retrieved successfully.", responseCategories)

}

func GetCategory(c *gin.Context) {
	category_id := c.Param("category_id")

	var category models.Category
	if err := database.DB.Where("id = ?", category_id).First(&category).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Category not found", nil)
		return
	}

	response := responses.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	utils.CreateResponse(c, http.StatusOK, "Category retrieved successfully", response)

}

func CreateCategory(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.CreateResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var input requests.CreateCategoryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	category := models.Category{
		Name:      input.Name,
		CreatedBy: user.(models.User).ID,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not create category", nil)
		return
	}
	response := responses.CategoryResponse{
		Name: category.Name,
	}
	utils.CreateResponse(c, http.StatusOK, "Category created successfully", response)
}
