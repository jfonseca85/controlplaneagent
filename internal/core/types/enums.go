package types

type HandlerErrorCode string

// Enum values for HandlerErrorCode
const (
	HandlerErrorCodeNotUpdatable            HandlerErrorCode = "NotUpdatable"
	HandlerErrorCodeInvalidRequest          HandlerErrorCode = "InvalidRequest"
	HandlerErrorCodeAccessDenied            HandlerErrorCode = "AccessDenied"
	HandlerErrorCodeInvalidCredentials      HandlerErrorCode = "InvalidCredentials"
	HandlerErrorCodeAlreadyExists           HandlerErrorCode = "AlreadyExists"
	HandlerErrorCodeNotFound                HandlerErrorCode = "NotFound"
	HandlerErrorCodeResourceConflict        HandlerErrorCode = "ResourceConflict"
	HandlerErrorCodeThrottling              HandlerErrorCode = "Throttling"
	HandlerErrorCodeServiceLimitExceeded    HandlerErrorCode = "ServiceLimitExceeded"
	HandlerErrorCodeNotStabilized           HandlerErrorCode = "NotStabilized"
	HandlerErrorCodeGeneralServiceException HandlerErrorCode = "GeneralServiceException"
	HandlerErrorCodeServiceInternalError    HandlerErrorCode = "ServiceInternalError"
	HandlerErrorCodeServiceTimeout          HandlerErrorCode = "ServiceTimeout"
	HandlerErrorCodeNetworkFailure          HandlerErrorCode = "NetworkFailure"
	HandlerErrorCodeInternalFailure         HandlerErrorCode = "InternalFailure"
)

// Values returns all known values for HandlerErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HandlerErrorCode) Values() []HandlerErrorCode {
	return []HandlerErrorCode{
		"NotUpdatable",
		"InvalidRequest",
		"AccessDenied",
		"InvalidCredentials",
		"AlreadyExists",
		"NotFound",
		"ResourceConflict",
		"Throttling",
		"ServiceLimitExceeded",
		"NotStabilized",
		"GeneralServiceException",
		"ServiceInternalError",
		"ServiceTimeout",
		"NetworkFailure",
		"InternalFailure",
	}
}

type Operation string

// Enum values for Operation
const (
	OperationCreate Operation = "CREATE"
	OperationDelete Operation = "DELETE"
	OperationUpdate Operation = "UPDATE"
)

// Values returns all known values for Operation. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operation) Values() []Operation {
	return []Operation{
		"CREATE",
		"DELETE",
		"UPDATE",
	}
}

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusPending          OperationStatus = "PENDING"
	OperationStatusInProgress       OperationStatus = "IN_PROGRESS"
	OperationStatusSuccess          OperationStatus = "SUCCESS"
	OperationStatusFailed           OperationStatus = "FAILED"
	OperationStatusCancelInProgress OperationStatus = "CANCEL_IN_PROGRESS"
	OperationStatusCancelComplete   OperationStatus = "CANCEL_COMPLETE"
)

// Values returns all known values for OperationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		"PENDING",
		"IN_PROGRESS",
		"SUCCESS",
		"FAILED",
		"CANCEL_IN_PROGRESS",
		"CANCEL_COMPLETE",
	}
}
