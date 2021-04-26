package dto

type CreateIssuesRequest struct {
	Issues []CreateIssueRequest `json:"issues"`
}
