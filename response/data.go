package response

import "github.com/geniusrabbit/opennative"

// Data - Corresponds to the Data Object in the request, with the value filled in.
// The Data Object is to be used for all miscellaneous elements of the native unit
// such as Brand Name, Ratings, Review Count, Stars, Downloads, Price count etc.
// It is also generic for future native elements not contemplated at the time of
// the writing of this document.
type Data struct {
	// The optional formatted string name of the data type to be displayed
	// Not used in v1.2 and above
	Label string `json:"label,omitempty"`

	// Type ID of the element supported by the publisher.
	// The publisher can display this information in an appropriate format.
	// (v1.2 >= REQUIRED) for asseturl responses, NOT REQUIRED for embedded asset responses.
	// The type of data element being submitted from the Data Asset Types table.
	TypeID native.DataTypeID `json:"type,omitempty"`

	// (v1.2 >= REQUIRED) for asseturl responses, not required for embedded asset responses.
	// The length of the data element being submitted. Where applicable,
	// must comply with the recommended maximum lengths in the **Data Asset Types table**.
	Len int `json:"len,omitempty"`

	// The formatted string of data to be displayed.
	// Can contain a formatted value such as “5 stars” or “$10” or “3.4 stars out of 5”
	Value string `json:"value"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
