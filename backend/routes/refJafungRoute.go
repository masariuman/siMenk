package routes

import (
	"backend/configs"
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func refJafung(superRoute *gin.RouterGroup) {
	refJafungRepository := repositories.NewRefJafungRepository(configs.Connect)
	refJafungService := services.NewRefJafungService(refJafungRepository)
	refJafungHandler := controllers.NewRefJafungHandler(refJafungService)
	refJafung := superRoute.Group("/referensi/jafung")
	{
		refJafung.GET("", refJafungHandler.RootHandler)
	}
}
