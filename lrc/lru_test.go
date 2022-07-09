package lrc

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var total = 10
	cache := New(total)

	if cache.Cap() != total {
		t.Errorf("new lruCache error,wrong cache.cap:%d != total:%d", cache.Cap(), total)
	}
}

type TestStruct struct {
	Cap int
	Ops []func(l *LruCache) error
}

func TestLruCache_Put(t *testing.T) {
	d := TestStruct{
		Cap: 4,
		Ops: []func(cache *LruCache) error{
			func(cache *LruCache) error {
				cache.Put("1", 1)
				cache.Put("2", 2)
				cache.Put("3", 3)
				cache.Put("4", 4)
				cache.Put("1", 5)
				if cache.Total() != 4 {
					return fmt.Errorf("cache item total should be %d", 4)
				}
				cache.Get("2")
				ok, v := cache.Get("1")
				if !ok || v.Value.(int) != 5 {
					return fmt.Errorf("cache item should exists and value should be %d", 5)
				}
				return nil
			},
		},
	}

	check(d, t)
}

func check(data TestStruct, t *testing.T) {
	c := New(data.Cap)
	for _, op := range data.Ops {
		if err := op(c); err != nil {
			t.Error(err)
		}
	}
}
