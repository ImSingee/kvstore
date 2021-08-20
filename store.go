package kvstore

import (
	"sync"
)

type Store interface {
	Getter
	Setter
	Deleter
	Applier

	UnsafeUnderlying() *Map
	UnsafeUnderlyingPointer() **Map
}

type UnsafeStore interface {
	Store

	UnsafeGetter
	UnsafeApplier
}

type store struct {
	*Provider

	lock sync.RWMutex
}

var _ Store = (*store)(nil)
var _ UnsafeStore = (*store)(nil)

func (s *store) UnsafeUnderlying() *Map {
	return s.Provider
}
func (s *store) UnsafeUnderlyingPointer() **Map {
	return &s.Provider
}
