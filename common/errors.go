package common

import "net/http"

type AppError struct{
	StatusCode int `json:"status_code"`
	RootErr error `json:"-"`
	Message string `json:"message"`
	Log string `json:"log"`
	Key string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int,root error, message,log,key string) * AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
// 400
func BadRequestError(root error,message,log,key string) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
// 401
func UnauthorizedError(root error,message,log,key string) *AppError{
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
// 403
func ForbiddenError(root error,message,log,key string) *AppError{
	return &AppError{
		StatusCode: http.StatusForbidden,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
// 404
func NotFoundError(root error,message,log,key string) *AppError{
	return &AppError{
		StatusCode: http.StatusNotFound,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
// 405
func MethodNotAllowedError(root error,message,log,key string) *AppError{
	return &AppError{
		StatusCode: http.StatusMethodNotAllowed,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
func ConflictError(root error,message,log,key string) *AppError{
	return &AppError{
		StatusCode: http.StatusConflict,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
// 500
func InternalServerError(root error,message,log,key string) *AppError{
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		RootErr: root,
		Message: message,
		Log: log,
		Key: key,
	}
}
// DB
func DB_ERROR(err error) * AppError{
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"Something wrongs",
		err.Error(),
		"DB_ERROR",
	)
}
func (e* AppError) RootError() error {
	err,ok := e.RootErr.(*AppError);
	if ok {
		return err.RootError()
	}
	return e.RootErr
}
func (e *AppError) Error() string {
	return e.RootError().Error()
}


