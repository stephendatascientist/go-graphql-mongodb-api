package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stephendatascientist/go-graphql-mongodb-api/graph/generated"
	"github.com/stephendatascientist/go-graphql-mongodb-api/graph/model"
)

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*model.Customer, error) {
	return r.CustomerRepository.CreateCustomer(input), nil
}

// Customer is the resolver for the customer field.
func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	return r.GetCustomer(id), nil
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return r.CustomerRepository.GetCustomers(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
