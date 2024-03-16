package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Tier struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	MinPoints   int    `json:"min_points" validate:"required"`
	MaxPoints   int    `json:"max_points" validate:"required"`
	Benefits    string `json:"benefits"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TierCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	MinPoints   int    `json:"min_points" validate:"required"`
	MaxPoints   int    `json:"max_points" validate:"required,gtfield=MinPoints"`
	Benefits    string `json:"benefits"`
}

type TierUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	MinPoints   int    `json:"min_points"`
	MaxPoints   int    `json:"max_points" validate:"gtfield=MinPoints"`
	Benefits    string `json:"benefits"`
}

func GetTier(tx *sql.Tx, id string) (*Tier, error) {
	tier := new(Tier)
	err := tx.QueryRow("SELECT ID, Name, Description, MinPoints, MaxPoints, Benefits, CreatedAt, UpdatedAt FROM `Tier` WHERE ID = ?", id).Scan(&tier.ID, &tier.Name, &tier.Description, &tier.MinPoints, &tier.MaxPoints, &tier.Benefits, &tier.CreatedAt, &tier.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return tier, nil
}

func GetTiers(tx *sql.Tx) ([]*Tier, error) {
	rows, err := tx.Query("SELECT ID, Name, Description, MinPoints, MaxPoints, Benefits, CreatedAt, UpdatedAt FROM `Tier`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tiers := make([]*Tier, 0)
	for rows.Next() {
		tier := new(Tier)
		err := rows.Scan(&tier.ID, &tier.Name, &tier.Description, &tier.MinPoints, &tier.MaxPoints, &tier.Benefits, &tier.CreatedAt, &tier.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tiers = append(tiers, tier)
	}

	return tiers, nil
}

func CreateTier(tx *sql.Tx, tier *TierCreateRequest) (*Tier, error) {
	id := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)

	_, err := tx.Exec("INSERT INTO `Tier` (ID, Name, Description, MinPoints, MaxPoints, Benefits, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", id, tier.Name, tier.Description, tier.MinPoints, tier.MaxPoints, tier.Benefits, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}

	return &Tier{
		ID:          id,
		Name:        tier.Name,
		Description: tier.Description,
		MinPoints:   tier.MinPoints,
		MaxPoints:   tier.MaxPoints,
		Benefits:    tier.Benefits,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func UpdateTier(tx *sql.Tx, id string, tier *TierUpdateRequest) (*Tier, error) {
	updatedAt := time.Now().Format(time.RFC3339)

	_, err := tx.Exec("UPDATE `Tier` SET Name = ?, Description = ?, MinPoints = ?, MaxPoints = ?, Benefits = ?, UpdatedAt = ? WHERE ID = ?", tier.Name, tier.Description, tier.MinPoints, tier.MaxPoints, tier.Benefits, updatedAt, id)
	if err != nil {
		return nil, err
	}

	return GetTier(tx, id)
}

func DeleteTier(tx *sql.Tx, id string) error {
	_, err := tx.Exec("DELETE FROM `Tier` WHERE ID = ?", id)
	return err
}