package event

import (
	"time"

	"github.com/k0kubun/pp"
)

type Event struct {
	text        string
	description string
	time        time.Time
	humanTime   string
}

func NewEvent(text string, description string, time time.Time, humanTime string) Event {
	return Event{
		text:        text,
		description: description,
		time:        time,
		humanTime:   humanTime,
	}
}

func AddEvent(arr *[]Event, new_event Event) {
	*arr = append(*arr, new_event)
}

func GetAllEvents(arr *[]Event) {
	pp.Println(arr)
}
