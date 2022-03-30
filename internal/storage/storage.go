package storage

import (
	"sync"
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
	items map[string]string
	list  []string
}

type intCache struct {
	items map[string]int64
	list  []int64
}

type floatCache struct {
	items map[string]float64
	list  []float64
}

func NewCacheService() *Cache {
	cstr := stringCache{
		map[string]string{},
		[]string{},
	}
	cint := intCache{
		map[string]int64{},
		[]int64{},
	}
	cflt := floatCache{
		map[string]float64{},
		[]float64{},
	}
	C := &Cache{
		Strings: &cstr,
		Ints:    &cint,
		Floats:  &cflt,
	}
	return C
}
