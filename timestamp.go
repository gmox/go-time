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

// Implement interface for unmarshaling JSON to Timestamp via json.Unmarshal
func (ts *Timestamp) UnmarshalJSON(value []byte) error {
    // start by seeing if we can convert the byte value to an integer
    str, err := strconv.Atoi(string(value))

    if err != nil {
        // try with parent unmarshal
        return ts.Time.UnmarshalJSON(value)
    }

    ts.Time = t.Unix(int64(str), 0)

    return nil
}

// Implement interface for unmarshaling JSON to Timestamp via json.Unmarshal
func (ts *Timestamp) MarshalJSON() ([]byte, error) {
    return ts.Time.MarshalJSON()
}

// Implement interface for unmarshaling JSON to Timestamp via Decoder.Decode
func (ts *Timestamp) Decode(value interface {}) error {
    return ts.UnmarshalJSON(value.([]byte))
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
