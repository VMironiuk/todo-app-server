package handler

import (
	"net/http"

	todoappserver "github.com/VMironiuk/todo-app-server"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input todoappserver.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
