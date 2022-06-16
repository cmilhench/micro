package model

import "errors"

var (
	ErrPersonNotFound      = errors.New("the person was not found")
	ErrPersonAlreadyExists = errors.New("the person was not found")
)
