package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

type Blog struct {
	Username      string    `json:"username"`
	Title         string    `json:"title"`
	Tag           string    `json:"tag"`
	Description   string    `json:"description"`
	Likes         uint      `json:"likes"`
	LikedUsername []string  `json:"likedUsername"`
	Comments      []Comment `json:"comments"`
	ImagePath     string    `json:"imagepath"`
}

type Comment struct {
	Username    string `json:"username"`
	Description string `json:"description"`
}
