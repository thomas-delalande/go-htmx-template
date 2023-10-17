package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go-htmx-template/src/templates"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Compress(5))

	router.Get("/", helloWorld)
	value, exists := os.LookupEnv("PORT")
	if !exists {
		value = "8080"
	}

	server := &http.Server{Addr: fmt.Sprintf(":%s", value), Handler: router}
	log.Default().Println(fmt.Sprintf("Listening on port %s...", value))
	server.ListenAndServe()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	server.Shutdown(context.Background())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	component := templates.HelloWorld()
	component.Render(context.Background(), w)
}
