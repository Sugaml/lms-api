package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListLibraryDashboardStats	godoc
// @Summary 		List LibraryDashboard
// @Description 	List LibraryDashboard
// @Tags 			Report
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Success 		200 {object} domain.LibraryDashboardStats
// @Router 			/reports/dashboard-stats	[get]
func (h *Handler) GetLibraryDashboardStats(ctx *gin.Context) {
	result, err := h.svc.GetLibraryDashboardStats()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListBorrowStats	godoc
// @Summary 		List Borrow
// @Tags 			Report
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

// ListBookProgramstats	godoc
// @Summary 			List Borrow
// @Tags 				Report
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Success 			200 {array} domain.BookProgramstats
// @Router 				/reports/program-stats	[get]
func (h *Handler) GetBookProgramstats(ctx *gin.Context) {
	result, err := h.svc.GetBookProgramstats()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
