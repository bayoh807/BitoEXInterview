package providders

import "os"

type provider struct {
}

var Provider provider

func (p *provider) GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
