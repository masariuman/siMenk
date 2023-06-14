package routes

import (
	"backend/configs"
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func refSubRumpunJabatan(superRoute *gin.RouterGroup) {
	refSubRumpunJabatanRepository := repositories.NewRefSubRumpunJabatanRepository(configs.Connect)
	refSubRumpunJabatanService := services.NewRefSubRumpunJabatanService(refSubRumpunJabatanRepository)
	refSubRumpunJabatanHandler := controllers.NewRefSubRumpunJabatanHandler(refSubRumpunJabatanService)
	refSubRumpunJabatan := superRoute.Group("/referensi/subrumpun_jabatan")
	{
		refSubRumpunJabatan.GET("", refSubRumpunJabatanHandler.RootHandler)
	}
}
