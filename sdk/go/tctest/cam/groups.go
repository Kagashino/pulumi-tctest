// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Groups(ctx *pulumi.Context, args *GroupsArgs, opts ...pulumi.InvokeOption) (*GroupsResult, error) {
	var rv GroupsResult
	err := ctx.Invoke("tctest:Cam/groups:Groups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Groups.
type GroupsArgs struct {
	GroupId          *string `pulumi:"groupId"`
	Name             *string `pulumi:"name"`
	Remark           *string `pulumi:"remark"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by Groups.
type GroupsResult struct {
	GroupId    *string           `pulumi:"groupId"`
	GroupLists []GroupsGroupList `pulumi:"groupLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Name             *string `pulumi:"name"`
	Remark           *string `pulumi:"remark"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GroupsOutput(ctx *pulumi.Context, args GroupsOutputArgs, opts ...pulumi.InvokeOption) GroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GroupsResult, error) {
			args := v.(GroupsArgs)
			r, err := Groups(ctx, &args, opts...)
			var s GroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GroupsResultOutput)
}

// A collection of arguments for invoking Groups.
type GroupsOutputArgs struct {
	GroupId          pulumi.StringPtrInput `pulumi:"groupId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	Remark           pulumi.StringPtrInput `pulumi:"remark"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupsArgs)(nil)).Elem()
}

// A collection of values returned by Groups.
type GroupsResultOutput struct{ *pulumi.OutputState }

func (GroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupsResult)(nil)).Elem()
}

func (o GroupsResultOutput) ToGroupsResultOutput() GroupsResultOutput {
	return o
}

func (o GroupsResultOutput) ToGroupsResultOutputWithContext(ctx context.Context) GroupsResultOutput {
	return o
}

func (o GroupsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o GroupsResultOutput) GroupLists() GroupsGroupListArrayOutput {
	return o.ApplyT(func(v GroupsResult) []GroupsGroupList { return v.GroupLists }).(GroupsGroupListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GroupsResultOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupsResult) *string { return v.Remark }).(pulumi.StringPtrOutput)
}

func (o GroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupsResultOutput{})
}
