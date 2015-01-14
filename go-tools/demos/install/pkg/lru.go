package lru

type Cache struct {
	MaxEntries int
}

func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
	}
}
