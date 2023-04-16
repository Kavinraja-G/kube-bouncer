package utils

import "os"

// helper function for fetching envs with defaults
func GetEnv(env, defaults string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}
	return defaults
}
