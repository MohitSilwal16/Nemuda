package models

type User struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,max=20"`
	Token    string `json:"token,omitempty"`
}

type Tweet struct {
	// ID       string `json:"id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}
