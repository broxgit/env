package env

// override the log level
func LogLevel() string {
	return String("LOG_LEVEL", "")
}
