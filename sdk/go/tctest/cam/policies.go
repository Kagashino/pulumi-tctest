// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Policies(ctx *pulumi.Context, args *PoliciesArgs, opts ...pulumi.InvokeOption) (*PoliciesResult, error) {
	var rv PoliciesResult
	err := ctx.Invoke("tctest:Cam/policies:Policies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Policies.
type PoliciesArgs struct {
	CreateMode       *int    `pulumi:"createMode"`
	Description      *string `pulumi:"description"`
	Name             *string `pulumi:"name"`
	PolicyId         *string `pulumi:"policyId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Type             *int    `pulumi:"type"`
}

// A collection of values returned by Policies.
type PoliciesResult struct {
	CreateMode  *int    `pulumi:"createMode"`
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id               string               `pulumi:"id"`
	Name             *string              `pulumi:"name"`
	PolicyId         *string              `pulumi:"policyId"`
	PolicyLists      []PoliciesPolicyList `pulumi:"policyLists"`
	ResultOutputFile *string              `pulumi:"resultOutputFile"`
	Type             *int                 `pulumi:"type"`
}

func PoliciesOutput(ctx *pulumi.Context, args PoliciesOutputArgs, opts ...pulumi.InvokeOption) PoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (PoliciesResult, error) {
			args := v.(PoliciesArgs)
			r, err := Policies(ctx, &args, opts...)
			var s PoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(PoliciesResultOutput)
}

// A collection of arguments for invoking Policies.
type PoliciesOutputArgs struct {
	CreateMode       pulumi.IntPtrInput    `pulumi:"createMode"`
	Description      pulumi.StringPtrInput `pulumi:"description"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	PolicyId         pulumi.StringPtrInput `pulumi:"policyId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Type             pulumi.IntPtrInput    `pulumi:"type"`
}

func (PoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PoliciesArgs)(nil)).Elem()
}

// A collection of values returned by Policies.
type PoliciesResultOutput struct{ *pulumi.OutputState }

func (PoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoliciesResult)(nil)).Elem()
}

func (o PoliciesResultOutput) ToPoliciesResultOutput() PoliciesResultOutput {
	return o
}

func (o PoliciesResultOutput) ToPoliciesResultOutputWithContext(ctx context.Context) PoliciesResultOutput {
	return o
}

func (o PoliciesResultOutput) CreateMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PoliciesResult) *int { return v.CreateMode }).(pulumi.IntPtrOutput)
}

func (o PoliciesResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PoliciesResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o PoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o PoliciesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PoliciesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PoliciesResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PoliciesResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o PoliciesResultOutput) PolicyLists() PoliciesPolicyListArrayOutput {
	return o.ApplyT(func(v PoliciesResult) []PoliciesPolicyList { return v.PolicyLists }).(PoliciesPolicyListArrayOutput)
}

func (o PoliciesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PoliciesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o PoliciesResultOutput) Type() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PoliciesResult) *int { return v.Type }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PoliciesResultOutput{})
}
