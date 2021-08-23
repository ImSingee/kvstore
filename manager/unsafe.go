package manager

import "github.com/ImSingee/kvstore"

type UnsafeManager interface {
	UnsafeLock()
	UnsafeUnlock()

	UnsafeStore() kvstore.Store
}

func (m *manager) Unsafe() UnsafeManager {
	return m
}

var _ UnsafeManager = (*manager)(nil)

func (m *manager) UnsafeLock() {
	m.Store.(kvstore.UnsafeStore).UnsafeLock()
}

func (m *manager) lock() {
	m.Store.(kvstore.UnsafeStore).UnsafeLock()
}

func (m *manager) UnsafeUnlock() {
	m.Store.(kvstore.UnsafeStore).UnsafeUnlock()
}

func (m *manager) unlock() {
	m.Store.(kvstore.UnsafeStore).UnsafeUnlock()
}

func (m *manager) UnsafeStore() kvstore.Store {
	return m.Store
}
