// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Cloud Security Client Go contributors
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/sap/cloud-security-client-go/auth"
	"github.com/sap/cloud-security-client-go/env"
)

// Main class for demonstration purposes.
func main() {
	r := mux.NewRouter()

	config, err := env.ParseIdentityConfig()
	if err != nil {
		panic(err)
	}
	authMiddleware := auth.NewMiddleware(config, auth.Options{})
	r.Use(authMiddleware.AuthenticationHandler)

	r.HandleFunc("/helloWorld", helloWorld).Methods(http.MethodGet)

	address := ":8080"
	log.Println("Starting server on address", address)
	err = http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, r))
	if err != nil {
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	user := auth.TokenFromCtx(r)
	_, _ = fmt.Fprintf(w, "Hello world!\nYou're logged in as %s", user.Email())
}
