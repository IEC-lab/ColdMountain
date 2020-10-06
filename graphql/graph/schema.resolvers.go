package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ColdMountain/graphql/graph/generated"
	"ColdMountain/streams"
	"context"
)

func (r *queryResolver) GetFrameStreams(ctx context.Context) ([]string, error) {
	return streams.DiscoverStreams(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
