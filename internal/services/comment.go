package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
)

func SaveComment(ctx context.Context, id string, input models.CommentInput) (string, error) {
	return datastore.AddComment(ctx, id, input)
}

func UpdateComment(ctx context.Context, id string, value string) (bool, error) {
	return datastore.UpdateComment(ctx, id, value)
}

func DeleteComment(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteComment(ctx, id)
}

func AllComment(ctx context.Context, userId string, filterUserId *string) ([]*models.Comment, error) {
	if filterUserId != nil {
		userId = *filterUserId
	}
	return datastore.GetComments(ctx, userId)
}

func SetCommentDefault(ctx context.Context, userId string) (bool, error) {

	err := datastore.DeleteComments(ctx, userId)
	if err != nil {
		return false, err
	}
	err = initComments(ctx, userId)
	if err != nil {
		return false, err
	}

	return true, nil
}
