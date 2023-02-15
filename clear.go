package kvstore

type Clearer interface {
	Clear()
}

type UnsafeClearer interface {
	UnsafeClear()
}

var _ Clearer = (*store)(nil)
var _ UnsafeClearer = (*store)(nil)

// Clear 用于清空现有所有数据
func (s *store) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.clear()
}
func (s *store) Reset() {
	s.Clear()
}
func (s *store) UnsafeClear() {
	s.clear()
}
func (s *store) clear() {
	s.Provider.Reset()
}
