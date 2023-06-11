package controllers

import (
	"backend/services"
	"backend/universalfunctions"
	"net/http"

	"github.com/gin-gonic/gin"
)

type refJafungHandler struct {
	refJafungService services.RefJafungService
}

func NewRefJafungHandler(refJafungService services.RefJafungService) *refJafungHandler {
	return &refJafungHandler{refJafungService}
}

func (h *refJafungHandler) RootHandler(c *gin.Context) {
	refJafung, err := h.refJafungService.FindAll()

	universalfunctions.PanicErr(err)

	c.JSON(http.StatusOK, refJafung)
}
