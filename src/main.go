package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"time"
	"log"
	"config"

)

func main() {

	/*Router*/
	r := mux.NewRouter()
	adminRouter := mux.NewRouter().PathPrefix("/admin").Subrouter().StrictSlash(true)

	r.HandleFunc("/tocs", TocsHandler).Methods("GET")
	r.HandleFunc("/tocs/{id}", TocsHandler).Methods("GET")
	r.HandleFunc("/articles/{id}", TocsHandler).Methods("GET")
	r.HandleFunc("/images/{id}", TocsHandler).Methods("GET")
	r.HandleFunc("/files/{id}", TocsHandler).Methods("GET")
	r.HandleFunc("/tokens", NewToken).Methods("POST")
	r.HandleFunc("/nodes/{kind}", GetNodes).Methods("GET")


	adminRouter.HandleFunc("/{id}", TocsHandler).Methods("GET")

	adminRouter.HandleFunc("/tocs", TocsHandler).Methods("POST")
	adminRouter.HandleFunc("/tocs/{id}", PointHandler).Methods("PUT", "DELETE", "PATCH")

	adminRouter.HandleFunc("/articles", TocsHandler).Methods("POST")
	adminRouter.HandleFunc("/articles/{id}", PointHandler).Methods("PUT", "DELETE", "PATCH")

	adminRouter.HandleFunc("/images", TocsHandler).Methods("POST")

	adminRouter.HandleFunc("/files", TocsHandler).Methods("POST")

	adminRouter.HandleFunc("/users/{id}", TocsHandler).Methods("GET","PATCH","DELETE")
	adminRouter.HandleFunc("/users", TocsHandler).Methods("POST")



	// Create a new negroni for the admin middleware
	r.PathPrefix("/admin").Handler(negroni.New(
		negroni.HandlerFunc(ValidateToken),
		negroni.Wrap(adminRouter),
	))

	srv := &http.Server{
		Addr:    config.WebPort,
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())

}
