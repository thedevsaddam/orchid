package routes

import (
	"github.com/gin-gonic/gin"

	"app/middlewares/cors"
	"app/controllers/users"
	"app/middlewares/example"
	"app/controllers/welcome"
)

func InitRoutes(router gin.Engine) {
	//use middleware
	router.Use(cors.CORSMiddleware())

	//root route
	router.GET("/", welcome.Welcome)
	router.GET("/about", welcome.About)
	router.GET("/users", users.FetchUsers)
	router.GET("/api/v1/welcome", welcome.WelcomeApi)
	//route group
	v1 := router.Group("/api/v1").Use(example.Example())
	{
		v1.GET("/", welcome.Hello)
	}
}
