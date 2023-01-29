package graph

import "Module.com/GraphQL-Server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Login               *model.Session
	RSAPublicKeyBackend *model.RSAPublicKey
	ValidateToken       *model.DataToken
}
