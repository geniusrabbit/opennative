package response

import (
	"encoding/json"

	"github.com/geniusrabbit/opennative"
)

// ResponseNative - response in 'native' field
type responseWrapper struct {
	Native *Response `json:"native"`
}

// Response - The native object is the top level JSON object which identifies a native response
type Response struct {
	// Version of the Native Markup
	Ver opennative.Version `json:"ver,omitempty"`

	// An array of Asset Objects
	// v1.2 < REQUIRED
	// v1.2 >= it's optional param
	Assets []Asset `json:"assets,omitempty"`

	// (v1.2 >=) URL of an alternate source for the assets object.
	// The expected response is a JSON object mirroring the assets
	// object in the bid response, subject to certain requirements
	// as specified in the individual objects.
	AssetsURL string `json:"assetsurl,omitempty"`

	// (v1.2 >=) BETA: URL where a dynamic creative specification may be found for
	// populating this ad, per the Dynamic Content Ads Specification.
	Dcourl string `json:"dcourl,omitempty"`

	// Destination Link. This is default link object for the ad
	Link Link `json:"link"`

	// @NOTE: Deprecated doesn't mean not used anymore
	ImpTrackers []string `json:"imptrackers,omitempty"` // Array of impression tracking URLs, expected to return a 1x1 image or 204 response
	JSTracker   string   `json:"jstracker,omitempty"`   // Optional JavaScript impression tracker. This is a valid HTML, Javascript is already wrapped in <script> tags. It should be executed at impression time where it can be supported

	// Array of tracking objects to run with the ad, in response to the
	// declared supported methods in the request. Replaces imptrackers and
	// jstracker, to be deprecated.
	EventTrackers []EventTracker `json:"eventtrackers,omitempty"`

	// If support was indicated in the request, URL of a page informing
	// the user about the buyer’s targeting activity.
	Privacy string `json:"privacy,omitempty"`

	// This object is a placeholder that may contain custom JSON
	// agreed to by the parties to support flexibility beyond the standard defined in this specification
	Ext opennative.Extension `json:"ext,omitempty"`
}

// DefaultResponse - return default object of response
func DefaultResponse() (out Response) {
	out.Ver = opennative.Version1_2
	return
}

// OpenRTBMarkup returns formated JSON data according with version of open native
func (r *Response) OpenRTBMarkup() ([]byte, error) {
	return json.Marshal(responseWrapper{Response: r})
}

// DecodeOpenRTBMarkup decodes wrapped response JSON
func DecodeOpenRTBMarkup(markup []byte) (*Response, error) {
	var (
		resp responseWrapper
		err  = json.Unmarshal(data, &resp)
	)
	return resp.Response, err
}

// AppendImpTrackers - add impression image trackers
func (r *Response) AppendImpTrackers(ss ...string) {
	if r.Ver.IsVer1_2() {
		for _, s := range ss {
			r.EventTrackers = append(r.EventTrackers, EventTracker{
				Event:   opennative.EventTypeImpression,
				Methods: []opennative.EventMethod{opennative.EventMethodImage},
				URL:     s,
			})
		}
	} else {
		r.ImpTrackers = append(r.ImpTrackers, ss...)
	}
}

// AppendOrReplaceJSTracker - append or replace JS tracker
func (r *Response) AppendOrReplaceJSTracker(s string) {
	if r.Ver.IsVer1_2() {
		r.EventTrackers = append(r.EventTrackers, EventTracker{
			Event:   opennative.EventTypeImpression,
			Methods: []opennative.EventMethod{opennative.EventMethodJS},
			URL:     s,
		})
	} else if s != "" {
		r.JSTracker = s
	}
}

// CopyTrackers - copy trackers imp and js
func (r *Response) CopyTrackers(resp *Response) {
	if resp.Ver.IsVer1_2() {
		if r.Ver.IsVer1_2() {
			r.EventTrackers = append([]EventTracker{}, resp.EventTrackers...)
		} else {
			for _, et := range resp.EventTrackers {
				for _, m := range et.Methods {
					switch m {
					case opennative.EventMethodImage:
						r.AppendImpTrackers(et.URL)
					case opennative.EventMethodJS:
						r.AppendOrReplaceJSTracker(et.URL)
					}
				}
			}
		} // end if
	}

	// It still can be used in version 1.2 ...
	r.AppendImpTrackers(resp.ImpTrackers...)
	r.AppendOrReplaceJSTracker(resp.JSTracker)
}

// GetImpTrackers - return impression trackers
func (r *Response) GetImpTrackers() []string {
	var trackers []string

	if r.Ver.IsVer1_2() {
		for _, et := range r.EventTrackers {
			if hasMethod(et.Methods, opennative.EventMethodImage) {
				trackers = append(trackers, et.URL)
			}
		}
		trackers = append(trackers, r.ImpTrackers...)
	} else {
		trackers = r.ImpTrackers
	}
	return trackers
}

// GetJSTrackers - get js trackers
func (r *Response) GetJSTrackers() []string {
	var trackers []string

	if r.Ver.IsVer1_2() {
		for _, et := range r.EventTrackers {
			if hasMethod(et.Methods, opennative.EventMethodJS) {
				trackers = append(trackers, et.URL)
			}
		}
	}

	if r.JSTracker != "" {
		trackers = append(trackers, r.JSTracker)
	}
	return trackers
}

// PrepareImpTrackers - all impression trackers
func (r *Response) PrepareImpTrackers(fn func(url string) string) {
	for i, t := range r.ImpTrackers {
		r.ImpTrackers[i] = fn(t)
	}

	for i, et := range r.EventTrackers {
		if hasMethod(et.Methods, opennative.EventMethodImage) {
			r.EventTrackers[i].URL = fn(et.URL)
		}
	}
}

// PrepareJSTrackers - all js trackers
func (r *Response) PrepareJSTrackers(fn func(url string) string) {
	if len(r.JSTracker) > 0 {
		r.JSTracker = fn(r.JSTracker)
	}

	for i, et := range r.EventTrackers {
		if hasMethod(et.Methods, opennative.EventMethodJS) {
			r.EventTrackers[i].URL = fn(et.URL)
		}
	}
}

// PrepareTrackers - all trackers url
func (r *Response) PrepareTrackers(fn func(url string) string) {
	r.ProcessImpTrackers(fn)
	r.ProcessJSTrackers(fn)
}

func hasMethod(mm []opennative.EventMethod, m opennative.EventMethod) bool {
	for _, method := range mm {
		if method == m {
			return true
		}
	}
	return false
}
