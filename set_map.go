package kvstore

import "reflect"

func (s *store) SetMap(key string, value map[string]interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, map[string]interface{}(value))
}

// SetAnyMap
// Panic: if value is not a map[string]T, it will panic
func (s *store) SetAnyMap(key string, value interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	rv := reflect.ValueOf(value)

	m := make(map[string]interface{}, rv.Len())
	for _, key := range rv.MapKeys() {
		m[key.String()] = rv.MapIndex(key).Interface()
	}

	return s.setValueByAnyValue(key, map[string]interface{}(m))
}
