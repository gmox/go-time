package time

import (
    "database/sql/driver"
    "strconv"
    t "time"
)

type (
    Timestamp struct {
        t.Time
    }
)

// Implement interface for unmarshaling JSON to Timestamp
func (ts *Timestamp) UnmarshalJSON(value []byte) error {
    // start by seeing if we can convert the byte value to an integer
    str, err := strconv.Atoi(string(value))

    if err != nil {
        return err
    }

    ts.Time = t.Unix(int64(str), 0)

    return nil
}

// Implement the Scanner interface
func (ts *Timestamp) Scan(data interface{}) error {
    if data == nil {
        ts = nil
        return nil
    }

    if c, ok := data.(t.Time); ! ok {
        ts = &Timestamp{c}
    }

    return nil
}

// Implement the Valuer interface
func (ts Timestamp) Value() (driver.Value, error) {
    return ts.Time, nil
}
