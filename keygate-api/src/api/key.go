package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)


func CreateKey (c echo.Context) error {
	// Request body to Key
	var key model.Key
	if err := c.Bind(&key); err != nil {
		return err
	}

	if err := c.Validate(&key); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)

	// Check if key already exists
	exists, err := model.DoesKeyExistByName(tx, key.Name)
	if err != nil {
		return err
	}

	if exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Key already exists.",
		}
		c.JSON(http.StatusConflict, response)
		return nil
	}

	err = model.CreateKey(tx, &key)

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Key created successfully.",
		Data:    key,
	}

	c.JSON(http.StatusOK, response)
	return err
}

func GetKeys (c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	keys, err := model.GetKeys(tx)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Message: "Keys retrieved successfully.",
		Data:    keys,
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func DeleteKey (c echo.Context) error {
	id := c.Param("id")
	tx := c.Get("Tx").(*sql.Tx)

	exists, err := model.DoesKeyExist(tx, id)
	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Key does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}

	err = model.DeleteKey(tx, id)
	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Key deleted successfully.",
	}
	c.JSON(http.StatusOK, response)
	return err
}

func CreatePermission (c echo.Context) error {
	var permission model.Permission
	if err := c.Bind(&permission); err != nil {
		return err
	}

	if err := c.Validate(&permission); err != nil {
		return err
	}

	exists, err := model.DoesPermissionExistByName(c.Get("Tx").(*sql.Tx), permission.Name)
	if err != nil {
		return err
	}

	if exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Permission already exists.",
		}
		c.JSON(http.StatusConflict, response)
		return nil
	}

	tx := c.Get("Tx").(*sql.Tx)
	err = model.CreatePermission(tx, permission)
	return err
}

func GetPermissions (c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	permissions, err := model.GetPermissions(tx)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Message: "Permissions retrieved successfully.",
		Data:    permissions,
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func DeletePermission (c echo.Context) error {
	id := c.Param("id")

	exists, err := model.DoesPermissionExist(c.Get("Tx").(*sql.Tx), id)
	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Permission does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}
	
	tx := c.Get("Tx").(*sql.Tx)
	err = model.DeletePermission(tx, id)

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Permission deleted successfully.",
	}

	c.JSON(http.StatusOK, response)

	return err
}

// API checks if key exists
// API checks if user exists
// API creates link between key and user

func AttachKey (c echo.Context) error {
	// POST /users/{userId}/keys

	// get user id from params
	userID := c.Param("userId")

	// get key id from request body
	var attachment model.UserKeyAttachment

	if err := c.Bind(&attachment); err != nil {
		return err
	}

	if err := c.Validate(&attachment); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)

	// Check if key exists
	exists, err := model.DoesKeyExist(tx, attachment.KeyID)
	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Key does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}

	// Check if user exists
	exists, err = model.DoesUserExist(tx, userID)
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

	// Check if key is already attached to user
	attached, err := model.IsKeyAttachedToUser(tx, attachment.KeyID, userID)
	if err != nil {
		return err
	}

	if attached {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Key is already attached to user.",
		}
		c.JSON(http.StatusConflict, response)
		return nil
	}

	// Attach key to user
	err = model.AttachKeyToUser(tx, attachment.KeyID, userID)
	if err != nil {
		return err
	}

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Key attached to user successfully.",
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func AddPermissionToKey (c echo.Context) error {
	// POST /keys/{keyId}/permissions

	// get key id from params
	keyID := c.Param("keyId")

	// get permission id from request body
	var attachment model.KeyPermissionAttachment

	if err := c.Bind(&attachment); err != nil {
		return err
	}

	if err := c.Validate(&attachment); err != nil {
		return err
	}

	tx := c.Get("Tx").(*sql.Tx)

	// Check if key exists
	exists, err := model.DoesKeyExist(tx, keyID)
	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Key does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}

	// Check if permission exists
	exists, err = model.DoesPermissionExist(tx, attachment.PermissionID)
	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Permission does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}

	// Check if permission is already attached to key
	attached, err := model.DoesKeyContainPermission(tx, attachment.PermissionID, keyID)
	if err != nil {
		return err
	}

	if attached {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Permission is already attached to key.",
		}
		c.JSON(http.StatusConflict, response)
		return nil
	}

	// Attach permission to key
	err = model.AddPermissionToKey(tx, keyID, attachment.PermissionID)
	if err != nil {
		return err
	}

	response := JSendResponse{
		Status:  ResponseSuccess,
		Message: "Permission attached to key successfully.",
	}

	return c.JSON(http.StatusOK, response)
}

func GetKeyPermissions (c echo.Context) error {
	// GET /keys/{keyId}/permissions

	// get key id from params
	keyID := c.Param("keyId")

	tx := c.Get("Tx").(*sql.Tx)

	// Check if key exists
	exists, err := model.DoesKeyExist(tx, keyID)
	if err != nil {
		return err
	}

	if !exists {
		response := JSendResponse{
			Status:  ResponseFail,
			Message: "Key does not exist.",
		}
		c.JSON(http.StatusNotFound, response)
		return nil
	}

	permissions, err := model.GetKeyPermissions(tx, keyID)
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