package cmdx

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Init loads the .env file in the current working directory if exists
func Init() {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("unable to load .env file: `%s`", err.Error()))
	}
}

// EnvOrInt returns the value from the environment by the given key if provided,
// otherwise returns def value. It panics if the provided value could not be parsed as integer.
func EnvOrInt(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Errorf("unable to parse integer value from the environment variable %s -> `%s`: %s", key, v, err.Error()))
		}

		return i
	}

	return def
}

// EnvOrStr returns the value from the environment by the given key if provided,
// otherwise returns def value.
func EnvOrStr(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return def
}

// EnvOrBool returns the value from the environment by the given key if provided,
// otherwise returns def value. It panics if the provided value could not be parsed as integer.
func EnvOrBool(key string, def bool) bool {
	if v, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseBool(v)
		if err != nil {
			panic(fmt.Errorf("unable to parse boolean value from the environment variable %s -> `%s`: %s", key, v, err.Error()))
		}

		return i
	}

	return def
}

// EnvOrStrArr returns the array of values from the environment by the given key if provided.
// If the given key exact matches with an environment variable it tries to parse the value using encoding/csv package.
// If there are environment variables that have prefix matches with the given key and have suffix _[0-9+] they also appended the returned list.
func EnvOrStrArr(key string, def []string) []string {
	ret := make([]string, 0, 10)

	if v, ok := os.LookupEnv(key); ok {
		a, err := csv.NewReader(strings.NewReader(v)).Read()
		if err != nil {
			panic(fmt.Errorf("unable to parse CSV from the environment variable %s -> `%s`: %s", key, v, err.Error()))
		}

		ret = append(ret, a...)
	}

	pattern := regexp.MustCompile(fmt.Sprintf("^(%s_(\\d+))=.*", regexp.QuoteMeta(key)))

	vals := make(map[int64]string)
	indices := make([]int, 0, 10)
	for _, v := range os.Environ() {
		if v := pattern.FindAllStringSubmatch(v, 1); len(v) == 1 && len(v[0]) == 3 {
			index, _ := strconv.ParseInt(v[0][2], 10, 64)
			vals[index] = os.Getenv(v[0][1])
			indices = append(indices, int(index))
		}
	}

	sort.Ints(indices)
	for _, i := range indices {
		ret = append(ret, vals[int64(i)])
	}

	if len(ret) > 0 {
		return ret
	}

	return def
}

// EnvOrIntArr returns the array of values from the environment by the given key if provided.
// If the given key exact matches with an environment variable it tries to parse the value using encoding/csv package.
// If there are environment variables that have prefix matches with the given key and have suffix _[0-9+] they also appended the returned list.
// A value that cannot be parsed as an integer will cause panic.
func EnvOrIntArr(key string, def []int) []int {
	sDef := make([]string, len(def))
	for i, v := range def {
		sDef[i] = strconv.FormatInt(int64(v), 10)
	}

	vals := EnvOrStrArr(key, sDef)
	ret := make([]int, 0, len(vals))

	for _, v := range vals {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}

		iv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Errorf("unable to parse the integer value from the environment variable %s -> `%s`: %s", key, v, err.Error()))
		}

		ret = append(ret, int(iv))
	}

	return ret
}

// EnvOrBoolArr returns the array of values from the environment by the given key if provided.
// If the given key exact matches with an environment variable it tries to parse the value using encoding/csv package.
// If there are environment variables that have prefix matches with the given key and have suffix _[0-9+] they also appended the returned list.
// A value that cannot be parsed as an boolean will cause panic.
func EnvOrBoolArr(key string, def []bool) []bool {
	sDef := make([]string, len(def))
	for i, v := range def {
		sDef[i] = strconv.FormatBool(v)
	}

	vals := EnvOrStrArr(key, sDef)
	ret := make([]bool, 0, len(vals))

	for _, v := range vals {
		bv, err := strconv.ParseBool(v)
		if err != nil {
			panic(fmt.Errorf("unable to parse the integer value from the environment variable %s -> `%s`: %s", key, v, err.Error()))
		}

		ret = append(ret, bv)
	}

	return ret
}
