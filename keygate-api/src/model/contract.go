package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type status string
type address string

type Contract struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Address address `json:"address"`
	Type string `json:"contract_type" validate:"required"`
	Status status `json:"status"`
}

func (a *status) UnmarshalJSON(data []byte) error {
	a = nil
	return nil
}

func (a *address) UnmarshalJSON(data []byte) error {
	a = nil
	return nil
}

func NewContract (name string, description string, contractType string, status status) Contract {
	return Contract{
		Name: name,
		Description: description,
		Type: contractType,
		Status: status,
		Address: "0x056192239861881bC3d3edb668B32BF3c976da8e",
	}
}

func CreateContract (tx *sql.Tx, contract *Contract) error {
	ID := uuid.New().String()
	_, err := tx.Exec("INSERT INTO `Contract` (ID, Name, Description, Type, Status, Address) VALUES (?, ?, ?, ?, ?, ?)", ID, contract.Name, contract.Description, contract.Type, contract.Status, contract.Address)
	contract.ID = ID
	return err
}

func GetContracts (tx *sql.Tx) ([]Contract, error) {
	rows, err := tx.Query("SELECT ID, Name, Description, Type, Status, Address FROM `Contract`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []Contract = make([]Contract, 0)
	for rows.Next() {
		var contract Contract
		err := rows.Scan(&contract.ID, &contract.Name, &contract.Description, &contract.Type, &contract.Status, &contract.Address)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}