package env

import "time"

// override the log level
func LogLevel() string {
	return String("LOG_LEVEL", "")
}

// number of go routines to run
func Routines() int {
	return Int("GOROUTINES", 50)
}

// mongo address string
func MongoAddress() string {
	return String("MONGO_ADDRESS", "mongos:27017")
}

// mongo username
func MongoUser() string {
	return String("MONGO_USER", "admin")
}

// mongo password
func MongoPassword() string {
	return String("MONGO_PASSWORD", "mangotango")
}

// mongo authentication database
func MongoAuthenticationDatabase() string {
	return String("MONGO_AUTHENTICATION_DATABASE", "admin")
}

// maximum number of sockets to allow connecting to mongo
func MongoPool() int {
	return Int("MONGO_CONNECTION_POOL_SIZE", 200)
}

// port
func Port(def int) int {
	return Int("PORT", def)
}

// server bind address
func Address() string {
	return String("SERVER_ADDRESS", "0.0.0.0")
}

// debug and read data from stdin
func Debug() bool {
	return Bool("DEBUG", false)
}

// Port on which prometheus will listen and provide metrics
func PrometheusPort(def int) int {
	return Int("PROMETHEUS_PORT", def)
}

// Enables prometheus to record timing of gRPC handlers into a histogram
func EnablePrometheusHandlingTimeHistogram() bool {
	return Bool("ENABLE_PROMETHEUS_HANDLING_TIME_HISTOGRAM", true)
}

// Enables GRPC reflection to allow usage of GRPC debugging tools
func EnableGrpcReflection() bool {
	return Bool("ENABLE_GRPC_REFLECTION", false)
}

// The maximum duration a gRPC connection will be kept alive by a gRPC server before forcing a reconnect, to promote connection balancing
func MaxGrpcConnectionAge() time.Duration {
	return Duration("MAX_GRPC_CONNECTION_AGE", time.Second*360)
}

// The grace period the gRPC server will use when recycling connections due to the maximum connection age setting, before forcibly closing the connection.
func MaxGrpcConnectionAgeGracePeriod() time.Duration {
	return Duration("MAX_GRPC_CONNECTION_AGE_GRACE_PERIOD", time.Second*5)
}

// Whether to collect profiling data
func PprofEnabled() bool {
	return Bool("PPROF_ENABLED", true)
}

// Port to serve profiling data on
func PprofPort() int {
	return Int("PPROF_PORT", 6060)
}

// Port to service health and liveness endpoints on
func HealthPort() int {
	return Int("HEALTH_PORT", 2112)
}
