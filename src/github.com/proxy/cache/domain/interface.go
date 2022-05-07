package domain

type CacheRepository interface {
	Get(key string) CacheResponse
	Set(key string, data string) CacheResponse
	Del(key string) CacheResponse
	Connect() CacheConnection
	Error() error
}
