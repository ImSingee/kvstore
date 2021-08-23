package kvstore

import "strconv"

func JSONToActions(data []byte) ([]*Action, error) {
	store, err := LoadFromJSON(data)
	if err != nil {
		return nil, err
	}
	u := store.Unsafe().UnsafeUnderlying().Fields
	actions := make([]*Action, 0, len(u))

	for k, v := range u {
		actions = append(actions, NewSetAction(k, v))
	}

	return actions, nil
}

func (s *store) ApplyByJSON(data []byte) error {
	actions, err := JSONToActions(data)
	if err != nil {
		return err
	}
	return s.Apply(actions) // Apply 内部会上锁
}

// mapToFlat 将 data 拍平展开至只有基础类型，操作为直接修改 data
func mapToFlat(data *Map) *Map {
	for {
		toDelete := make([]string, 0, 64)

		for k, v := range data.Fields {
			switch vv := v.Unwrap().(type) {
			case nil, int64, float64, string, bool:
			case *List:
				toDelete = append(toDelete, k)
				for i, e := range vv.Values {
					data.Set(k+"."+strconv.Itoa(i), e)
				}
			case *Map:
				toDelete = append(toDelete, k)
				for sk, e := range vv.Fields {
					data.Set(k+"."+sk, e)
				}
			default:
			}
		}

		if len(toDelete) == 0 {
			break
		} else {
			for _, k := range toDelete {
				delete(data.Fields, k)
			}
		}
	}

	return data
}
