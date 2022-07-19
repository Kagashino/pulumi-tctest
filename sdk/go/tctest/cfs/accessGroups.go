// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func AccessGroups(ctx *pulumi.Context, args *AccessGroupsArgs, opts ...pulumi.InvokeOption) (*AccessGroupsResult, error) {
	var rv AccessGroupsResult
	err := ctx.Invoke("tctest:Cfs/accessGroups:AccessGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking AccessGroups.
type AccessGroupsArgs struct {
	AccessGroupId    *string `pulumi:"accessGroupId"`
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by AccessGroups.
type AccessGroupsResult struct {
	AccessGroupId    *string                       `pulumi:"accessGroupId"`
	AccessGroupLists []AccessGroupsAccessGroupList `pulumi:"accessGroupLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func AccessGroupsOutput(ctx *pulumi.Context, args AccessGroupsOutputArgs, opts ...pulumi.InvokeOption) AccessGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AccessGroupsResult, error) {
			args := v.(AccessGroupsArgs)
			r, err := AccessGroups(ctx, &args, opts...)
			var s AccessGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AccessGroupsResultOutput)
}

// A collection of arguments for invoking AccessGroups.
type AccessGroupsOutputArgs struct {
	AccessGroupId    pulumi.StringPtrInput `pulumi:"accessGroupId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (AccessGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessGroupsArgs)(nil)).Elem()
}

// A collection of values returned by AccessGroups.
type AccessGroupsResultOutput struct{ *pulumi.OutputState }

func (AccessGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessGroupsResult)(nil)).Elem()
}

func (o AccessGroupsResultOutput) ToAccessGroupsResultOutput() AccessGroupsResultOutput {
	return o
}

func (o AccessGroupsResultOutput) ToAccessGroupsResultOutputWithContext(ctx context.Context) AccessGroupsResultOutput {
	return o
}

func (o AccessGroupsResultOutput) AccessGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessGroupsResult) *string { return v.AccessGroupId }).(pulumi.StringPtrOutput)
}

func (o AccessGroupsResultOutput) AccessGroupLists() AccessGroupsAccessGroupListArrayOutput {
	return o.ApplyT(func(v AccessGroupsResult) []AccessGroupsAccessGroupList { return v.AccessGroupLists }).(AccessGroupsAccessGroupListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AccessGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccessGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o AccessGroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessGroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AccessGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessGroupsResultOutput{})
}
