package memstore

import (
	"sync"

	"github.com/albsen/chainstore"
	"github.com/albsen/chainstore/lrumgr"
	"golang.org/x/net/context"
)

var _ chainstore.Store = &memStore{}

type memStore struct {
	sync.RWMutex
	data map[string][]byte
}

// New creates and returns a memory based store.
func New(capacity int64) chainstore.Store {
	memStore := &memStore{
		data: make(map[string][]byte, 1000),
	}
	store := lrumgr.New(capacity, memStore)
	return store
}

func (s *memStore) Open() error {
	return nil
}

func (s *memStore) Close() error {
	return nil
}

func (s *memStore) Put(ctx context.Context, key string, val []byte) error {
	s.Lock()
	s.data[key] = val
	s.Unlock()
	return nil
}

func (s *memStore) Get(ctx context.Context, key string) ([]byte, error) {
	s.RLock()
	val := s.data[key]
	s.RUnlock()
	return val, nil
}

func (s *memStore) Del(ctx context.Context, key string) error {
	s.Lock()
	delete(s.data, key)
	s.Unlock()
	return nil
}
