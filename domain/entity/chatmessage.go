package entity

// ChatMessage entity
type ChatMessage struct {
	Identifier string `json:"-"`
	ClientID   string `json:"-"`
	IsSelf     bool   `json:"isSelf"`
	Message    string `json:"message"`
}
