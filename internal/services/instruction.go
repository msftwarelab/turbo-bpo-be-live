package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SaveInstruction(ctx context.Context, input models.SaveInstructionInput) (string, error) {

	fileUrl, err := awsS3.S3Uploader(input.File)
	if err != nil {
		return "", err
	}

	return datastore.SaveInstruction(ctx, input, fileUrl)
}

func DeleteInstruction(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteInstruction(ctx, id)
}

func AllInstruction(ctx context.Context, filter *models.InstructionFilterInput) (*models.InstructionResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.InstructionFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawInstructions, err := datastore.SearchInstructions(ctx, *filter.Offset, *filter.Limit, filter.Tag, nil, nil)
	if err != nil {
		log.Debug("Failed to retrieve instruction: %v", err)
		return nil, err
	}

	count, err := datastore.GetInstructionsCount(ctx, filter.Tag)
	if err != nil {
		log.Debug("Failed to retrieve count of instruction: %v", err)
		return nil, err
	}

	instructions := make([]*models.Instruction, 0)
	for _, u := range rawInstructions {
		instructions = append(instructions, u.ToModels())
	}
	toInt := int(*count)
	return &models.InstructionResult{
		TotalCount: &toInt,
		Results:    instructions,
	}, nil
}
