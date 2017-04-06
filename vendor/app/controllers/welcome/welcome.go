package welcome

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.HTML(200, "welcome", gin.H{
		"title": "Introduce to Orchid",
	})
}

func About(c *gin.Context) {
	c.HTML(200, "about", gin.H{
		"title": "About",
	})
}

func WelcomeApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Orchid!", "status": 2000})
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are requesting for some api..", "status": 2000})
}
