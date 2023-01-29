package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"log"

	"Module.com/GraphQL-Server/graph/generated"
	encrypt "Module.com/GraphQL-Server/graph/grpc-encryption"
	jwtToken "Module.com/GraphQL-Server/graph/grpc-token"
	"Module.com/GraphQL-Server/graph/model"
)

// Mutation.
func (r *mutationResolver) Login(ctx context.Context, input model.DataChiper) (*model.Session, error) {
	token, err := jwtToken.CreateToken(&jwtToken.DataTokenRequest{PublicKey: input.PublicKey, Data: input.DataChiper})
	if err != nil {
		return &model.Session{
			Token:  "",
			Result: "session auth failed",
		}, nil
	}
	if token == "session auth failed" {
		return &model.Session{
			Token:  "",
			Result: "session auth failed",
		}, nil
	}
	return &model.Session{
		Token:  token,
		Result: "session auth successful",
	}, nil
}

// Mutation.
func (r *mutationResolver) ValidateToken(ctx context.Context, input model.DataChiper) (*model.DataToken, error) {
	ValidToken, err := jwtToken.ValidateToken(&jwtToken.DataTokenRequest{PublicKey: input.PublicKey, Data: input.DataChiper})
	if err != nil {
		return &model.DataToken{
			Matricule: "",
			Result:    "valid token failed",
		}, nil
	}

	return &model.DataToken{
		Matricule: ValidToken.Matricule,
		Result:    ValidToken.Result,
	}, nil
}

// RSAPublicKeyBackend is the resolver for the rSAPublicKeyBackend field.
func (r *queryResolver) RSAPublicKeyBackend(ctx context.Context) (*model.RSAPublicKey, error) {
	KeysRSA, err := encrypt.GetKeysRSA(&encrypt.Empty{})
	if err != nil {
		log.Fatalf("error generating keys RSA: %v", err)
	}
	return &model.RSAPublicKey{
		PublicKey: KeysRSA.PublicKey,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }