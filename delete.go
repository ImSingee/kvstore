package kvstore

import (
	"strings"
)

type Deleter interface {
	Delete(key string) error
}

var _ Deleter = (*store)(nil)

// Delete 用于删除一个 key
// 如果 key 不存在，返回 nil
// 如果 key 的中间索引不存在，返回 error
// 可以安全的忽略返回值，如果只需要确保这一 key 不可能被 Get 成功
//
// 数组元素：不支持删除一个数组元素的某一项的整体
// 例如 a.b 是一个数组，则删除 a.b.1 的行为是不合法的，但支持删除数组的子元素，例如支持删除 a.b.1.c
// 如需明确删除数组的某一项，可以根据实际需求选择 将相应项设置为 None / 整体替换数组元素 / 使用未来或许支持的数组操作
func (s *store) Delete(key string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.delete(key, true)
}
func (s *store) delete(key string, allowNotExist bool) error {
	last, parent, err := s.valueForChange(key, false)
	if err != nil {
		return err
	}

	switch p := parent.(type) {
	case *Map:
		_, ok := p.Fields[last]
		if ok {
			delete(p.Fields, last)
		} else if !allowNotExist {
			return ErrKeyNotExist{Key: key}
		}

		return nil
	case *List:
		return ErrKeyTypeNotMatch{
			Key:    key,
			On:     strings.TrimSuffix(key, "."+last),
			Expect: "map",
			Got:    "list",
		}
	default:
		return ImpossibleError()
	}
}
