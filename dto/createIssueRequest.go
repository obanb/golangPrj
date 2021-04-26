package dto

type CreateIssueRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AccountId   string `json:"account_id"`
}
