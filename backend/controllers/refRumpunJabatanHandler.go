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

func (h *refRumpunJabatanHandler) EditHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	refRumpunJabatan, err := h.refRumpunJabatanService.GetById(ID)
	universals.PanicErr(err)
	c.JSON(http.StatusOK, gin.H{"Sukses": "Data Berhasil Disimpan.", "data": refRumpunJabatan})
}

func (h *refRumpunJabatanHandler) UpdateHandler(c *gin.Context) {
	var rumpunJabatan requests.RefRumpunJabatanRequest
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	err := c.ShouldBindJSON(&rumpunJabatan)
	err = h.refRumpunJabatanService.Update(ID, rumpunJabatan)
	universals.PanicErr(err)
	h.RootHandler(c)
}
