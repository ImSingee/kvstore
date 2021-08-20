package kvstore

type UnsafeGetter interface {
	UnsafeGet(key string) (AnyValue, error)
}

var _ UnsafeGetter = (*store)(nil)

func (s *store) UnsafeGet(key string) (AnyValue, error) {
	return s.readValue(key)
}
