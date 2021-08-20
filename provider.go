package kvstore

import (
	"fmt"
	"github.com/ImSingee/structpb"
)

import _ "unsafe"

type Provider = structpb.Struct

/*
//	*Value_NullValue
	//	*Value_IntValue
	//	*Value_FloatValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_ListValue
*/

// Any 标识着来源于 Value.Unwrap() 的可能返回值
// 其类型可能为 nil, int64, float64, string, bool, *Map, *List
type Any interface{}

// AnyValue 标识着来源于 Value.AsInterface() 的可能返回值
// 其类型可能为 nil, int64, float64, string, bool, map[string]interface{}, []interface{}
type AnyValue interface{}

type Value = structpb.Value

type Map = structpb.Struct
type List = structpb.ListValue

var (
	NULL  = structpb.NewNullValue()
	TRUE  = structpb.NewBoolValue(true)
	FALSE = structpb.NewBoolValue(false)
)

//go:linkname NewValue github.com/ImSingee/structpb.NewValue
func NewValue(v interface{}) (*Value, error)

//go:linkname NewIntValue github.com/ImSingee/structpb.NewIntValue
func NewIntValue(v int64) *Value

//go:linkname NewEmptyMap github.com/ImSingee/structpb.NewEmptyStruct
func NewEmptyMap() *Map

func NewEmptyMapValue() *Value {
	return structpb.NewStructValue(NewEmptyMap())
}

//go:linkname NewMap github.com/ImSingee/structpb.NewStruct
func NewMap(v map[string]interface{}) (*Map, error)

//go:linkname NewList github.com/ImSingee/structpb.NewList
func NewList(v []interface{}) (*List, error)

//go:linkname NewEmptyList github.com/ImSingee/structpb.NewEmptyList
func NewEmptyList() (*List, error)

// TypeName for Any or AnyValue
// 返回值可能为 "null", "int", "float", "string", "bool", "map", "list"
// 对于非 Any 和 AnyValue 的类型均返回 "unknown"
func TypeName(v interface{}) string {
	switch v.(type) {
	// nil, int64, float64, string, bool, *Struct, *List
	case nil:
		return "null"
	case int64:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case *Map, map[string]interface{}:
		return "map"
	case *List, []interface{}:
		return "list"
	default:
		return "unknown"
	}
}

// AnyToValue convert Any to AnyValue, 转换为一一对应
// 需要注意的是，与 Value 的 AsInterface 不同，这一函数并不会特别处理 float infinity 和 NaN
func AnyToValue(from Any) AnyValue {
	switch v := from.(type) {
	case nil, int64, float64, string, bool:
		return v
	case *Map:
		return v.AsMap()
	case *List:
		return v.AsSlice()
	default:
		return v
	}
}

// ValueToAny convert AnyValue to Any, 转换为一一对应
// 需要注意的是，AnyValue 为字符串形式的 float64 (NaN, Infinity) 则会出现不一致
func ValueToAny(from AnyValue) (Any, error) {
	switch v := from.(type) {
	case nil, int64, float64, string, bool:
		return v, nil
	case map[string]interface{}:
		return NewMap(v)
	case []interface{}:
		return NewList(v)
	default:
		return nil, fmt.Errorf("invalid type %T", from)
	}
}

// AnyValueToValue convert AnyValue to *Value, 转换为一一对应
// 相比于 NewValue 而言这一函数更加严格（同时效率也更高）
func AnyValueToValue(from AnyValue) (*Value, error) {
	switch v := from.(type) {
	case nil:
		return NULL, nil
	case bool:
		if v {
			return TRUE, nil
		} else {
			return FALSE, nil
		}
	case string:
		return structpb.NewStringValue(v), nil
	case int64:
		return structpb.NewIntValue(v), nil
	case float64:
		return structpb.NewFloatValue(v), nil
	case map[string]interface{}:
		s, err := structpb.NewStruct(v)
		if err != nil {
			return nil, err
		}
		return structpb.NewStructValue(s), nil
	case []interface{}:
		l, err := structpb.NewList(v)
		if err != nil {
			return nil, err
		}
		return structpb.NewListValue(l), nil
	default:
		return nil, fmt.Errorf("invalid type %T", from)
	}

}
