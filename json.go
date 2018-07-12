package sqlx

import (
	"database/sql"
	"database/sql/driver"
	ej "encoding/json"
	"fmt"
)

type json struct {
	v interface{}
}

// JSON is convenient wrapper to auto marshal and unmarshal JSON data type as string
// v should be a pointer. It works similar to json.Unmarshal and Marshal
func JSON(v interface{}) interface {
	driver.Valuer
	sql.Scanner
} {
	return &json{v}
}

func (j *json) Scan(src interface{}) error {
	dat, ok := src.(string)
	if !ok {
		return fmt.Errorf("sqlx: error while scanning json: expected string go %T", src)
	}

	return ej.Unmarshal([]byte(dat), j.v)
}

func (j *json) Value() (driver.Value, error) {
	dat, err := ej.Marshal(j.v)
	return string(dat), err
}

type jsonb struct {
	v interface{}
}

// JSONB is convenient wrapper to auto marshal and unmarshal JSON data type as []byte
// v should be a pointer. It works similar to json.Unmarshal and Marshal
func JSONB(v interface{}) interface {
	driver.Valuer
	sql.Scanner
} {
	return &json{v}
}

func (j *jsonb) Scan(src interface{}) error {
	dat, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("sqlx: error while scanning jsonb: expected []byte go %T", src)
	}

	return ej.Unmarshal(dat, j.v)
}

func (j *jsonb) Value() (driver.Value, error) {
	return ej.Marshal(j.v)
}
