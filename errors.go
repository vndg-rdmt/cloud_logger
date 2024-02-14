package cloudlogger

import (
	"errors"
	"fmt"
)

func errBadKVString(pos int, raw string) error {
	return errors.New(
		fmt.Sprintf("cloudlogger: bad key-value string, position %d raw '%s'. Expected format - key1:ENV_VARIABLE1,key2:ENV_VARIABLE2,...", pos, raw),
	)
}
