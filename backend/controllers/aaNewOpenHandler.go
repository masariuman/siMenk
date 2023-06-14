package controllers

import "backend/services"

func NewRefJafungHandler(refJafungService services.RefJafungService) *refJafungHandler {
	return &refJafungHandler{refJafungService}
}

func NewRefRumpunJabatanHandler(refRumpunJabatanService services.RefRumpunJabatanService) *refRumpunJabatanHandler {
	return &refRumpunJabatanHandler{refRumpunJabatanService}
}

func NewRefSubRumpunJabatanHandler(refSubRumpunJabatanService services.RefSubRumpunJabatanService) *refSubRumpunJabatanHandler {
	return &refSubRumpunJabatanHandler{refSubRumpunJabatanService}
}
