package util

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RFCTime struct {
	time.Time
}

func (t *RFCTime) MarsahlJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, time.RFC3339)
	b = append(b, '"')
	return b, nil
}

func (t *RFCTime) UnmarshalJSON(b []byte) error {
	t2, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}

	*t = RFCTime{t2}
	return nil
}

func (t RFCTime) Value() (driver.Value, error) {
	return t.Time, nil
}

func (t *RFCTime) Scan(src interface{}) error {
	if t2, ok := src.(time.Time); ok {
		*t = RFCTime{t2}
		return nil
	}
	return fmt.Errorf("Scan error")
}

type UUID struct {
	uuid.UUID
}

func (u *UUID) MarshalJSON() ([]byte, error) {
	b, err := u.MarshalText()
	if err != nil {
		return nil, err
	}

	b2 := make([]byte, 0, len(b)+2)
	b2 = append(b2, '"')
	b2 = append(b2, b...)
	b2 = append(b2, '"')
	return b2, nil
}

func (u *UUID) UnmarshalJSON(b []byte) error {
	b = b[1 : len(b)-1]
	return u.UnmarshalText(b)
}

func (u UUID) Value() (driver.Value, error) {
	b, _ := u.MarshalBinary()
	return b, nil
}

func (u *UUID) Scan(src interface{}) error {
	if src == nil {
		*u = UUID{}
		return nil
	}

	if u2, ok := src.([]byte); ok {
		err := u.UUID.UnmarshalBinary(u2)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Scan error")
}
