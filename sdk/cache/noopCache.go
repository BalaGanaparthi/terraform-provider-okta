// DO NOT EDIT LOCAL SDK - USE v3 okta-sdk-golang FOR API CALLS THAT DO NOT EXIST IN LOCAL SDK
package cache

import "net/http"

type NoOpCache struct{}

func NewNoOpCache() Cache {
	return NoOpCache{}
}

func (c NoOpCache) Get(key string) *http.Response {
	return nil
}

func (c NoOpCache) Set(key string, value *http.Response) {
}

func (c NoOpCache) GetString(key string) string {
	return ""
}

func (c NoOpCache) SetString(key, value string) {
}

func (c NoOpCache) Delete(key string) {
}

func (c NoOpCache) Clear() {
}

func (c NoOpCache) Has(key string) bool {
	return false
}
