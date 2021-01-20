package users

type User struct {
	Active bool `json:"active"`
	CreationDate uint `json:"creationDate,omitempty"`
	DeactivationDate int64 `json:"deactivationDate,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	Id string `json:"-"`
	LastLogin uint `json:"lastLogin,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Password string `json:"password,omitempty"`
	Reports  string `json:"reports,omitempty"` // this is a list of report IDs
	Username string `json:"username,omitempty"`
}