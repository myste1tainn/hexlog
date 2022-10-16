package log_test

import (
	"testing"

	"github.com/myste1tainn/hexlog"
)

func TestDestroy_whenGot5Child_removeOne_remains4(t *testing.T) {
	l := log.NewLogger(log.DefaultConfig)
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()
	c := l.NewChildLogger()
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()

	if l.GetChildLoggerCount() != 5 {
		t.Fatalf("expect child logger count to equal 5 got %d\n", l.GetChildLoggerCount())
	}

	c.Destroy()

	if l.GetChildLoggerCount() != 4 {
		t.Fatalf("expect child logger count to equal 4 got %d\n", l.GetChildLoggerCount())
	}
}

func TestDestroy_whenGot5Child_removeLastOne_remains4(t *testing.T) {
	l := log.NewLogger(log.DefaultConfig)
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()
	c := l.NewChildLogger()

	if l.GetChildLoggerCount() != 5 {
		t.Fatalf("expect child logger count to equal 5 got %d\n", l.GetChildLoggerCount())
	}

	c.Destroy()

	if l.GetChildLoggerCount() != 4 {
		t.Fatalf("expect child logger count to equal 4 got %d\n", l.GetChildLoggerCount())
	}
}

func TestDestroy_whenGot5Child_removeFirstOne_remains4(t *testing.T) {
	l := log.NewLogger(log.DefaultConfig)
	c := l.NewChildLogger()
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()
	_ = l.NewChildLogger()

	if l.GetChildLoggerCount() != 5 {
		t.Fatalf("expect child logger count to equal 5 got %d\n", l.GetChildLoggerCount())
	}

	c.Destroy()

	if l.GetChildLoggerCount() != 4 {
		t.Fatalf("expect child logger count to equal 4 got %d\n", l.GetChildLoggerCount())
	}
}
