package cloudlogger

import (
	"fmt"
	"testing"
)

type caseParseKeyValueString struct {
	Raw     string
	IsError bool
	Result  []KV
}

func getCasesTestParseKeyValueString() []caseParseKeyValueString {
	return []caseParseKeyValueString{
		{
			Raw:     "key:VALUE,aboba:ABOBA,BIBIBI:BABABA",
			IsError: false,
			Result: []KV{
				{
					Key:   "key",
					Value: "VALUE",
				},
				{
					Key:   "aboba",
					Value: "ABOBA",
				},
				{
					Key:   "BIBIBI",
					Value: "BABABA",
				},
			},
		},
		{
			Raw:     "key:VALUE,:ABOBA,BIBIBI:BABABA",
			IsError: true,
			Result: []KV{
				{
					Key:   "key",
					Value: "VALUE",
				},
				{
					Key:   "aboba",
					Value: "ABOBA",
				},
				{
					Key:   "BIBIBI",
					Value: "BABABA",
				},
			},
		},
		{
			Raw:     "",
			IsError: false,
			Result:  []KV{},
		},
		{
			Raw:     "key,",
			IsError: true,
			Result:  []KV{},
		},
		{
			Raw:     ",",
			IsError: true,
			Result:  []KV{},
		},
		{
			Raw:     ":,:,",
			IsError: true,
			Result:  []KV{},
		},
		{
			Raw:     "1yf1y12ybu12bdyudb12u1bdudbudbu12buud2:1,::",
			IsError: true,
			Result:  []KV{},
		},
	}
}

func TestParseKeyValueString(t *testing.T) {

	failed := func(i int, msg string) {
		t.Errorf("[case %d] failed: %s", i, msg)
	}

	done := func(i int) {
		t.Logf("[case %d] done", i)
	}

	for i, v := range getCasesTestParseKeyValueString() {
		res, err := parseKeyValueString(v.Raw)

		if (err != nil) != v.IsError {
			failed(i, fmt.Sprintf("err '%v' != case.IsErorr '%v'", err, v.IsError))
			continue
		}

		if v.IsError {
			done(i)
			continue
		}

		if len(v.Result) != len(res) {
			failed(i, fmt.Sprintf("expected len(res) == %d", len(v.Result)))
			continue
		}

		var INT bool
		for j, kv := range v.Result {
			if (kv.Key != res[j].Key) || (kv.Value != res[j].Value) {
				failed(i, fmt.Sprintf("expected - %v, got - %v", kv, res[j]))
				INT = true
				break
			}
		}
		if INT {
			continue
		}

		done(i)
	}
}
