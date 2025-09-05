package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type MonthYear time.Time

func (m MonthYear) MarshalJSON() ([]byte, error) {
	t := time.Time(m)
	formatted := fmt.Sprintf("\"%02d-%d\"", t.Month(), t.Year())
	return []byte(formatted), nil
}

func (m *MonthYear) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("01-2006", s)
	if err != nil {
		return err
	}
	*m = MonthYear(t)
	return nil
}

func (m MonthYear) Value() (driver.Value, error) {
	return time.Time(m), nil
}

func (m *MonthYear) Scan(value interface{}) error {
	if value == nil {
		*m = MonthYear(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*m = MonthYear(v)
		return nil
	default:
		return fmt.Errorf("cannot convert %T to MonthYear", value)
	}
}
