package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID string `json:"id"`
	ExternalID string `json:"external_id" validate:"required"`
	Points int `json:"points"`
}

func NewUser (externalID string) User {
	return User{
		ExternalID: externalID,
	}
}

func CreateUser (tx *sql.Tx, user *User) error {
	ID := uuid.New().String()
	_, err := tx.Exec("INSERT INTO User (ID, ExternalID) VALUES (?, ?)", ID, user.ExternalID)

	user.ID = ID
	return err
}

func DoesUserExist (tx *sql.Tx, id string) (bool, error) {
	rows, err := tx.Query("SELECT ID FROM User WHERE ID = ?", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func DoesUserExistByExternalID (tx *sql.Tx, externalID string) (bool, error) {
	rows, err := tx.Query("SELECT ID FROM User WHERE ExternalID = ?", externalID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func GetUserByExternalID (tx *sql.Tx, externalID string) (*User, error) {
	user := new(User)
	row := tx.QueryRow("SELECT ID, ExternalID FROM User WHERE ExternalID = ?", externalID)
	if row.Err() != nil {
		return nil, row.Err()
	}
	
	err := row.Scan(&user.ID, &user.ExternalID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return user, err
}

func GetUsers (tx *sql.Tx) ([]User, error) {
	rows, err := tx.Query("SELECT ID, ExternalID FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User = make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.ExternalID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func DeleteUser (tx *sql.Tx, userID string) error {
	_, err := tx.Exec("DELETE FROM User WHERE ID = ?", userID)
	return err
}

func IsKeyAttachedToUser (tx *sql.Tx, keyID string, userID string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM UserKey WHERE KeyID = ? AND UserID = ?", keyID, userID).Scan(&count)
	return count > 0, err
}

func AttachKeyToUser (tx *sql.Tx, keyID string, userID string) error {
	attachmentID := uuid.New().String()
	_, err := tx.Exec("INSERT INTO UserKey (ID, UserID, KeyID) VALUES (?, ?, ?)", attachmentID, userID, keyID)
	return err
}

func GetUserPermissions (tx *sql.Tx, userID string) ([]Permission, error) {
	rows, err := tx.Query("SELECT p.ID, p.Name FROM Permission p JOIN KeyPermission kp ON p.ID = kp.PermissionID JOIN UserKey uk ON kp.KeyID = uk.KeyID WHERE uk.UserID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []Permission = make([]Permission, 0)
	for rows.Next() {
		var permission Permission
		err := rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func GetUserKeys (tx *sql.Tx, userID string) ([]Key, error) {
	rows, err := tx.Query("SELECT k.ID, k.Name FROM `Key` k JOIN UserKey uk ON k.ID = uk.KeyID WHERE uk.UserID = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []Key = make([]Key, 0)
	for rows.Next() {
		var key Key
		err := rows.Scan(&key.ID, &key.Name)
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}

	for i, key := range keys {
		permissions, err := GetAttachedPermissions(tx, key.ID)
		if err != nil {
			return nil, err
		}
		keys[i].Permissions = permissions
	}

	return keys, nil
}