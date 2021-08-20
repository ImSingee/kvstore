package kvstore

type UnsafeLocker interface {
	UnsafeLock()
	UnsafeUnlock()
	UnsafeRLock()
	UnsafeRUnlock()
}

var _ UnsafeLocker = (*store)(nil)

func (s *store) UnsafeLock() {
	s.lock.Lock()
}
func (s *store) UnsafeUnlock() {
	s.lock.Unlock()
}
func (s *store) UnsafeRLock() {
	s.lock.RLock()
}
func (s *store) UnsafeRUnlock() {
	s.lock.RUnlock()
}
