package dir

import (
	"errors"
	"os"
)

func Create() error {
	err := os.MkdirAll(Get(), 0777)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}

		return err
	}

	return nil
}
