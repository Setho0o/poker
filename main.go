package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"web_proj/handler"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
  if err := InitEverything(); err != nil {
    log.Fatal(err)
  }
  port := os.Getenv("HTTP_LISTEN_ADDR")
  slog.Info("application running", "port", port)

  r := chi.NewMux()
  r.Get("/", handler.HandlerHomeIndex) 
	log.Fatal(http.ListenAndServe(port, r))
  
}
func InitEverything() error {
  if err := godotenv.Load(); err != nil {
    return err
  }
  return nil 
}
