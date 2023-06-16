package controllers

import (
	"backend/models/requests"
	"backend/universals"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *refRumpunJabatanHandler) RootHandler(c *gin.Context) {
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	refRumpunJabatan, err := h.refRumpunJabatanService.GetAllPaginatedRumpunJabatan(page)
	universals.PanicErr(err)
	c.JSON(http.StatusOK, refRumpunJabatan)
}

func (h *refRumpunJabatanHandler) StoreHandler(c *gin.Context) {
	var rumpunJabatan requests.RefRumpunJabatanRequest
	err := c.ShouldBindJSON(&rumpunJabatan)
	refRumpunJabatan, err := h.refRumpunJabatanService.Store(rumpunJabatan, err, c)
	universals.PanicErr(err)
	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Disimpan.", "data": refRumpunJabatan})
}