package handler

import (
	"time"

	"github.com/Nemuda/client/constants"
)

const SERVICE_BASE_URL = constants.SERVICE_BASE_URL

const LONG_CONTEXT_TIMEOUT = 10 * time.Second
const CONTEXT_TIMEOUT = 5 * time.Second
const SHORT_CONTEXT_TIMEOUT = time.Second

const INTERNAL_SERVER_ERROR_MESSAGE = "<script>alert('Internal Server Error');</script>"
const REQUEST_TIMED_OUT_MESSAGE = "<script>alert('Request Timed Out');</script>"
const BLOG_NOT_FOUND_MESSAGE = "<script>alert('Blog Not Found');</script>"

// Tags' slice
var tagsList = []string{"Political", "Technical", "Educational", "Geographical", "Programming", "Other"}
