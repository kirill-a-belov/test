package models

import (
	"errors"
	"fmt"
	"strings"
)

type Event struct {
	Name      string
	StartTime string
	EndTime   string
}

func (e *Event) String() string {
	return fmt.Sprintf("%v;%v;%v", e.Name, e.StartTime, e.EndTime)
}

func (e *Event) Validate() error {
	if strings.Compare(e.StartTime, e.EndTime) > 0 {
		return errors.New("invalid event duration")

	}
	return nil
}

func (e *Event) FillFromCSVString(input string) error {
	inputs := strings.Split(input, ";")
	if len(inputs) != 3 {
		return errors.New("invalid CSV string")
	}
	e.Name = inputs[0]
	e.StartTime = inputs[1]
	e.EndTime = inputs[2]

	return e.Validate()
}
