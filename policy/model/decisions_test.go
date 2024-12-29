package model

import (
	"reflect"
	"testing"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

func TestBoolDecisionValue_And(t *testing.T) {
	tests := []struct {
		name   string
		value  types.Bool
		ands   []ref.Val
		result ref.Val
	}{
		{
			name:   "init_false_end_false",
			value:  types.False,
			ands:   []ref.Val{types.NewErr("err"), types.True},
			result: types.False,
		},
		{
			name:   "init_true_end_false",
			value:  types.True,
			ands:   []ref.Val{types.NewErr("err"), types.False},
			result: types.False,
		},
		{
			name:   "init_true_end_err",
			value:  types.True,
			ands:   []ref.Val{types.True, types.NewErr("err")},
			result: types.NewErr("err"),
		},
		{
			name:   "init_true_end_unk",
			value:  types.True,
			ands:   []ref.Val{types.True, types.NewUnknown(1, nil), types.NewErr("err"), types.NewUnknown(2, nil)},
			result: types.MergeUnknowns(types.NewUnknown(1, nil), types.NewUnknown(2, nil)),
		},
	}
	for _, tst := range tests {
		tc := tst
		t.Run(tc.name, func(tt *testing.T) {
			v := NewBoolDecisionValue(tc.name, tc.value)
			for _, av := range tc.ands {
				v = v.And(av)
			}
			v.Finalize(nil, nil)
			if !reflect.DeepEqual(v.Value(), tc.result) {
				tt.Errorf("decision AND failed. got %v, wanted %v", v.Value(), tc.result)
			}
		})
	}
}

func TestBoolDecisionValue_Or(t *testing.T) {
	tests := []struct {
		name   string
		value  types.Bool
		ors    []ref.Val
		result ref.Val
	}{
		{
			name:   "init_false_end_true",
			value:  types.False,
			ors:    []ref.Val{types.NewErr("err"), types.NewUnknown(1, nil), types.True},
			result: types.True,
		},
		{
			name:   "init_true_end_true",
			value:  types.True,
			ors:    []ref.Val{types.NewErr("err"), types.False},
			result: types.True,
		},
		{
			name:   "init_false_end_err",
			value:  types.False,
			ors:    []ref.Val{types.False, types.NewErr("err1"), types.NewErr("err2")},
			result: types.NewErr("err1"),
		},
		{
			name:   "init_false_end_unk",
			value:  types.False,
			ors:    []ref.Val{types.False, types.NewUnknown(1, nil), types.NewErr("err"), types.NewUnknown(2, nil)},
			result: types.MergeUnknowns(types.NewUnknown(1, nil), types.NewUnknown(2, nil)),
		},
	}
	for _, tst := range tests {
		tc := tst
		t.Run(tc.name, func(tt *testing.T) {
			v := NewBoolDecisionValue(tc.name, tc.value)
			for _, av := range tc.ors {
				v = v.Or(av)
			}
			// Test finalization
			v.Finalize(nil, nil)
			// Ensure that calling string on the value doesn't error.
			_ = v.String()
			// Compare the output result
			if !reflect.DeepEqual(v.Value(), tc.result) {
				tt.Errorf("decision OR failed. got %v, wanted %v", v.Value(), tc.result)
			}
		})
	}
}
