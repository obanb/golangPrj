package app

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

//SignUpUser godoc
//@Summary SignUpUser
//@Produce json
//@Accept  json
//@Param   body body dto.RegisterUserRequest true  "body"
//@Success 200 {object} domain.User
//@Failure 500 {object} errs.AppError
//@Router /issue [post]
func (h *UserHandler) SignUpUser(c *gin.Context) {
	req, err := dto.RawBodyUnmarshal(&c.Request.Body)
	validationErrors := req.Validate()
	if validationErrors != nil {
		c.JSON(500, validationErrors)
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Message)
	} else {
		issue, appError := h.service.SignUpUser(*req)
		if appError != nil {
			c.JSON(appError.Code, appError.AsMessage())
		} else {
			c.JSON(http.StatusCreated, issue)
		}
	}
}
