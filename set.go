package kvstore

import (
	"math"
	"strconv"
)

type Setter interface {
	SetNull(key string) error
	SetInt64(key string, value int64) error
	SetUint64(key string, value uint64) error
	SetFloat64(key string, value float64) error
	SetBool(key string, value bool) error
	SetString(key string, value string) error
	SetTrue(key string) error
	SetFalse(key string) error
	SetList(key string, value []interface{}) error
	SetMap(key string, value map[string]interface{}) error

	SetAnyList(key string, value interface{}) error
	SetStringList(key string, value []string) error
	SetAnyMap(key string, value interface{}) error
}

var _ Store = (*store)(nil)

func (s *store) SetNull(key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, nil)
}

func (s *store) SetInt64(key string, value int64) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, int64(value))
}

// SetUint64 可能将结果使用 int64 或 float64 存储
func (s *store) SetUint64(key string, value uint64) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if value > math.MaxInt64 {
		return s.setValueByAnyValue(key, float64(value))
	} else {
		return s.setValueByAnyValue(key, int64(value))
	}
}

func (s *store) SetFloat64(key string, value float64) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, float64(value))
}

func (s *store) SetBool(key string, value bool) error {
	if value {
		return s.SetTrue(key)
	} else {
		return s.SetFalse(key)
	}
}

func (s *store) SetString(key string, value string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, string(value))
}

func (s *store) SetTrue(key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, true)
}

func (s *store) SetFalse(key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.setValueByAnyValue(key, false)
}

func (s *store) setValueByAnyValue(key string, newValue_ AnyValue) error {
	newValue, err := AnyValueToValue(newValue_)
	if err != nil {
		return err // 应当在外部提前保证好不可能出现
	}

	return s.setValueByValue(key, newValue)
}

func (s *store) setValueByValue(key string, newValue *Value) error {
	last, parent, err := s.valueForChange(key, true)
	if err != nil {
		return err
	}

	switch p := parent.(type) {
	case *Map:
		p.Fields[last] = newValue
		return nil
	case *List:
		if last == "+" {
			p.Values = append(p.Values, newValue)
			return nil
		} else {
			index, err := strconv.Atoi(last)
			if err != nil || index < 0 {
				return ErrKeyIndexNotValid{
					Key:   key,
					Index: -1,
				}
			}
			if index >= len(p.Values) {
				return ErrKeyIndexNotValid{
					Key:   key,
					Index: index,
					Max:   len(p.Values) - 1,
				}
			}
			p.Values[index] = newValue
			return nil
		}
	default:
		return ImpossibleError()
	}
}
