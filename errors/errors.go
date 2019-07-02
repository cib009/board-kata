package errors

import "github.com/pkg/errors"

type fileNotFound struct {
	error
}

type couldNotWriteFile struct {
	error
}

type emptyMessage struct {
	error
}

func WrapFileNotFound(err error, format string, args ...interface{}) error {
	return &fileNotFound{errors.Wrapf(err, format, args...)}
}

func IsFileNotFound(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*fileNotFound)

	return ok
}

func WrapCouldNotWriteFile(err error, format string, args ...interface{}) error {
	return &couldNotWriteFile{errors.Wrapf(err, format, args...)}
}

func IsCouldNotWriteFile(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*couldNotWriteFile)

	return ok
}

func WrapEmptyMessage(err error, format string, args ...interface{}) error {
	return &emptyMessage{errors.Wrapf(err, format, args...)}
}

func IsEmptyMessage(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*emptyMessage)

	return ok
}