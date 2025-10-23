package store_test

import (
	"testing"

	"github.com/sanchayj/redis-store/store"
)

func TestInsertValue(t *testing.T) {
	store.Insert("foo", "val")

	val, err := store.Fetch("foo")
	if err != nil {
		t.Error(err)
	}

	if val != "val" {
		t.Errorf("got %q, want val", val)
	}
}

// this is dependent on a shared store not having the val
func TestInsertNotFoundVal(t *testing.T) {
	_, err := store.Fetch("bar")
	if err == nil {
		t.Error(err)
	}
}
