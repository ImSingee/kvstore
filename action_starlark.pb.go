// Code generated by protoc-gen-starlark-go. DO NOT EDIT.

package kvstore

import (
	fmt "fmt"
	_ "github.com/ImSingee/structpb"
	starlark "go.starlark.net/starlark"
)

type ActionStarlark struct {
	Action  starlark.Value
	Set     *SetStarlark
	Delete  *DeleteStarlark
	Replace *ReplaceStarlark
}

var _ starlark.Value = (*ActionStarlark)(nil)
var _ starlark.HasAttrs = (*ActionStarlark)(nil)

func (x *ActionStarlark) String() string {
	if x == nil {
		return "Action{}"
	}
	return fmt.Sprintf("Action{"+
		"action= %s"+
		"set= %s"+
		", delete= %s"+
		", replace= %s"+
		"}",
		x.Action,
		x.Set,
		x.Delete,
		x.Replace,
	)
}
func (x *ActionStarlark) Type() string          { return "singee.kvstore.Action" }
func (x *ActionStarlark) Freeze()               {}
func (x *ActionStarlark) Truth() starlark.Bool  { return true }
func (x *ActionStarlark) Hash() (uint32, error) { return 0, fmt.Errorf("un-hashable") }
func (x *ActionStarlark) AttrNames() []string {
	return []string{
		"action",
		"set",
		"delete",
		"replace",
	}
}
func (x *ActionStarlark) Attr(name string) (starlark.Value, error) {
	if x == nil {
		return nil, nil
	}
	switch name {
	case "action":
		return x.Action, nil
	case "set":
		return x.Set, nil
	case "delete":
		return x.Delete, nil
	case "replace":
		return x.Replace, nil
	default:
		return nil, nil
	}
}
func (x *Action) ToStarlark() *ActionStarlark {
	if x == nil {
		return nil
	}

	var oneof_Action starlark.Value
	switch x.Action.(type) {
	case *Action_Set:
		oneof_Action = x.GetSet().ToStarlark()
	case *Action_Delete:
		oneof_Action = x.GetDelete().ToStarlark()
	case *Action_Replace:
		oneof_Action = x.GetReplace().ToStarlark()
	default:
		oneof_Action = starlark.None
	}

	return &ActionStarlark{
		Action:  oneof_Action,
		Set:     x.GetSet().ToStarlark(),
		Delete:  x.GetDelete().ToStarlark(),
		Replace: x.GetReplace().ToStarlark(),
	}
}

type SetStarlark struct {
	Key   starlark.String
	Value starlark.Value
}

var _ starlark.Value = (*SetStarlark)(nil)
var _ starlark.HasAttrs = (*SetStarlark)(nil)

func (x *SetStarlark) String() string {
	if x == nil {
		return "Set{}"
	}
	return fmt.Sprintf("Set{"+
		"key= %s"+
		", value= %s"+
		"}",
		x.Key,
		x.Value,
	)
}
func (x *SetStarlark) Type() string          { return "singee.kvstore.Set" }
func (x *SetStarlark) Freeze()               {}
func (x *SetStarlark) Truth() starlark.Bool  { return true }
func (x *SetStarlark) Hash() (uint32, error) { return 0, fmt.Errorf("un-hashable") }
func (x *SetStarlark) AttrNames() []string {
	return []string{
		"key",
		"value",
	}
}
func (x *SetStarlark) Attr(name string) (starlark.Value, error) {
	if x == nil {
		return nil, nil
	}
	switch name {
	case "key":
		return x.Key, nil
	case "value":
		return x.Value, nil
	default:
		return nil, nil
	}
}
func (x *Set) ToStarlark() *SetStarlark {
	if x == nil {
		return nil
	}

	return &SetStarlark{
		Key:   starlark.String(x.GetKey()),
		Value: x.GetValue().ToStarlark(),
	}
}

type DeleteStarlark struct {
	Key starlark.String
}

var _ starlark.Value = (*DeleteStarlark)(nil)
var _ starlark.HasAttrs = (*DeleteStarlark)(nil)

func (x *DeleteStarlark) String() string {
	if x == nil {
		return "Delete{}"
	}
	return fmt.Sprintf("Delete{"+
		"key= %s"+
		"}",
		x.Key,
	)
}
func (x *DeleteStarlark) Type() string          { return "singee.kvstore.Delete" }
func (x *DeleteStarlark) Freeze()               {}
func (x *DeleteStarlark) Truth() starlark.Bool  { return true }
func (x *DeleteStarlark) Hash() (uint32, error) { return 0, fmt.Errorf("un-hashable") }
func (x *DeleteStarlark) AttrNames() []string {
	return []string{
		"key",
	}
}
func (x *DeleteStarlark) Attr(name string) (starlark.Value, error) {
	if x == nil {
		return nil, nil
	}
	switch name {
	case "key":
		return x.Key, nil
	default:
		return nil, nil
	}
}
func (x *Delete) ToStarlark() *DeleteStarlark {
	if x == nil {
		return nil
	}

	return &DeleteStarlark{
		Key: starlark.String(x.GetKey()),
	}
}

type ReplaceStarlark struct {
	New *starlark.Dict
}

var _ starlark.Value = (*ReplaceStarlark)(nil)
var _ starlark.HasAttrs = (*ReplaceStarlark)(nil)

func (x *ReplaceStarlark) String() string {
	if x == nil {
		return "Replace{}"
	}
	return fmt.Sprintf("Replace{"+
		"new= %s"+
		"}",
		x.New,
	)
}
func (x *ReplaceStarlark) Type() string          { return "singee.kvstore.Replace" }
func (x *ReplaceStarlark) Freeze()               {}
func (x *ReplaceStarlark) Truth() starlark.Bool  { return true }
func (x *ReplaceStarlark) Hash() (uint32, error) { return 0, fmt.Errorf("un-hashable") }
func (x *ReplaceStarlark) AttrNames() []string {
	return []string{
		"new",
	}
}
func (x *ReplaceStarlark) Attr(name string) (starlark.Value, error) {
	if x == nil {
		return nil, nil
	}
	switch name {
	case "new":
		return x.New, nil
	default:
		return nil, nil
	}
}
func (x *Replace) ToStarlark() *ReplaceStarlark {
	if x == nil {
		return nil
	}

	return &ReplaceStarlark{
		New: x.GetNew().ToStarlark(),
	}
}