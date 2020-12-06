package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DBMap map[string]interface{}

func (m DBMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan not supported")
	}
	if bytes == nil || value == nil {
		return nil
	}
	return json.Unmarshal(bytes, &m)
}

func (m DBMap) Value() (driver.Value, error) {
	j, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal to map[string]interface")
	}
	return driver.Value(j), nil
}
