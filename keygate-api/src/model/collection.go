package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Image string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CollectionCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Image string `json:"image"`
}

func NewCollection (name string, description string, image string) *Collection {
	return &Collection{
		ID: uuid.New().String(),
		Name: name,
		Description: description,
		Image: image,
	}
}

func CreateCollection (tx *sql.Tx, collection *CollectionCreateRequest) (error, *Collection) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)
	
	_, err := tx.Exec("INSERT INTO `Collection` (ID, Name, Description, Image, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?)", id, collection.Name, collection.Description, collection.Image, createdAt, updatedAt)
	if err != nil {
		return err, nil
	}

	// return full collection
	return nil, &Collection{
		ID: id,
		Name: collection.Name,
		Description: collection.Description,
		Image: collection.Image,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func GetCollections (tx *sql.Tx) ([]*Collection, error) {
	rows, err := tx.Query("SELECT * FROM `Collection`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := make([]*Collection, 0)
	for rows.Next() {
		collection := new(Collection)
		err := rows.Scan(&collection.ID, &collection.Name, &collection.Description, &collection.Image, &collection.CreatedAt, &collection.UpdatedAt)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return collections, nil
}

func DeleteCollection(tx *sql.Tx, collectionID string) error {
    _, err := tx.Exec("DELETE FROM `Collection` WHERE ID = ?", collectionID)
    if err != nil {
        return err
    }

    return nil
}