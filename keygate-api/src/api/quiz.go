package api

import (
	"database/sql"
	"keygate/api/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetQuizzes(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	quizzes, err := model.GetQuizzes(tx)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    quizzes,
		Message: "Quizzes retrieved",
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func CreateQuiz(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	req := new(model.QuizCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	err, quiz := model.CreateQuiz(tx, req)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    quiz,
		Message: "Quiz created",
	}

	c.JSON(http.StatusCreated, response)
	return nil
}

func CreateQuestion(c echo.Context) error {
	tx := c.Get("Tx").(*sql.Tx)
	quizID := c.Param("id")

	req := new(model.QuestionCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	err, question := model.CreateQuestion(tx, quizID, req)
	if err != nil {
		return err
	}

	var response = JSendResponse{
		Status:  ResponseSuccess,
		Data:    question,
		Message: "Question added",
	}

	c.JSON(http.StatusCreated, response)
	return nil
}