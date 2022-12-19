package web

type WebResponse struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Data   any    `json:"data,omitempty"`
}
