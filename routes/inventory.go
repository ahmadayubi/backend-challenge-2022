package routes

import (
	"backend-challenge-2022/controllers/inventory"
	"github.com/go-chi/chi/v5"
)

func InventoryRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Put("/create", inventory.Create)
	router.Put("/edit/{itemID}", inventory.Edit)
	router.Delete("/delete/{itemID}", inventory.Delete)
	router.Put("/undo-delete/{itemID}", inventory.UndoDelete)
	router.Get("/view/all", inventory.ViewAll)
	router.Get("/view/{itemID}", inventory.View)

	return router
}
