package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)



func GetSignedURL(c echo.Context) error {
	var request model.ImageURLRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)
	urlData, err := model.GetSignedURL(tx, &request)
	if err != nil {
		return err
	}

	response := JSendResponse{
		Status: ResponseSuccess,
		Data:   urlData,
		Message: "Presigned URL created.",
	}

	c.JSON(http.StatusOK, response)
	return nil
}