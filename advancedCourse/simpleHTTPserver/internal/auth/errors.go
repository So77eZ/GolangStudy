package auth

import "errors"

const (
// ErrUserNameExists  = "user with such name already exists"
// ErrUserEmailExists = "user with such email already exists"
// ErrRegistrationErr = "occured error while registration"
// ErrConnection      = "occured connection error"
// ErrWrongCredetials = "wrong email or password"
// ErrNotEnougthData  = "input email or name and password"
)

var ErrUserNameExists = errors.New("user with such name already exists")
var ErrUserEmailExists = errors.New("user with such email already exists")
var ErrRegistrationErr = errors.New("occured error while registration")
var ErrConnection = errors.New("occured connection error")
var ErrWrongCredetials = errors.New("wrong email or password")
var ErrNotEnougthData = errors.New("input email or name and password")
