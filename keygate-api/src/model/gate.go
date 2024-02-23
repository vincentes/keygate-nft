package model


type Gate struct {
	ID string `json:"id"`
	Name string `json:"name" validate:"required"`
	RequiredPermissions []Permission `json:"required_permissions" validate:"required"`
}