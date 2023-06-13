package controllers

import (
	"backend/services"
	"backend/universals"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type refJafungHandler struct {
	refJafungService services.RefJafungService
}

func NewRefJafungHandler(refJafungService services.RefJafungService) *refJafungHandler {
	return &refJafungHandler{refJafungService}
}

func (h *refJafungHandler) RootHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	refJafung, err := h.refJafungService.GetAllPaginatedJafung(page)
	universals.PanicErr(err)
	c.JSON(http.StatusOK, refJafung)
}
