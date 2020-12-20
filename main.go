/*
 * Swagger Blog
 *
 * Simple Blog
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"
	sw "github.com/610masters/Backend/go"
)

func main() {
	log.Printf("Backend started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
