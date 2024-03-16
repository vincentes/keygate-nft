package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetContracts(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	contracts, err := model.GetContracts(tx)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    contracts,
	}
	c.JSON(http.StatusOK, response)
	return nil
}

func CreateContract(c echo.Context) error {
	var contract model.Contract
	if err := c.Bind(&contract); err != nil {
		return err
	}

	if err := c.Validate(&contract); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)
	err := model.CreateContract(tx, &contract)

	return err
}
