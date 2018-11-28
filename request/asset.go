package request

import "github.com/geniusrabbit/opennative"

// Asset - The main container object for each asset requested or supported by Exchange
// on behalf of the rendering client.  Only one of the {title,img,video,data}
// objects should be present in each object.  The id is to be unique within the
// AssetObject array so that the response can be aligned.
type Asset struct {
	ID       int    `json:"id"`                 // Unique asset ID, assigned by exchange
	Required int    `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *Title `json:"title,omitempty"`    // Title object for title assets
	Image    *Image `json:"img,omitempty"`      // Image object for image assets
	Video    *Video `json:"video,omitempty"`    // Video object for video assets
	Data     *Data  `json:"data,omitempty"`     // Data object for brand name, description, ratings, prices etc.

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
