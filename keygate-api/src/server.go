package main

import (
	"errors"
	"fmt"
	"keygate/api/api"
	"keygate/api/db"
	kgmiddleware "keygate/api/middleware"
	"net/http"
	"strings"

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

type ApiError struct {
	Param string `json:"param"`
	Message string `json:"message"`
}

func tagToMessage(fe validator.FieldError) string {
	switch fe.Tag() {
		case "required":
			return "This field is required"
		case "email":
			return "Invalid email"
		case "gtfield":
			return "This field must be greater than " + fe.Param()
	}

	return fe.Tag()
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
	var ve validator.ValidationErrors
	var out []ApiError
	if errors.As(err, &ve) {
		out = make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{strings.ToLower(fe.Field()), tagToMessage(fe)}
		}
	} else {
		out = append(out, ApiError{"unknown", err.Error()})
	}

	res := &api.JSendResponse{
		Status:  api.ResponseError,
		Message: "Validation error",
		Data:    out,
	}

	return echo.NewHTTPError(http.StatusBadRequest, res)
  }

  return nil
}

func main() {
	loadEnv()
	db.Connect()
	
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(kgmiddleware.TransactionHandler(db.GetDB()))

	// Validate
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &api.JSendResponse{ Message: "Keygate API v1.0", Status: "success" })
	});

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

	e.HEAD("/users/:userId/permissions", func(c echo.Context) error {
		err := api.CheckUserPermissionByName(c)
		return err
	})

	e.HEAD("/users/:userId/permissions/:permissionId", func(c echo.Context) error {
		err := api.CheckUserPermission(c)
		return err
	})

	e.HEAD("/ext/users/:extUserId/permissions/:permissionId", func(c echo.Context) error {
		err := api.CheckUserPermissionByExternalID(c)
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

	e.POST("/collections", func(c echo.Context) error {
		err := api.CreateCollection(c)
		return err
	});

	e.GET("/collections", func(c echo.Context) error {
		err := api.GetCollections(c)
		return err
	});

	e.DELETE("/collections/:id", func(c echo.Context) error {
		err := api.DeleteCollection(c)
		return err
	})

	e.GET("/quizzes", func(c echo.Context) error {
		err := api.GetQuizzes(c)
		return err
	});

	e.POST("/quizzes", func(c echo.Context) error {
		err := api.CreateQuiz(c)
		return err
	});

	e.POST("/quizzes/:id/questions", func(c echo.Context) error {
		err := api.CreateQuestion(c)
		return err
	});

	e.GET("/tiers", func(c echo.Context) error {
		err := api.GetTiers(c)
		return err
	})

	e.GET("/tiers/:id", func(c echo.Context) error {
		err := api.GetTier(c)
		return err
	})

	e.POST("/tiers", func(c echo.Context) error {
		err := api.CreateTier(c)
		return err
	})

	e.PUT("/tiers/:id", func(c echo.Context) error {
		err := api.UpdateTier(c)
		return err
	})

	e.DELETE("/tiers/:id", func(c echo.Context) error {
		err := api.DeleteTier(c)
		return err
	})

	e.POST("/images/upload-url", func(c echo.Context) error {
		err := api.GetSignedURL(c)
		return err
	})

	e.Logger.Fatal(e.Start(":8080"))
}