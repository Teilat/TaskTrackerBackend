package helpers

import (
	"log"
	"main/internal/server/api/v1/models"
	"strings"
)

func CheckUserPass(credentials models.Login) bool {
	userpass := make(map[string]string)
	userpass["hello"] = "itsme"
	userpass["john"] = "doe"
	userpass["user"] = "user"

	log.Println("checkUserPass", credentials.Username, credentials.Password, userpass)

	if val, ok := userpass[credentials.Username]; ok {
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
