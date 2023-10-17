package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/internal/cache"
	"server/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	cache    *cache.Cache
	validate *validator.Validate
}

type Response struct {
	Status int
	Data   interface{}
}

type Request struct {
	Data      string `json:"data"`
	Sequence  uint64 `json:"sequence"`
	Timestamp int64  `json:"timestamp"`
}

func New(c *cache.Cache, v *validator.Validate) *Handler {
	return &Handler{cache: c, validate: v}
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid := chi.URLParam(r, "id")

	ord, err := h.cache.Get(uid)
	if err != nil {
		render.JSON(w, r, Response{
			Status: 500,
			Data:   err.Error(),
		})
		return
	}

	render.JSON(w, r, Response{
		Status: 200,
		Data:   ord,
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
			Data:   errors.New("validation failed, please check essential parameters").Error(),
		})
		return
	}

	err = h.cache.PutInCacheAndDB(data.OrderUid, data)
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
