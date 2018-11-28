package request

import "github.com/geniusrabbit/opennative"

// Video object to be used for all video elements supported in the Native Ad. This corresponds to the Video object of OpenRTB. Exchange implementers can impose their own specific restrictions. Here are the required attributes of the Video Object. For optional attributes please refer to OpenRTB.
type Video struct {
	Mimes       []string `json:"mimes"`       // Whitelist of content MIME types supported
	MinDuration int      `json:"minduration"` // Minimum video ad duration in seconds
	MaxDuration int      `json:"maxduration"` // Maximum video ad duration in seconds

	// An array of video protocols the integers publisher can accept in the bid response.
	// See OpenRTB Table ‘Video Bid Response Protocols’ for a list of possible values.
	Protocols []int `json:"protocols"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
