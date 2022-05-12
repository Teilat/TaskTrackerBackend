package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"       example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	RoleId   uuid.UUID `json:"roleId"   example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name     string    `json:"name"     example:"John"`
	Surname  string    `json:"surname"  example:"Joe"`
	Nickname string    `json:"nickname" example:"Nickname"`
}

type UserCredentials struct {
	Id       uuid.UUID `json:"id"       example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Nickname string    `json:"nickname" example:"Nickname"`
	Password string    `json:"password" example:"password"`
}

type AddUser struct {
	RoleId   uuid.UUID `json:"projectId" form:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name     string    `json:"name"     example:"John"`
	Surname  string    `json:"surname"  example:"Doe"`
	Nickname string    `json:"nickname" example:"Joe"`
	Password string    `json:"password" example:"password"`
}

type UpdateUser struct {
	Id       uuid.UUID `json:"id"       example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	RoleId   uuid.UUID `json:"projectId" form:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name     string    `json:"name"     example:"John"`
	Surname  string    `json:"surname"  example:"Joe"`
	Nickname string    `json:"nickname" example:"Nickname"`
}

type DeleteUser struct {
	Id uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type UpdateUserPassword struct {
	Id       uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Password string    `json:"password" example:"password"`
}
