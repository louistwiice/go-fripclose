package entity

import "errors"

// -- Global errors --
var  ErrNotFound = errors.New("not found")


var ErrInvalidModel = errors.New("invalid model")

var ErrCanNotBeDeleted = errors.New("can not be deleted")

var ErrNullField = errors.New("some fields should not be null")

var ErrUnauthorizedAction = errors.New("you do not have right to perform this action")

var ErrNotFoundInRedis = errors.New("data not found in redis")

// -- Errors related to users --
var ErrInvalidPassword = errors.New("password is empty or invalid")

var  ErrUserNotFound = errors.New("user does not exist")

var  ErrUserDeactivated = errors.New("user is not activated")


