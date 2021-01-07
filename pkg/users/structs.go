package users

type User struct {
	CreationDate int64 `json:"creation_date,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastLogin int64 `json:"last_login,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Password int64 `json:"password,omitempty"`
	Reports  string `json:"reports,omitempty"` // this is a list of report IDs
	Username string `json:"username,omitempty"`
}