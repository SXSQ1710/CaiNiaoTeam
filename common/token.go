package common

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

var setTime = SetTime

//加密token
func TokenEncrypt(user_id string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Unix() + setTime,
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err.Error())
		return tokenString
	} else {
		return tokenString
	}
}

//解密token
func TokenParse(tokenString string) interface{} {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return mySigningKey, nil
	})

	user_id := token.Claims.(jwt.MapClaims)["user_id"]

	if err != nil {
		log.Println(err.Error())
		return user_id
	}
	return user_id
}
