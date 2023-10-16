package cache

import (
	"container/list"
	"errors"
	"server/internal/db"
	"server/internal/models"
)

const (
	CacheFillingMode = "cfm"
	CacheWorkingMode = "cwm"
)

func New(cv uint8, s *db.Store) (*Cache, error) {
	nc := Cache{queue: list.New(), records: make(map[string]*CacheRecord), capacity: cv, store: s}

	orders, err := s.ReadAllOrders(cv)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		nc.PutInCache(order.Data.OrderUid, order.Data)
	}

	return &nc, nil
}

func (c *Cache) Get(oid string) (models.OrderStruct, error) {
	if val, ok := c.records[oid]; ok {
		c.queue.MoveToFront(val.keyPtr)
		return val.data, nil
	}

	ord, err := c.store.ReadOrder(oid)
	if err != nil {
		return models.OrderStruct{}, err
	}

	c.PutInCache(oid, ord.Data)
	return ord.Data, nil
}

func (c *Cache) PutInCache(oid string, val models.OrderStruct) {
	if rec, ok := c.records[oid]; !ok {
		if c.capacity == uint8(len(c.records)) {
			back := c.queue.Back()
			c.queue.Remove(back)
			delete(c.records, back.Value.(string))
		}
		c.records[oid] = &CacheRecord{data: val, keyPtr: c.queue.PushFront(oid)}
	} else {
		rec.data = val
		c.records[oid] = rec
		c.queue.MoveToFront(rec.keyPtr)
	}
}

func (c *Cache) PutInCacheAndDB(oid string, val models.OrderStruct) error {
	if !c.store.CheckOrderExists(oid) {
		err := c.store.CreateOrder(val)
		if err != nil {
			return err
		}
		c.PutInCache(oid, val)
		return nil
	}
	return errors.New("order with this id already exists")
}
