package models

import (
	"backend-challenge-2022/utils/database"
	"context"
	"errors"
)

type Inventory struct {
	ID           string  `json:"id,omitempty"`
	Title        string  `json:"title,omitempty"`
	Description  string  `json:"description,omitempty"`
	Quantity     int64   `json:"quantity,omitempty"`
	Price        float64 `json:"price,omitempty"`
	DeleteReason string  `json:"deleteReason,omitempty"`
	Deleted      bool    `json:"deleted,omitempty"`
}

const newInv = "INSERT INTO items (id, title, item_description, quantity, price, item_deleted) VALUES ($1, $2, $3, $4, $5, false)"

func (i *Inventory) Create() error {
	_, err := database.DB.Exec(context.Background(), newInv, i.ID, i.Title, i.Description, i.Quantity, i.Price)
	return err
}

const getInv = "SELECT * FROM items WHERE id=$1"

func (i *Inventory) Get() error {
	if i.ID == "" {
		return errors.New("missing ID")
	}
	if err := database.DB.QueryRow(context.Background(), getInv, i.ID).Scan(&i.ID, &i.Title, &i.Description, &i.Quantity, &i.Price, &i.DeleteReason, &i.Deleted); err != nil {
		return err
	}
	return nil
}

const getAllInv = "SELECT * FROM items WHERE item_deleted=false"

func GetAllInventory() ([]Inventory, error) {
	rows, err := database.DB.Query(context.Background(), getAllInv)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Inventory

	for rows.Next() {
		var inv Inventory
		if err = rows.Scan(&inv.ID, &inv.Title, &inv.Description, &inv.Quantity, &inv.Price, &inv.DeleteReason, &inv.Deleted); err != nil {
			return nil, err
		}
		all = append(all, inv)
	}

	return all, nil
}

const editInv = "UPDATE items SET title=$1, item_description=$2, quantity=$3, price=$4 WHERE id=$5"

func (i *Inventory) Update() error {
	if i.ID == "" {
		return errors.New("missing ID")
	}
	_, err := database.DB.Exec(context.Background(), editInv, i.Title, i.Description, i.Quantity, i.Price, i.ID)
	return err
}

const deleteInv = "UPDATE items SET item_deleted=true, delete_reason=$2 WHERE id=$1"

func (i *Inventory) Delete(reason string) error {
	if i.ID == "" {
		return errors.New("missing ID")
	}
	_, err := database.DB.Exec(context.Background(), deleteInv, i.ID, reason)
	return err
}

const undoDelete = "UPDATE items SET item_deleted=false, delete_reason='' WHERE id=$1"

func (i Inventory) UndoDelete() error {
	if i.ID == "" {
		return errors.New("missing ID")
	}
	_, err := database.DB.Exec(context.Background(), undoDelete, i.ID)
	return err
}
