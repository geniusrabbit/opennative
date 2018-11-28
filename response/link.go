package response

import "github.com/geniusrabbit/opennative"

// Link - Used for ‘call to action’ assets, or other links from the Native ad.
// This Object should be associated to its peer object in the parent Asset Object
// or as the master link in the top level Native Ad response object. When that
// peer object is activated (clicked) the action should take the user to the
// location of the link.
type Link struct {
	// Landing URL of the clickable link
	URL string `json:"url"`

	// List of third-party tracker URLs to be fired on click of the URL
	ClickTrackers []string `json:"clicktrackers,omitempty"`

	// Fallback URL for deeplink. To be used if the URL given in url is not supported by the device.
	FallbackURL string `json:"fallback,omitempty"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}
