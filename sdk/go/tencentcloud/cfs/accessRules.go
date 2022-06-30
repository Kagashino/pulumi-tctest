// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func AccessRules(ctx *pulumi.Context, args *AccessRulesArgs, opts ...pulumi.InvokeOption) (*AccessRulesResult, error) {
	var rv AccessRulesResult
	err := ctx.Invoke("tencentcloud:Cfs/accessRules:AccessRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking AccessRules.
type AccessRulesArgs struct {
	AccessGroupId    string  `pulumi:"accessGroupId"`
	AccessRuleId     *string `pulumi:"accessRuleId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by AccessRules.
type AccessRulesResult struct {
	AccessGroupId   string                      `pulumi:"accessGroupId"`
	AccessRuleId    *string                     `pulumi:"accessRuleId"`
	AccessRuleLists []AccessRulesAccessRuleList `pulumi:"accessRuleLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func AccessRulesOutput(ctx *pulumi.Context, args AccessRulesOutputArgs, opts ...pulumi.InvokeOption) AccessRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AccessRulesResult, error) {
			args := v.(AccessRulesArgs)
			r, err := AccessRules(ctx, &args, opts...)
			var s AccessRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AccessRulesResultOutput)
}

// A collection of arguments for invoking AccessRules.
type AccessRulesOutputArgs struct {
	AccessGroupId    pulumi.StringInput    `pulumi:"accessGroupId"`
	AccessRuleId     pulumi.StringPtrInput `pulumi:"accessRuleId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (AccessRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRulesArgs)(nil)).Elem()
}

// A collection of values returned by AccessRules.
type AccessRulesResultOutput struct{ *pulumi.OutputState }

func (AccessRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRulesResult)(nil)).Elem()
}

func (o AccessRulesResultOutput) ToAccessRulesResultOutput() AccessRulesResultOutput {
	return o
}

func (o AccessRulesResultOutput) ToAccessRulesResultOutputWithContext(ctx context.Context) AccessRulesResultOutput {
	return o
}

func (o AccessRulesResultOutput) AccessGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessRulesResult) string { return v.AccessGroupId }).(pulumi.StringOutput)
}

func (o AccessRulesResultOutput) AccessRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessRulesResult) *string { return v.AccessRuleId }).(pulumi.StringPtrOutput)
}

func (o AccessRulesResultOutput) AccessRuleLists() AccessRulesAccessRuleListArrayOutput {
	return o.ApplyT(func(v AccessRulesResult) []AccessRulesAccessRuleList { return v.AccessRuleLists }).(AccessRulesAccessRuleListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AccessRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccessRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o AccessRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessRulesResultOutput{})
}
