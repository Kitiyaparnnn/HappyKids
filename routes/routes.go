package routes

import (
	"example/happy-kids/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/add-product", controllers.ProductViewAdmin())
	incomingRoutes.GET("/users/product-view", controllers.ProductViewUser())
	incomingRoutes.GET("/users/search", controllers.SearchProduct())
}
