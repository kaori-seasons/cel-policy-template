package parser

import (
	"github.com/google/cel-go/cel"

	"github.com/google/cel-policy-templates-go/policy/model"
	"github.com/google/cel-policy-templates-go/policy/parser/yml"
)

// ParseYaml decodes a YAML source to a model.ParsedValue.
//
// If errors are encountered during decode, they are returned via the Errors object.
func ParseYaml(src *model.Source) (*model.ParsedValue, *cel.Issues) {
	return yml.Parse(src)
}
