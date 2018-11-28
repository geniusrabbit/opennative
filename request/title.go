package request

import "github.com/geniusrabbit/opennative"

// Title - Corresponds to the Title Object in the request, with the value filled in.
// If using asseturl response rather than embedded asset response, it is recommended that three title objects be provided, the length of each of which is less than or equal to the three recommended maximum title lengths (25,90,140).
type Title struct {
	// Maximum length of the text in the title element. Recommended to be 25, 90, or 140.
	Length int `json:"len"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
