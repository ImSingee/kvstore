package kvstore

import (
	"strconv"
	"strings"
)

func indexByKeyParts(parent Any, originalKey string, keyParts []string, createMapKeyIfPossible bool) (Any, error) {
	prefixParts := keyParts[:0]

	for len(keyParts) > 0 {
		k := keyParts[0]
		keyParts = keyParts[1:]

		switch p := parent.(type) {
		case nil, int64, float64, string, bool:
			return nil, ErrKeyTypeNotMatch{
				Key:    originalKey,
				On:     strings.Join(prefixParts, "."),
				Expect: "map or list",
				Got:    TypeName(p),
			}
		case *Map:
			parent = p.Get(k).Unwrap() // (new parent <-> value from struct)
			if parent == nil && createMapKeyIfPossible {
				// parent 不存在 k 对应的值，创建
				p.Set(k, NewEmptyMapValue())
				parent = p.Get(k).Unwrap()
			}
			if parent == nil {
				return nil, ErrKeyNotExist{Key: originalKey, On: strings.Join(prefixParts, ".")}
			}
		case *List:
			index, err := strconv.Atoi(k)
			if err != nil || index < 0 {
				return nil, ErrKeyIndexNotValid{
					Key:   originalKey,
					On:    strings.Join(prefixParts, "."),
					Index: -1,
				}
			}
			if index >= len(p.Values) {
				return nil, ErrKeyIndexNotValid{
					Key:   originalKey,
					On:    strings.Join(prefixParts, "."),
					Index: index,
					Max:   len(p.Values) - 1,
				}
			}
			parent = p.Values[index].Unwrap()
		default:
			return nil, ImpossibleError()
		}

		prefixParts = prefixParts[:len(prefixParts)+1]
	}

	return parent, nil
}

// 返回 key 对应的值
func (s *store) readValue(key string) (AnyValue, error) {
	keyParts := strings.Split(key, ".")

	any, err := indexByKeyParts(s.Provider, key, keyParts, false)
	if err != nil {
		return nil, err
	}

	return AnyToValue(any), nil
}

// 返回 key 的上层所对应的值
// 返回的第一个参数为 key 的最后一个部分
// 返回的第二个参数为 key 的倒数第二个部分所对应的 Any 值，且其类型只可能为 *Map 或 *List
// 返回的第三个参数为可能出现的错误
func (s *store) valueForChange(key string, createMapKeyIfPossible bool) (string, Any, error) {
	keyParts := strings.Split(key, ".")
	last := len(keyParts) - 1
	lastKey := keyParts[last]

	any, err := indexByKeyParts(s.Provider, key, keyParts[:last], createMapKeyIfPossible)
	if err != nil {
		return lastKey, nil, err
	}

	switch any.(type) {
	case *Map, *List:
		return lastKey, any, nil
	default:
		return lastKey, nil, ErrKeyTypeNotMatch{
			Key:    key,
			On:     strings.Join(keyParts[:last], "."),
			Expect: "map or list",
			Got:    TypeName(any),
		}
	}
}
