package handler

import (
	"errors"
)

var (
	ErrIgnore      = errors.New("ignore message from user in config")
	ErrNoCondition = errors.New("no condition")
	ErrNoMessage   = errors.New("no message")
)
