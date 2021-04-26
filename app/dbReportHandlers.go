package app

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DbReportHandlers struct {
	service service.DbReportService
}

func (h DbReportHandlers) CreateDbReport(c *gin.Context) {
	var request *dto.CreateDbReportRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		dbReport, appError := h.service.CreateDbReport(request)
		if appError != nil {
			c.JSON(appError.Code, appError.AsMessage())
		} else {
			c.JSON(http.StatusCreated, dbReport)
		}
	}
}
