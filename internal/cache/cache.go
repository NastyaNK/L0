package cache

import "L0/internal/entity"

type Cache struct {
	data map[string]*entity.Model
}

func NewCache() *Cache {
	c := Cache{map[string]*entity.Model{}}
	return &c
}
func (c *Cache) Get(str string) *entity.Model {
	return c.data[str]
}
func (c *Cache) Set(model *entity.Model) {
	c.data[model.OrderUid] = model
}
func (c *Cache) SetAll(models []*entity.Model) {
	for _, model := range models {
		c.Set(model)
	}
}
