package cloudlogger

import (
	"os"
	"testing"
)

func TestWritersCount(t *testing.T) {
	a := []Config{
		{
			// FileWriters: []FileLoggingConfig{
			// 	{},
			// 	{},
			// 	{},
			// },
			GraylogAddr: "",
			Stdout:      true,
		},
		{
			// FileWriters: []FileLoggingConfig{
			// 	{},
			// 	{},
			// },
			GraylogAddr: "",
			Stdout:      false,
		},
		{
			// FileWriters: []FileLoggingConfig{
			// 	{},
			// 	{},
			// },
			GraylogAddr: "123",
			Stdout:      true,
		},
	}

	for _, v := range a {
		// count := len(v.FileWriters)
		var count int

		if v.Stdout {
			count++
		}

		if v.GraylogAddr != "" {
			count++
		}

		if v.getWritersCount() != count {
			t.Error("[failed]")
		} else {
			t.Log("[passed]")
		}
	}
}

func TestGetMessageField(t *testing.T) {
	conf := DefaultConfig()
	res := conf.getMessageField()
	if res != DefaultFieldMessage {
		t.Error("[failed]")
	}

	conf.FieldMessage = ""

	if res := conf.getMessageField(); res != DefaultFieldMessage {
		t.Error("[failed]")
	} else {
		t.Log("[passed]")
	}
}

func TestGetNameField(t *testing.T) {
	conf := DefaultConfig()
	res := conf.getNameField()
	if res != DefaultFieldName {
		t.Error("[failed]")
	}

	conf.FieldName = ""

	if res := conf.getNameField(); res != DefaultFieldName {
		t.Error("[failed]")
	} else {
		t.Log("[passed]")
	}
}

func TestGetFieldLevel(t *testing.T) {
	conf := DefaultConfig()
	res := conf.getFieldLevel()
	if res != DefaultFieldLevel {
		t.Error("[failed]")
	}

	conf.FieldLevel = ""

	if res := conf.getFieldLevel(); res != DefaultFieldLevel {
		t.Error("[failed]")
	} else {
		t.Log("[passed]")
	}
}

func TestGetFieldTimestamp(t *testing.T) {
	conf := DefaultConfig()
	res := conf.getFieldTimestamp()
	if res != DefaultFieldTimestamp {
		t.Error("[failed]")
	}

	conf.FieldTimestamp = ""

	if res := conf.getFieldTimestamp(); res != DefaultFieldTimestamp {
		t.Error("[failed]")
	} else {
		t.Log("[passed]")
	}
}

func TestGetEndline(t *testing.T) {
	conf := DefaultConfig()
	res := conf.getEndline()
	if res != DefaultEndline {
		t.Error("[failed]")
	}

	conf.Endline = ""

	if res := conf.getEndline(); res != DefaultEndline {
		t.Error("[failed]")
	} else {
		t.Log("[passed]")
	}
}

func TestGetStaticFields(t *testing.T) {
	c := DefaultConfig()
	c.StaticFields = []ConfShardStaticField{
		{Key: "a", Value: "b"},
		{Key: "aa", Value: "bb"},
	}

	a := c.getStaticFields()
	if len(a) != len(c.StaticFields) {
		t.Error("[failed]")
	}
}

// func TestGetFileWriters(t *testing.T) {
// 	c := DefaultConfig()
// 	c.FileWriters = []FileLoggingConfig{
// 		{
// 			Path:  "./dqwudbwyqbdgyqwbdyuwqbdqw",
// 			Level: "info",
// 		},
// 	}

// 	a, err := c.getFileWriters()
// 	if err != nil {
// 		t.Fatal("[failed]")
// 	}
// 	defer os.Remove("./dqwudbwyqbdgyqwbdyuwqbdqw")
// 	if len(a) != len(c.FileWriters) {
// 		t.Error("[failed]")
// 	}
// }

func TestGetEnvFields(t *testing.T) {
	c := DefaultConfig()

	os.Setenv("VALUE1", "VAL_1")
	os.Setenv("VALUE2", "VAL_2")

	defer func() {
		if err := recover(); err == nil {
			t.Fail()
		}
	}()

	{
		c.EnvFields = "key1:VALUE1,key2:VALUE2"

		a := c.getEnvFields()
		// if !((a != nil && err == nil) || (a == nil && err != nil)) {
		// 	t.Fatal()
		// }

		// if err != nil {
		// 	return
		// }

		if (a[0].Key != "key1") ||
			(a[0].String != "VAL_1") ||
			(a[1].Key != "key2") ||
			(a[1].String != "VAL_2") {
			t.Fatal()
		}
	}

	{
		c.EnvFields = "key1:,key2:VALUE2"
		a := c.getEnvFields()

		// if !((a != nil && err == nil) || (a == nil && err != nil)) {
		// 	t.Fatal()
		// }

		// if err != nil {
		// 	return
		// }

		if (a[0].Key != "key1") ||
			(a[0].String != "VAL_1") ||
			(a[1].Key != "key2") ||
			(a[1].String != "VAL_2") {
			t.Fatal()
		}
	}
}
