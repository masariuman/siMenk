package controllers

import (
	"backend/universals"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *refSubRumpunJabatanHandler) RootHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	refJafung, err := h.refSubRumpunJabatanService.GetAllPaginatedSubRumpunJabatan(page)
	universals.PanicErr(err)
	c.JSON(http.StatusOK, refJafung)
}
