package env

import (
	"os"
)

type lookerUpper interface {
	lookup(key string) (string, bool)
}

type osLookup struct{}

// lookup returns a value and bool that is true if the environment variable exists.
// this is a wrapper around os.LookupEnv for convenience and testability
func (o *osLookup) lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}
