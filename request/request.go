package request

import (
	"encoding/json"

	"github.com/geniusrabbit/opennative"
)

// NativeRequest - wrapper around Request - support old version
type NativeRequest struct {
	Native Request `json:"native"`
}

// Request - This object represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as
// a native type impression. At the publisher’s discretion, that same impression may also be offered as
// banner and/or video by also including as Imp subordinates the Banner and/or Video objects,
// respectively. However, any given bid for the impression must conform to one of the offered types.
type Request struct {
	// Version of the Native Markup
	Ver opennative.Version `json:"ver"`

	// The Layout ID of the native ad
	// v1.0 - Recommended
	// v1.2 - Not used
	LayoutID int `json:"layout"`

	// The Ad unit ID of the native ad
	// v1.0 - Recommended
	// v1.2 - Not used
	AdUnitID int `json:"adunit"`

	// The context in which the ad appears; RECOMMENDED
	ContextTypeID ContextTypeID `json:"context,omitempty"`

	// A more detailed context in which the ad appears
	ContextSubTypeID ContextSubTypeID `json:"contextsubtype,omitempty"`

	// The design/format/layout of the ad unit being offered; RECOMMENDED
	PlacementTypeID PlacementTypeID `json:"plcmttype,omitempty"`

	// The number of identical placements in this Layout; Default 1
	PlacementCount int `json:"plcmtcnt,omitempty"`

	// 0 for the first ad, 1 for the second ad, and so on.
	// Note this would generally NOT be used in combination with plcmtcnt - either
	// you are auctioning multiple identical placements (in which case plcmtcnt>1, seq=0)
	// or you are holding separate auctions for distinct items in the feed (in which case plcmtcnt=1, seq=>=1)
	Sequence int `json:"seq,omitempty"`

	// An array of Asset Objects. Any bid response must comply with
	// the array of elements expressed in the bid request.
	Assets []Asset `json:"assets"`

	// (v1.2 >=) Whether the supply source / impression supports returning an
	// assetsurl instead of an asset object. 0 or the absence of the field indicates no such support.
	AURLSupport int `json:"aurlsupport,omitempty"`

	// (v1.2 >=) Whether the supply source / impression supports returning a
	// dco url instead of an asset object. 0 or the absence of the field indicates no such support. Beta
	DURLSupport int `json:"durlsupport,omitempty"`

	// (v1.2 >=) Specifies what type of event tracking is supported – see Event Trackers Request Object
	EventTrackers []EventTracker `json:"eventtrackers,omitempty"`

	// (v1.2 >=) Set to 1 when the native ad supports buyer-specific privacy notice
	// (see http://youradchoices.com/ for example). Set to 0 (or field absent
	// when the native ad doesn’t support custom privacy links or if support is unknown.
	Privacy int `json:"privacy,omitempty"`

	// This object is a placeholder that may contain custom JSON agreed
	// to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}

// OpenRTBMarkup returns formated JSON data according with version of open native
func (r *Request) OpenRTBMarkup() ([]byte, error) {
	return json.Marshal(struct {
		Native *Request `json:"native"`
	}{Native: r})
}

// ContextTypeID - The context in which the ad appears - what type of content is surrounding the ad on the page at a high level. This maps directly to the new Deep Dive on In-Feed Ad Units. This denotes the primary context, but does not imply other content may not exist on the page - for example it's expected that most content platforms have some social components, etc.
type ContextTypeID int

// Context types
const (
	ContextTypeContent ContextTypeID = 1 // newsfeed, article, image gallery, video gallery
	ContextTypeSocial  ContextTypeID = 2 // social network feed, email, chat
	ContextTypeProduct ContextTypeID = 3 // product listings, details, recommendations, reviews
)

// ContextSubTypeID - Next-level context in which the ad appears.
// Again this reflects the primary context, and does not imply no presence
// of other elements. For example, an article is likely to contain images
// but is still first and foremost an article. SubType should only be combined
// with the primary context type as indicated (ie for a context type of 1,
// only context subtypes that start with 1 are valid).
type ContextSubTypeID int

// Context sub types
const (
	ContextSubTypeGeneral        ContextSubTypeID = 10
	ContextSubTypeArticle        ContextSubTypeID = 11
	ContextSubTypeVideo          ContextSubTypeID = 12
	ContextSubTypeAudio          ContextSubTypeID = 13
	ContextSubTypeImage          ContextSubTypeID = 14
	ContextSubTypeUserGenerated  ContextSubTypeID = 15
	ContextSubTypeSocial         ContextSubTypeID = 20
	ContextSubTypeEmail          ContextSubTypeID = 21
	ContextSubTypeChat           ContextSubTypeID = 22
	ContextSubTypeSelling        ContextSubTypeID = 30
	ContextSubTypeAppStore       ContextSubTypeID = 31
	ContextSubTypeProductReviews ContextSubTypeID = 32
)

// PlacementTypeID - The FORMAT of the ad you are purchasing, separate from the surrounding context
type PlacementTypeID int

// Placement types
const (
	PlacementTypeInFeed         PlacementTypeID = 1 // In the feed of content - for example as an item inside the organic feed/grid/listing/carousel
	PlacementTypeAtomic         PlacementTypeID = 2 // In the atomic unit of the content - IE in the article page or single image page
	PlacementTypeOutside        PlacementTypeID = 3 // Outside the core content - for example in the ads section on the right rail, as a banner-style placement near the content, etc.
	PlacementTypeRecommendation PlacementTypeID = 4 // Recommendation widget, most commonly presented below the article content
)
