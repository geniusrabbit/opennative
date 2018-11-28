package request

import "github.com/geniusrabbit/opennative"

// Data - Object is to be used for all non-core elements of the native unit
// such as Brand Name, Ratings, Review Count, Stars, Download count, descriptions etc.
// It is also generic for future native elements not contemplated at the time of
// the writing of this document. In some cases, additional recommendations
// are also included in the Data Asset Types table.
type Data struct {
	// Type ID of the element supported by the publisher.
	// The publisher can display this information in an appropriate format
	TypeID native.DataTypeID `json:"type"`

	// Maximum length of the text in the element’s response
	Length int `json:"len,omitempty"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
