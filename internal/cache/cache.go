package cache

import "errors"

var (
	ErrDuplicateEntry = errors.New("cache: entry already exists")
	ErrNoRecord       = errors.New("cache: no matching record found")
)

type Cache struct {
	Engagements *engagementCache
}

func New() *Cache {
	return &Cache{
		Engagements: newEngagementCache(),
	}
}
