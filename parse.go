package env

import (
	"math/bits"
	"strconv"
	"time"
)

func parseInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func parseInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func parseInt32(value string) (int32, error) {
	v, err := strconv.ParseInt(value, 10, 32)
	return int32(v), err
}

func parseInt16(value string) (int16, error) {
	v, err := strconv.ParseInt(value, 10, 16)
	return int16(v), err
}

func parseInt8(value string) (int8, error) {
	v, err := strconv.ParseInt(value, 10, 8)
	return int8(v), err
}

func parseUint(value string) (uint, error) {
	v, err := strconv.ParseUint(value, 10, bits.UintSize)
	return uint(v), err
}

func parseUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

func parseUint32(value string) (uint32, error) {
	v, err := strconv.ParseUint(value, 10, 32)
	return uint32(v), err
}

func parseUint16(value string) (uint16, error) {
	v, err := strconv.ParseUint(value, 10, 16)
	return uint16(v), err
}

func parseUint8(value string) (uint8, error) {
	v, err := strconv.ParseUint(value, 10, 8)
	return uint8(v), err
}

func parseFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func parseFloat32(value string) (float32, error) {
	v, err := strconv.ParseFloat(value, 32)
	return float32(v), err
}

func parseBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func parseDuration(value string) (time.Duration, error) {
	return time.ParseDuration(value)
}
