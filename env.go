package goenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func CanGet(key string, def string) string {
	fileEnvTag := key + "_FILE"

	filename := os.Getenv(fileEnvTag)

	if filename != "" {
		return MustGetFromFile(filename)
	}
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func CanGetSlice(key string, def []string) []string {
	val := CanGet(key, "")
	if val == "" {
		return def
	}
	return strings.Split(val, ",")
}

func CanGetBool(key string, def bool) bool {
	val := CanGet(key, "")
	if val == "" {
		return def
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type bool, error: " + err.Error()))
	}
	return b
}

func CanGetInt(key string, def int) int {
	val := CanGet(key, "")
	if val == "" {
		return def
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type int, error: " + err.Error()))
	}
	return i
}

func CanGetInt32(key string, def int32) int32 {
	val := CanGet(key, "")
	if val == "" {
		return def
	}
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type int32, error: " + err.Error()))
	}
	return int32(i)
}

func CanGetInt64(key string, def int64) int64 {
	val := CanGet(key, "")
	if val == "" {
		return def
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type int64, error: " + err.Error()))
	}
	return int64(i)
}

/*
MustGet will return the env or panic if it is not present.

It will also check if the key exists with _FILE at the end,
if a key exists like that, the value from the file will be loaded.
*/
func MustGet(key string) string {
	fileEnvTag := key + "_FILE"

	val := CanGet(key, fileEnvTag)
	if strings.HasSuffix(key, "_FILE") && val == fileEnvTag {
		// panic if we are checking for key_FILE and no match was found
		panic(fmt.Errorf("ENV missing: " + key))
	}

	if val == fileEnvTag {
		return MustGetFromFile(MustGet(fileEnvTag))
	}

	return val
}

func MustGetSlice(key string) []string {
	v := MustGet(key)
	return strings.Split(v, ",")
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(key string) bool {
	val := MustGet(key)
	b, err := strconv.ParseBool(val)

	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type bool, error: " + err.Error()))
	}
	return b
}

// MustGetInt will return the env as int or panic if it is not present
func MustGetInt(key string) int {
	val := MustGet(key)
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type int, error: " + err.Error()))
	}
	return int(i)
}

// MustGetInt32 will return the env as int or panic if it is not present
func MustGetInt32(key string) int32 {
	val := MustGet(key)
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type int32, error: " + err.Error()))
	}
	return int32(i)
}

// MustGetInt64 will return the env as int64 or panic if it is not present
func MustGetInt64(key string) int64 {
	val := MustGet(key)
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(fmt.Errorf("ENV " + key + " not type int64, error: " + err.Error()))
	}
	return int64(i)
}

func MustGetFromFile(filename string) string {
	secret, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("ENV missing: " + filename))
	}
	return string(secret)
}
