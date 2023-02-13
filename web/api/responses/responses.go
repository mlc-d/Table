package responses

type Response struct {
	Code    int    `json:"code,omitempty"`
	Payload any    `json:"payload"`
	Comment string `json:"comment,omitempty"`
}

type Error struct {
	ErrorString string `json:"error"`
}
