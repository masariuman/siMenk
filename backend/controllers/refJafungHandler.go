package controllers

import (
	"backend/configs"
	"backend/repositories"
	"backend/services"
	"backend/universalfunctions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	refJafungService := services.NewRefJafungService(repositories.NewRefJafungRepository(configs.Connect))
	refJafung, err := refJafungService.FindAll()

	universalfunctions.PanicErr(err)

	c.JSON(http.StatusOK, refJafung)
}
