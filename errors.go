package chains

import "errors"

var (
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUnauthorized          = errors.New("access unauthorized")
	ErrBadRequest            = errors.New("bad request")
	ErrInvalidRequestPayload = errors.New("invalid request payload")
)
