package api

import (
	"context"

	"github.com/avag-sargsyan/stricache/internal/storage"
	"github.com/avag-sargsyan/stricache/proto/stricache"
)

// AddString(context.Context, *StringItem) (*StringItem, error)
// 	AddInt(context.Context, *IntItem) (*IntItem, error)
// 	AddFloat(context.Context, *FloatItem) (*FloatItem, error)
// 	UnshiftString(context.Context, *StringItem) (*StringItem, error)
// 	UnshiftInt(context.Context, *IntItem) (*IntItem, error)
// 	UnshiftFloat(context.Context, *FloatItem) (*FloatItem, error)
// 	GetString(context.Context, *GetKey) (*StringItem, error)
// 	GetInt(context.Context, *GetKey) (*IntItem, error)
// 	GetFloat(context.Context, *GetKey) (*FloatItem, error)
// 	DeleteString(context.Context, *GetKey) (*Success, error)
// 	DeleteInt(context.Context, *GetKey) (*Success, error)
// 	DeleteFloat(context.Context, *GetKey) (*Success, error)
// 	ShiftString(context.Context, *EmptyR) (*Success, error)
// 	ShiftInt(context.Context, *EmptyR) (*Success, error)
// 	ShiftFloat(context.Context, *EmptyR) (*Success, error)
// 	PopString(context.Context, *EmptyR) (*Success, error)
// 	PopInt(context.Context, *EmptyR) (*Success, error)
// 	PopFloat(context.Context, *EmptyR) (*Success, error)

func (c *storage.Cache) AddString(ctx context.Context, item *stricache.StringItem) (*stricache.StringItem, error) {
	c.mu.Lock()
	c.Strings.items[item.Key] = StringItem{
		Object:     item.Value,
	}
	c.Strings.list = append(c.Strings.list, item.Value)
	c.mu.Unlock()
	return item, nil
}

func (c *storage.Cache) AddInt(ctx context.Context, item *stricache.IntItem) (*stricache.IntItem, error) {
	c.mu.Lock()
	c.Ints.items[item.Key] = IntItem{
		Object:     item.Value,
	}
	c.Ints.list = append(c.Ints.list, item.Value)
	c.mu.Unlock()
	return item, nil
}

func (c *storage.Cache) AddFloat(ctx context.Context, item *stricache.FloatItem) (*stricache.FloatItem, error) {
	c.mu.Lock()
	c.Floats.items[item.Key] = FloatItem{
		Object:     item.Value,
	}
	c.Floats.list = append(c.Floats.list, item.Value)
	c.mu.Unlock()
	return item, nil
}

func (c *storage.Cache) UnshiftString(ctx context.Context, item *stricache.StringItem) (*stricache.StringItem, error) {
	c.mu.Lock()
	c.Strings.items[item.Key] = Item{
		Object:     item.Value,
	}
	c.Strings.list = append([]string{item.Value}, c.Strings.list...)

	c.mu.Unlock()
	return item, nil
}

func (c *storage.Cache) UnshiftInt(ctx context.Context, item *stricache.IntItem) (*stricache.IntItem, error) {
	c.mu.Lock()
	c.Ints.items[item.Key] = Item{
		Object:     item.Value,
	}
	c.Ints.list = append([]string{item.Value}, c.Ints.list...)
	c.mu.Unlock()
	return item, nil
}

func (c *storage.Cache) UnshiftFloat(ctx context.Context, item *stricache.FloatItem) (*stricache.FloatItem, error) {
	c.mu.Lock()
	c.Floats.items[item.Key] = Item{
		Object:     item.Value,
	}
	c.Floats.list = append([]string{item.Value}, c.Floats.list...)
	c.mu.Unlock()
	return item, nil
}

func (c *storage.Cache) GetString(ctx context.Context, args *stricache.GetKey) (*stricache.StringItem, error) {
	key := args.Key
	c.mu.RLock()
	value, exists := c.Strings.items[key]
	if !exists {
		c.mu.RUnlock()
		return nil, ErrNoKey
	}
	c.mu.RUnlock()
	return &api.Item{
		Key:        key,
		Value:      value.(Item).Object.(string),
	}, nil
}

func (c *storage.Cache) GetInt(ctx context.Context, args *stricache.GetKey) (*stricache.IntItem, error) {
	key := args.Key
	c.mu.RLock()
	value, exists := c.Ints.items[key]
	if !exists {
		c.mu.RUnlock()
		return nil, ErrNoKey
	}
	c.mu.RUnlock()
	return &api.Item{
		Key:        key,
		Value:      value.(Item).Object.(string),
	}, nil
}

func (c *storage.Cache) GetFloat(ctx context.Context, args *stricache.GetKey) (*stricache.FloatItem, error) {
	key := args.Key
	c.mu.RLock()
	value, exists := c.Floats.items[key]
	if !exists {
		c.mu.RUnlock()
		return nil, ErrNoKey
	}
	c.mu.RUnlock()
	return &api.Item{
		Key:        key,
		Value:      value.(Item).Object.(string),
	}, nil
}

func (c *storage.Cache) DeleteString(ctx context.Context, args *stricache.GetKey) (*stricache.Success, error) {
	c.mu.Lock()
	value, exists := c.Strings.items[key]
	if exists {
		delete(c.Strings.items, args.Key)
		for i, v := range c.Strings.list {
			if value == v {
				c.Strings.list = append(c.Strings.list[:i], c.Strings.list[i+1:]...)
				break
			}
		}
	}

	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) DeleteInt(ctx context.Context, args *stricache.GetKey) (*stricache.Success, error) {
	c.mu.Lock()
	value, exists := c.Ints.items[key]
	if exists {
		delete(c.Ints.items, args.Key)
		for i, v := range c.Ints.list {
			if value == v {
				c.Ints.list = append(c.Ints.list[:i], c.Ints.list[i+1:]...)
				break
			}
		}
	}

	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) DeleteFloat(ctx context.Context, args *stricache.GetKey) (*stricache.Success, error) {
	c.mu.Lock()
	value, exists := c.Floats.items[key]
	if exists {
		delete(c.Floats.items, args.Key)
		for i, v := range c.Floats.list {
			if value == v {
				c.Floats.list = append(c.Floats.list[:i], c.Floats.list[i+1:]...)
				break
			}
		}
	}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) ShiftString(ctx context.Context) (*stricache.Success, error) {
	var first string
	c.mu.Lock()
	first, c.Strings.list = c.Strings.list[0], c.Strings.list[1:len(c.Strings.list)-1]
	// Sync
	for i, v := range c.Strings.items {
		if first == v {
			delete(c.Strings.items, i)
		}
	}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) ShiftInt(ctx context.Context) (*stricache.Success, error) {
	var first string
	c.mu.Lock()
	first, c.Ints.list = c.Ints.list[0], c.Ints.list[1:len(c.Ints.list)-1]
	// Sync
	for i, v := range c.Ints.items {
		if first == v {
			delete(c.Ints.items, i)
		}
	}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) ShiftFloat(ctx context.Context) (*stricache.Success, error) {
	var first string
	c.mu.Lock()
	first, c.Floats.list = c.Floats.list[0], c.Floats.list[1:len(c.Floats.list)-1]
	// Sync
	for i, v := range c.Floats.items {
		if first == v {
			delete(c.Floats.items, i)
		}
	}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) PopString(ctx context.Context) (*stricache.Success, error) {
	var first string
	c.mu.Lock()
	c.Strings.list, last = c.Strings.list[:len(c.Strings.list)-1], c.Strings.list[len(c.Strings.list)-1]
	// Sync
	for i, v := range c.Strings.items {
		if last == v {
			delete(c.Strings.items, i)
		}
	}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) PopInt(ctx context.Context) (*stricache.Success, error) {
	var first string
	c.mu.Lock()
	c.Ints.list, last = c.Ints.list[:len(c.Ints.list)-1], c.Ints.list[len(c.Ints.list)-1]
	// Sync
	for i, v := range c.Ints.items {
		if last == v {
			delete(c.Ints.items, i)
		}
	}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}

func (c *storage.Cache) PopFloat(ctx context.Context) (*stricache.Success, error) {
	var first string
	c.mu.Lock()
	c.Floats.list, last = c.Floats.list[:len(c.Floats.list)-1], c.Floats.list[len(c.Floats.list)-1]
	// Sync
	for i, v := range c.Floats.items {
		if last == v {
			delete(c.Floats.items, i)
		}
	}
	c.mu.Unlock()
	return &api.Success{
		Success: true,
	}, nil
}





	

	
