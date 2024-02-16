package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"tutorials/gqlgen-users/databaseConnector"
	"tutorials/gqlgen-users/graph/model"
)

// SaveCustomer is the resolver for the saveCustomer field.
func (r *mutationResolver) SaveCustomer(ctx context.Context, input model.NewCustomer) (bool, error) {
	customer := &model.Customer{
		CustomerID: input.CustomerID,
		ZipCode:    input.ZipCode,
	}
	r.customer = append(r.customer, customer)
	return true, nil
}

// GetCustomer is the resolver for the getCustomer field.
func (r *queryResolver) GetCustomer(ctx context.Context, customerID string) (*model.Customer, error) {
	customer, err := databaseConnector.GetUserByID(4)
	if err != nil {
		// handle error
		return nil, err
	}
	c := model.Customer{
		CustomerID: "123",
		ZipCode:    customer.Email,
	}

	return &c, nil
	// return r.customer[0], nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
