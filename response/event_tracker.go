// 5.8 Event Tracker Response Object
// https://www.iab.com/wp-content/uploads/2018/03/OpenRTB-Native-Ads-Specification-Final-1.2.pdf

package response

import "github.com/geniusrabbit/opennative"

// EventTracker - The event trackers object specifies the types
// of events the bidder can request to be tracked in the bid response,
// and which types of tracking are available for each event type,
// and is included as an array in the request.
type EventTracker struct {
	// Type of event available for tracking. See Event Types table.
	Event opennative.EventType `json:"event"`

	// Array of the types of tracking available for the given event.
	// See Event Tracking Methods table.
	Methods []opennative.EventMethod `json:"methods"`

	// The URL of the impage or js. Required for image or js, optional for custom.
	URL string `json:"url,omitempty"`

	// To be agreed individually with the exchange, an array
	// of key:value objects for custom tracking, for example
	// the account number of the DSP with a tracking company.
	// IE {“accountnumber”:”123”}.
	CustomData map[string]string `json:"customdata,omitempty"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
