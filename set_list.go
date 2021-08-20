package kvstore

import "reflect"

func (s *store) SetList(key string, value []interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, []interface{}(value))
}

func (s *store) SetStringList(key string, value []string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	list := make([]interface{}, len(value))
	for i, e := range value {
		list[i] = e
	}

	return s.setValueByAnyValue(key, []interface{}(list))
}

// SetAnyList
// Panic: if value is not a []T, it will panic
func (s *store) SetAnyList(key string, value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	rv := reflect.ValueOf(value)

	list := make([]interface{}, rv.Len())
	for i := range list {
		list[i] = rv.Index(i).Interface()
	}

	return s.setValueByAnyValue(key, []interface{}(list))
}
