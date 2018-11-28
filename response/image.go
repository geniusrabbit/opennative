package response

import "github.com/geniusrabbit/opennative"

// Image - Corresponds to the Image Object in the request.
// The Image object to be used for all image elements of the Native
// ad such as Icons, Main Image, etc.
//
// It is recommended that if asseturl is being used rather than
// embedded assets, that an image of each recommended aspect
// ratio (per the Image Types table) be provided for image type 3.
type Image struct {
	// Type ID of the image element supported by the publisher
	TypeID opennative.ImageTypeID `json:"type,omitempty"`

	URL string `json:"url"` // URL of the image asset

	// Width of the image in pixels.
	// Recommended for embedded asset responses.
	// (v1.2 >=) Required for assetsurl/dcourlresponses if multiple assets of same type submitted.
	Width int `json:"w,omitempty"`

	// Height of the image in pixels.
	// Recommended for embedded asset responses.
	// (v1.2 >=) Required for assetsurl/dcourl responses if multiple assets of same type submitted.
	Height int `json:"h,omitempty"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
