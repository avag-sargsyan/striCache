package api

import (
	"context"
	"strings"
	"time"

	api "github.com/avag-sargsyan/striCache"
)

func (c *cache) AddString(ctx context.Context, item *api.StringItem) (*api.StringItem, error) {
	var expiration int64
	duration, _ := time.ParseDuration(item.Expiration)
	// d should be "5m30s"
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}
	c.mu.Lock()
	c.items[item.Key] = Item{
		Object:     item.Value,
		Expiration: expiration,
	}
	c.mu.Unlock()
	return item, nil
}