package utils

import (
    "database/sql/driver"
    "errors"
)

type NullToEmptyString string

func (s *NullToEmptyString) Scan(value interface{}) error {
    if value == nil {
        *s = ""
        return nil
    }

    strVal, ok := value.(string)
    if !ok {
        bs, ok := value.([]byte)
        if !ok {
            return errors.New("cannot convert to string")
        }
        strVal = string(bs)
    }

    *s = NullToEmptyString(strVal)
    return nil
}

func (s NullToEmptyString) Value() (driver.Value, error) {
    return string(s), nil
}
