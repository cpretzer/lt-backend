package goals

type Goal struct {
	Category string `json:"category"`
	Description string `json:"string,omitempty"`
	Id string `json:"id"`
	IsActive bool `json:"isActive"`
	IsSystem string `json:"isSystem,omitempty"`
	Name string `json:"name"`
}