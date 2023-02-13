package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/mlc-d/table/pkg/dto"
	"gitlab.com/mlc-d/table/web/api/responses"
	"gitlab.com/mlc-d/table/web/api/services"
)

type signUpResponse struct {
	ID int64 `json:"id"`
}

func signUp(c echo.Context) error {
	u := &dto.User{}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, &responses.Response{
			Code:    http.StatusBadRequest,
			Comment: "invalid json parameters",
			Payload: &responses.Error{ErrorString: err.Error()},
		})
	}

	id, err := services.UserService.Signup(c.Request().Context(), u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Comment: "internal server error",
			Payload: &responses.Error{ErrorString: err.Error()},
		})
	}
	return c.JSON(http.StatusCreated, &responses.Response{
		Code:    http.StatusCreated,
		Payload: signUpResponse{ID: id},
		Comment: "signed up",
	})
}
