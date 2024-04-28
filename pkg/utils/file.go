package utils

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/chein-huang/errorc"
	"github.com/pkg/errors"
)

var (
	ErrNotFile = fmt.Errorf("path not file")
)

func OpenFile(name string, flag int, perm fs.FileMode, overwrite bool) (*os.File, error) {
	stat, err := os.Stat(name)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, errorc.AddField(err, "file name", name)
	}

	if err == nil {
		if stat.IsDir() {
			return nil, errorc.AddField(ErrNotFile, "file name", name)
		}

		if overwrite {
			err = os.Remove(name)
			if err != nil {
				return nil, errorc.AddField(err, "file name", name)
			}
		}
	}

	return os.OpenFile(name, flag, perm)
}
