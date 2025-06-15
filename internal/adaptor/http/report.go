package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListBorrowStats	godoc
// @Summary 		List Borrow
// @Description 	Get Borrow from Id
// @Tags 			Borrow
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Success 		200 {object} domain.BorrowedBookStats
// @Router 			/reports/borrowedbookstats	[get]
func (h *Handler) GetBorrowedBookStats(ctx *gin.Context) {
	result, err := h.svc.GetBorrowedBookStats()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
