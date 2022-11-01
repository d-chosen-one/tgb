package enums

import (
	"bytes"
	"encoding/json"
)

type EFieldKind int64

const (
	Space EFieldKind = iota
	Anomaly
	System
)

var EFieldKinds = map[EFieldKind]string{
	Space:   "space",
	Anomaly: "anomaly",
	System:  "system",
}

// String is used to "write" a readable text instead of a number
func (k EFieldKind) String() string {

	if val, ok := EFieldKinds[k]; ok {
		return val
	}

	return "space"
}

// MarshalJSON marshals the enum as a quoted json string
func (k *EFieldKind) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(k.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (k *EFieldKind) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	for key, value := range EFieldKinds {
		if value == j {
			*k = key
			break
		}
	}

	return nil
}
