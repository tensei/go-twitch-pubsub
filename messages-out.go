package twitchpubsub

const (
	// TypeListen ...
	TypeListen = "LISTEN"
)

// ListenData ...
type ListenData struct {
	Topics    []string `json:"topics"`
	AuthToken string   `json:"auth_token,omitempty"`
}

// Listen ...
type Listen struct {
	Base

	Nonce string `json:"nonce,omitempty"`

	Data ListenData `json:"data"`
}
