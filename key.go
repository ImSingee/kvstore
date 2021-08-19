package kvstore

import (
	"strconv"
	"strings"
)

func indexByKeyParts(parent Any, originalKey string, keyParts []string) (Any, error) {
	prefixParts := keyParts[:0]

	for len(keyParts) > 0 {
		k := keyParts[0]
		keyParts = keyParts[1:]
		prefixParts = prefixParts[:len(prefixParts)+1]

		switch p := parent.(type) {
		case nil, int64, float64, string, bool:
			if len(keyParts) != 0 { // not last part, disallow basic types
				return nil, ErrKeyTypeNotMatch{
					Key:    originalKey,
					On:     strings.Join(prefixParts, "."),
					Expect: "map or list",
					Got:    TypeName(p),
				}
			}

			parent = p
		case *Map:
			parent = p.Get(k).Unwrap() // (new parent <-> value from struct)
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
	}

	return parent, nil
}

// 返回 key 对应的值
func (s *store) readValue(key string) (AnyValue, error) {
	keyParts := strings.Split(key, ".")

	any, err := indexByKeyParts(s.Provider, key, keyParts)
	if err != nil {
		return nil, err
	}

	return AnyToValue(any), nil
}

// 返回 key 对应的值，但保证返回的值只可能是 Map 或 List
func (s *store) valueForChange(key string) (Any, error) {
	keyParts := strings.Split(key, ".")

	any, err := indexByKeyParts(s.Provider, key, keyParts)
	if err != nil {
		return nil, err
	}

	switch any.(type) {
	case *Map, *List:
		return any, nil
	default:
		return nil, ErrKeyTypeNotMatch{
			Key:    key,
			Expect: "map or list",
			Got:    TypeName(any),
		}
	}
}
