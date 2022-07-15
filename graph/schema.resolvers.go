package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/NicholasR77/starfield/graph/generated"
	"github.com/NicholasR77/starfield/graph/model"
)

// CreateShip is the resolver for the createShip field.
func (r *mutationResolver) CreateShip(ctx context.Context, input model.NewShip) (*model.Ship, error) {
	panic(fmt.Errorf("not implemented"))
}

// Ships is the resolver for the ships field.
func (r *queryResolver) Ships(ctx context.Context) ([]*model.Ship, error) {
	panic(fmt.Errorf("not implemented"))
}

// Ship is the resolver for the ship field.
func (r *queryResolver) Ship(ctx context.Context, id string) (*model.Ship, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
