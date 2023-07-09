package exceptions

import (
	"fmt"
	"github.com/labstack/echo"
)

type DWSErrorResponse struct {
	errorCode        string
	errorDescription string
	debugReferenceID string
}

func (e *DWSErrorResponse) Error() string {
	return fmt.Sprintf("Error: %s \n %s \n %s", e.errorCode, e.errorDescription, e.debugReferenceID)
}

func (e *DWSErrorResponse) UpdateRefID(ctx echo.Context) DWSErrorResponse {
	requestIDVal, ok := ctx.Get("requestID").(string)
	if ok {
		e.debugReferenceID = requestIDVal
	}
	return *e
}

var ProjectCreationError = DWSErrorResponse{errorCode: "projectCreationError", debugReferenceID: "",
	errorDescription: "There is some issue in creating your project",
}
var UnexpectedError = DWSErrorResponse{errorCode: "unexpectedError", debugReferenceID: "",
	errorDescription: "Some unexpected error has occurred",
}
