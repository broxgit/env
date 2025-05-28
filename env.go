package env

import "time"

var (
	// used to override in testing
	defaultEnvSet = &environment{
		lookerUpper: &osLookup{},
	}
)

// String returns a string that is populated by the environment variable or given a default value.
func String(key, value string) string {
	return defaultEnvSet.String(key, value)
}

// Int returns an int  that is populated by the environment variable or given a default value.
func Int(key string, value int) int {
	return defaultEnvSet.Int(key, value)
}

// Bool returns an bool  that is populated by the environment variable or given a default value.
func Bool(key string, value bool) bool {
	return defaultEnvSet.Bool(key, value)
}

// Duration returns a duration that is populated by the environment variable or given a default value.
func Duration(key string, value time.Duration) time.Duration {
	return defaultEnvSet.Duration(key, value)
}

// Date returns a date that is populated by the format YYYY-MM-DD. Empty string returns a zero value time object
func Date(key string) (time.Time, error) {
	return defaultEnvSet.Date(key)
}

// MustDate returns a date that is populated by the format YYYY-MM-DD.
// Empty string returns a zero value time object panics if parse fails
func MustDate(key string) time.Time {
	t, err := Date(key)
	if err != nil {
		panic("Parse date failed " + err.Error())
	}
	return t
}
