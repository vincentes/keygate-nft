package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Collectible struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Image string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CollectibleCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Image string `json:"image"`
}

func NewCollectible (name string, description string, image string) *Collectible {
	return &Collectible{
		ID: uuid.New().String(),
		Name: name,
		Description: description,
		Image: image,
	}
}

func CreateCollectible (tx *sql.Tx, collectible *CollectibleCreateRequest) (error, *Collectible) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)
	
	_, err := tx.Exec("INSERT INTO `Collectible` (ID, Name, Description, Image, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?)", id, collectible.Name, collectible.Description, collectible.Image, createdAt, updatedAt)
	if err != nil {
		return err, nil
	}

	// return full collectible
	return nil, &Collectible{
		ID: id,
		Name: collectible.Name,
		Description: collectible.Description,
		Image: collectible.Image,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func GetCollectibles (tx *sql.Tx) ([]*Collectible, error) {
	rows, err := tx.Query("SELECT * FROM `Collectible`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collectibles := make([]*Collectible, 0)
	for rows.Next() {
		collectible := new(Collectible)
		err := rows.Scan(&collectible.ID, &collectible.Name, &collectible.Description, &collectible.Image, &collectible.CreatedAt, &collectible.UpdatedAt)
		if err != nil {
			return nil, err
		}
		collectibles = append(collectibles, collectible)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return collectibles, nil
}