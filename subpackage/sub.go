package subpackage

import "github.com/go-errors/errors"

func Deeper() error {
	return errors.New("hello from subpackage")
}
