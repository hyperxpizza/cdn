package customErrors

import "errors"

const (
	ErrBucketNotFound      = "Bucket with provided name has not been found"
	ErrBucketAlreadyExists = "Bucket with provided name already exists"
	ErrFileNotFound        = "File with provided name was not found"
)

func Wrap(msg string) error {
	return errors.New(msg)
}
