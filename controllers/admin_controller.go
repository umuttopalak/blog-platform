package controllers

import (
	"blog-platform/database"
	"blog-platform/models"
	"blog-platform/requests"
	"blog-platform/responses"
	"blog-platform/utils"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRole godoc
// @Summary Create a new role
// @Description Create a role with a given name and description
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body requests.CreateRoleRequest true "Role information"
// @Success 200 {object} responses.RoleResponse
// @Failure 400 {object} responses.ErrorResponse "Invalid input"
// @Failure 502 {object} responses.ErrorResponse "Database error"
// @Router /admin/role/create [post]
func CreateRole(c *gin.Context) {
	var input requests.CreateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	role := models.Role{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := database.DB.Create(&role).Error; err != nil {
		utils.CreateResponse(c, http.StatusBadGateway, "Database error", nil)
		return
	}

	response := responses.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
	}

	utils.CreateResponse(c, http.StatusOK, "Role created successfully.", response)
}

// RemoveRole godoc
// @Summary Remove an existing role
// @Description Delete a role by its ID
// @Tags Roles
// @Param role_id path int true "Role ID"
// @Success 200 {string} string "Role deleted successfully"
// @Failure 400 {object} responses.ErrorResponse "Invalid role ID"
// @Failure 404 {object} responses.ErrorResponse "Role not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /admin/role/remove/{role_id} [delete]
func RemoveRole(c *gin.Context) {
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid role ID format", nil)
		return
	}

	var role models.Role
	if err := database.DB.Where("id = ?", roleID).First(&role).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Role not found", nil)
		return
	}

	if err := database.DB.Delete(&role).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not delete role", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Role deleted successfully", nil)
}

// AddRoleToUser godoc
// @Summary Assign a role to a user
// @Description Add a role to a specific user by user ID and role ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body requests.AddRoleRequest true "Add role to user"
// @Success 200 {string} string "Role added successfully"
// @Failure 400 {object} responses.ErrorResponse "Invalid input"
// @Failure 404 {object} responses.ErrorResponse "User or role not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /admin/role/add [post]
func AddRoleToUser(c *gin.Context) {
	var input requests.AddRoleRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", input.UserID).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	var role models.Role
	if err := database.DB.First(&role, "id = ?", input.RoleID).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Role not found", nil)
		return
	}

	if err := database.DB.Model(&user).Association("Roles").Append(&role); err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not add role", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Role added successfully", nil)
}

// RemoveRoleFromUser godoc
// @Summary Remove a role from a user
// @Description Delete a role from a specific user by user ID and role ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body requests.RemoveRoleRequest true "Remove role from user"
// @Success 200 {string} string "Role removed successfully"
// @Failure 400 {object} responses.ErrorResponse "Invalid input"
// @Failure 404 {object} responses.ErrorResponse "User or role not found"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /admin/role/remove-from-user [delete]
func RemoveRoleFromUser(c *gin.Context) {
	var input requests.RemoveRoleRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", input.UserID).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	var role models.Role
	if err := database.DB.First(&role, "id = ?", input.RoleID).Error; err != nil {
		utils.CreateResponse(c, http.StatusNotFound, "Role not found", nil)
		return
	}

	if err := database.DB.Model(&user).Association("Roles").Delete(&role); err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not remove role", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Role removed successfully", nil)
}
