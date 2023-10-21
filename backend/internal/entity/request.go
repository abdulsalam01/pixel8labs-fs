package entity

type BaseRequestParams struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Result struct {
	StatusCode int
	Data       []byte
}
