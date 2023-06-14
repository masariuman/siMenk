package routes

import "github.com/gin-gonic/gin"

func AddRoutes(superRoute *gin.RouterGroup) {
	refJafung(superRoute)
	refRumpunJabatan(superRoute)
	refSubRumpunJabatan(superRoute)
}
