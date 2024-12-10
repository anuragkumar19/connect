package env

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func Str(key string, ptr *string) {
	if v := os.Getenv(key); v != "" {
		*ptr = v
	}
}

func Duration(key string, ptr *time.Duration) error {
	if v := os.Getenv(key); v != "" {
		parsed, err := time.ParseDuration(v)
		if err != nil {
			return fmt.Errorf("%w : %w", parseErr(key), err)
		}
		*ptr = parsed
	}
	return nil
}

func Int(key string, ptr *int) error {
	if v := os.Getenv(key); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("%w : %w", parseErr(key), err)
		}
		*ptr = parsed
	}
	return nil
}

func Int32(key string, ptr *int32) error {
	if v := os.Getenv(key); v != "" {
		parsed, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return fmt.Errorf("%w : %w", parseErr(key), err)
		}
		if parsed < math.MinInt32 || parsed > math.MaxInt32 {
			return fmt.Errorf("%w : value not in range of int32", parseErr(key))
		}
		*ptr = int32(parsed)
	}
	return nil
}

func Int64(key string, ptr *int64) error {
	if v := os.Getenv(key); v != "" {
		parsed, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return fmt.Errorf("%w : %w", parseErr(key), err)
		}
		*ptr = parsed
	}
	return nil
}

func Uint32(key string, ptr *uint32) error {
	if v := os.Getenv(key); v != "" {
		parsed, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return fmt.Errorf("%w : %w", parseErr(key), err)
		}
		if parsed < 0 || parsed > math.MaxUint32 {
			return fmt.Errorf("%w : value not in range of uint32", parseErr(key))
		}
		*ptr = uint32(parsed)
	}
	return nil
}

func Uint8(key string, ptr *uint8) error {
	if v := os.Getenv(key); v != "" {
		parsed, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return fmt.Errorf("%w : %w", parseErr(key), err)
		}
		if parsed < 0 || parsed > math.MaxUint8 {
			return fmt.Errorf("%w : value not in range of int8", parseErr(key))
		}
		*ptr = uint8(parsed)
	}
	return nil
}

func Bool(key string, ptr *bool) error {
	if v := os.Getenv(key); v != "" {
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("%w : %w", parseErr(key), err)
		}
		*ptr = parsed
	}
	return nil
}

func parseErr(key string) error {
	return fmt.Errorf("failed to parse env variable %s", key)
}
