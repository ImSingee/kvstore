package kvstore

import (
	"strconv"
	"strings"
)

// FlatForHighLevel 返回 key 的所有上层 key
// 该函数不会检查返回值是否在 store 中为合法的 key，但这一约束未来可能会添加
// 如果打算手动 opt-out 这一约束请使用独立的 FlatForHighLevel 函数（而非当前方法）
func (s *store) FlatForHighLevel(key string) []string {
	// 当前不涉及到内部查询，因此不上锁
	return FlatForHighLevel(key)
}

// FlatForLowLevel 返回 key 的所有子 key（拍平展开）
// 如果 key 不存或不合法会返回空数组（或 nil）
// 返回结果包括所有层级的 key
func (s *store) FlatForLowLevel(key string) []string {
	v, err := s.Get(key) // 这一函数不上锁，锁由 Get 完成
	if err != nil {
		return nil
	}

	switch data := v.(type) {
	case nil, int64, float64, string, bool:
		return []string{key}
	case []interface{}:
		keys := make([]string, 1, 16)
		keys[0] = key
		return flatForLowLevelList(keys, key+".", data)
	case map[string]interface{}:
		keys := make([]string, 1, 16)
		keys[0] = key
		return flatForLowLevelMap(keys, key+".", data)
	default:
		return []string{key}
	}
}

func flatForLowLevelMap(keys []string, prefix string, data map[string]interface{}) []string {
	for len(data) != 0 {
		for k, v := range data {
			keys = append(keys, prefix+k)
			delete(data, k)

			switch vv := v.(type) {
			case nil, int64, float64, string, bool:
			case []interface{}:
				for i, e := range vv {
					data[k+"."+strconv.Itoa(i)] = e
				}
			case map[string]interface{}:
				for sk, e := range vv {
					data[k+"."+sk] = e
				}
			default:
			}
		}
	}

	return keys
}

func flatForLowLevelList(keys []string, prefix string, data []interface{}) []string {
	for i, v := range data {
		index := strconv.Itoa(i)
		keys = append(keys, prefix+index)

		switch vv := v.(type) {
		case nil, int64, float64, string, bool:
		case []interface{}:
			keys = flatForLowLevelList(keys, prefix+index+".", vv)
		case map[string]interface{}:
			keys = flatForLowLevelMap(keys, prefix+index+".", vv)
		default:
		}
	}

	return keys
}

func FlatForHighLevel(key string) []string {
	if key == "" {
		return nil
	}

	parts := make([]string, 0, 8)

	findAndDo(key, '.', func(i int) {
		parts = append(parts, key[:i])
	})

	parts = append(parts, key)

	return parts
}

func findAndDo(s string, c byte, f func(int)) {
	skip := 0
	sub := s

	for {
		next := strings.IndexByte(sub, c)
		if next == -1 {
			return
		}
		f(skip + next)

		skip += next + 1
		sub = s[skip:]
	}
}
