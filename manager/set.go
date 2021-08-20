package manager

import (
	"github.com/ImSingee/kvstore"
	"math"
)

func (m *manager) setValueByValue(key string, value *kvstore.Value) error {
	return m.writeLogAndDoE(kvstore.NewSetAction(key, value), func(s kvstore.UnsafeStore) error {
		return s.UnsafeSetValueByValue(key, value)
	})
}

func (m *manager) setValueByAnyValue(key string, newValue_ kvstore.AnyValue) error {
	newValue, err := kvstore.AnyValueToValue(newValue_)
	if err != nil {
		return err // 应当在外部提前保证好不可能出现
	}

	return m.setValueByValue(key, newValue)
}

func (m *manager) Set(key string, value *kvstore.Value) error {
	return m.setValueByValue(key, value)
}

func (m *manager) SetNull(key string) error {
	return m.setValueByAnyValue(key, nil)
}

func (m *manager) SetInt64(key string, value int64) error {
	return m.setValueByAnyValue(key, int64(value))
}

func (m *manager) SetUint64(key string, value uint64) error {
	if value > math.MaxInt64 {
		return m.setValueByAnyValue(key, float64(value))
	} else {
		return m.setValueByAnyValue(key, int64(value))
	}
}

func (m *manager) SetFloat64(key string, value float64) error {
	return m.setValueByAnyValue(key, float64(value))
}

func (m *manager) SetBool(key string, value bool) error {
	if value {
		return m.SetTrue(key)
	} else {
		return m.SetFalse(key)
	}
}

func (m *manager) SetString(key string, value string) error {
	return m.setValueByAnyValue(key, string(value))
}

func (m *manager) SetTrue(key string) error {
	return m.setValueByAnyValue(key, true)
}

func (m *manager) SetFalse(key string) error {
	return m.setValueByAnyValue(key, false)
}

func (m *manager) SetList(key string, value []interface{}) error {
	return m.setValueByAnyValue(key, []interface{}(value))
}

func (m *manager) SetMap(key string, value map[string]interface{}) error {
	return m.setValueByAnyValue(key, map[string]interface{}(value))
}
