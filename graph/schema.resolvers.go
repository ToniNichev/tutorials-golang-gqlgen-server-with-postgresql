package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"strconv"
	"tutorials/gqlgen-users/databaseConnector"
	"tutorials/gqlgen-users/graph/model"
)

// SaveCustomer is the resolver for the saveCustomer field.
func (r *mutationResolver) SaveCustomer(ctx context.Context, input model.NewCustomer) (bool, error) {
	databaseConnector.CreateUser(input.Username, input.Email, input.Age, input.MetaData)
	return true, nil
}

// CreateDb is the resolver for the createDB field.
func (r *mutationResolver) CreateDb(ctx context.Context) (bool, error) {
	err := databaseConnector.CreateDB()

	if err != nil {
		// handle error
		return false, err
	}
	return true, nil
}

// AddBookmark is the resolver for the addBookmark field.
func (r *mutationResolver) AddBookmark(ctx context.Context, userID string, name string, group string, metaData *string) (bool, error) {
	databaseConnector.AddBookmark(userID, name, group, *metaData)
	return true, nil
}

// GetCustomer is the resolver for the getCustomer field.
func (r *queryResolver) GetCustomer(ctx context.Context, customerID string) (*model.Customer, error) {
	cid, _ := strconv.Atoi(customerID)
	var customer *databaseConnector.User
	var err error
	customer, err = databaseConnector.GetUserByID(uint(cid))

	userId := strconv.FormatUint(uint64(customer.ID), 10)

	if err != nil {
		// handle error
		return nil, err
	}

	// get the underlying byte slice.
	jsonbText, _ := customer.MetaData.Value()
	// Convert byte slice to string
	jsonString := string(jsonbText.([]byte))

	bookmarks, _ := databaseConnector.GetBookmarks(userId)

	var result []*model.Bookmark
	for _, v := range *bookmarks {

		// get the underlying byte slice.
		jsonbText, _ := v.MetaData.Value()
		// Convert byte slice to string
		jsonString := string(jsonbText.([]byte))

		b := model.Bookmark{
			BookmarkID: strconv.FormatUint(uint64(v.ID), 10),
			UserID:     v.UserId,
			Name:       v.Name,
			Group:      v.Group,
			MetaData:   jsonString,
		}
		result = append(result, &b)
	}

	nextPage := "123"
	previousPage := "1122"
	pageInfo := model.PageInfo{
		NextPage:     &nextPage,
		PreviousPage: &previousPage,
	}

	bookmarksPaginated := model.BookmarksPaginated{
		Data:     result,
		PageInfo: &pageInfo,
	}

	// map returned customer structure from the DB into the model
	c := model.Customer{
		CustomerID: userId,
		Username:   customer.Username,
		Email:      customer.Email,
		Age:        customer.Age,
		MetaData:   jsonString,
		Bookmarks:  &bookmarksPaginated,
	}

	return &c, nil
}

// GetCustomerByMetaData is the resolver for the getCustomerByMetaData field.
func (r *queryResolver) GetCustomerByMetaData(ctx context.Context, metaData string) ([]*model.Customer, error) {
	customers, err := databaseConnector.GetUserByMetaData(metaData)

	if err != nil {
		// handle error
		return nil, err
	}

	var result []*model.Customer

	for _, user := range *customers {

		// get the underlying byte slice.
		jsonbText, _ := user.MetaData.Value()
		// Convert byte slice to string
		jsonString := string(jsonbText.([]byte))

		// map returned customer structure from the DB into the model
		c := model.Customer{
			CustomerID: strconv.FormatUint(uint64(user.ID), 10),
			Username:   user.Username,
			Email:      user.Email,
			Age:        user.Age,
			MetaData:   jsonString,
		}
		result = append(result, &c)
	}

	return result, nil
}

func (r *Resolver) Bookmarks(ctx context.Context, metaData string) ([]*model.BookmarksPaginated, error) {
	// var result []*model.BookmarksPaginated

	//result

	return nil, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
