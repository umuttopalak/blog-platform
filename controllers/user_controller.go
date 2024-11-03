package controllers

import (
	"blog-platform/database"
	"blog-platform/models"
	"blog-platform/requests"
	"blog-platform/responses"
	"blog-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser godoc
// @Summary Kullanıcı girişi
// @Description Kullanıcı email ve şifre ile giriş yapar
// @Tags User
// @Accept json
// @Produce json
// @Param user body requests.UserLoginRequest true "Giriş bilgileri"
// @Success 200 {object} responses.LoginResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 401 {object} responses.ErrorResponse "Yetkisiz"
// @Router /users/login [post]
func LoginUser(c *gin.Context) {
	var input requests.UserLoginRequest
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

	loginResponse := responses.LoginResponse{
		Message: "Login successful",
		Token:   token,
	}
	utils.CreateResponse(c, http.StatusOK, "Login successful", loginResponse)
}

// RegisterUser godoc
// @Summary Yeni kullanıcı kaydı
// @Description Yeni bir kullanıcı kaydı oluşturur
// @Tags User
// @Accept json
// @Produce json
// @Param user body requests.UserRegisterRequest true "Kayıt bilgileri"
// @Success 200 {object} responses.RegisterResponse
// @Failure 400 {object} responses.ErrorResponse "Geçersiz veri"
// @Failure 500 {object} responses.ErrorResponse "Sunucu hatası"
// @Router /users/register [post]
func RegisterUser(c *gin.Context) {
	var input requests.UserRegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.CreateResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not hash password", nil)
		return
	}

	user := models.User{
		ID:        uuid.New(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.CreateResponse(c, http.StatusInternalServerError, "Could not create user", nil)
		return
	}

	userResponse := responses.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	registerResponse := responses.RegisterResponse{
		Message: "User registered successfully",
		User:    userResponse,
	}
	utils.CreateResponse(c, http.StatusOK, "User registered successfully", registerResponse)
}
