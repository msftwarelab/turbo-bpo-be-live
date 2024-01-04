package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
)

func UpdatePipelineNeighborhood(ctx context.Context, pipelineId string, input models.UpdatePipelineNeighborhoodInput) (bool, error) {
	piplineNeighborhoodRaw, err := datastore.GetPipelineNeighborhood(ctx, pipelineId)
	if err != nil || piplineNeighborhoodRaw == nil {
		//add empty set on PipelineNeighborhood DB collection
		_, err = datastore.SavePipelineNeighborhood(ctx, pipelineId)
		if err != nil {
			return false, err
		}
	}

	return datastore.UpdatePiplineNeighborhood(ctx, pipelineId, input)
}

func PipelineNeighborhood(ctx context.Context, pipelineId string) (*models.PipelineNeighborhood, error) {
	return datastore.GetPipelineNeighborhood(ctx, pipelineId)
}
