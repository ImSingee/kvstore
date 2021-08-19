package kvstore

import (
	"sync"
)

type Store interface {
	UnsafeUnderlying() *Map
	UnsafeUnderlyingPointer() **Map
}

type store struct {
	*Provider

	lock sync.RWMutex
}

var _ Store = (*store)(nil)

func (s *store) UnsafeUnderlying() *Map {
	return s.Provider
}
func (s *store) UnsafeUnderlyingPointer() **Map {
	return &s.Provider
}
