package enums

import (
	"bytes"
	"encoding/json"
)

type ERace int64

const (
	Human ERace = iota
	Wilderian
	Vigan
	Bugged
	Juss
	Otts
	Mongs
	Razus
	Euterian
)

var ERaces = map[ERace]string{
	Human:     "human",
	Wilderian: "wilderian",
	Vigan:     "vigan",
	Bugged:    "bugged",
	Juss:      "juss",
	Otts:      "otts",
	Mongs:     "mongs",
	Razus:     "razus",
	Euterian:  "euterian",
}

// String is used to "write" a readable text instead of a number
func (r ERace) String() string {

	if val, ok := ERaces[r]; ok {
		return val
	}

	return "unknown"
}

// MarshalJSON marshals the enum as a quoted json string
func (r *ERace) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(r.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (r *ERace) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	for key, value := range ERaces {
		if value == j {
			*r = key
			break
		}
	}

	return nil
}
