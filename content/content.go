package content

// Content is a structure that represents a piece of content
type Content struct {
	ID                  string `json:"id,omitempty"`
	Type                string `json:"type,omitempty"`
	Version             int    `json:"version,omitempty"`
	ActivationTimeStamp int64  `json:"activationtimestamp,omitempty"`
	Body                string `json:"body,omitempty"`
}
