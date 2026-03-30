package domain

import "errors"

var ErrInvalidRefreshToken = errors.New("invalid refresh token")
var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrInvalidRole = errors.New("Error Invalid Role status")
var ErrInvalidCredentials = errors.New("Error invalid Credetials")
var ErrResumeNotFound = errors.New("Resume not found")
var ErrInvalidData = errors.New("invalid data")
var ErrForbidden = errors.New("forbidden")
var ErrNotFound = errors.New("not found")
