package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

type Tweet struct {
	// ID       string `json:"id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}
