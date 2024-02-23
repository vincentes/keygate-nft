package main

import (
	"fmt"
	"keygate/api/api"
	"keygate/api/db"
	kgmiddleware "keygate/api/middleware"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load .env file: %s", err))
	}
	fmt.Println("Loaded environment variables.")
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}

func main() {
	loadEnv()
	db.Connect()
	
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(kgmiddleware.TransactionHandler(db.GetDB()))

	// Validate
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/users", func(c echo.Context) error {
		return api.CreateUser(c)
	})

	e.GET("/users", func(c echo.Context) error {
		err := api.GetUsers(c)
		return err
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		err := api.DeleteUser(c)
		return err
	})

	e.GET("/users/:userId/permissions", func(c echo.Context) error {
		err := api.GetUserPermissions(c)
		return err
	})

	e.POST("/users/:userId/keys", func(c echo.Context) error {
		err := api.AttachKey(c)
		return err
	})

	e.GET("/users/:userId/keys", func(c echo.Context) error {
		err := api.GetUserKeys(c)
		return err
	})

	e.POST("/keys", func(c echo.Context) error {
		err := api.CreateKey(c)
		return err
	})

	e.GET("/keys", func(c echo.Context) error {
		err := api.GetKeys(c)
		return err
	})

	e.DELETE("/keys/:id", func(c echo.Context) error {
		err := api.DeleteKey(c)
		return err
	})

	e.GET("/keys/:keyId/permissions", func(c echo.Context) error {
		err := api.GetKeyPermissions(c)
		return err
	})

	e.POST("/keys/:keyId/permissions", func(c echo.Context) error {
		err := api.AddPermissionToKey(c)
		return err
	})

	e.POST("/permissions", func(c echo.Context) error {
		err := api.CreatePermission(c)
		return err
	})

	e.GET("/permissions", func(c echo.Context) error {
		err := api.GetPermissions(c)
		return err
	})

	e.DELETE("/permissions/:id", func(c echo.Context) error {
		err := api.DeletePermission(c)
		return err
	})

	e.Logger.Fatal(e.Start(":8080"))
}