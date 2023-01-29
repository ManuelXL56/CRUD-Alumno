package graph

import "Module.com/GraphQL-Server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SearchUser   *model.DataUser
	RegisterUser *model.Result
	UpdateUser   *model.Result
	DeleteUser   *model.Result
}
