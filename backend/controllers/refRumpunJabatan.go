package controllers

import (
	"backend/universals"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *refRumpunJabatanHandler) RootHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	refJafung, err := h.refRumpunJabatanService.GetAllPaginatedRumpunJabatan(page)
	universals.PanicErr(err)
	c.JSON(http.StatusOK, refJafung)
}
