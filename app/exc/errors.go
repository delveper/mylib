package exc

import "errors"

var ErrDuplicateEmail = errors.New("email is already taken")