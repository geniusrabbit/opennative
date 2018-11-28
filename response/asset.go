package response

import "github.com/geniusrabbit/opennative"

// Asset - Corresponds to the Asset Object in the request. The main container object for each asset requested or supported by Exchange on behalf of the rendering client. Any object that is required is to be flagged as such. Only one of the {title,img,video,data} objects should be present in each object. All others should be null/absent. The id is to be unique within the AssetObject array so that the response can be aligned.
type Asset struct {
	ID       opennative.NumberOrString `json:"id,omitempty"`       // Unique asset ID, assigned by exchange
	Required opennative.NumberOrString `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *Title                    `json:"title,omitempty"`    // Title object for title assets
	Image    *Image                    `json:"img,omitempty"`      // Image object for image assets
	Video    *Video                    `json:"video,omitempty"`    // Video object for video assets
	Data     *Data                     `json:"data,omitempty"`     // Data object for brand name, description, ratings, prices etc.
	Link     *Link                     `json:"link,omitempty"`     // Link object for call to actions. The link object applies if the asset item is activated (clicked). If there is no link object on the asset, the parent link object on the bid response applies.

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}

// DefaultAsset - default asset response object
func DefaultAsset() (out Asset) {
	out.Required = 0
	return
}
