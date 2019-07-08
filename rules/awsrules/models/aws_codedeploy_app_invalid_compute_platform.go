// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCodedeployAppInvalidComputePlatformRule checks the pattern is valid
type AwsCodedeployAppInvalidComputePlatformRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodedeployAppInvalidComputePlatformRule returns new rule with default attributes
func NewAwsCodedeployAppInvalidComputePlatformRule() *AwsCodedeployAppInvalidComputePlatformRule {
	return &AwsCodedeployAppInvalidComputePlatformRule{
		resourceType:  "aws_codedeploy_app",
		attributeName: "compute_platform",
		enum: []string{
			"Server",
			"Lambda",
			"ECS",
		},
	}
}

// Name returns the rule name
func (r *AwsCodedeployAppInvalidComputePlatformRule) Name() string {
	return "aws_codedeploy_app_invalid_compute_platform"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodedeployAppInvalidComputePlatformRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCodedeployAppInvalidComputePlatformRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCodedeployAppInvalidComputePlatformRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodedeployAppInvalidComputePlatformRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`compute_platform is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}