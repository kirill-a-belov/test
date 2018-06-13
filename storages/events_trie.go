package storages

import (
	"regexp"

	"awesomeProject1/test/models"
)

var re = regexp.MustCompile(`[-:TZ]*`)

type PrefixTree struct {
	children map[rune]*PrefixTree
	events   []*models.Event
}

func (pt *PrefixTree) Add(event *models.Event) {
	addRune(pt, []rune(re.ReplaceAllString(event.StartTime, "")), 0, event)
}

func addRune(pt *PrefixTree, runes []rune, index int, event *models.Event) {
	if index == len(runes) {
		return
	}

	if pt.children == nil {
		pt.children = make(map[rune]*PrefixTree)
	}

	_, ok := pt.children[runes[index]]
	if !ok {
		if index == len(runes)-1 {
			pt.children[runes[index]] = &PrefixTree{events: []*models.Event{event}}

		} else {
			pt.children[runes[index]] = &PrefixTree{}
		}
		addRune(pt.children[runes[index]], runes, index+1, event)

		return
	} else {
		if index == len(runes)-1 {
			if len(pt.children[runes[index]].events) == 0 {

				pt.children[runes[index]].events = []*models.Event{event}
			} else {
				pt.children[runes[index]].events = append(pt.children[runes[index]].events, event)
			}
		}
	}

	addRune(pt.children[runes[index]], runes, index+1, event)
}

func (pt *PrefixTree) Find(value string) []*models.Event {
	return findRune(pt, []rune(re.ReplaceAllString(value, "")), 0)
}

func findRune(pt *PrefixTree, runes []rune, index int) []*models.Event {
	if index == len(runes) {
		return nil
	}

	_, ok := pt.children[runes[index]]
	if !ok {
		return nil
	}

	if index == len(runes)-1 {
		return pt.children[runes[index]].events
	}

	return findRune(pt.children[runes[index]], runes, index+1)
}

func (pt *PrefixTree) FindBefore(value string) []*models.Event {
	return findRunesBefore(pt, []rune(re.ReplaceAllString(value, "")), 0, []*models.Event{})
}

func findRunesBefore(pt *PrefixTree, runes []rune, index int, result []*models.Event) []*models.Event {
	if index == len(runes) {
		result = append(result, pt.events...)
		return result
	}

	_, ok := pt.children[runes[index]]
	if !ok {
		for k, v := range pt.children {
			if k <= runes[index] {
				result = append(result, collectResults(v)...)
			}
		}

		return result
	}

	for k, v := range pt.children {
		if k < runes[index] {
			result = append(result, collectResults(v)...)
		}
	}

	return findRunesBefore(pt.children[runes[index]], runes, index+1, result)
}

func collectResults(pt *PrefixTree) []*models.Event {
	if pt.children == nil || len(pt.children) == 0 {
		return pt.events
	}

	result := []*models.Event{}
	for _, v := range pt.children {
		result = append(result, collectResults(v)...)
	}

	return result
}
