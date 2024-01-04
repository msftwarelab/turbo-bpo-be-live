package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
)

func UpdatePipelineState(ctx context.Context, input models.UpdatePipelineStateInput) (bool, error) {
	return datastore.UpdatePiplineState(ctx, input)
}

func PipelineState(ctx context.Context) (*models.PipelineState, error) {
	return datastore.GetPipelineState(ctx)
}
