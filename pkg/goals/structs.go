package goals

type Goal struct {
	Category string `json:"category"`
	Description string `json:"description,omitempty"`
	GoalId string `json:"gid,omitempty"`
	Id string `json:"-"`
	IsActive bool `json:"isActive"`
	IsSystem bool `json:"isSystem,omitempty"`
	Name string `json:"name"`
}