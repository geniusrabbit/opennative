package request

import "github.com/geniusrabbit/opennative"

// Image object to be used for all image elements of the Native ad such as Icons, Main Image, etc.
// Recommended sizes and aspect ratios are included in the **Image Asset Types** section.
type Image struct {
	// Type ID of the image element supported by the publisher.
	// The publisher can display this information in an appropriate format.
	TypeID opennative.ImageTypeID `json:"type,omitempty"`

	Width     int `json:"w,omitempty"`    // Width of the image in pixels
	WidthMin  int `json:"wmin,omitempty"` // The minimum requested width of the image in pixels
	Height    int `json:"h,omitempty"`    // Height of the image in pixels
	HeightMin int `json:"hmin,omitempty"` // The minimum requested height of the image in pixels
	// Either h/w or hmin/wmin should be transmitted. If only h/w is included, it
	// should be considered an exact requirement

	// supported. Popular MIME types include, but are not limited to “image/jpg” “image/gif”
	Mimes []string `json:"mimes,omitempty"` // Whitelist of content MIME types supported

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
