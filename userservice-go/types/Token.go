package types

type Token struct {
	AccessToken string `json:"access_token,omitempty"`

	ExpiresIn int32 `json:"expires_in,omitempty"`

	RefreshExpiresIn int32 `json:"refresh_expires_in,omitempty"`

	RefreshToken string `json:"refresh_token,omitempty"`

	TokenType string `json:"token_type,omitempty"`

	NotBeforePolicy int32 `json:"not-before-policy,omitempty"`

	SessionState string `json:"session_state,omitempty"`

	Scope string `json:"scope,omitempty"`
}
