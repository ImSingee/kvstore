package kvstore

func (x *Action) Unwrap() interface{} {
	switch action := x.Action.(type) {
	case *Action_Set:
		return action.Set
	case *Action_Delete:
		return action.Delete
	case *Action_Replace:
		return action.Replace
	default:
		return action
	}
}

func (x *Action) IsSet() bool {
	return x.GetSet() != nil
}

func (x *Action) IsDelete() bool {
	return x.GetDelete() != nil
}

func (x *Action) IsReplace() bool {
	return x.GetReplace() != nil
}

func NewSetAction(key string, value *Value) *Action {
	return &Action{
		Action: &Action_Set{Set: &Set{Key: key, Value: value}},
	}
}

func NewDeleteAction(key string) *Action {
	return &Action{
		Action: &Action_Delete{Delete: &Delete{Key: key}},
	}
}

func NewReplaceAction(new *Map) *Action {
	return &Action{
		Action: &Action_Replace{Replace: &Replace{New: new}},
	}
}
