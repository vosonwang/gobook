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
	"github.com/satori/go.uuid"
)

type Token struct {
	Token    string    `json:"token"`
	Username string    `json:"username"`
	UserId   uuid.UUID `json:"userId"`
}

func NewNode(w http.ResponseWriter, r *http.Request) {
	if node, err := models.ParseNode(r.Body); err == nil {
		if id, err := models.AddNode(node); err == nil {
			w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(id); err != nil {
				panic(err)
			}
		} else {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "新建节点错误！")
		}
	} else {
		fmt.Print(err)
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "内容解析错误！")
		ErrorResponse(err, w, http.StatusBadRequest)
	}

}

func UpdateNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	switch {
	case r.Method == "PATCH":
		if node, err := models.FindNode(id); err == nil {
			if m, err := models.ParseNode(r.Body); err == nil {
				node.Title = m["title"].(string)
				if err := models.UpdateNode(node); err == nil {
					JsonResponse(1, w)
				} else {
					ErrorResponse(err, w, http.StatusBadRequest)
				}
			} else {
				ErrorResponse(err, w, http.StatusBadRequest)
			}
		} else {
			ErrorResponse(err, w, http.StatusBadRequest)
		}

	case r.Method == "DELETE":
		if node, err := models.FindNode(id); err == nil {
			if err := models.DeleteNode(node); err == nil {
				JsonResponse(1, w)
			} else {

				ErrorResponse(err, w, http.StatusBadRequest)
			}
		} else {
			w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}

	case r.Method == "PUT":

	}
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
	if a, b := models.FindUser(user); b != false {
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

		response := Token{tokenString, a.Username, a.ID}
		etcd.Set(user.Username, "")
		JsonResponse(response, w)

	} else {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		fmt.Fprintln(w, 0)
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

func ErrorResponse(err error, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(code)
	fmt.Fprintln(w, err)
}

func GetNodes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kind := vars["kind"]
	a, err := strconv.Atoi(kind)
	util.CheckErr(err)
	b, c := models.GetNode(a)
	util.CheckErr(c)
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if d := json.NewEncoder(w).Encode(b); err != nil {
		panic(d)
	}

}

func ExchangeNode(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

}
