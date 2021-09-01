package dummy

import (
	"github.com/go-errors/errors"
	"github.com/grongor/tmp-dummy-package/subpackage"
)

func HelloFromAnotherPackage() error {
	return errors.Errorf("something happened: %w", subpackage.Deeper())
}
