package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
}

type Blog struct {
	// BlogId      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Tags        []string  `json:"tags"`
	Description string    `json:"description"`
	Likes       uint      `json:"likes"`
	Comments    []Comment `json:"comments"`
}

type Comment struct {
	// CommentId   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string `json:"username"`
	Description string `json:"description"`
}
