package api

import (
	"context"
	"errors"
	"sync"

	"github.com/avag-sargsyan/stricache/proto/stricache"
)

type StringItem struct {
	Value string
}

type IntItem struct {
	Value int64
}

type FloatItem struct {
	Value float64
}

type Cache struct {
	Strings *stringCache
	Ints    *intCache
	Floats  *floatCache
	mu      sync.RWMutex
}

type stringCache struct {
	items map[string]StringItem
	list  []string
}

type intCache struct {
	items map[string]IntItem
	list  []int64
}

type floatCache struct {
	items map[string]FloatItem
	list  []float64
}

func NewCacheService() *Cache {
	cstr := stringCache{
		map[string]StringItem{},
		[]string{},
	}
	cint := intCache{
		map[string]IntItem{},
		[]int64{},
	}
	cflt := floatCache{
		map[string]FloatItem{},
		[]float64{},
	}
	C := &Cache{
		Strings: &cstr,
		Ints:    &cint,
		Floats:  &cflt,
	}
	return C
}

func (c *Cache) AddString(ctx context.Context, item *stricache.StringItem) (*stricache.StringItem, error) {
	c.mu.Lock()
	c.Strings.items[item.Key] = StringItem{
		Value: item.Value,
	}
	c.Strings.list = append(c.Strings.list, item.Value)
	c.mu.Unlock()
	return item, nil
}

func (c *Cache) AddInt(ctx context.Context, item *stricache.IntItem) (*stricache.IntItem, error) {
	c.mu.Lock()
	c.Ints.items[item.Key] = IntItem{
		Value: item.Value,
	}
	c.Ints.list = append(c.Ints.list, item.Value)
	c.mu.Unlock()
	return item, nil
}

func (c *Cache) AddFloat(ctx context.Context, item *stricache.FloatItem) (*stricache.FloatItem, error) {
	c.mu.Lock()
	c.Floats.items[item.Key] = FloatItem{
		Value: item.Value,
	}
	c.Floats.list = append(c.Floats.list, item.Value)
	c.mu.Unlock()
	return item, nil
}

func (c *Cache) UnshiftString(ctx context.Context, item *stricache.StringItem) (*stricache.StringItem, error) {
	c.mu.Lock()
	c.Strings.items[item.Key] = StringItem{
		Value: item.Value,
	}
	c.Strings.list = append([]string{item.Value}, c.Strings.list...)

	c.mu.Unlock()
	return item, nil
}

func (c *Cache) UnshiftInt(ctx context.Context, item *stricache.IntItem) (*stricache.IntItem, error) {
	c.mu.Lock()
	c.Ints.items[item.Key] = IntItem{
		Value: item.Value,
	}
	c.Ints.list = append([]int64{item.Value}, c.Ints.list...)
	c.mu.Unlock()
	return item, nil
}

func (c *Cache) UnshiftFloat(ctx context.Context, item *stricache.FloatItem) (*stricache.FloatItem, error) {
	c.mu.Lock()
	c.Floats.items[item.Key] = FloatItem{
		Value: item.Value,
	}
	c.Floats.list = append([]float64{item.Value}, c.Floats.list...)
	c.mu.Unlock()
	return item, nil
}

func (c *Cache) GetString(ctx context.Context, args *stricache.GetKey) (*stricache.StringItem, error) {
	key := args.Key
	c.mu.RLock()
	value, exists := c.Strings.items[key]
	if !exists {
		c.mu.RUnlock()
		return nil, errors.New("No key found")
	}
	c.mu.RUnlock()
	return &stricache.StringItem{
		Key:   key,
		Value: value.Value,
	}, nil
}

func (c *Cache) GetInt(ctx context.Context, args *stricache.GetKey) (*stricache.IntItem, error) {
	key := args.Key
	c.mu.RLock()
	value, exists := c.Ints.items[key]
	if !exists {
		c.mu.RUnlock()
		return nil, errors.New("No key found")
	}
	c.mu.RUnlock()
	return &stricache.IntItem{
		Key:   key,
		Value: value.Value,
	}, nil
}

func (c *Cache) GetFloat(ctx context.Context, args *stricache.GetKey) (*stricache.FloatItem, error) {
	key := args.Key
	c.mu.RLock()
	value, exists := c.Floats.items[key]
	if !exists {
		c.mu.RUnlock()
		return nil, errors.New("No key found")
	}
	c.mu.RUnlock()
	return &stricache.FloatItem{
		Key:   key,
		Value: value.Value,
	}, nil
}

func (c *Cache) DeleteString(ctx context.Context, args *stricache.GetKey) (*stricache.Success, error) {
	c.mu.Lock()
	value, exists := c.Strings.items[args.Key]
	if exists {
		delete(c.Strings.items, args.Key)
		for i, v := range c.Strings.list {
			if value.Value == v {
				c.Strings.list = append(c.Strings.list[:i], c.Strings.list[i+1:]...)
				break
			}
		}
	}

	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) DeleteInt(ctx context.Context, args *stricache.GetKey) (*stricache.Success, error) {
	c.mu.Lock()
	value, exists := c.Ints.items[args.Key]
	if exists {
		delete(c.Ints.items, args.Key)
		for i, v := range c.Ints.list {
			if value.Value == v {
				c.Ints.list = append(c.Ints.list[:i], c.Ints.list[i+1:]...)
				break
			}
		}
	}

	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) DeleteFloat(ctx context.Context, args *stricache.GetKey) (*stricache.Success, error) {
	c.mu.Lock()
	value, exists := c.Floats.items[args.Key]
	if exists {
		delete(c.Floats.items, args.Key)
		for i, v := range c.Floats.list {
			if value.Value == v {
				c.Floats.list = append(c.Floats.list[:i], c.Floats.list[i+1:]...)
				break
			}
		}
	}
	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) ShiftString(ctx context.Context, e *stricache.EmptyR) (*stricache.Success, error) {
	var first string
	c.mu.Lock()
	first, c.Strings.list = c.Strings.list[0], c.Strings.list[1:len(c.Strings.list)-1]
	// Sync
	for i, v := range c.Strings.items {
		if first == v.Value {
			delete(c.Strings.items, i)
		}
	}
	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) ShiftInt(ctx context.Context, e *stricache.EmptyR) (*stricache.Success, error) {
	var first int64
	c.mu.Lock()
	first, c.Ints.list = c.Ints.list[0], c.Ints.list[1:len(c.Ints.list)-1]
	// Sync
	for i, v := range c.Ints.items {
		if first == v.Value {
			delete(c.Ints.items, i)
		}
	}
	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) ShiftFloat(ctx context.Context, e *stricache.EmptyR) (*stricache.Success, error) {
	var first float64
	c.mu.Lock()
	first, c.Floats.list = c.Floats.list[0], c.Floats.list[1:len(c.Floats.list)-1]
	// Sync
	for i, v := range c.Floats.items {
		if first == v.Value {
			delete(c.Floats.items, i)
		}
	}
	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) PopString(ctx context.Context, e *stricache.EmptyR) (*stricache.Success, error) {
	var last string
	c.mu.Lock()
	c.Strings.list, last = c.Strings.list[:len(c.Strings.list)-1], c.Strings.list[len(c.Strings.list)-1]
	// Sync
	for i, v := range c.Strings.items {
		if last == v.Value {
			delete(c.Strings.items, i)
		}
	}
	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) PopInt(ctx context.Context, e *stricache.EmptyR) (*stricache.Success, error) {
	var last int64
	c.mu.Lock()
	c.Ints.list, last = c.Ints.list[:len(c.Ints.list)-1], c.Ints.list[len(c.Ints.list)-1]
	// Sync
	for i, v := range c.Ints.items {
		if last == v.Value {
			delete(c.Ints.items, i)
		}
	}
	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}

func (c *Cache) PopFloat(ctx context.Context, e *stricache.EmptyR) (*stricache.Success, error) {
	var last float64
	c.mu.Lock()
	c.Floats.list, last = c.Floats.list[:len(c.Floats.list)-1], c.Floats.list[len(c.Floats.list)-1]
	// Sync
	for i, v := range c.Floats.items {
		if last == v.Value {
			delete(c.Floats.items, i)
		}
	}
	c.mu.Unlock()
	return &stricache.Success{
		Success: true,
	}, nil
}
