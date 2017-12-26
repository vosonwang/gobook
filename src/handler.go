package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"models"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"config"
	"util"
	"etcd"
	"strconv"
)

type Token struct {
	Token string `json:"token"`
}

func TocsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(2); err != nil {
		panic(err)
	}
}

func PointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(id); err != nil {
		panic(err)
	}
}

func NewToken(w http.ResponseWriter, r *http.Request) {
	user := models.ParseUser(r.Body)
	if models.FindUser(user) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * time.Duration(8)).Unix(), //8小时后过期
			"iat": time.Now().Unix(),
			"sub": user.Username,
		})

		tokenString, err := token.SignedString([]byte(config.SecretKey))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
			util.CheckErr(err)
		}

		response := Token{tokenString}
		etcd.Set(user.Username, "")
		JsonResponse(response, w)

	} else {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		fmt.Fprintln(w,0)
	}

}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetNodes(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	kind := vars["kind"]
	a,err:=strconv.Atoi(kind)
	util.CheckErr(err)
	b,c:=models.GetNode(a)
	util.CheckErr(c)
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if d := json.NewEncoder(w).Encode(b); err != nil {
		panic(d)
	}

}