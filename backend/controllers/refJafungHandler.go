package controllers

import (
	"backend/universals"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *refJafungHandler) RootHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	refJafung, err := h.refJafungService.GetAllPaginatedJafung(page)
	universals.PanicErr(err)
	c.JSON(http.StatusOK, refJafung)
}
