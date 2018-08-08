package sqlx

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type NullString string

func (j *NullString) Scan(src interface{}) error {
	if src == nil {
		*j = ""
		return nil
	}

	if val, ok := src.(string); ok {
		*j = NullString(val)
		return nil
	}

	return fmt.Errorf("expected: %T, got: %T", j, src)
}

func (j *NullString) Value() (driver.Value, error) {
	return string(*j), nil
}

type NullInt64 int64

func (j *NullInt64) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}

	if val, ok := src.(int64); ok {
		*j = NullInt64(val)
		return nil
	}

	return fmt.Errorf("expected: %T, got: %T", j, src)
}

func (j *NullInt64) Value() (driver.Value, error) {
	return int64(*j), nil
}

type NullInt32 int32

func (j *NullInt32) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}

	if val, ok := src.(int64); ok {
		*j = NullInt32(val)
		return nil
	}

	return fmt.Errorf("expected: %T, got: %T", j, src)
}

func (j *NullInt32) Value() (driver.Value, error) {
	return int64(*j), nil
}

type NullFloat32 float32

func (j *NullFloat32) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}

	if val, ok := src.(float64); ok {
		*j = NullFloat32(val)
		return nil
	}

	return fmt.Errorf("expected: %T, got: %T", j, src)
}

func (j *NullFloat32) Value() (driver.Value, error) {
	return float64(*j), nil
}

type NullFloat64 float64

func (j *NullFloat64) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}

	if val, ok := src.(float64); ok {
		*j = NullFloat64(val)
		return nil
	}

	return fmt.Errorf("expected: %T, got: %T", j, src)
}

func (j *NullFloat64) Value() (driver.Value, error) {
	return float64(*j), nil
}

type NullTime time.Time

func (j *NullTime) Scan(src interface{}) error {
	if src == nil {
		*j = NullTime(time.Time{})
		return nil
	}

	if val, ok := src.(time.Time); ok {
		*j = NullTime(val)
		return nil
	}

	return fmt.Errorf("expected: %T, got: %T", j, src)
}

func (j *NullTime) Value() (driver.Value, error) {
	return time.Time(*j), nil
}

type NullInt int

func (j *NullInt) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}

	if val, ok := src.(int64); ok {
		*j = NullInt(val)
		return nil
	}

	return fmt.Errorf("expected: %T, got: %T", j, src)
}

func (j *NullInt) Value() (driver.Value, error) {
	return int64(*j), nil
}
