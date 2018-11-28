// 4.7 Event Trackers Request Object
// https://www.iab.com/wp-content/uploads/2018/03/OpenRTB-Native-Ads-Specification-Final-1.2.pdf

package request

import "github.com/geniusrabbit/opennative"

// EventTracker - (v1.2 >=) The event trackers object specifies the types of events
// the bidder can request to be tracked in the bid response, and which types
// of tracking are available for each event type, and is included as an array in the request.
type EventTracker struct {
	// Type of event available for tracking. See Event Types table.
	Event opennative.EventType `json:"event"`

	// Array of the types of tracking available for the given event. See **Event Tracking** Methods table.
	Methods []opennative.EventMethod `json:"methods"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
