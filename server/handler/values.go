package handler

import "errors"

var (
	// Auth
	ErrInvalidUsernameFormat    = errors.New("USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC")
	ErrInvalidPasswordFormat    = errors.New("PASSWORD MUST BE B'TWIN 8-20 CHARS, ATLEAST 1 UPPER, LOWER CASE & SPECIAL CHARACTER")
	ErrUsernameAlreadyUsedError = errors.New("USERNAME IS ALREADY USED")
	ErrInvalidCredentials       = errors.New("INVALID CREDENTIALS")

	// Common
	ErrInternalServerError = errors.New("INTERNAL SERVER ERROR")
	ErrInvalidSessionToken = errors.New("INVALID SESSION TOKEN")
	ErrUserNotFound        = errors.New("USER NOT FOUND")
	ErrXSSDetected         = errors.New("XSS DETECTED")
	ErrInvalidOffset       = errors.New("INVALID OFFSET") // Used in Messages & Blogs Pagination

	// Blogs
	ErrInvalidBlogTag                      = errors.New("INVALID BLOG TAG")
	ErrInvalidBlogTitleFormat              = errors.New("TITLE MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC")
	ErrInvalidBlogDescriptionFormat        = errors.New("DESCRIPTION MUST BE B'TWIN 4-50 CHARS")
	ErrInvalidBlogCommentDescriptionFormat = errors.New("COMMENT MUST BE B'TWIN 4-50 CHARS")
	ErrInvalidBlogImageFileType            = errors.New("INVALID FILE TYPE, ONLY JPG, JPEG & PNG ARE ACCEPTED")
	ErrInvalidBlogImageSize                = errors.New("IMAGE SIZE EXCEEDS 2 MB")
	ErrInvalidBlogTitleAlreadyUsed         = errors.New("BLOG TITLE IS ALREADY USED")
	ErrBlogNotFound                        = errors.New("BLOG NOT FOUND")
)

const (
	MESSAGE_LIMIT       = 9
	BLOG_LIMIT          = 3
	MAX_BLOG_IMAGE_SIZE = 2 * 1024 * 1024 // 2 MB
)

// Tags' slice
var tagsList = []string{"Political", "Technical", "Educational", "Geographical", "Programming", "Other"}
