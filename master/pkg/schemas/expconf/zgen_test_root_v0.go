// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (t TestRootV0) ValX() int {
	return t.RawValX
}

func (t *TestRootV0) SetValX(val int) {
	t.RawValX = val
}

func (t TestRootV0) SubObj() TestSubV0 {
	if t.RawSubObj == nil {
		panic("You must call WithDefaults on TestRootV0 before .SubObj")
	}
	return *t.RawSubObj
}

func (t *TestRootV0) SetSubObj(val TestSubV0) {
	t.RawSubObj = &val
}

func (t TestRootV0) SubUnion() *TestUnionV0 {
	return t.RawSubUnion
}

func (t *TestRootV0) SetSubUnion(val *TestUnionV0) {
	t.RawSubUnion = val
}

func (t TestRootV0) DefaultedArray() []string {
	return t.RawDefaultedArray
}

func (t *TestRootV0) SetDefaultedArray(val []string) {
	t.RawDefaultedArray = val
}

func (t TestRootV0) NodefaultArray() []string {
	return t.RawNodefaultArray
}

func (t *TestRootV0) SetNodefaultArray(val []string) {
	t.RawNodefaultArray = val
}

func (t TestRootV0) RuntimeDefaultable() TestRuntimeDefaultable {
	return t.RawRuntimeDefaultable
}

func (t *TestRootV0) SetRuntimeDefaultable(val TestRuntimeDefaultable) {
	t.RawRuntimeDefaultable = val
}

func (t TestRootV0) ParsedSchema() interface{} {
	return schemas.ParsedTestRootV0()
}

func (t TestRootV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/test-root.json")
}

func (t TestRootV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/test-root.json")
}
