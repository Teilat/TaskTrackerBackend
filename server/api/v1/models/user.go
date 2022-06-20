package models

type User struct {
	Id       int32  `json:"id"       example:"15" format:"integer"`
	RoleId   int32  `json:"roleId"   example:"2" format:"integer"`
	Name     string `json:"name"     example:"John"`
	Surname  string `json:"surname"  example:"Joe"`
	Nickname string `json:"nickname" example:"Nickname"`
}

type UserCredentials struct {
	Id       int32  `json:"id"       example:"15" format:"integer"`
	Nickname string `json:"nickname" example:"Nickname"`
	Password string `json:"password" example:"password"`
}

type AddUser struct {
	RoleId   int32  `json:"roleId" form:"roleId" example:"2" format:"integer"`
	Name     string `json:"name" form:"name" example:"John"`
	Surname  string `json:"surname" form:"surname" example:"Doe"`
	Nickname string `json:"nickname" form:"nickname" example:"Joe"`
	Password string `json:"password" form:"password" example:"password"`
}

type UpdateUser struct {
	Id       int32  `json:"id"       example:"15" format:"integer"`
	RoleId   int32  `json:"roleId" form:"roleId" example:"2" format:"integer"`
	Name     string `json:"name"     example:"John"`
	Surname  string `json:"surname"  example:"Joe"`
	Nickname string `json:"nickname" example:"Nickname"`
}

type DeleteUser struct {
	Id int32 `json:"id" example:"15" format:"integer"`
}

type UpdateUserPassword struct {
	Id       int32  `json:"id" example:"15" format:"integer"`
	Password string `json:"password" example:"password"`
}
