package models

import (
	"testing"
)

type TestData struct {
	IsError bool
	Input   *Event
}

func TestEvent_Validate(t *testing.T) {
	testData := []*TestData{
		{
			IsError: false,
			Input: &Event{
				StartTime: "2018-06-13T15:49:51Z",
				EndTime:   "2018-06-13T15:50:51Z",
			},
		},
		{
			IsError: true,
			Input: &Event{
				StartTime: "2018-06-13T15:50:51Z",
				EndTime:   "2018-06-13T15:49:51Z",
			},
		},
	}

	for _, td := range testData {
		err := td.Input.Validate()
		if (err != nil && !td.IsError) ||
			(err == nil && td.IsError) {
			t.Error(
				"For event:", td.Input,
				"validation error expectation", td.IsError,
				"got", err != nil,
			)
		}
	}
}
