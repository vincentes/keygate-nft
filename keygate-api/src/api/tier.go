package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetTiers(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	tiers, err := model.GetTiers(tx)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    tiers,
		Message: "Tiers retrieved",
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func GetTier(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	id := c.Param("id")

	tier, err := model.GetTier(tx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, JSendResponse{
				Status:  ResponseFail,
				Data:    nil,
				Message: "Tier not found",
			})
		}
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    tier,
		Message: "Tier retrieved",
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func CreateTier(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	req := new(model.TierCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	tier, err := model.CreateTier(tx, req)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    tier,
		Message: "Tier created",
	}

	c.JSON(http.StatusCreated, response)
	return nil
}

func UpdateTier(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	id := c.Param("id")

	req := new(model.TierUpdateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	tier, err := model.UpdateTier(tx, id, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, JSendResponse{
				Status:  ResponseFail,
				Data:    nil,
				Message: "Tier not found",
			})
		}
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    tier,
		Message: "Tier updated",
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func DeleteTier(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	id := c.Param("id")

	err := model.DeleteTier(tx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, JSendResponse{
				Status:  ResponseFail,
				Data:    nil,
				Message: "Tier not found",
			})
		}
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    nil,
		Message: "Tier deleted",
	}

	c.JSON(http.StatusOK, response)
	return nil
}