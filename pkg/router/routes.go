package router

import (
	"net/http"
	"github.com/cpretzer/lt-backend/pkg/structs"
	"github.com/cpretzer/lt-backend/pkg/handlers"
	"github.com/cpretzer/lt-backend/pkg/users"
	"github.com/cpretzer/lt-backend/pkg/goals"
)

type Routes []structs.Route

// Array of Route objects, each associated with a unique HTTP endpoint
var routes = Routes{
	// Handler listening for GET at "/" URI
	// Returns specified string
	structs.Route{
		Name: "Home",
		Method: http.MethodGet,
		Pattern: "/",
		Function: handlers.HandleHome,
	},
	// TODO: Think about separating these out into a separate array of
	// routes that can be appended
	structs.Route{
		Name: "GetUser",
		Method: http.MethodGet,
		Pattern: "/users",
		Function: users.HandleGetUser,
	},
	structs.Route{
		Name: "AddUser",
		Method: http.MethodPost,
		Pattern: "/users/add",
		Function: users.HandleAddUser,
	},
	structs.Route{
		Name: "UpdateUser",
		Method: http.MethodPatch,
		Pattern: "/users/update",
		Function: users.HandleUpdateUser,
	},
	structs.Route{
		Name: "DeleteUser",
		Method: http.MethodDelete,
		Pattern: "/users/delete",
		Function: users.HandleDeleteUser,
	},
	// TODO: Think about separating these out into a separate array of
	// routes that can be appended
	structs.Route{
		Name: "GetGoal",
		Method: http.MethodGet,
		Pattern: "/goals",
		Function: goals.HandleGetGoal,
	},
	structs.Route{
		Name: "CreateGoal",
		Method: http.MethodPost,
		Pattern: "/goals/add",
		Function: goals.HandleCreateGoal,
	},
	structs.Route{
		Name: "UpdateGoal",
		Method: http.MethodPatch,
		Pattern: "/goals/update",
		Function: goals.HandleUpdateGoal,
	},
	structs.Route{
		Name: "DeleteUser",
		Method: http.MethodDelete,
		Pattern: "/goals/delete",
		Function: goals.HandleDeleteGoal,
	},
}
