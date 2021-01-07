package testdata

import (
	"github.com/cpretzer/lt-backend/pkg/users"
)

func GenerateTestUser() *User {
	return &User{
		FirstName: "Charles",
		LastName: "Pretzer",
		EmailAddress: "c@chabrina.com",
		Username: "cpretzer",
	}
}

