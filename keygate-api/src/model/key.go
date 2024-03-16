package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type Key struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
	Permissions []Permission `json:"permissions"`
	Status string `json:"-"`
}

type Permission struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
}

type UserKeyAttachment struct {
	KeyID string `json:"key_id" validate:"required"`
}

type KeyPermissionAttachment struct {
	PermissionID string `json:"permission_id" validate:"required"`
}

func NewKey (name string, permissions []Permission, status string) Key {
	return Key{
		Name: name,
		Permissions: permissions,
		Status: status,
	}
}

func CreateKey (tx *sql.Tx, key *Key) error {
	ID := uuid.New().String()
	_, err := tx.Exec("INSERT INTO `Key` (ID, Name) VALUES (?, ?)", ID, key.Name)

	for _, permission := range key.Permissions {
		_, err = tx.Exec("INSERT INTO KeyPermission (KeyID, PermissionID) VALUES (?, ?)", ID, permission.ID)
	}

	key.ID = ID

	return err
}

func GetAttachedPermissions (tx *sql.Tx, keyID string) ([]Permission, error) {
	rows, err := tx.Query("SELECT p.ID, p.Name FROM Permission p JOIN KeyPermission kp ON p.ID = kp.PermissionID WHERE kp.KeyID = ?", keyID)
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

func GetKeys (tx *sql.Tx) ([]Key, error) {
	rows, err := tx.Query("SELECT ID, Name, status FROM `Key`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []Key = make([]Key, 0)
	for rows.Next() {
		var key Key
		err := rows.Scan(&key.ID, &key.Name, &key.Status)
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

func DeleteKey (tx *sql.Tx, id string) error {
	_, err := tx.Exec("DELETE FROM `Key` WHERE ID = ?", id)
	return err
}

func DoesKeyExist (tx *sql.Tx, id string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM `Key` WHERE id = ?", id).Scan(&count)
	return count > 0, err
}

func DoesKeyExistByName (tx *sql.Tx, name string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM `Key` WHERE name = ?", name).Scan(&count)
	return count > 0, err
}

func CreatePermission (tx *sql.Tx, permission Permission) error {
	ID := uuid.New().String()
	_, err := tx.Exec("INSERT INTO Permission (ID, Name) VALUES (?, ?)", ID, permission.Name)
	permission.ID = ID
	return err
}

func GetPermissions (tx *sql.Tx) ([]Permission, error) {
	rows, err := tx.Query("SELECT * FROM Permission")
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

func DeletePermission (tx *sql.Tx, id string) error {
	_, err := tx.Exec("DELETE FROM Permission WHERE ID = ?", id)
	return err
}

func DoesPermissionExist (tx *sql.Tx, id string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM Permission WHERE ID = ?", id).Scan(&count)
	return count > 0, err
}

func DoesPermissionExistByName (tx *sql.Tx, name string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM Permission WHERE Name = ?", name).Scan(&count)
	return count > 0, err
}

func GrantKeyToUser (tx *sql.Tx, keyID string, userID string) error {
	_, err := tx.Exec("INSERT INTO UserKey (UserID, KeyID) VALUES (?, ?)", userID, keyID)
	return err
}

func AddPermissionToKey (tx *sql.Tx, keyID string, permissionID string) error {
	_, err := tx.Exec("INSERT INTO KeyPermission (KeyID, PermissionID) VALUES (?, ?)", keyID, permissionID)
	return err
}

func DoesKeyContainPermission (tx *sql.Tx, keyID string, permissionID string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM KeyPermission WHERE KeyID = ? AND PermissionID = ?", keyID, permissionID).Scan(&count)
	return count > 0, err
}

func GetKeyPermissions (tx *sql.Tx, keyID string) ([]Permission, error) {
	rows, err := tx.Query("SELECT p.ID, p.Name FROM Permission p JOIN KeyPermission kp ON p.ID = kp.PermissionID WHERE kp.KeyID = ?", keyID)
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

func CheckUserPermissionByName (tx *sql.Tx, userID string, name string) (bool, error) {
	log.Info("Checking permission for user ", userID, " with name ", name)
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM UserKey uk JOIN KeyPermission kp ON uk.KeyID = kp.KeyID JOIN Permission p ON kp.PermissionID = p.ID WHERE uk.UserID = ? AND p.Name = ?", userID, name).Scan(&count)
	return count > 0, err
}

func CheckUserPermission (tx *sql.Tx, userID string, permissionID string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM UserKey uk JOIN KeyPermission kp ON uk.KeyID = kp.KeyID WHERE uk.UserID = ? AND kp.PermissionID = ?", userID, permissionID).Scan(&count)
	return count > 0, err
}

func CheckUserPermissionByExternalID (tx *sql.Tx, externalID string, permissionID string) (bool, error) {
	var count int
	err := tx.QueryRow("SELECT COUNT(*) FROM UserKey uk JOIN KeyPermission kp ON uk.KeyID = kp.KeyID WHERE uk.UserID = (SELECT ID FROM User WHERE ExternalID = ?) AND kp.PermissionID = ?", externalID, permissionID).Scan(&count)
	return count > 0, err
}