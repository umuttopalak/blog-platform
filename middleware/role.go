package middleware

import (
	"blog-platform/models"
	"blog-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(roleName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Kullanıcıyı oturumdan veya JWT token'dan alabilirsiniz (örnek için basit bir kullanıcı modeli varsayıyoruz)
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı oturum açmamış"})
			c.Abort()
			return
		}

		// Kullanıcının belirtilen role sahip olup olmadığını kontrol et
		if !utils.HasRole(user.(models.User), roleName) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Yetkisiz erişim"})
			c.Abort()
			return
		}

		c.Next()
	}
}
