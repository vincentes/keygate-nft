package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetCollectibles(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	collectibles, err := model.GetCollectibles(tx)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    collectibles,
		Message: "Collectibles retrieved",
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func CreateCollectible(c echo.Context) error {
	var request model.CollectibleCreateRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)
	err, collectible := model.CreateCollectible(tx, &request)
	if err != nil {
		return err
	}

	response := JSendResponse{
		Status: ResponseSuccess,
		Data:   collectible,
		Message: "Collectible created",
	}


	c.JSON(http.StatusCreated, response)

	return err
}
