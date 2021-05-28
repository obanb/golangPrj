package dto

type CreateDbReportRequestLookup struct {
	LocalKey   string `json:"localKey"`
	ForeignKey string `json:"foreignKey"`
}

type CreateDbReportOrderBy struct {
	Key string `json:"key"`
	Asc bool   `json:"asc"`
}

type CreateDbReportRequest struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Query       string                        `json:"query"`
	Lookup      []CreateDbReportRequestLookup `json:"lookup"`
	Projections []string                      `json:"projections"`
	OrderBy     CreateDbReportOrderBy         `json:"orderBy"`
	limit       int                           `json:"limit"`
	Source      string                        `json:"source"`
}
