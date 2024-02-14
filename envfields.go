package cloudlogger

import (
	"strings"
)

const (
	EnvFieldIteratorSep = ","
	EnvFieldKeyValueSep = ":"
)

type KV struct {
	Key   string
	Value string
}

func parseKeyValueString(s string) ([]KV, error) {
	if s == "" {
		return []KV{}, nil
	}

	rawarr := strings.Split(s, EnvFieldIteratorSep)
	res := make([]KV, len(rawarr))

	for i, v := range rawarr {
		rawkv := strings.Split(v, EnvFieldKeyValueSep)
		if malformedRawKeyValue(rawkv) {
			return nil, errBadKVString(i, v)
		}

		res[i] = KV{
			Key:   rawkv[0],
			Value: rawkv[1],
		}
	}

	return res, nil
}

func malformedRawKeyValue(s []string) bool {
	return (len(s) != 2) ||
		(s[0] == "") ||
		(s[1] == "")
}
