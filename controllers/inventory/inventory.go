package inventory

import (
	"backend-challenge-2022/models"
	"backend-challenge-2022/utils/request"
	"backend-challenge-2022/utils/response"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func View(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "itemID")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "Missing item ID.")
		return
	}

	inv := &models.Inventory{
		ID: id,
	}

	if err := inv.Get(); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, inv)
}

func ViewAll(w http.ResponseWriter, r *http.Request) {
	allInv, err := models.GetAllInventory()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, allInv)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "itemID")
	var body map[string]string

	if err := request.ParseBody(r, &body); err != nil || id == "" || body["reason"] == "" {
		response.Error(w, http.StatusBadRequest, "Missing item ID or reason.")
		return
	}

	inv := &models.Inventory{
		ID: id,
	}

	if err := inv.Delete(body["reason"]); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.String(w, http.StatusAccepted, fmt.Sprintf("%s was deleted.", inv.ID))
}

func UndoDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "itemID")
	var body map[string]string

	if err := request.ParseBody(r, &body); err != nil || id == "" || body["reason"] == "" {
		response.Error(w, http.StatusBadRequest, "Missing item ID or reason.")
		return
	}

	inv := &models.Inventory{
		ID: id,
	}

	if err := inv.UndoDelete(); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, inv)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "itemID")
	var updated models.Inventory

	if err := request.ParseBody(r, &updated); err != nil || id == "" {
		response.Error(w, http.StatusBadRequest, "Missing item ID.")
		return
	}

	inv := &models.Inventory{
		ID: id,
	}

	if err := inv.Get(); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if updated.Title != "" {
		inv.Title = updated.Title
	}

	if updated.Description != "" {
		inv.Description = updated.Description
	}

	if updated.Price != 0 {
		inv.Price = updated.Price
	}

	if updated.Quantity != 0 {
		inv.Quantity = updated.Quantity
	}

	if err := inv.Update(); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusAccepted, inv)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var inv models.Inventory

	if err := request.ParseBody(r, &inv); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	inv.ID = uuid.New().String()
	inv.ID = strings.Replace(inv.ID, "-", "", -1)
	if err := inv.Create(); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, inv)
}
