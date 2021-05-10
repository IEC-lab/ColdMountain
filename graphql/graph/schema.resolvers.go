package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ColdMountain/graphql/graph/generated"
	"ColdMountain/graphql/graph/model"
	"ColdMountain/pkg"
	"context"
)

func (r *queryResolver) FrameStreams(ctx context.Context) ([]*model.FrameStream, error) {
	return pkg.DiscoverStreams(), nil
}

func (r *queryResolver) IntelligentMsgs(ctx context.Context, timeStampStart *string, timeStampEnd *string, vehicleLp *string, vehicleColor *string, taskID *string) ([]*model.IntelligentMsg, error) {
	return pkg.GetIntelligentMsgs(timeStampStart, timeStampEnd, vehicleLp, vehicleColor, taskID), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
