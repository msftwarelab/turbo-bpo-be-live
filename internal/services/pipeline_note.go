package services

import (
	"context"
	"errors"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SavePipelineNote(ctx context.Context, pipelineId string, input models.SavePipelineNoteInput, userId string) (string, error) {
	if *input.OrderNotes == "" {
		return "", errors.New("Empty Order Notes")
	}
	id, err := datastore.AddPipelineNote(ctx, pipelineId, input, userId)
	if err != nil {
		return "", err
	}
	//Todo, add repair
	pipelineRepairRaw, _ := datastore.GetPipelineRepair(ctx, pipelineId)

	if pipelineRepairRaw == nil {
		datastore.AddPipelineRepair(ctx, pipelineId, input)
	} else {
		updatePipelineRepair := models.UpdatePipelineRepairInput{
			ExteriorRepairDescription1:  input.ExteriorRepairDescription1,
			ExteriorRepairPrice1:        input.ExteriorRepairPrice1,
			ExteriorRepairDescription2:  input.ExteriorRepairDescription2,
			ExteriorRepairPrice2:        input.ExteriorRepairPrice2,
			ExteriorRepairDescription3:  input.ExteriorRepairDescription3,
			ExteriorRepairPrice3:        input.ExteriorRepairPrice3,
			ExteriorRepairDescription4:  input.ExteriorRepairDescription4,
			ExteriorRepairPrice4:        input.ExteriorRepairPrice4,
			ExteriorRepairDescription5:  input.ExteriorRepairDescription5,
			ExteriorRepairPrice5:        input.ExteriorRepairPrice5,
			ExteriorRepairDescription6:  input.ExteriorRepairDescription6,
			ExteriorRepairPrice6:        input.ExteriorRepairPrice6,
			ExteriorRepairDescription7:  input.ExteriorRepairDescription7,
			ExteriorRepairPrice7:        input.ExteriorRepairPrice7,
			ExteriorRepairDescription8:  input.ExteriorRepairDescription8,
			ExteriorRepairPrice8:        input.ExteriorRepairPrice8,
			ExteriorRepairDescription9:  input.ExteriorRepairDescription9,
			ExteriorRepairPrice9:        input.ExteriorRepairPrice9,
			ExteriorRepairDescription10: input.ExteriorRepairDescription10,
			ExteriorRepairPrice10:       input.ExteriorRepairPrice10,
			ExteriorRepairPriceTotal:    input.ExteriorRepairPriceTotal,
			InteriorRepairDescription1:  input.InteriorRepairDescription1,
			InteriorRepairPrice1:        input.InteriorRepairPrice1,
			InteriorRepairDescription2:  input.InteriorRepairDescription2,
			InteriorRepairPrice2:        input.InteriorRepairPrice2,
			InteriorRepairDescription3:  input.InteriorRepairDescription3,
			InteriorRepairPrice3:        input.InteriorRepairPrice3,
			InteriorRepairDescription4:  input.InteriorRepairDescription4,
			InteriorRepairPrice4:        input.InteriorRepairPrice4,
			InteriorRepairDescription5:  input.InteriorRepairDescription5,
			InteriorRepairPrice5:        input.InteriorRepairPrice5,
			InteriorRepairDescription6:  input.InteriorRepairDescription6,
			InteriorRepairPrice6:        input.InteriorRepairPrice6,
			InteriorRepairDescription7:  input.InteriorRepairDescription7,
			InteriorRepairPrice7:        input.InteriorRepairPrice7,
			InteriorRepairDescription8:  input.InteriorRepairDescription8,
			InteriorRepairPrice8:        input.InteriorRepairPrice8,
			InteriorRepairDescription9:  input.InteriorRepairDescription9,
			InteriorRepairPrice9:        input.InteriorRepairPrice9,
			InteriorRepairDescription10: input.InteriorRepairDescription10,
			InteriorRepairPrice10:       input.InteriorRepairPrice10,
			InteriorRepairPriceTotal:    input.InteriorRepairPriceTotal,
		}
		datastore.UpdatePipelineRepair(ctx, pipelineId, updatePipelineRepair, userId)
	}

	return id, nil
}

func AllPipelineNote(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineNoteResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {
		filter = &models.FilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawPipelineNotes, err := datastore.SearchPipelineNotes(ctx, pipelineId, *filter.Offset, *filter.Limit)
	if err != nil {
		log.Debug("Failed to retrieve pipelineNote: %v", err)
		return nil, err
	}

	count, err := datastore.GetPipelineNotesCount(ctx, pipelineId)
	if err != nil {
		log.Debug("Failed to retrieve count of pipelineNote: %v", err)
		return nil, err
	}

	pipelineQualityNotes := make([]*models.PipelineNote, 0)
	for _, u := range rawPipelineNotes {
		pipelineQualityNotes = append(pipelineQualityNotes, u.ToModels())
	}
	toInt := int(*count)
	return &models.PipelineNoteResult{
		TotalCount: &toInt,
		Results:    pipelineQualityNotes,
	}, nil
}
