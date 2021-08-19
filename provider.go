package kvstore

import "github.com/ImSingee/structpb"

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
// 其类型可能为 nil, int64, float64, string, bool, map[string]interface(), []interface{}
type AnyValue interface{}

type Value = structpb.Value

type Map = structpb.Struct
type List = structpb.ListValue

var (
	NULL  = structpb.NewNullValue()
	TRUE  = structpb.NewBoolValue(true)
	FALSE = structpb.NewBoolValue(false)
)

//go:linkname NewIntValue github.com/ImSingee/structpb.NewIntValue
func NewIntValue(v int64) *Value

//go:linkname NewEmptyMap github.com/ImSingee/structpb.NewEmptyStruct
func NewEmptyMap() *Map

//go:linkname NewMap github.com/ImSingee/structpb.NewStruct
func NewMap(v map[string]interface{}) (*Map, error)

// TypeName for Any or AnyValue
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
