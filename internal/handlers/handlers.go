package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/db"
	"server/internal/models"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store    *db.Store
	validate *validator.Validate
}

type Response struct {
	Status int
	Data   string
}

type Request struct {
	Data      string `json:"data"`
	Sequence  uint64 `json:"sequence"`
	Timestamp int64  `json:"timestamp"`
}

func New(db *db.Store, v *validator.Validate) *Handler {
	return &Handler{store: db, validate: v}
}

func (h *Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, Response{
		Status: 200,
		Data:   "world",
	})
}

func (h *Handler) SaveOrder(w http.ResponseWriter, r *http.Request) {
	var req Request

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		render.JSON(w, r, Response{
			Status: 500,
			Data:   err.Error(),
		})
		return
	}

	var data models.OrderStruct
	err = json.Unmarshal([]byte(req.Data), &data)
	if err != nil {
		render.JSON(w, r, Response{
			Status: 500,
			Data:   err.Error(),
		})
		return
	}

	if err := h.validate.Struct(data); err != nil {
		fmt.Println(err)
		render.JSON(w, r, Response{
			Status: 500,
			Data:   err.Error(),
		})
		return
	}

	err = h.store.CreateOrder(data)
	if err != nil {
		render.JSON(w, r, Response{
			Status: 500,
			Data:   err.Error(),
		})
		return
	}

	render.JSON(w, r, Response{
		Status: 200,
		Data:   "Order record added successfully",
	})

}
