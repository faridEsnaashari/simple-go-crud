package main

import (
	"fmt"
	"log"
	"net/http"

	"a.com/test/internal/database"
	"a.com/test/internal/httpx/middlewares"
	"a.com/test/internal/list"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	database.Open(database.GetDatabaseConfigs())

	mux.HandleFunc("GET /list", middlewares.Config(list.ReadAllListHandler))
	mux.HandleFunc("POST /list", middlewares.Config(list.SaveListHandler))
	mux.HandleFunc("PUT /list/{id}", middlewares.Config(list.UpdateListHandler))
	mux.HandleFunc("DELETE /list/{id}", middlewares.Config(list.DeleteListHandler))

	fmt.Println("running")
	fmt.Println(database.GetDatabaseConfigs())
	err := http.ListenAndServe(":8000", mux)
	fmt.Println(err)

}
