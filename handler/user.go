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
        c.JSON(http.StatusBadRequest, helper.ApiResponse("Register failed", 400, "FAILED", err.Error()))
		return
    }

	token := "tokenexample"
	formatter := user.UserFormatter(userService, token)

	c.JSON(http.StatusOK, helper.ApiResponse("Register successfully", 201, "SUCCESS", formatter))
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessages := gin.H{
            "errors": errors,
        }
		c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "FAILED", errorMessages))
		return
	}

	userService, err := h.userService.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusBadGateway, helper.ApiResponse("Login failed", http.StatusBadGateway, "FAILED", err.Error()))
		return
	}
	
	c.JSON(http.StatusOK, helper.ApiResponse("Login successfully", http.StatusOK, "SUCCESS", user.UserFormatter(userService, "tokentoken")))
}

func (h *UserHandler) CheckEmailAvaibility(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errMessage := gin.H{
			"errors": errors,
		}

		c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Check email failed", http.StatusUnprocessableEntity, "FAILED", errMessage))
		return
	}

	isEmailAvaibility, err := h.userService.CheckEmailInput(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ApiResponse("Internal Server Error", http.StatusInternalServerError, "FAILED", nil))
		return
	}

	message := "Email is not available"
	if isEmailAvaibility {
		message = "Email is available"
	}

	data := gin.H{
		"is_available": isEmailAvaibility,
	}

	c.JSON(http.StatusOK, helper.ApiResponse(message, http.StatusOK, "SUCCESS", data))
}