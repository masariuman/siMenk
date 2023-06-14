package controllers

import "backend/services"

type refJafungHandler struct {
	refJafungService services.RefJafungService
}

type refRumpunJabatanHandler struct {
	refRumpunJabatanService services.RefRumpunJabatanService
}

type refSubRumpunJabatanHandler struct {
	refSubRumpunJabatanService services.RefSubRumpunJabatanService
}
