package error

type RestError struct {
	Message    string `json: "message"`
	StatusCode int    `json: "status_code"`
	Error      string `json: "error"`
}
