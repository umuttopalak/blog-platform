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

// GetCategories godoc
// @Summary Retrieve all categories
// @Description Get a list of all categories
// @Tags Categories
// @Produce json
// @Success 200 {object} []responses.CategoryResponse
// @Failure 500 {object} responses.ErrorResponse "Could not retrieve categories"
// @Router /category [get]
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

// GetCategory godoc
// @Summary Retrieve a single category
// @Description Get a category by its ID
// @Tags Categories
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {object} responses.CategoryResponse
// @Failure 404 {object} responses.ErrorResponse "Category not found"
// @Router /category/{category_id} [get]
func GetCategory(c *gin.Context) {
	categoryID := c.Param("category_id")

	var category models.Category
	if err := database.DB.Where("id = ?", categoryID).First(&category).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Category not found", nil)
		return
	}

	response := responses.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	utils.CreateResponse(c, http.StatusOK, "Category retrieved successfully", response)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with a given name
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body requests.CreateCategoryRequest true "Category information"
// @Success 200 {object} responses.CategoryResponse
// @Failure 400 {object} responses.ErrorResponse "Invalid input"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Could not create category"
// @Router /category [post]
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
		ID:   category.ID,
		Name: category.Name,
	}
	utils.CreateResponse(c, http.StatusOK, "Category created successfully", response)
}
