package helpers

import (
	"log"
	"main/database/sql"
	"main/server/api/v1/models"
	"strings"
)

func CheckUserPass(credentials models.Login) bool {
	db := sql.GetDb()
	users, err := db.GetUsersCredentials()
	if err != nil {
		return false
	}
	userPass := make(map[string]string)
	for _, user := range users {
		userPass[user.Nickname] = user.Password
	}

	log.Println("checkUserPass", credentials.Username, credentials.Password, userPass)

	if val, ok := userPass[credentials.Username]; ok {
		log.Println(val, ok)
		if val == credentials.Password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func EmptyUserPass(credentials models.Login) bool {
	return strings.Trim(credentials.Username, " ") == "" || strings.Trim(credentials.Password, " ") == ""
}
