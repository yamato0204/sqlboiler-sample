package cache

import "github.com/yamato0204/sqlboiler-sample/internal/infra/datamodel"

type Cache struct {
	datamodel.User
}

func NewCache(user datamodel.User) *Cache {
	return &Cache{
		User: user,
	}
}

func (c *Cache) Table() string {
	return "users"
}
