package cloudlogger

import (
	"os"
	"testing"
)

// func TestFileLoggingCreate(t *testing.T) {
// 	conf := DefaultConfig()
// 	conf.FileWriters = append(conf.FileWriters, FileLoggingConfig{
// 		Path:  "./temp.log",
// 		Level: "debug",
// 	})
// 	defer os.Remove("./temp.log")

// 	l := New(conf)
// 	l.Info("test")

// 	f, err := os.Stat("./temp.log")
// 	if err != nil {

// 		t.Error("[failed]")
// 	}

// 	if f.Size() > 0 {
// 		t.Log("[passed]")
// 	} else {
// 		t.Error("[passed]")
// 	}
// }

// func TestFileLoggingExists(t *testing.T) {
// 	f, err := os.CreateTemp("", "")
// 	if err != nil {
// 		t.Fatalf("[failed] - %s", err.Error())
// 	}
// 	defer f.Close()
// 	defer os.Remove(f.Name())

// 	conf := DefaultConfig()
// 	conf.FileWriters = append(conf.FileWriters, FileLoggingConfig{
// 		Path:  f.Name(),
// 		Level: "debug",
// 	})

// 	l := New(conf)
// 	l.Info("test")

// 	fa, err := os.Stat(f.Name())
// 	if err != nil {
// 		t.Error("[failed]")
// 	}

// 	if fa.Size() > 0 {
// 		t.Log("[passed]")
// 	} else {
// 		t.Error("[passed]")
// 	}
// }

func TestFileLoggingExists(t *testing.T) {

	conf := DefaultConfig()

	if addr := os.Getenv("TEST_GRAYLOG_ADDR"); addr != "" {
		conf.GraylogAddr = addr
	}

	conf.getWriterLoggerCore(os.Stdout, "info")
	l := New(conf)
	l.Info("THIS IS A GO.CLOUDLOGGER TEST")
}
