package dto

type CreateDbReportRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Query       string `json:"query"`
	Source      string `json:"source"`
}
