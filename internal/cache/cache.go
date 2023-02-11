package cache

import (
	"L0/internal/entity"
	"sync"
)

type Cache struct {
	data map[string]*entity.Model
	mu   sync.Mutex
}

func NewCache() *Cache {
	c := Cache{map[string]*entity.Model{}, sync.Mutex{}}
	return &c
}
func (c *Cache) Get(str string) *entity.Model {
	return c.data[str]
}
func (c *Cache) Set(model *entity.Model) {
	c.mu.Lock()
	c.data[model.OrderUid] = model
	c.mu.Unlock()

}
func (c *Cache) SetAll(models []*entity.Model) {
	for _, model := range models {
		c.Set(model)
	}
}
