package kvstore

type Checker interface {
	CheckExist(key string) error
	CheckType(key string, typeName string) error
}

var _ Checker = (*store)(nil)

// CheckExist 检查 key 是否存在，如果不存在或有中间错误会返回 error
func (s *store) CheckExist(key string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.checkExist(key)
}
func (s *store) checkExist(key string) error {
	_, err := s.readValue(key)
	// 不检查 value 是否为 nil
	// 因为如果 key 不存在 readValue 一定会返回 error
	// 而 key 对应 None 值则不会返回 error 但 value 为 nil
	return err
}

// CheckType 返回 key 对应的类型是否合法
// 类型名称支持包括 TypeName 所返回的名称
// 或是 "number" 代表 "int" 或 "float"
func (s *store) CheckType(key string, typeName string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.checkType(key, typeName)
}
func (s *store) checkType(key string, typeName string) error {
	v, err := s.readValue(key)
	if err != nil {
		return err
	}

	expectName := TypeName(v)
	if expectName == typeName {
		return nil
	}
	if typeName == "any" || typeName == "" {
		return nil
	}

	switch expectName {
	case "int", "float":
		if typeName == "number" {
			return nil
		}
		return ErrKeyTypeNotMatch{
			Key:    key,
			Expect: typeName,
			Got:    expectName + " (number)",
		}
	case "list", "map":
		// TODO
	}

	return ErrKeyTypeNotMatch{
		Key:    key,
		Expect: typeName,
		Got:    expectName,
	}
}

// TypeNameCompare 检查两个类型的包含关系
// 0 代表相同，-1 代表前者范围更小，1 代表后者返回更小
// -2 代表检查不通过（二者无交集）
func TypeNameCompare(previous, current string) int {
	if previous == current {
		return 0
	}
	if previous == "any" { // previous > current
		return 1
	}
	if current == "any" { // current > previous
		return -1
	}

	if previous == "number" && (current == "int" || current == "float") { // previous > current
		return 1
	}
	if current == "number" && (previous == "int" || previous == "float") { // current > previous
		return -1
	}

	return -2 // 无关
}
