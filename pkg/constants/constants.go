package constants

const UserStatusDeleted = "DELETED"
const UserStatusActive = "ACTIVE"
const ActiveTurbo = "ACTIVE-TURBO"

const UserRoleClient = "CLIENT"
const UserRoleAdmin = "ADMIN"
const UserRoleContractor = "CONTRACTOR"
const UserRoleQualityControl = "QUALITY_CONTROL"

const IformProcessingFee = 1.00

const AppActionSave = "ADD"
const AppActionUpdate = "UPDATE"

const PipelineStatusActive = "ACTIVE"
const PipelineStatusStandBy = "STANDBY"
const PipelineStatusComplete = "COMPLETE"
const PipelineStatusSubmit = "SUBMIT"
const PipelineStatusCancelled = "CANCEL"
const PipelineStatusPaid = "PAID"
const PipelineStatusQC = "QC"
const PipelineStatusLate = "LATE"
const PipelineStatushold = "HOLD"
const PipelineStatusUnhold = "UNHOLD"
const PipelineStatusActivePhotos = "ACTIVE_PHOTOS"

const QualityControlStatusForQCRequest = "FOR_QC_REQUEST"
const QualityControlStatusActive = "ACTIVE"
const QualityControlStatusPending = "PENDING"
const QualityControlStatusHold = "HOLD"
const QualityControlStatusCancelled = "CANCELLED"

const QualityControlRequestTypeDataDisCrepancy = "DATA_DISCREPANCY"
const QualityControlRequestTypeDataDisCrepancyNqc = "DATA_DISCREPANCY_NQC"

const QualityControlRequestStatusPending = "HOLD"
const QualityControlRequestStatusApproved = "COMPLETE"
const QualityControlRequestStatusDD = "DD"

const QcTypeDataDiscrepancy = "Data Discripancy"
const QcTypeDataDiscrepancyNqc = "Data Discripancy NQC" // count
const QcTypesSubmit = "Submit"
const QcTypeSubmitNqc = "Submit NQC"   // count
const QcTypeNormal = "Normal"          // count
const QcTypefullRecomp = "Full Recomp" // count
const QcTypeCancel = "Cancel"
const QcTypeAlreadyAddressed = "Already Addressed"

const PipelineOrderTypeInterior = "INTERIOR"
const PipelineOrderTypeExterior = "EXTERIOR"
const PipelineOrderTypeDataEntry = "DATA ENTRY"
const PipelineOrderTypeConditionReport = "CONDITION REPORT"

const BillingStatusPending = "PENDING"
const BillingStatusPaid = "PAID"

const EmailTemplateOrderAssignment = "ORDER_ASSIGNMENT"
const EmailTemplateAccountVerification = "ACCOUNT_VERIFICATION"
const EmailTemplateAccountVerificationSuccess = "ACCOUNT_VERIFICAION_SUCCESS"
const EmailTemplateAccountForgetPassword = "ACCOUNT_FORGET_PASSWORD"
const EmailTemplateAccountChangePasswordSuccess = "ACCOUNT_CHANGE_PASSWORD_SUCCESS"
const EmailTemplateOrderActive = "ORDER_ACTIVE"
const EmailTemplateOrderComplateForStandby = "ORDER_COMPLETE_FROM_STANDBY"
const EmailTemplateOrderRush = "ORDER_RUSH"
const EmailTemplateOrderComplete = "ORDER_COMPLETE"
const EmailTemplateOrderCancel = "ORDER_CANCEL"
const EmailTemplateOrderNew = "ORDER_NEW"
const EmailTemplateOrderStandby = "ORDER_STANDBY"

const EmailTemplateOrderSuperRush = "ORDER_SUPER_RUSH"

const RequestStatusPending = "PENDING"

const EmailParseKeyClientId = "{clientId}"
const EmailParseKeyOrderNumber = "{orderNumber}"
const EmailParseKeyAccountResetPassword = "{accountPasswordResetLink}"

const SuperAdminId = "5dc12cf84f26a8f5c2501d92"
const DefaultClientPermissionGroupID = "5dea4e55c35ae1b222148908"

//ALtisrouce mkt id
const ALtisrouceMktID = "5dd6317d778617cbe9f32926"

const PipelineDocTypeCompsMls = "COMPS MLS"

const DefaultTheme = "theme-classic"

//comment
const CommentSectionDefaultValue = "FINAL_VALUES"
