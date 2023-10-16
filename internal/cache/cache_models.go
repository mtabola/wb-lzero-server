package cache

import (
	"container/list"
	"server/internal/db"
	"server/internal/models"
)

type Cache struct {
	records  map[string]*CacheRecord
	queue    *list.List
	capacity uint8
	store    *db.Store
}

type CacheRecord struct {
	data   models.OrderStruct
	keyPtr *list.Element
}
