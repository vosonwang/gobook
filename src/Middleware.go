package main

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"fmt"
	"config"
	"etcd"
)

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.SecretKey), nil
		})

	if err == nil {
		claims := token.Claims.(jwt.MapClaims)
		username := claims["sub"].(string)
		b := etcd.Get(username)
		if token.Valid && b == nil {
			etcd.Set(username, "")
			next(w, r)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, 0)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}
