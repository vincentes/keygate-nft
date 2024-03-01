package api

import (
	"database/sql"
	"keygate/api/db"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)

	// External ID from the query params
	externalID := c.QueryParam("external_id")

	var users []model.User
	var err error
	if externalID != "" {
		user, err := model.GetUserByExternalID(tx, externalID)
		if err != nil {
			return err
		}

		if user == nil {	
			return c.JSON(http.StatusOK, JSendResponse{
				Status:  ResponseSuccess,
				Message: "Users retrieved succesfully.",
				Data:    []model.User{},
			})
		}

		users = append(users, *user)
	} else {
		users, err = model.GetUsers(tx)
	}

	if err != nil {
		return err
	}
	
	var response = JSendResponse{
		Status:  ResponseSuccess,
		Message: "Users retrieved successfully.",
		Data:    users,
	}

	c.JSON(http.StatusOK, response)

	return err
}

func DoesUserExist(c echo.Context, externalID string) (bool, error) {
	conn, err := db.Conn(c.Request().Context())
	if err != nil {
		return false, err
	}

	rows, err := conn.QueryContext(c.Request().Context(), "SELECT ID FROM User WHERE ExternalID = ?", externalID)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func CreateUser(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}


	tx := c.Get("Tx").(*sql.Tx)
	exists, err := model.DoesUserExist(tx, user.ExternalID)

	if err != nil {
		return err
	}

	if exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "User already exists.",
		}
		c.JSON(http.StatusConflict, response)
		return nil
	}

	err = model.CreateUser(tx, &user)
	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "User created successfully.",
		Data:    user,
	}

	c.JSON(http.StatusOK, response)
	return err
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	tx := c.Get("Tx").(*sql.Tx)

	exists, err := DoesUserExist(c, id)

	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "User does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}

	err = model.DeleteUser(tx, id)

	if err != nil {
		return err
	}

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "User deleted successfully.",
	}

	return c.JSON(http.StatusOK, response)
}



func GetUserPermissions (c echo.Context) error {
	userID := c.Param("userId")

	tx := c.Get("Tx").(*sql.Tx)

	permissions, err := model.GetUserPermissions(tx, userID)

	if err != nil {
		return err
	}

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Permissions retrieved successfully.",
		Data:    permissions,
	}

	c.JSON(http.StatusOK, response)

	return nil
}



func GetUserKeys (c echo.Context) error {
	userID := c.Param("userId")

	tx := c.Get("Tx").(*sql.Tx)

	exists, err := model.DoesUserExist(tx, userID)

	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "User does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}

	keys, err := model.GetUserKeys(tx, userID)

	if err != nil {
		return err
	}

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Keys retrieved successfully.",
		Data:    keys,
	}

	c.JSON(http.StatusOK, response)

	return nil
}

func CheckUserPermissionByExternalID (c echo.Context) error {
	extUserId := c.Param("extUserId")
	permissionID := c.Param("permissionId")

	tx := c.Get("Tx").(*sql.Tx)

	exists, err := model.DoesUserExistByExternalID(tx, extUserId)

	if err != nil {
		return err
	}

	if !exists {
		c.NoContent(http.StatusNotFound)
		return nil
	}

	hasPermission, err := model.CheckUserPermissionByExternalID(tx, extUserId, permissionID)

	if err != nil {
		return err
	}

	if hasPermission {
		c.NoContent(http.StatusOK)
		return nil
	}

	c.NoContent(http.StatusNotFound)
	return nil
}