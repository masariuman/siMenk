package routes

import (
	"backend/configs"
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func refRumpunJabatan(superRoute *gin.RouterGroup) {
	refRumpunJabatanRepository := repositories.NewRefRumpunJabatanRepository(configs.Connect)
	refRumpunJabatanService := services.NewRefRumpunJabatanService(refRumpunJabatanRepository)
	refRumpunJabatanHandler := controllers.NewRefRumpunJabatanHandler(refRumpunJabatanService)
	refRumpunJabatan := superRoute.Group("/referensi/rumpun_jabatan")
	{
		refRumpunJabatan.GET("", refRumpunJabatanHandler.RootHandler)
		refRumpunJabatan.POST("", refRumpunJabatanHandler.StoreHandler)
		refRumpunJabatan.GET("/:id", refRumpunJabatanHandler.EditHandler)
		refRumpunJabatan.PUT("/:id", refRumpunJabatanHandler.UpdateHandler)
		refRumpunJabatan.POST("/search", refRumpunJabatanHandler.SearchHandler)
	}
}
