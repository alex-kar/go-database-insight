package insight

import (
	"database/sql/driver"
)

type Insight struct {
}

func (i Insight) Wrap() (driver.Driver, error) {
	return nil, nil
}
