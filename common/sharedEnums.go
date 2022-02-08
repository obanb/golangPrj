package common

import (
	"bytes"
	"encoding/json"
)

type Position string

const (
	Tester  Position = "tester"
	Lead             = "lead"
	Analyst          = "analyst"
	Dev              = "dev"
	Other            = "other"
)

type AccountType string

const (
	Common AccountType = "common"
	Admin              = "admin"
)

type AccountStatus int

const (
	Active AccountStatus = iota
	Inactive
	RegistationInProgress
)

var toString = map[Position]string{
	Tester:  "Tester",
	Analyst: "Analyst",
	Dev:     "Dev",
	Lead:    "Lead",
	Other:   "Other",
}

var toIota = map[string]Position{
	"Tester":  Tester,
	"Analyst": Analyst,
	"Dev":     Dev,
	"Lead":    Lead,
	"Other":   Other,
}

func (p Position) UnmarshalPositionJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	p = toIota[j]
	return nil
}

func (p Position) MarshalPositionJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[p])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
