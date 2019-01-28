package models

type Summary struct {
	Model

	Image string `json:"image"`
}

func (Summary) TableName() string {
	return "summary_summary"
}
