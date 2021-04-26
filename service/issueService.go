package service

import (
	"awesomeProject/domain"
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"time"
)

type IssueService interface {
	GetAllIssues() (*[]domain.Issue, *errs.AppError)
	CreateIssue(request dto.CreateIssueRequest) (*dto.CreateIssueResponse, *errs.AppError)
	CreateIssueMongo(request dto.CreateIssueRequest) (*domain.IssueMongo, *errs.AppError)
	CreateIssues(request dto.CreateIssuesRequest) (*[]dto.MayFailWithIdResponse, *errs.AppError)
}

type DefaultIssueService struct {
	repo domain.IssueRepository
}

func (s DefaultIssueService) GetAllIssues() (*[]domain.Issue, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultIssueService) CreateIssue(req dto.CreateIssueRequest) (*dto.CreateIssueResponse, *errs.AppError) {
	issue := domain.Issue{
		IssueId:     "",
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		AccountId:   "1234567891",
		Status:      1,
	}
	newIssue, err := s.repo.Save(issue)

	if err != nil {
		return nil, err
	}
	response := newIssue.ToCreateIssueResponseDto()

	return &response, nil
}

func (s DefaultIssueService) CreateIssueMongo(req dto.CreateIssueRequest) (*domain.IssueMongo, *errs.AppError) {
	issue := domain.Issue{
		IssueId:     "",
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		AccountId:   "1234567891",
		Status:      1,
	}
	newIssue, err := s.repo.SaveMongo(issue)

	if err != nil {
		return nil, err
	}

	return newIssue, nil
}

func (s DefaultIssueService) CreateIssues(req dto.CreateIssuesRequest) (*[]dto.MayFailWithIdResponse, *errs.AppError) {
	l := len(req.Issues)
	responses := make([]dto.MayFailWithIdResponse, 0)

	c := make(chan dto.MayFailWithIdResponse, l)

	for index, issue := range req.Issues {
		go func(issue dto.CreateIssueRequest) {
			i, err := s.repo.Save(domain.Issue{
				IssueId:     string(index),
				Name:        issue.Name,
				Description: issue.Description,
				CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				AccountId:   "1234567891",
				Status:      1,
			})
			if err != nil {
				c <- dto.MayFailWithIdResponse{
					Success: false,
					Error: err,
				}
			} else {
				id := i.ToCreateIssueResponseDto().IssueId
				c <-  dto.MayFailWithIdResponse{
					Success: true,
					Id:      &id,
				}
			}
		}(issue)
	}

	for range req.Issues {
		responses = append(responses, <-c)
	}
	return &responses, nil
}

func NewIssueService(repository domain.IssueRepository) DefaultIssueService {
	return DefaultIssueService{repository}
}
