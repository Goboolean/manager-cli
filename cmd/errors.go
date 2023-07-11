package cmd

import "errors"

var ErrInsufficientArgs = errors.New("insufficient args")
var ErrTooManyArgs = errors.New("too many args")
