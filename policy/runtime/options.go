package runtime

import (
	"github.com/google/cel-policy-templates-go/policy/limits"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/interpreter/functions"
)

// TemplateOption is a functional optoin for configuring template evaluation.
type TemplateOption func(*Template) (*Template, error)

// DecisionAggregator registers an Aggregator for a given decision name.
func DecisionAggregator(decision string, agg Aggregator) TemplateOption {
	return func(t *Template) (*Template, error) {
		t.decAggMap[decision] = agg
		return t, nil
	}
}

// ExprOptions configues a set of options for use with constructing CEL programs within the
// template.
func ExprOptions(opts ...cel.ProgramOption) TemplateOption {
	return func(t *Template) (*Template, error) {
		t.exprOpts = append(t.exprOpts, opts...)
		return t, nil
	}
}

// Functions configures the template runtime with function implementations that correspond with
// the compilation environment specification.
func Functions(funcs ...*functions.Overload) TemplateOption {
	return func(t *Template) (*Template, error) {
		t.exprOpts = append(t.exprOpts, cel.Functions(funcs...))
		return t, nil
	}
}

// Limits configures limits which should be enforced during runtime evaluation.
func Limits(l *limits.Limits) TemplateOption {
	return func(t *Template) (*Template, error) {
		t.limits = l
		return t, nil
	}
}
