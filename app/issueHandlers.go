package app

import (
	"awesomeProject/dto"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IssueHandlers struct {
	service service.IssueService
}

func (h *IssueHandlers) getAllIssues(c *gin.Context) {
	customers, err := h.service.GetAllIssues()

	if err != nil {
		c.JSON(err.Code, err.AsMessage())

	} else {
		c.JSON(http.StatusOK, customers)
	}
}

func (h *IssueHandlers) CreateIssue(c *gin.Context) {
	var request dto.CreateIssueRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		issue, appError := h.service.CreateIssue(request)
		if appError != nil {
			c.JSON(appError.Code, appError.AsMessage())
		} else {
			c.JSON(http.StatusCreated, issue)
		}
	}
}

func (h *IssueHandlers) CreateIssueMongo(c *gin.Context) {
	var request dto.CreateIssueRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		issue, appError := h.service.CreateIssue(request)
		if appError != nil {
			c.JSON(appError.Code, appError.AsMessage())
		} else {
			c.JSON(http.StatusCreated, issue)
		}
	}
}


func (h *IssueHandlers) CreateIssues(c *gin.Context) {
	var request dto.CreateIssuesRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		response, appError := h.service.CreateIssues(request)
		if appError != nil {
			c.JSON(appError.Code, appError.AsMessage())
		} else {
			c.JSON(http.StatusCreated, response)
		}
	}
}
