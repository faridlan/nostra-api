package web

type UserUpdateReq struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
