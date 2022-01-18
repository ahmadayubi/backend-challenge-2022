package main

import (
	"backend-challenge-2022/routes"
	"backend-challenge-2022/utils/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func CreateRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Throttle(20), // limit due to connections the database allows
	)
	router.Route("/", func(r chi.Router) {
		r.Mount("/inventory", routes.InventoryRoutes())
	})
	return router
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env File")
		return
	}
	router := CreateRoutes()
	err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error Connecting To Database")
		return
	}
	walkF := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s, %s\n", method, route)
		return nil
	}
	if err = chi.Walk(router, walkF); err != nil {
		log.Fatalf("Logging Error: %s", err.Error())
	}
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
