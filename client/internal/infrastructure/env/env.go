package env

import (
	"os"
	"strconv"
)

// All variables for project
var (
	DataServiceHost = Getter("DATA_SERVICE_HOST", "http://host.docker.internal:8000/api/%s")
)

// Getter -
func Getter(key, defaultValue string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return defaultValue
}

// GetterInt -
func GetterInt(key string, defaultValue int) int {
	env, ok := os.LookupEnv(key)
	if ok {
		res, err := strconv.ParseInt(env, 10, 32)
		if err == nil {
			return int(res)
		}
	}
	return defaultValue
}

func GetterBool(key string, defaultValue bool) bool {
	env, ok := os.LookupEnv(key)
	if ok {
		res, err := strconv.ParseBool(env)
		if err == nil {
			return res
		}
	}
	return defaultValue
}
