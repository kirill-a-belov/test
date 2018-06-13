package storages

import (
	"testing"

	"awesomeProject1/test/models"
)

func TestPrefixTree_AddFind(t *testing.T) {
	testData := []*models.Event{
		{
			Name:      "Event1",
			StartTime: "2018-06-13T15:48:00Z",
		},
		{
			Name:      "Event2",
			StartTime: "2018-06-13T15:48:00Z",
		},
		{
			Name:      "Event3",
			StartTime: "2018-06-13T15:49:00Z",
		},
	}

	trie := PrefixTree{}

	for _, td := range testData {
		trie.Add(td)
	}
	absentEventName := "Event_None"
	testData = append(testData, &models.Event{
		Name:      absentEventName,
		StartTime: "2018-07-13T15:49:00Z",
	})

LOOP:
	for _, td := range testData {
		events := trie.Find(td.StartTime)

		for _, e := range events {
			if e.Name == td.Name {
				continue LOOP
			}
		}

		if td.Name == absentEventName && len(events) == 0 {
			continue LOOP
		}

		t.Error("For event:", td, "not found element in trie")

	}
}

func TestPrefixTree_FindBefore(t *testing.T) {
	testData := []*models.Event{
		{
			StartTime: "2018-06-13T15:48:00Z",
		},
		{
			StartTime: "2018-06-13T15:49:00Z",
		},
		{
			StartTime: "2018-06-13T15:50:00Z",
		},
	}

	trie := PrefixTree{}

	for _, td := range testData {
		trie.Add(td)
	}

	for i := 0; i < len(testData); i++ {
		res := trie.FindBefore(testData[i].StartTime)

		if len(res) != i+1 {
			t.Error("For event:", testData[i], "wrong before elements: ", res)
		}

	}

}
