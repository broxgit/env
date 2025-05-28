package env

import (
	"testing"
	"time"
)

type fakeLookup struct {
	env map[string]string
}

func (f *fakeLookup) lookup(key string) (string, bool) {
	v, ok := f.env[key]
	return v, ok
}

func TestDefaultEnvSet(t *testing.T) {
	defaultEnvSet.lookerUpper = &fakeLookup{env: map[string]string{
		"test_str":      "bar",
		"test_int":      "14",
		"test_bool":     "true",
		"test_date":     "1999-03-15",
		"test_duration": "100ms",
		"empty":         "",
	}}

	if String("test_str", "xxx") != "bar" {
		t.Errorf("Invalid string environment variable returned")
	}
	if String("missing_test_str", "xxx") != "xxx" {
		t.Errorf("Invalid default string environment variable returned")
	}
	if v := String("empty", "abc"); v != "" {
		t.Errorf("Invalid default string environment variable returned %v", v)
	}

	if v := Int("test_int", 100); v != 14 {
		t.Errorf("Invalid int environment variable returned %v", v)
	}
	if v := Int("missing_test_int", 100); v != 100 {
		t.Errorf("Invalid default int environment variable returned %v", v)
	}
	if v := Int("empty", 100); v != 100 {
		t.Errorf("Invalid default int environment variable returned %v", v)
	}

	if v := Bool("test_bool", false); v != true {
		t.Errorf("Invalid bool environment variable returned %v", v)
	}
	if v := Bool("missing_test_bool", false); v != false {
		t.Errorf("Invalid default bool environment variable returned %v", v)
	}
	if v := Bool("missing_test_bool", true); v != true {
		t.Errorf("Invalid default bool environment variable returned %v", v)
	}
	if v := Bool("empty", true); v != true {
		t.Errorf("Invalid default bool environment variable returned %v", v)
	}
	if v := Bool("empty", false); v != false {
		t.Errorf("Invalid default bool environment variable returned %v", v)
	}

	if v := Duration("test_duration", 10*time.Millisecond); v != 100*time.Millisecond {
		t.Errorf("Invalid duration environment variable returned %v", v)
	}
	if v := Duration("missing_test_duration", 1*time.Millisecond); v != 1*time.Millisecond {
		t.Errorf("Invalid default duration environment variable returned %v", v)
	}
	if v := Duration("empty", 1*time.Millisecond); v != 1*time.Millisecond {
		t.Errorf("Invalid default duration environment variable returned %v", v)
	}

	date, err := time.Parse(time.DateOnly, "1999-03-15")
	if err != nil {
		t.Errorf("Failed to parse test date")
	}

	if v := MustDate("test_date"); !v.Equal(date) {
		t.Errorf("Invalid date environment variable returned %v", v)
	}
	if v := MustDate("empty"); !v.IsZero() {
		t.Errorf("Invalid empty date environment variable returned %v", v)
	}
}
