package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Token struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Address string `json:"address"`
}

func NewToken (name string, description string) Token {
	return Token{
		Name: name,
		Description: description,
	}
}

func CreateToken (tx *sql.Tx, token *Token) error {
	ID := uuid.New().String()
	_, err := tx.Exec("INSERT INTO `Token` (ID, Name, Description) VALUES (?, ?, ?)", ID, token.Name, token.Description)
	token.ID = ID
	return err
}
