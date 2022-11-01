package enums

import (
	"bytes"
	"encoding/json"
)

type EFieldSize int64

const (
	Small EFieldSize = iota
	Medium
	Large
	XLarge
)

var EFieldSizes = map[EFieldSize]string{
	Small:  "small",
	Medium: "medium",
	Large:  "large",
	XLarge: "xLarge",
}

// String is used to "write" a readable text instead of a number
func (f EFieldSize) String() string {

	if val, ok := EFieldSizes[f]; ok {
		return val
	}

	return "unknown"
}

// MarshalJSON marshals the enum as a quoted json string
func (f *EFieldSize) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(f.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (f *EFieldSize) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	for key, value := range EFieldSizes {
		if value == j {
			*f = key
			break
		}
	}

	return nil
}
