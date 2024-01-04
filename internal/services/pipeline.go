package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func SavePipeline(ctx context.Context, userId string, input models.PipelineInput) (string, error) {
	// Validate if user has overdue billing
	list, err := datastore.GetPipelineByOrderNumber(ctx, input.OrderNumber)
	if err != nil {
		return "", err
	}
	if len(list) > 0 {
		return "", errs.OrderNumberAlreadyExist
	}

	if input.AuthorID != nil {
		userId = *input.AuthorID
	}

	hasDueBilling, err := datastore.SearchBillings(ctx, 0, 1, nil, nil, &userId, pointers.Bool(true), nil)
	if len(hasDueBilling) > 0 {
		return "", errs.AccountOverDue
	}
	if err != nil {
		return "", err
	}

	pipelineId, err := datastore.AddPipeline(ctx, userId, input)
	if err != nil {
		return "", err
	}
	pipelineTotals, err := datastore.GetPipelinesCount(ctx, datastore.EmptyFilter())
	if err != nil {
		return "", err
	}
	err = datastore.AddUserOrder(ctx, userId, pipelineTotals)
	if err != nil {
		return "", err
	}

	//add empty set on PipelineNeighborhood DB collection
	_, err = datastore.SavePipelineNeighborhood(ctx, pipelineId)
	if err != nil {
		return "", err
	}

	return pipelineId, nil
}

func UpdatePipeline(ctx context.Context, id string, input models.UpdatePipelineInput, myFullname, myID string) (bool, error) {
	isSuccess, err := datastore.UpdatePipeline(ctx, id, input, myFullname, myID, nil)
	if err != nil {
		return false, err
	}
	if isSuccess {

		//save pipeline note
		message := ""
		if input.IsHold != nil {
			//me hold
			if *input.IsHold {
				message = fmt.Sprintf("Pipeline Status update :Hold %s", strings.ObjectTOString(input.HoldRemarks))
			} else {
				message = fmt.Sprintf("Pipeline Status update :UnHold %s", strings.ObjectTOString(input.UnHoldRemarks))
			}

			noteData := models.SavePipelineNoteInput{
				OrderNotes: strings.ToObject(message),
			}

			datastore.AddPipelineNote(ctx, id, noteData, myFullname)

		} else if strings.ObjectTOString(input.Status) == constants.PipelineStatusCancelled {
			message = fmt.Sprintf("%s", strings.ObjectTOString(input.CancelRemarks))

			noteData := models.SavePipelineNoteInput{
				OrderNotes: strings.ToObject(message),
			}

			datastore.AddPipelineNote(ctx, id, noteData, myFullname)
		}

	}

	return isSuccess, err
}

func DeletePipeline(ctx context.Context, id string) (bool, error) {
	panic("Todo")
	//todo
}

func Pipeline(ctx context.Context, id string) (*models.Pipeline, error) {
	return datastore.GetPipelineById(ctx, id)
}

func AllPipeline(ctx context.Context, userId string, roles []*string, filter *models.PipelineFilterInput) (*models.PipelineResult, error) {
	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {
		filter = &models.PipelineFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	roleFilter := datastore.SearchFilterDefault(filter)
	if contains(roles, constants.UserRoleClient) {
		roleFilter = datastore.SearchFilterForClient(userId, filter)
	}
	if contains(roles, constants.UserRoleContractor) {
		roleFilter = datastore.SearchFilterForContractor(userId, filter)
	}
	if contains(roles, constants.UserRoleAdmin) {
		roleFilter = datastore.SearchFilterForAdmin(filter)
	}
	rawPipelines, err := datastore.SearchPipelines(ctx, roleFilter, *filter.Offset, *filter.Limit)
	if err != nil {
		log.Debug("Failed to retrieve pipeline: %v", err)
		return nil, err
	}

	count, err := datastore.GetPipelinesCount(ctx, roleFilter)
	if err != nil {
		log.Debug("Failed to retrieve count of pipeline: %v", err)
		return nil, err
	}

	var pipelineIds []string
	var userIds []string

	for _, pipeline := range rawPipelines {
		pipelineIds = append(pipelineIds, pipeline.ID.Hex())
		if pipeline.UserId != nil {
			userIds = append(userIds, *pipeline.UserId)
		}
	}

	pipelines := make([]*models.Pipeline, 0)
	for key, u := range rawPipelines {
		pipelines = append(pipelines, u.ToModels())
		pipelines[key].AuthorID = u.UserId
	}

	if len(pipelineIds) > 0 {

		pipelineQualityControlsTotals, err := datastore.GetPipelineQualityControlsCountGroupByPipelineId(ctx, pipelineIds)
		if err != nil {
			return nil, err
		}

		pipelineNotesTotals, err := datastore.GetPipelineNotesCountGroupByPipelineId(ctx, pipelineIds)
		if err != nil {
			return nil, err
		}

		pipelineAuthors, err := datastore.GetUsers(ctx, datastore.FilterByIds(userIds))
		if err != nil {
			return nil, err
		}

		for key, pipeline := range pipelines {
			var pipelineQualityControlTotal int
			var pipelineNoteTotal int
			var authorName string

			for _, rawQualityControl := range pipelineQualityControlsTotals {
				if rawQualityControl.PipelineId == pipeline.ID {
					pipelineQualityControlTotal = rawQualityControl.Count
				}
			}
			for _, rawNote := range pipelineNotesTotals {
				if rawNote.PipelineId == pipeline.ID {
					pipelineNoteTotal = rawNote.Count
				}
			}
			for _, rawAuthor := range pipelineAuthors {
				if pipeline.AuthorID != nil {
					if rawAuthor.ID.Hex() == *pipeline.AuthorID {
						authorName = Name(rawAuthor.FirstName, rawAuthor.LastName)
					}
				}
			}

			isQc, err := datastore.CheckQualityControlStatus(ctx, pipeline.ID)
			if err != nil {
				return nil, err
			}

			pipelines[key].IsQc = &isQc
			pipelines[key].PipelineQualityControlTotal = &pipelineQualityControlTotal
			pipelines[key].PipelineNoteTotal = &pipelineNoteTotal
			pipelines[key].AuthorName = &authorName

		}
	}

	toInt := int(*count)
	return &models.PipelineResult{
		TotalCount: &toInt,
		Results:    pipelines,
	}, nil

}

//For API Endpoints
func SearchPipelineByOrderNumberOrORderAddress(ctx context.Context, orderNumber, orderAddress string, limit int) ([]*datastore.Pipeline, error) {

	return datastore.GetPipelineSearchbyOrderNumberOrAddress(ctx, orderNumber, orderAddress, 0, limit)
}

func Name(firstname string, lastname string) string {
	return fmt.Sprintf("%s %s", firstname, lastname)
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func BatchUpdatePipelineNotes(ctx context.Context) error {
	log.Info("Running on service!!")
	AllEmptyPipelineNotes, err := datastore.GetAllEmptyPipelineNotes(ctx)
	if err != nil {
		return err
	}
	for _, i := range AllEmptyPipelineNotes {
		log.Info(PrettyPrint(i.CreatedDateTime))
		isSuccess, err := datastore.DeletePipelineNote(ctx, i.ID)
		if err != nil {
			return err
		}
		if isSuccess != true {
			return errors.New("Failed to update status")
		} else {
			log.Info("Success")
		}
	}
	return nil

}
