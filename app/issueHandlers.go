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

// getAllIssues godoc
// @Summary Retrieves list of Issues
// @Produce json
// @Success 200 {object} []domain.Issue
// @Failure 500 {object} errs.AppError
// @Router /issues [get]
func (h *IssueHandlers) getAllIssues(c *gin.Context) {
	customers, err := h.service.GetAllIssues()

	if err != nil {
		c.JSON(err.Code, err.AsMessage())

	} else {
		c.JSON(http.StatusOK, customers)
	}
}

// CreateIssue godoc
// @Summary Create issue
// @Produce json
// @Accept  json
// @Param   body body dto.CreateIssueRequest true  "body"
// @Success 200 {object} domain.Issue
// @Failure 500 {object} errs.AppError
// @Router /issue [post]
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
