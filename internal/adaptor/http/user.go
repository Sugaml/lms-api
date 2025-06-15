package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddUser			godoc
// @Summary			Add a new User
// @Description		Add a new User
// @Tags			User
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			UserRequest			body		domain.UserRequest		true		"Add User Request"
// @Success			200					{object}	domain.UserResponse					"User created"
// @Router			/users 				[post]
func (h *Handler) CreateUser(ctx *gin.Context) {
	var req *domain.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateUser(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// LoginUser		godoc
// @Summary			Login User
// @Description		Login User
// @Tags			User
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			UserRequest			body		domain.LoginRequest		true		"Login User Request"
// @Success			200					{object}	domain.LoginUserResponse			"LoginUser Reponse"
// @Router			/users/login 				[post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	var req *domain.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.LoginUser(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListUser 		godoc
// @Summary 		List User
// @Description 	List User
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			query 						query 		string 		false 	"query"
// @Success 		200 		{array} 		domain.UserResponse
// @Router 			/users	 	[get]
func (h *Handler) ListUser(ctx *gin.Context) {
	var req domain.UserListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	result, count, err := h.svc.ListUser(&req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetUser 			godoc
// @Summary 		Get User
// @Description 	Get User from Id
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "User id"
// @Success 		200 {object} domain.UserResponse
// @Router 			/users/{id} [get]
func (h *Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.GetUser(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateUser		godoc
// @Summary 			Update User
// @Description 		Update User from Id
// @Tags 				User
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 							path 		string 								true 	"User id"
// @Param 				UserUpdateRequest	 		body 		domain.UserUpdateRequest 	true 	"Update User Response request"
// @Success 			200 						{object} 	domain.UserResponse
// @Router 				/users/{id} 				[put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.UpdateUser(id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// DeleteUser 		godoc
// @Summary 			Delete User
// @Description 		Delete User from Id
// @Tags 				User
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			UserAuth
// @Param 				id 						path 		string 						true 	"User id"
// @Success 			200 					{object} 	domain.UserResponse
// @Router 				/users/{id} 	[delete]
func (ch *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required User id"))
		return
	}
	result, err := ch.svc.DeleteUser(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
