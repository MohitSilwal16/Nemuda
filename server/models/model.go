package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

type Blog struct {
	// ID       string `json:"id"`
	Username    string   `json:"username"`
	Title       string   `json:"title"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}
