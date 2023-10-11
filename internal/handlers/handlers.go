package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/models"

	"github.com/go-chi/render"
)

type Response struct {
	Status int
	Data   string
}

type Request struct {
	Data      string `json:"data"`
	Sequence  uint64 `json:"sequence"`
	Timestamp int64  `json:"timestamp"`
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, Response{
		Status: 200,
		Data:   "world",
	})
}

func SaveOrder(w http.ResponseWriter, r *http.Request) {
	var req Request
	var data models.OrderStruct

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		render.JSON(w, r, Response{
			Status: 500,
			Data:   err.Error(),
		})
	}

	err = json.Unmarshal([]byte(req.Data), &data)
	if err != nil {
		render.JSON(w, r, Response{
			Status: 500,
			Data:   err.Error(),
		})
	}

	render.JSON(w, r, Response{
		Status: 200,
		Data:   "Order record added successfully",
	})

}
