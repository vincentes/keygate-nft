package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID          string `json:"id"`
	Content    string `json:"content" validate:"required"`
	QuizID      string `json:"-"`
	Options    []Answer `json:"options"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Answer struct {
	ID          string `json:"id"`
	Answer      string `json:"answer" validate:"required"`
	IsCorrect     bool `json:"is_correct"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Quiz struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Questions   []Question `json:"questions"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type QuizCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
}


type AnswerCreateRequest struct {
	Answer      string `json:"answer" validate:"required"`
	IsCorrect     bool `json:"is_correct"`
}

type QuestionCreateRequest struct {
	Content    string `json:"content" validate:"required"`
	Options	[]AnswerCreateRequest `json:"options"`
}

func GetQuestions(tx *sql.Tx, quizID string) ([]Question, error) {
	rows, err := tx.Query("SELECT ID, Content, QuizID, CreatedAt, UpdatedAt FROM `Question` WHERE QuizID = ?", quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	questions := make([]Question, 0)
	for rows.Next() {
		question := new(Question)
		err := rows.Scan(&question.ID, &question.Content, &question.QuizID, &question.CreatedAt, &question.UpdatedAt)
		if err != nil {
			return nil, err
		}
		questions = append(questions, *question)
	}

	for i, question := range questions {
		answers, err := GetAnswers(tx, question.ID)
		if err != nil {
			return nil, err
		}
		
		questions[i].Options = answers
	}

	return questions, nil
}

func GetAnswers(tx *sql.Tx, questionID string) ([]Answer, error) {
	rows, err := tx.Query("SELECT ID, Answer, IsCorrect, CreatedAt, UpdatedAt FROM `Answer` WHERE QuestionID = ?", questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	answers := make([]Answer, 0)
	for rows.Next() {
		answer := new(Answer)
		err := rows.Scan(&answer.ID, &answer.Answer, &answer.IsCorrect, &answer.CreatedAt, &answer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		answers = append(answers, *answer)
	}

	return answers, nil
}

func GetQuiz(tx *sql.Tx, id string) (*Quiz, error) {
	quiz := new(Quiz)
	err := tx.QueryRow("SELECT ID, Name, Description, Image, CreatedAt, UpdatedAt FROM `Quiz` WHERE ID = ?", id).Scan(&quiz.ID, &quiz.Name, &quiz.Description, &quiz.Image, &quiz.CreatedAt, &quiz.UpdatedAt)
	if err != nil {
		return nil, err
	}

	questions, err := GetQuestions(tx, id)
	if err != nil {
		return nil, err
	}
	
	quiz.Questions = questions

	return quiz, nil
}

func GetQuizzes(tx *sql.Tx) ([]*Quiz, error) {
	rows, err := tx.Query("SELECT ID, Name, Description, Image, CreatedAt, UpdatedAt FROM `Quiz`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quizzes := make([]*Quiz, 0)
	for rows.Next() {
		quiz := new(Quiz)
		err := rows.Scan(&quiz.ID, &quiz.Name, &quiz.Description, &quiz.Image, &quiz.CreatedAt, &quiz.UpdatedAt)
		if err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}

	for _, quiz := range quizzes {
		questions, err := GetQuestions(tx, quiz.ID)
		if err != nil {
			return nil, err
		}

		quiz.Questions = questions
	}

	return quizzes, nil
}

func NewQuiz(name string, description string, image string, questions []Question) *Quiz {
	return &Quiz{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Image:       image,
		Questions:   questions,
	}
}

func CreateQuiz(tx *sql.Tx, quiz *QuizCreateRequest) (error, *Quiz) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)

	_, err := tx.Exec("INSERT INTO `Quiz` (ID, Name, Description, Image, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?)", id, quiz.Name, quiz.Description, quiz.Image, createdAt, updatedAt)
	if err != nil {
		return err, nil
	}

	// return full quiz
	return nil, &Quiz{
		ID:          id,
		Name:        quiz.Name,
		Description: quiz.Description,
		Questions:  []Question{},
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func CreateAnswer(tx *sql.Tx, questionID string , answer *AnswerCreateRequest) (error, *Answer) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)

	_, err := tx.Exec("INSERT INTO `Answer` (ID, Answer, IsCorrect, QuestionID, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?)", id, answer.Answer, answer.IsCorrect, questionID, createdAt, updatedAt)
	if err != nil {
		return err, nil
	}

	// return full answer
	return nil, &Answer{
		ID:          id,
		Answer:    answer.Answer,
		IsCorrect: answer.IsCorrect,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func CreateQuestion(tx *sql.Tx, quizID string , question *QuestionCreateRequest) (error, *Question) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)

	_, err := tx.Exec("INSERT INTO `Question` (ID, Content, QuizID, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?)", id, question.Content, quizID, createdAt, updatedAt)
	if err != nil {
		return err, nil
	}

	var answers []Answer
	for _, answer := range question.Options {
		err, answer := CreateAnswer(tx, id, &answer)
		if err != nil {
			return err, nil
		}

		answers = append(answers, *answer)
	}

	// return full question
	return nil, &Question{
		ID:          id,
		Content:    question.Content,
		QuizID: quizID,
		Options: answers,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
