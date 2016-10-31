package main

import (
	"sort"
	"testing"
	"time"
)

func TestMessageSorting(t *testing.T) {
	unsorted := messageBatch{
		logEvent{timestamp: 2},
		logEvent{timestamp: 1},
	}

	sort.Sort(unsorted)

	for i, elem := range []int64{1, 2} {
		if unsorted[i].timestamp != elem {
			t.Errorf("Timestamps should be equal. Got: %v Expected: %v", unsorted[i].timestamp, elem)
		}
	}
}

func TestEventSize(t *testing.T) {
	event := logEvent{msg: "123", timestamp: 123}
	expected := 3 + eventSizeOverhead
	if event.size() != expected {
		t.Errorf("Event size shoud be %v. Got: %v", expected, event.size())
	}
}

func TestBatchSize(t *testing.T) {
	events := messageBatch{
		logEvent{msg: "123456"},
		logEvent{msg: "12345"},
		logEvent{msg: "123"},
	}

	result := events.size()
	expected := (6 + 5 + 3) + (eventSizeOverhead * 3)

	if result != expected {
		t.Errorf("Batch size shoud be %v. Got: %v", expected, result)
	}
}

func TestEventValidateTooBig(t *testing.T) {
	event := logEvent{msg: RandomString(maxEventSize + 1)}
	if err := event.validate(); err != errMessageTooBig {
		t.Errorf("Should return %q. Got: %q", errMessageTooBig, err)
	}
}

func TestDestinationString(t *testing.T) {
	dst := destination{group: "group", stream: "stream"}
	expected := "group: group stream: stream"
	if str := dst.String(); str != expected {
		t.Errorf("Should return '%s'. Got: '%s'", expected, str)
	}
}
