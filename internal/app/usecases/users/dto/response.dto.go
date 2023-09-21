package dto

// UserDefault DTO default response type for structure User.
type UserDefault struct {
	ID               int     `json:"id,omitempty"`
	UserCode         string  `json:"user_code,omitempty"`
	EnteID           *string `json:"ente_id,omitempty"`
	EnteType         string  `json:"ente_type,omitempty"`
	Role             string  `json:"role,omitempty"`
	Username         string  `json:"username,omitempty"`
	Password         string  `json:"password,omitempty"`
	SecurityQuestion *string `json:"security_question,omitempty"`
	SecurityAnswer   *string `json:"security_answer,omitempty"`
	State            string  `json:"state,omitempty"`
}
