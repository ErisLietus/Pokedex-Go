package pokecache_test

import (
	"testing"
	"time"

	pokecache "github.com/ErisLietus/pokedex_go/internal"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key     string
		wantVal []byte
		wantOK  bool
	}{
		{key: "a", wantVal: []byte{1}, wantOK: true},
	}
	interval := 10 * time.Microsecond
	r := pokecache.NewCache(interval)
	r.Add("key", []byte{1})

	for _, c := range cases {
		r.Get(c.key)
	}
}
