package kvstore

import (
	"github.com/golang/protobuf/proto"
	"strconv"
)

// Clone 会克隆一个新的 Store 对象，其数据与原始相同但修改互不影响
func (s *store) Clone() Store {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &store{Provider: proto.Clone(s.Provider).(*Map)}
}

// Export 将返回底层数据结构，其会对底层数据进行完全拷贝
// 返回的 value 类型等同于 AnyValue，即  nil, int64, float64, string, bool, map[string]interface{}, []interface{}
func (s *store) Export() map[string]interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return proto.Clone(s.Provider).(*Map).AsMap()
}

// ExportFlat 将返回底层数据的拍平拷贝
// 返回的 key 为多个 `.` 的组合
// 返回的 value 类型只可能为基本类型（nil, int64, float64, string, bool）
func (s *store) ExportFlat() map[string]interface{} {
	// 该函数不上锁，唯一读在 Export 内上锁
	data := s.Export()

	for {
		toDelete := make([]string, 0, 64)

		for k, v := range data {
			switch vv := v.(type) {
			case nil, int64, float64, string, bool:
			case []interface{}:
				toDelete = append(toDelete, k)
				for i, e := range vv {
					data[k+"."+strconv.Itoa(i)] = e
				}
			case map[string]interface{}:
				toDelete = append(toDelete, k)
				for sk, e := range vv {
					data[k+"."+sk] = e
				}
			default:
			}
		}

		if len(toDelete) == 0 {
			break
		} else {
			for _, k := range toDelete {
				delete(data, k)
			}
		}
	}

	return data
}

//// ExportAllLevel 获得所有可以利用 Get 获得的 k-v 集合
//// 例如对于 a.b.c.d ，其会同时返回 a, a.b, a.b.c, a.b.c.d 的内容
//func (s *store) ExportAllLevel() map[string]interface{} {
//
//}

func Filter(dataset map[string]interface{}, shouldRemain func(k string, v interface{}) bool) map[string]interface{} {
	newSet := make(map[string]interface{}, len(dataset))
	for k, v := range dataset {
		if shouldRemain(k, v) {
			newSet[k] = v
		}
	}
	return newSet
}

// FilterSubSet
// 将 dataset 中所有不在 includes 或 includes 子元素中的元素筛选掉
func FilterSubSet(dataset map[string]interface{}, includes []string) map[string]interface{} {
	trie := NewTrie(includes)
	return Filter(dataset, func(k string, _ interface{}) bool {
		return trie.CheckExist(k)
	})
}
