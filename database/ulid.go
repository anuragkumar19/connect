package database

import (
	"database/sql/driver"

	"github.com/oklog/ulid/v2"
)

type ULIDValue struct {
	Valid bool
	ULID  ulid.ULID
}

// Scan implements the database/sql Scanner interface.
func (id *ULIDValue) Scan(src any) error {
	if src == nil {
		*id = ULIDValue{}
		return nil
	}

	return id.ULID.Scan(src)
}

// Value implements the database/sql/driver Valuer interface.
func (id ULIDValue) Value() (driver.Value, error) {
	if !id.Valid {
		return nil, nil
	}

	return id.ULID, nil
}
