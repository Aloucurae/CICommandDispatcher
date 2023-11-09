package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"localDeployer/src/release"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"poseur.com/dotenv"
)

var envfile = flag.String("env", ".env", "environment file")

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		if request.RequestURI == "/check" {
			next.ServeHTTP(response, request)
			return
		}

		SERVER_TOKEN := "Bearer " + os.Getenv("SERVER_TOKEN")

		if SERVER_TOKEN != "" {

			token := request.Header.Get("Authorization")

			if SERVER_TOKEN != token {
				http.Error(response, "não autorizado", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(response, request)
	})
}

type msg struct {
	Ok string `json:"ok"`
}

func checkHandler(response http.ResponseWriter, request *http.Request) {

	event := msg{Ok: "ok"}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(event)
}

func main() {
	_ = dotenv.SetenvFile(*envfile)

	router := mux.NewRouter()

	router.Use(AuthMiddleware)

	router.HandleFunc("/check", checkHandler).Methods("GET")

	router.HandleFunc("/release", release.Handler).Methods("POST")

	http.Handle("/", router)

	SERVER_PORT := os.Getenv("SERVER_PORT")

	if SERVER_PORT == "" {
		SERVER_PORT = "80"
	}

	fmt.Printf("got SERVER_PORT \t %v \n", SERVER_PORT)

	http.ListenAndServe(":"+SERVER_PORT, nil)
}
