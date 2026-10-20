package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"a.com/test/internal/database"
	logx "a.com/test/internal/httpx/logs"
	"a.com/test/internal/httpx/middlewares"
	"a.com/test/internal/list"
)

func main() {
	logx.InitLogger()

	if err := godotenv.Load(); err != nil {
		logx.Error("MAIN_ENV_EXCEPTION", "err", err)
	}

	ctx := context.Background()

	_, dbErr := database.Open(ctx)
	if dbErr != nil {
		logx.Error("MAIN_DATABASE_ERROR", "err", dbErr)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /list", middlewares.Config(list.ReadAllListHandler))
	mux.HandleFunc("POST /list", middlewares.Config(list.SaveListHandler))
	mux.HandleFunc("PUT /list/{id}", middlewares.Config(list.UpdateListHandler))
	mux.HandleFunc("DELETE /list/{id}", middlewares.Config(list.DeleteListHandler))

	logx.InfoString("MAIN", "running")
	server := http.Server{
		Addr:              ":8000",
		Handler:           mux,
		ReadHeaderTimeout: 60 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logx.Error("MAIN_SERVER_ERROR", "err", err)
		os.Exit(1)
	}
}
