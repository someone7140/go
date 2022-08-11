package errorConstants

import "errors"

// ErrInternalServer Internal Server Error
var ErrInternalServer = errors.New("Internal Server Error")

// ErrBadRequest Bad Request
var ErrBadRequest = errors.New("Bad Request Error")

// ErrNotFound Not Found Error
var ErrNotFound = errors.New("Not Found Error")

// ErrUnauthorized Unauthorized Error
var ErrUnauthorized = errors.New("Unauthorized Error")

// ErrForbidden Forbidden Error
var ErrForbidden = errors.New("Forbidden Error")
