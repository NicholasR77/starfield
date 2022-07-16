package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/NicholasR77/starfield/database"
	"github.com/NicholasR77/starfield/graph/generated"
	"github.com/NicholasR77/starfield/graph/model"
)

// CreateShip is the resolver for the createShip field.
func (r *mutationResolver) CreateShip(ctx context.Context, input *model.NewShip) (*model.Ship, error) {
	return db.Save(input), nil
}

// Ship is the resolver for the ship field.
func (r *queryResolver) Ship(ctx context.Context, id string) (*model.Ship, error) {
	return db.FindByID(id), nil
}

// Ships is the resolver for the ships field.
func (r *queryResolver) Ships(ctx context.Context) ([]*model.Ship, error) {
	return db.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
