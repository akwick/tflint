// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLambdaFunctionInvalidDescriptionRule checks the pattern is valid
type AwsLambdaFunctionInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsLambdaFunctionInvalidDescriptionRule returns new rule with default attributes
func NewAwsLambdaFunctionInvalidDescriptionRule() *AwsLambdaFunctionInvalidDescriptionRule {
	return &AwsLambdaFunctionInvalidDescriptionRule{
		resourceType:  "aws_lambda_function",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsLambdaFunctionInvalidDescriptionRule) Name() string {
	return "aws_lambda_function_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaFunctionInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaFunctionInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaFunctionInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaFunctionInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
