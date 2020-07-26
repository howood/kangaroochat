package entity

// ChatMessage entity
type ChatMessage struct {
	Identifier string `json:"-"`
	ClientID   string `json:"-"`
	IsSelf     bool   `json:"isSelf"`
	UserName   string `json:"username"`
	Message    string `json:"message"`
}
