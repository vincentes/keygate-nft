package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)

func CreateToken(c echo.Context) error {
	var token model.Token
	if err := c.Bind(&token); err != nil {
		return err
	}

	if err := c.Validate(&token); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)
	model.CreateToken(tx, &token);

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Token created successfully.",
	};

	c.JSON(http.StatusOK, response);
	return nil;
}