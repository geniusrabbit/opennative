package response

// Video - Corresponds to the Video Object in the request, yet containing
// a value of a conforming VAST tag as a value.
type Video struct {
	VASTTag string `json:"vasttag"` // VAST XML
}
