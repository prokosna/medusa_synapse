package route

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/prokosna/medusa_synapse/exception"
)

type ErrorResp struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

func HandleError(err error, c echo.Context) {
	var code int
	var desc string
	switch err := err.(type) {
	case *exception.BadRequestError:
		code = 400
		desc = err.Error()
	case *exception.UnauthorizedError:
		code = 401
		desc = err.Error()
	case *exception.ServiceUnavailableError:
		code = 503
		desc = err.Error()
	case *echo.HTTPError:
		code = err.Code
		desc = fmt.Sprintf("%v", err.Message)
	default:
		code = 500
		desc = err.Error()
	}
	c.JSON(code, ErrorResp{Status: "error", Description: desc})
}
