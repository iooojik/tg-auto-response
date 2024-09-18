package handler

import (
	"errors"
)

var (
	ErrIgnore = errors.New("ignore message from user in config")
)
