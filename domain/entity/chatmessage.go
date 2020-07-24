package entity

// ChatMessage entity
type ChatMessage struct {
	Identifier string `json:"-"`
	Message    string `json:"message"`
}
