package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marmotedu/errors"
)

type ErrResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
}

type SucResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func WriteResponse(c echo.Context, err error, data interface{}) error {
	if err != nil {
		coder := errors.ParseCoder(err)
		return c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})

	}

	return c.JSON(http.StatusOK, SucResponse{
		Code:    http.StatusOK,
		Message: "success",
		Result:  data,
	})
}
