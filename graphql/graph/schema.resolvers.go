package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ColdMountain/graphql/graph/generated"
	"ColdMountain/pkg"
	"context"
)

func (r *queryResolver) FrameStreams(ctx context.Context) ([]string, error) {
	return pkg.DiscoverStreams(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
