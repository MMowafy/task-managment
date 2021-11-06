package utils

const (
	ErrorInvalidRequestPayload          = "invalid requests payload"
	ErrorCreateUser                     = "create user failed"
	ErrorCreateTask                     = "create task failed"
	ErrorCreateNotification             = "create notifcation failed"
	ErrorUserNotFound                   = "User not found"
	ErrorTaskNotFound                   = "Task not found"
	ErrorNotificationNotFound           = "Task not found"
	ErrorTaskOverlappingFound           = "Task overlapping for same user"
	ErrorGeneralInternalError           = "something went wrong while gettling list of users"
	ErrorCreateUserSystemCode           = 1000
	ErrorCreateTaskSystemCode           = 1001
	ErrorCreateNotificationSystemCode   = 1002
	ErrorUserNotFoundSystemCode         = 2000
	ErrorTaskNotFoundSystemCode         = 2001
	ErrorNotificationNotFoundSystemCode = 2002
	ErrorTaskOverlappingFoundSystemCode = 1002
	ErrorGeneralInternalErrorSystemCode = 3000
	EmailAlreadyExisted                 = "email already existed"
	ErrorWithUnknownSystemCode          = 4000
	ErrorValidationSystemCode           = 5000
)

const (
	FILTER_OPERATOR_EQUAL = "="
)
