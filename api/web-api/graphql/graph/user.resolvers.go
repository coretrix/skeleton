package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"coretrix/skeleton/api/web-api/graphql/graph/model/custom"
)

func (r *queryResolver) Me(ctx context.Context) (*custom.User, error) {
	return &custom.User{ID: "1234", Username: "Me", Age: 69}, nil
}
