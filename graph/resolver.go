package graph

import (
	"context"
	"tutorials/gqlgen-users/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	customer []*model.Customer
}

func (r *Resolver) GetCustomerFromDB(ctx context.Context, id string) (*Customer, error) {
	customer := Customer{
		CustomerId: "123",
		ZipCode:    "07096",
	}
	return &customer, nil
}

type Customer struct {
	CustomerId string
	ZipCode    string
}
