package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
)

// func SavePipelineRepair(ctx context.Context, pipelineId string, input models.PipelineRepairInput, userId string, userName string) (string, error) {
// 	fileUrl, err := s3Uploader(input.Doc)
// 	if err != nil {
// 		return "", err
// 	}

// 	return datastore.AddPipelineRepair(ctx, pipelineId, input, userId, *fileUrl, userName)
// }

func UpdatePipelineRepair(ctx context.Context, pipelineId string, input models.UpdatePipelineRepairInput, updatedBy string) (bool, error) {

	pipelineRepairRaw, _ := datastore.GetPipelineRepair(ctx, pipelineId)

	if pipelineRepairRaw != nil {
		return datastore.UpdatePipelineRepair(ctx, pipelineId, input, updatedBy)
	} else {
		addPipelineRepair := models.SavePipelineNoteInput{
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
		_, err := datastore.AddPipelineRepair(ctx, pipelineId, addPipelineRepair)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errs.UnknownError

}
func GetPipelineRepair(ctx context.Context, pipelineId string) (*models.PipelineRepair, error) {

	return datastore.GetPipelineRepair(ctx, pipelineId)

}
