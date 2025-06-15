package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddBorrow		godoc
// @Summary			Add a new Borrow
// @Description		Add a new Borrow
// @Tags			Borrow
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			BorrowRequest			body		domain.BorrowRequest		true		"Add Borrow Request"
// @Success			200					{object}	domain.BorrowResponse					"Borrow created"
// @Router			/borrows 				[post]
func (h *Handler) CreateBorrow(ctx *gin.Context) {
	var req *domain.BorrowedBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateBorrow(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListBorrow 		godoc
// @Summary 		List Borrow
// @Description 	List Borrow
// @Tags 			Borrow
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			query 						query 		string 		false 	"query"
// @Success 		200 		{array} 		domain.BorrowResponse
// @Router 			/borrows	 	[get]
func (h *Handler) ListBorrow(ctx *gin.Context) {
	var req domain.ListBorrowedBookRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	result, count, err := h.svc.ListBorrow(&req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetBorrow 		godoc
// @Summary 		Get Borrow
// @Description 	Get Borrow from Id
// @Tags 			Borrow
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "Borrow id"
// @Success 		200 {object} domain.BorrowResponse
// @Router 			/borrows/{id} [get]
func (h *Handler) GetBorrow(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.GetBorrow(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateBorrow			godoc
// @Summary 			Update Borrow
// @Description 		Update Borrow from Id
// @Tags 				Borrow
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 							path 		string 								true 	"Borrow id"
// @Param 				BorrowUpdateRequest	 		body 		domain.BorrowUpdateRequest 	true 	"Update Borrow Response request"
// @Success 			200 						{object} 	domain.BorrowResponse
// @Router 				/borrows/{id} 				[put]
func (h *Handler) UpdateBorrow(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.UpdateBorrowedBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.UpdateBorrow(id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// DeleteBorrow 		godoc
// @Summary 			Delete Borrow
// @Description 		Delete Borrow from Id
// @Tags 				Borrow
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			BorrowAuth
// @Param 				id 						path 		string 						true 	"Borrow id"
// @Success 			200 					{object} 	domain.BorrowResponse
// @Router 				/borrows/{id} 	[delete]
func (ch *Handler) DeleteBorrow(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required Borrow id"))
		return
	}
	result, err := ch.svc.DeleteBorrow(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
