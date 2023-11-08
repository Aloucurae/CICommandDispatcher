package main

import (
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

func main() {
	_ = dotenv.SetenvFile(*envfile)

	release.CodeRunner("datafrete/datafrete:v2.14.02")

	return

	router := mux.NewRouter()

	router.Use(AuthMiddleware)

	router.HandleFunc("/release", release.Handler).Methods("POST")

	http.Handle("/", router)

	fmt.Printf("got SERVER_PORT \t %v \n", os.Getenv("SERVER_PORT"))

	SERVER_PORT := os.Getenv("SERVER_PORT")

	http.ListenAndServe(":"+SERVER_PORT, nil)
}
