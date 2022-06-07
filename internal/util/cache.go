package util

import "time"

type cacheData[T any] struct {
	Expires time.Time
	Data    T
}

type Cache[T any] struct {
	data map[string]cacheData[T]
}

func makeCacheData[T any]() map[string]cacheData[T] {
	return make(map[string]cacheData[T])
}

func MakeCache[T any]() Cache[T] {
	return Cache[T]{
		data: makeCacheData[T](),
	}
}

func (cache Cache[T]) Size() int {
	return len(cache.data)
}

func (cache *Cache[T]) Put(key string, data T, duration int) {
	cache.data[key] = cacheData[T]{
		Data:    data,
		Expires: time.Now().Add(time.Duration(duration) * time.Second),
	}
}

func (cache Cache[T]) Get(key string) (T, bool) {
	var empty T
	data, ok := cache.data[key]
	if !ok {
		return empty, false
	}

	if time.Now().After(data.Expires) {
		return empty, false
	}

	return data.Data, true
}

func (cache Cache[T]) Remove(key string) {
	delete(cache.data, key)
}

func (cache Cache[T]) Flush() {
	cache.data = makeCacheData[T]()
}
