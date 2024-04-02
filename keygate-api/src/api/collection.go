package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetCollections(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	collections, err := model.GetCollections(tx)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    collections,
		Message: "Collections retrieved",
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func CreateCollection(c echo.Context) error {
	var request model.CollectionCreateRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)
	err, collection := model.CreateCollection(tx, &request)
	if err != nil {
		return err
	}

	response := JSendResponse{
		Status: ResponseSuccess,
		Data:   collection,
		Message: "Collection created",
	}


	c.JSON(http.StatusCreated, response)

	return err
}

func DeleteCollection(c echo.Context) error {
    collectionID := c.Param("id")
    tx := c.Get("Tx").(*sql.Tx)

    err := model.DeleteCollection(tx, collectionID)
    if err != nil {
        return err
    }

    response := JSendResponse{
        Status:  ResponseSuccess,
        Message: "Collection deleted",
    }

    c.JSON(http.StatusOK, response)
    return nil
}