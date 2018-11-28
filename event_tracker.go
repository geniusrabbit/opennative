package opennative

// EventType - type of events
type EventType int

// Event types
const (
	EventTypeImpression      EventType = 1 // Impression
	EventTypeViewableMRC50   EventType = 2 // Visible impression using MRC definition at 50% in view for 1 second.
	EventTypeViewableMRC100  EventType = 3 // 100% in view for 1 second (ie GroupM standard)
	EventTypeViewableVideo50 EventType = 4 // Visible impression for video using MRC definition at 50% in view for 2 seconds.
)

// EventMethod - method of events
type EventMethod int

// Event methods
const (
	EventMethodImage EventMethod = 1 // Image-pixel tracking – URL provided will be inserted as a 1x1 pixel at the time of the event.
	EventMethodJS    EventMethod = 2 // Javascript-based tracking – URL provided will be inserted as a js tag at the time of the event.
)
