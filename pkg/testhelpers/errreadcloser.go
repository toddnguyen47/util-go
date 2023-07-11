package testhelpers

import "errors"

// ErrReadCloser - Ref: https://stackoverflow.com/a/45126402/6323360
type ErrReadCloser int

func (m ErrReadCloser) Read(_ []byte) (n int, err error) {
	return 0, errors.New("some error")
}

func (m ErrReadCloser) Close() error {
	return nil
}
