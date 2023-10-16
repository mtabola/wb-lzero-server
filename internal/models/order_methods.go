package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

func (os OrderStruct) Value() (driver.Value, error) {
	return json.Marshal(os)
}

func (os *OrderStruct) Scan(value interface{}) error {
	b, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &os)
}
