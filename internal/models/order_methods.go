package models

import (
	"database/sql/driver"
	"encoding/json"
)

func (os OrderStruct) Value() (driver.Value, error) {
	return json.Marshal(os)
}
