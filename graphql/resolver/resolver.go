package resolver

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/generated"
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/internal/middleware"
	"github.com/lonmarsDev/bpo-golang-grahpql/internal/services"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*models.Token, error) {
	ctxIPAddress := middleware.GetUserIPAddress(ctx)
	return services.Login(ctx, email, password, ctxIPAddress)

}

func (r *mutationResolver) RegisterUser(ctx context.Context, input models.RegisterInput) (string, error) {
	return services.RegisterUser(ctx, input)
}

func (r *mutationResolver) ForgetPassword(ctx context.Context, email string) (bool, error) {
	return services.ForgetPassword(ctx, email)
}

func (r *mutationResolver) ResetPassword(ctx context.Context, token string, newPassword string) (bool, error) {
	return services.ResetPassword(ctx, token, newPassword)
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, input models.ProfileInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateProfile(ctx, ctxUser.UserId, input)
}

func (r *mutationResolver) SaveProfileDoc(ctx context.Context, input models.ProfileDocInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.UploadProfileDoc(ctx, ctxUser.UserId, input)

}

func (r *mutationResolver) DeleteProfileDoc(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteProfileDoc(ctx, id)
}

func (r *mutationResolver) SaveAccount(ctx context.Context, input models.AccountInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveAccount(ctx, ctxUser.UserId, input, ctxUser.Name)
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, id string, input models.AccountInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateAccount(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteAccount(ctx, id)
}

func (r *mutationResolver) UpdateAdjustment(ctx context.Context, id string, value float64) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateAdjustment(ctx, id, value)
}

func (r *mutationResolver) SetAdjustmentDefault(ctx context.Context) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.SetAdjustmentDefault(ctx, ctxUser.UserId)
}

func (r *mutationResolver) SaveComment(ctx context.Context, input models.CommentInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveComment(ctx, ctxUser.UserId, input)
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, value string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateComment(ctx, id, value)
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteComment(ctx, id)
}

func (r *mutationResolver) SetCommentDefault(ctx context.Context) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.SetCommentDefault(ctx, ctxUser.UserId)
}

func (r *mutationResolver) UpdateDefault(ctx context.Context, input models.DefaultInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateDefault(ctx, ctxUser.UserId, ctxUser.Role, input)
}

func (r *mutationResolver) SavePipeline(ctx context.Context, input models.PipelineInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePipeline(ctx, ctxUser.UserId, input)
}

func (r *mutationResolver) UpdatePipeline(ctx context.Context, id string, input models.UpdatePipelineInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdatePipeline(ctx, id, input, ctxUser.Name, ctxUser.UserId)
}

func (r *mutationResolver) SavePipelineQualityControl(ctx context.Context, pipelineId string, orderNotes string) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePipelineQualityControl(ctx, pipelineId, orderNotes, ctxUser.Name)
}

func (r *mutationResolver) SavePipelineDoc(ctx context.Context, pipelineId string, input models.PipelineDocInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePipelineDoc(ctx, pipelineId, input, ctxUser.Name)
}

func (r *mutationResolver) SavePipelinePhoto(ctx context.Context, pipelineId string, input models.PipelinePhotoInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePipelinePhoto(ctx, pipelineId, input, ctxUser.Id, ctxUser.Name)
}

func (r *mutationResolver) SubmitPipelinePhoto(ctx context.Context, id string, isSubmitPipelinePhoto bool) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.SubmitPipelinePhoto(ctx, id, ctxUser.UserId, ctxUser.Name, isSubmitPipelinePhoto)
}

func (r *mutationResolver) DeletePipelinePhoto(ctx context.Context, pipelinePhotoID string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeletePipelinePhoto(ctx, pipelinePhotoID, ctxUser.UserId, ctxUser.Name)
}

func (r *mutationResolver) DeletePipelineDoc(ctx context.Context, pipelineId string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeletePipelineDoc(ctx, pipelineId, ctxUser.UserId)
}

func (r *mutationResolver) SavePipelineNote(ctx context.Context, pipelineId string, input models.SavePipelineNoteInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePipelineNote(ctx, pipelineId, input, ctxUser.Name)
}

func (r *mutationResolver) SaveCredit(ctx context.Context, input models.SaveCreditInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveCredit(ctx, ctxUser.UserId, input, ctxUser.Name)
}
func (r *mutationResolver) AddCreditLedger(ctx context.Context, input models.AddCreditLedgerInput) (string, error) {
	return services.SaveCreditLedger(ctx, input)
}

func (r *mutationResolver) SaveCompany(ctx context.Context, input models.CompanyInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveCompany(ctx, ctxUser.UserId, input)
}

func (r *mutationResolver) UpdateCompany(ctx context.Context, id string, input models.CompanyInput) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateCompany(ctx, id, input)
}

func (r *mutationResolver) DeleteCompany(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteCompany(ctx, id)
}

func (r *mutationResolver) SaveUser(ctx context.Context, input models.SaveUserInput) (string, error) {
	return services.SaveUser(ctx, input)
}

func (r *mutationResolver) UpdatePipelineState(ctx context.Context, input models.UpdatePipelineStateInput) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdatePipelineState(ctx, input)
}

func (r *mutationResolver) SaveEmailTemplate(ctx context.Context, input models.SaveEmailTemplateInput) (string, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveEmailTemplate(ctx, input)
}

func (r *mutationResolver) UpdateEmailTemplate(ctx context.Context, id string, input models.UpdateEmailTemplateInput) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateEmailTemplate(ctx, id, input)
}

func (r *mutationResolver) DeleteEmailTemplate(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteEmailTemplate(ctx, id)
}

func (r *mutationResolver) SaveHeader(ctx context.Context, name string) (string, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveHeader(ctx, name)
}

func (r *mutationResolver) SaveHeaderDetail(ctx context.Context, parentId string, name string) (string, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveHeaderDetail(ctx, parentId, name)
}

func (r *mutationResolver) UpdateHeader(ctx context.Context, id string, name string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateHeader(ctx, id, name)
}

func (r *mutationResolver) DeleteHeader(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteHeader(ctx, id)
}

func (r *mutationResolver) SaveInstruction(ctx context.Context, input models.SaveInstructionInput) (string, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveInstruction(ctx, input)
}

func (r *mutationResolver) DeleteInstruction(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteInstruction(ctx, id)
}

func (r *mutationResolver) SaveQualityControl(ctx context.Context, pipelineID string) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveQualityControl(ctx, pipelineID, ctxUser.Name)
}

func (r *mutationResolver) UpdateQualityControl(ctx context.Context, id string, input models.UpdateQualityControlInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateQualityControl(ctx, id, input, ctxUser.Name, ctxUser.Id)
}

func (r *mutationResolver) UpdatePipelineNeighborhood(ctx context.Context, pipelineID string, input models.UpdatePipelineNeighborhoodInput) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdatePipelineNeighborhood(ctx, pipelineID, input)
}

func (r *mutationResolver) SaveReview(ctx context.Context, input models.SaveReviewInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveReview(ctx, ctxUser.Name, input)
}

func (r *mutationResolver) UpdateReview(ctx context.Context, id string, input models.UpdateReviewInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateReview(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) DeleteReview(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteReview(ctx, id)
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, id string, input models.UpdateRequestInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateRequest(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) SaveRequest(ctx context.Context, pipelineID string) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveRequest(ctx, pipelineID, ctxUser.Name, ctxUser.UserId)
}

func (r *mutationResolver) SaveInvoice(ctx context.Context, input models.SaveInvoiceInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveInvoice(ctx, input, ctxUser.Name, ctxUser.UserId, nil, nil)
}

func (r *mutationResolver) UpdateInvoice(ctx context.Context, id string, input models.UpdateInvoiceInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateInvoice(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) CancelInvoice(ctx context.Context, id string, reason *string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.CancelInvoice(ctx, id, reason, ctxUser.Name)
}

func (r *mutationResolver) UpdateIform(ctx context.Context, pipelineId string, input models.UpdateIformInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateIform(ctx, pipelineId, input, ctxUser.Name)
}

func (r *mutationResolver) UpdateIformTemp(ctx context.Context, pipelineId string, input models.UpdateIformTempInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateIformTemp(ctx, pipelineId, input, ctxUser.Name)
}

func (r *mutationResolver) UpdatePipelineRepair(ctx context.Context, pipelineId string, input models.UpdatePipelineRepairInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdatePipelineRepair(ctx, pipelineId, input, ctxUser.Name)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UpdateUserInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateUser(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) SaveAnnouncement(ctx context.Context, input models.AnnouncementInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.AddAnnouncement(ctx, input, ctxUser.Name)
}

func (r *mutationResolver) UpdateAnnouncement(ctx context.Context, id string, input models.AnnouncementInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateAnnouncement(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) DeleteAnnouncement(ctx context.Context, id string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteAnnouncement(ctx, id)
}

func (r *mutationResolver) SaveSession(ctx context.Context, userID string, invoiceDate string) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveSession(ctx, userID, invoiceDate, ctxUser.Name)
}

func (r *mutationResolver) StopSession(ctx context.Context, userID string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.StopSession(ctx, userID, ctxUser.Name)
}

func (r *mutationResolver) SavePermissionGroup(ctx context.Context, input models.PermissionGroupInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePermissionGroup(ctx, input, ctxUser.Name)
}

func (r *mutationResolver) UpdatePermissionGroup(ctx context.Context, id string, input models.PermissionGroupInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdatePermissionGroup(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) DeletePermissionGroup(ctx context.Context, id string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeletePermissionGroup(ctx, id, ctxUser.Name)
}

func (r *mutationResolver) SavePipelineComparable(ctx context.Context, pipelineID string, input models.SavePipelineComparableInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePipelineComparable(ctx, pipelineID, input, ctxUser.Name)
}

func (r *mutationResolver) UpdatePipelineComparable(ctx context.Context, id string, mls string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdatePipelineComparable(ctx, id, mls, ctxUser.Name)
}

func (r *mutationResolver) DeletePipelineComparable(ctx context.Context, id string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeletePipelineComparable(ctx, id, ctxUser.Name)
}

func (r *mutationResolver) SaveBilling(ctx context.Context, input models.SaveBillingInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveBilling(ctx, input, ctxUser.Name)
}

func (r *mutationResolver) UpdateBilling(ctx context.Context, id string, input models.UpdateBillingInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateBilling(ctx, id, input, ctxUser.Name)
}

func (r *mutationResolver) DeleteBilling(ctx context.Context, id string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteBilling(ctx, id, ctxUser.Name)
}

func (r *mutationResolver) SaveIformGrid(ctx context.Context, pipelineID string, input models.SaveIformGridInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SaveIformGrid(ctx, pipelineID, input, ctxUser.Name)
}

func (r *mutationResolver) DeleteIformGrid(ctx context.Context, id string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.DeleteIformGrid(ctx, id, ctxUser.Name)
}

func (r *mutationResolver) VerifyPaypalTransaction(ctx context.Context, paypalOrderID string, billingOrderID *string) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.VerifyPaypalTransaction(ctx, paypalOrderID, billingOrderID, ctxUser.Name)
}

func (r *mutationResolver) SavePipelineQualityControlAndNote(ctx context.Context, pipelineID string, input models.SavePipelineQualityControlAndNoteInput) (string, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.SavePipelineQualityControlAndNote(ctx, pipelineID, input, ctxUser.Name)
}

func (r *mutationResolver) UpdateSession(ctx context.Context, sessionID string, input models.UpdateSessionInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateSession(ctx, sessionID, input, ctxUser.Name)
}

func (r *mutationResolver) CreateBillingExcel(ctx context.Context, billingID string) (string, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return "", err
	}
	return services.CreateBillingExcel(ctx, billingID)
}
func (r *mutationResolver) ContinueSession(ctx context.Context, userID string) (bool, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.ContinueSession(ctx, userID)
}

func (r *mutationResolver) UpdateQcRequest(ctx context.Context, ID string, input models.UpdateQcRequestInput) (bool, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return false, err
	}
	return services.UpdateQcRequest(ctx, ID, input, ctxUser.Name)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) AllUser(ctx context.Context, filter *models.UserFilterInput) (*models.UserResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllUser(ctx, filter)
}

func (r *queryResolver) Profile(ctx context.Context) (*models.User, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.GetUser(ctx, ctxUser.UserId)
}

func (r *queryResolver) AllProfileDoc(ctx context.Context, userId *string, filter *models.FilterInput) (*models.ProfileDocResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllProfileDoc(ctx, ctxUser.UserId, userId, filter)
}

func (r *queryResolver) AllAccount(ctx context.Context, filter *models.AccountFilterInput) (*models.AccountResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllAccount(ctx, ctxUser.UserId, filter)
}

func (r *queryResolver) AllAdjustment(ctx context.Context, userID *string) ([]*models.Adjustment, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllAdjustment(ctx, ctxUser.UserId, userID, ctxUser.Role)
}

func (r *queryResolver) AllComment(ctx context.Context, userId *string) ([]*models.Comment, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllComment(ctx, ctxUser.UserId, userId)
}

func (r *queryResolver) Default(ctx context.Context, filterUserId *string) (*models.Default, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.Default(ctx, ctxUser.UserId, filterUserId)
}

func (r *queryResolver) AllPipeline(ctx context.Context, filter *models.PipelineFilterInput) (*models.PipelineResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPipeline(ctx, ctxUser.UserId, ctxUser.Role, filter)
}

func (r *queryResolver) AllPipelineQualityControl(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineQualityControlResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPipelineQualityControl(ctx, pipelineId, filter)
}

func (r *queryResolver) AllPipelineDoc(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineDocResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPipelineDoc(ctx, pipelineId, filter)
}

func (r *queryResolver) AllPipelinePhoto(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelinePhotoResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPipelinePhoto(ctx, pipelineId, filter)
}

func (r *queryResolver) AllPipelineNote(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineNoteResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPipelineNote(ctx, pipelineId, filter)
}

func (r *queryResolver) AllCredit(ctx context.Context) ([]*models.Credit, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllCredit(ctx, ctxUser.UserId)
}

func (r *queryResolver) AllCompany(ctx context.Context, filter *models.CompanyFilterInput) (*models.CompanyResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllCompany(ctx, ctxUser.Id, filter)
}

func (r *queryResolver) Company(ctx context.Context, id string) (*models.Company, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.Company(ctx, id)
}

func (r *queryResolver) PipelineState(ctx context.Context) (*models.PipelineState, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.PipelineState(ctx)
}

func (r *queryResolver) AllEmailTemplate(ctx context.Context, filter *models.EmailTemplateFilterInput) (*models.EmailTemplateResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllEmailTemplate(ctx, filter)
}

func (r *queryResolver) AllHeader(ctx context.Context, filter *models.HeaderFilterInput) (*models.HeaderResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllHeader(ctx, filter, nil)
}

func (r *queryResolver) AllHeaderDetail(ctx context.Context, parentId string, filter *models.HeaderFilterInput) (*models.HeaderResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllHeader(ctx, filter, &parentId)
}

func (r *queryResolver) HeaderCode(ctx context.Context, codes []*string) ([]*models.HeaderCode, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.HeaderCode(ctx, codes)
}

func (r *queryResolver) AllInstruction(ctx context.Context, filter *models.InstructionFilterInput) (*models.InstructionResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllInstruction(ctx, filter)
}

func (r *queryResolver) AllSalesAnalytics(ctx context.Context, filter models.SalesAnalyticsFilterInput) ([]*models.SalesAnalytics, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllSalesAnalytics(ctx, filter)
}

func (r *queryResolver) AllOrderAnalytics(ctx context.Context, filter models.OrderAnalyticsFilterInput) ([]*models.OrderAnalytics, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllOrderAnalytics(ctx, filter)
}

func (r *queryResolver) AllQualityControl(ctx context.Context, filter *models.FilterInput) (*models.QualityControlResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllQualityControl(ctx, filter)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.GetUser(ctx, id)
}

func (r *queryResolver) AllPipelineQualityControlAndNote(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineQualityControlAndNoteResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPipelineQualityControlAndNote(ctx, pipelineId, filter)
}

func (r *queryResolver) PipelineNeighborhood(ctx context.Context, pipelineId string) (*models.PipelineNeighborhood, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.PipelineNeighborhood(ctx, pipelineId)
}

func (r *queryResolver) AllBalance(ctx context.Context, filter *models.BalanceFilterInput) (*models.BalanceResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllBalance(ctx, filter)
}

func (r *queryResolver) AllCheckout(ctx context.Context, filter *models.CheckoutFilterInput) (*models.CheckoutResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllCheckout(ctx, filter)
}

func (r *queryResolver) AllCredits(ctx context.Context, filter *models.CreditsFilterInput) (*models.CreditsResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllCredits(ctx, filter)
}

func (r *queryResolver) AllQcRating(ctx context.Context, year int, _type *string) ([]*models.QcRating, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllQcRating(ctx, year, _type)
}

func (r *queryResolver) AllOrderSubmit(ctx context.Context, year int) ([]*models.OrderSubmit, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllOrderSubmit(ctx, year)
}

func (r *queryResolver) AllQcHistory(ctx context.Context, filter *models.QcHistoryFilterInput) (*models.QcHistoryResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllQcHistory(ctx, filter, ctxUser.UserId, ctxUser.Role)
}

func (r *queryResolver) AllQcRequest(ctx context.Context, filter *models.QcRequestFilterInput) (*models.QcRequestResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllQcRequest(ctx, filter)
}

func (r *queryResolver) AllReview(ctx context.Context, filter *models.FilterInput) (*models.ReviewResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllReview(ctx, filter)
}

func (r *queryResolver) AllRequest(ctx context.Context, filter *models.RequestFilterInput) (*models.RequestResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllRequest(ctx, filter)
}

func (r *queryResolver) AllQcCompleted(ctx context.Context, year int) ([]*models.QcCompleted, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllQcCompleted(ctx, year)
}

func (r *queryResolver) Pipeline(ctx context.Context, id string) (*models.Pipeline, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.Pipeline(ctx, id)
}

func (r *queryResolver) AllInvoice(ctx context.Context, filter *models.InvoiceFilterInput) ([]*models.Invoice, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllInvoice(ctx, filter, ctxUser.Role, ctxUser.UserId)

}

func (r *queryResolver) Iform(ctx context.Context, pipelineId string) (*models.Iform, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.Iform(ctx, pipelineId, ctxUser.UserId, ctxUser.Name, ctxUser.Role)
}

func (r *queryResolver) IformTemp(ctx context.Context, pipelineId string) (*models.IformTemp, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.IformTemp(ctx, pipelineId, ctxUser.UserId, ctxUser.Name, ctxUser.Role)
}

func (r *queryResolver) PipelineRepair(ctx context.Context, pipelineId string) (*models.PipelineRepair, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.GetPipelineRepair(ctx, pipelineId)
}

func (r *queryResolver) AllAnnouncement(ctx context.Context, filter *models.AnnouncementFilterInput) (*models.AnnouncementResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllAnnouncement(ctx, filter)
}

func (r *queryResolver) AllSession(ctx context.Context, filter *models.SessionFilterInput) (*models.SessionResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllSession(ctx, ctxUser.UserId, filter)
}

func (r *queryResolver) AllPermissionGroup(ctx context.Context, filter *models.PermissionGroupFilterInput) (*models.PermissionGroupResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPermissionGroup(ctx, filter)
}

func (r *queryResolver) AllLoginLog(ctx context.Context, filter *models.LoginLogFilterInput) (*models.LoginLogResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllLoginLog(ctx, filter)
}

func (r *queryResolver) AllPipelineComparable(ctx context.Context, pipelineID string, filter *models.PipelineComparableFilterInput) (*models.PipelineComparableResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllPipelineComparable(ctx, pipelineID, filter)
}

func (r *queryResolver) AllBilling(ctx context.Context, filter *models.BillingFilterInput) (*models.BillingResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllBilling(ctx, filter, ctxUser.Role, ctxUser.UserId)
}

func (r *queryResolver) AllIformGrid(ctx context.Context, pipelineID string, filter *models.IformGridFilterInput) (*models.IformGridResult, error) {
	_, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllIformGrid(ctx, pipelineID, filter)
}

func (r *queryResolver) AllInvoiceRequest(ctx context.Context, filter *models.InvoiceRequestFilterInput) (*models.InvoiceRequestResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllInvoiceRequest(ctx, ctxUser.UserId, filter)
}

func (r *queryResolver) AllInvoiceRequestHistory(ctx context.Context, filter *models.FilterInput) (*models.InvoiceRequestHistoryResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllInvoiceRequestHistory(ctx, ctxUser.UserId, filter)
}

func (r *queryResolver) PriceModule(ctx context.Context) (*models.PriceModule, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.PriceModule(ctx, ctxUser.UserId)
}

func (r *queryResolver) AllCreditLedger(ctx context.Context, userId *string, filter *models.FilterInput) (*models.CreditLedgerResult, error) {
	ctxUser, err := middleware.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}
	return services.AllCreditLedger(ctx, filter, ctxUser.UserId, userId)
}
