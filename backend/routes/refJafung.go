package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func refJafung(superRoute *gin.RouterGroup) {
	refJafung := superRoute.Group("/referensi/jafung")
	{
		refJafung.GET("", controllers.RootHandler)
	}
}
