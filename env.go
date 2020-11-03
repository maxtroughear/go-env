package goenv

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CanGet(k string, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}

func CanGetSlice(k string, def []string) []string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return strings.Split(v, ",")
}

func CanGetInt32(k string, def int) int {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(fmt.Errorf("ENV err: [" + k + "]" + err.Error()))
	}
	return int(i)
}

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Errorf("ENV missing, key: " + k))
	}
	return v
}

func MustGetSlice(k string) []string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Errorf("ENV missing, key: " + k))
	}
	return strings.Split(v, ",")
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Errorf("ENV missing, key: " + k))
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		panic(fmt.Errorf("ENV err: [" + k + "]" + err.Error()))
	}
	return b
}

// MustGetInt32 will return the env as int32 or panic if it is not present
func MustGetInt32(k string) int {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Errorf("ENV missing, key: " + k))
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(fmt.Errorf("ENV err: [" + k + "]" + err.Error()))
	}
	return int(i)
}

// MustGetInt64 will return the env as int64 or panic if it is not present
func MustGetInt64(k string) int64 {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Errorf("ENV missing, key: " + k))
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(fmt.Errorf("ENV err: [" + k + "]" + err.Error()))
	}
	return i
}
