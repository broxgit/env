package env

import (
	"fmt"
	"time"
)

// environmentSet represents a set of methods to import various types from environment variables
type environmentSet interface {
	Bool(key string, value bool) bool
	Duration(key string, value time.Duration) time.Duration
	Date(key string, value time.Duration) time.Duration
	Int(key string, value int) int
	String(key, value string) string
}

// environment represents a structure to import environment variables with a lookup
type environment struct {
	environmentSet
	lookerUpper
}

// String returns a string that is populated by the environment variable or given a default value.
func (e *environment) String(key, value string) string {
	if v, ok := e.lookup(key); ok {
		return v
	}
	return value
}

// Int returns an int  that is populated by the environment variable or given a default value.
func (e *environment) Int(key string, value int) int {
	v := e.String(key, "")
	ret, err := parseInt(v)
	if err != nil {
		return value
	}
	return ret
}

// Bool returns a bool  that is populated by the environment variable or given a default value.
func (e *environment) Bool(key string, value bool) bool {
	v := e.String(key, "")
	ret, err := parseBool(v)
	if err != nil {
		return value
	}
	return ret
}

// Duration returns a duration that is populated by the environment variable or given a default value.
func (e *environment) Duration(key string, value time.Duration) time.Duration {
	v := e.String(key, fmt.Sprint(value))
	ret, err := time.ParseDuration(v)
	if err != nil {
		return value
	}
	return ret
}

// Date returns a date that is populated by the format YYYY-MM-DD
func (e *environment) Date(key string) (time.Time, error) {
	v := e.String(key, "")
	if v == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(time.DateOnly, v)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
