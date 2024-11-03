package controllers

import (
	"blog-platform/database"
	"blog-platform/models"
	"blog-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.CreateResponse(c, http.StatusUnauthorized, "User not found", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.CreateResponse(c, http.StatusUnauthorized, "Incorrect password", nil)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not generate token", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not hash password", nil)
		return
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not create user", nil)
		return
	}

	utils.CreateResponse(c, http.StatusOK, "User registered successfully", user)
}
