package goals

type Report struct {
	Date     string  `json:"date,omitempty"` // TODO: change this to a date type or write a converter
	Id       string  `json:"-"`
	IsActive bool    `json:"isActive"`
	Name     string  `json:"name,omitempty"`
	Notes    string  `json:"notes,omitempty"`
	Owner    string  `json:"owner"`
	Rating   float64 `json:"rating,omitempty"`
	ReportId string  `json:"gid,omitempty"`
	Status   string  `json:"status,omitempty"`
}
