package handler

import (
	"net/http"
	"server-vanstartup/helper"
	"server-vanstartup/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(us user.Service) *UserHandler {
	return &UserHandler{us}
}

func (h *UserHandler) RegisteUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err!= nil {
		errors := helper.FormatError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

        c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Register failed", http.StatusUnprocessableEntity, "FAILED", errorMessages))
		return
    }

	userService, err := h.userService.RegisteUser(input)
	if err!= nil {
        c.JSON(http.StatusBadRequest, helper.ApiResponse("Register failed", 400, "FAILED", nil))
		return
    }

	token := "tokenexample"
	formatter := user.UserFormatter(userService, token)

	c.JSON(http.StatusOK, helper.ApiResponse("Register successfully", 201, "SUCCESS", formatter))
}