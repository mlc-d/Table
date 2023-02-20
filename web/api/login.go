package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/mlc-d/table/dto"
	"gitlab.com/mlc-d/table/pkg/errs"
	"gitlab.com/mlc-d/table/web/api/responses"
	"gitlab.com/mlc-d/table/web/api/services"
)

type loginResponse struct {
	User *dto.User `json:"user"`
}

func login(c echo.Context) error {
	u := &dto.User{}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, &responses.Response{
			Code:    http.StatusBadRequest,
			Payload: &responses.Error{ErrorString: err.Error()},
			Comment: errs.ErrBadRequest.Error(),
		})
	}

	u, err := services.UserService.Login(c.Request().Context(), u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &responses.Response{
			Code:    http.StatusInternalServerError,
			Payload: &responses.Error{ErrorString: err.Error()},
			Comment: "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, &responses.Response{
		Code:    http.StatusOK,
		Payload: &loginResponse{User: u},
		Comment: "logged in",
	})
}
