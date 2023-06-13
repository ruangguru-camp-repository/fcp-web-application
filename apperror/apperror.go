package apperror

import "errors"

var (
	ErrInvalidUserIdOrCategoryId = errors.New("invalid user id or category id")
)