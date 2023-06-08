package lrumap_test

import (
	"testing"
	kvstore "travisroad/gotracker/kvstore"
	"travisroad/gotracker/kvstore/lrumap"

	"github.com/stretchr/testify/assert"
)

func TestLruMap(t *testing.T) {
	var kv kvstore.KVStore = lrumap.New(1000)

	kv.Set("foo", "bar")
	kv.Set("foofoo", "barbar")

	val, ok := kv.Get("foo")

	assert.Equal(t, val, "bar")
	assert.Equal(t, ok, true)

	_, ok = kv.Get("bar")

	assert.Equal(t, ok, false)

	kv.Delete("foo")
	val, ok = kv.Get("foo")
	assert.Equal(t, val, nil)
	assert.Equal(t, ok, false)
}
