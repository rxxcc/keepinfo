package utils

import "errors"

var (
	ErrInvalidToken        = errors.New("token is invalid")
	ErrExpiredToken        = errors.New("token has expired")
	ErrAuthHeader          = errors.New("authorization header not provided")
	ErrSqlNoRows           = errors.New("user does not exist")
	ErrInvalidAuthHeader   = errors.New("invalid authorization header format")
	ErrUnsupportedAuthType = errors.New("unsupported authorization type")
)
