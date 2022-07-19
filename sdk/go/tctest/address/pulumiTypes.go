// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package address

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TemplateGroupsGroupList struct {
	Id          string   `pulumi:"id"`
	Name        string   `pulumi:"name"`
	TemplateIds []string `pulumi:"templateIds"`
}

// TemplateGroupsGroupListInput is an input type that accepts TemplateGroupsGroupListArgs and TemplateGroupsGroupListOutput values.
// You can construct a concrete instance of `TemplateGroupsGroupListInput` via:
//
//          TemplateGroupsGroupListArgs{...}
type TemplateGroupsGroupListInput interface {
	pulumi.Input

	ToTemplateGroupsGroupListOutput() TemplateGroupsGroupListOutput
	ToTemplateGroupsGroupListOutputWithContext(context.Context) TemplateGroupsGroupListOutput
}

type TemplateGroupsGroupListArgs struct {
	Id          pulumi.StringInput      `pulumi:"id"`
	Name        pulumi.StringInput      `pulumi:"name"`
	TemplateIds pulumi.StringArrayInput `pulumi:"templateIds"`
}

func (TemplateGroupsGroupListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateGroupsGroupList)(nil)).Elem()
}

func (i TemplateGroupsGroupListArgs) ToTemplateGroupsGroupListOutput() TemplateGroupsGroupListOutput {
	return i.ToTemplateGroupsGroupListOutputWithContext(context.Background())
}

func (i TemplateGroupsGroupListArgs) ToTemplateGroupsGroupListOutputWithContext(ctx context.Context) TemplateGroupsGroupListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateGroupsGroupListOutput)
}

// TemplateGroupsGroupListArrayInput is an input type that accepts TemplateGroupsGroupListArray and TemplateGroupsGroupListArrayOutput values.
// You can construct a concrete instance of `TemplateGroupsGroupListArrayInput` via:
//
//          TemplateGroupsGroupListArray{ TemplateGroupsGroupListArgs{...} }
type TemplateGroupsGroupListArrayInput interface {
	pulumi.Input

	ToTemplateGroupsGroupListArrayOutput() TemplateGroupsGroupListArrayOutput
	ToTemplateGroupsGroupListArrayOutputWithContext(context.Context) TemplateGroupsGroupListArrayOutput
}

type TemplateGroupsGroupListArray []TemplateGroupsGroupListInput

func (TemplateGroupsGroupListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateGroupsGroupList)(nil)).Elem()
}

func (i TemplateGroupsGroupListArray) ToTemplateGroupsGroupListArrayOutput() TemplateGroupsGroupListArrayOutput {
	return i.ToTemplateGroupsGroupListArrayOutputWithContext(context.Background())
}

func (i TemplateGroupsGroupListArray) ToTemplateGroupsGroupListArrayOutputWithContext(ctx context.Context) TemplateGroupsGroupListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateGroupsGroupListArrayOutput)
}

type TemplateGroupsGroupListOutput struct{ *pulumi.OutputState }

func (TemplateGroupsGroupListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateGroupsGroupList)(nil)).Elem()
}

func (o TemplateGroupsGroupListOutput) ToTemplateGroupsGroupListOutput() TemplateGroupsGroupListOutput {
	return o
}

func (o TemplateGroupsGroupListOutput) ToTemplateGroupsGroupListOutputWithContext(ctx context.Context) TemplateGroupsGroupListOutput {
	return o
}

func (o TemplateGroupsGroupListOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateGroupsGroupList) string { return v.Id }).(pulumi.StringOutput)
}

func (o TemplateGroupsGroupListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateGroupsGroupList) string { return v.Name }).(pulumi.StringOutput)
}

func (o TemplateGroupsGroupListOutput) TemplateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TemplateGroupsGroupList) []string { return v.TemplateIds }).(pulumi.StringArrayOutput)
}

type TemplateGroupsGroupListArrayOutput struct{ *pulumi.OutputState }

func (TemplateGroupsGroupListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplateGroupsGroupList)(nil)).Elem()
}

func (o TemplateGroupsGroupListArrayOutput) ToTemplateGroupsGroupListArrayOutput() TemplateGroupsGroupListArrayOutput {
	return o
}

func (o TemplateGroupsGroupListArrayOutput) ToTemplateGroupsGroupListArrayOutputWithContext(ctx context.Context) TemplateGroupsGroupListArrayOutput {
	return o
}

func (o TemplateGroupsGroupListArrayOutput) Index(i pulumi.IntInput) TemplateGroupsGroupListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TemplateGroupsGroupList {
		return vs[0].([]TemplateGroupsGroupList)[vs[1].(int)]
	}).(TemplateGroupsGroupListOutput)
}

type TemplatesTemplateList struct {
	Addresses []string `pulumi:"addresses"`
	Id        string   `pulumi:"id"`
	Name      string   `pulumi:"name"`
}

// TemplatesTemplateListInput is an input type that accepts TemplatesTemplateListArgs and TemplatesTemplateListOutput values.
// You can construct a concrete instance of `TemplatesTemplateListInput` via:
//
//          TemplatesTemplateListArgs{...}
type TemplatesTemplateListInput interface {
	pulumi.Input

	ToTemplatesTemplateListOutput() TemplatesTemplateListOutput
	ToTemplatesTemplateListOutputWithContext(context.Context) TemplatesTemplateListOutput
}

type TemplatesTemplateListArgs struct {
	Addresses pulumi.StringArrayInput `pulumi:"addresses"`
	Id        pulumi.StringInput      `pulumi:"id"`
	Name      pulumi.StringInput      `pulumi:"name"`
}

func (TemplatesTemplateListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplatesTemplateList)(nil)).Elem()
}

func (i TemplatesTemplateListArgs) ToTemplatesTemplateListOutput() TemplatesTemplateListOutput {
	return i.ToTemplatesTemplateListOutputWithContext(context.Background())
}

func (i TemplatesTemplateListArgs) ToTemplatesTemplateListOutputWithContext(ctx context.Context) TemplatesTemplateListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatesTemplateListOutput)
}

// TemplatesTemplateListArrayInput is an input type that accepts TemplatesTemplateListArray and TemplatesTemplateListArrayOutput values.
// You can construct a concrete instance of `TemplatesTemplateListArrayInput` via:
//
//          TemplatesTemplateListArray{ TemplatesTemplateListArgs{...} }
type TemplatesTemplateListArrayInput interface {
	pulumi.Input

	ToTemplatesTemplateListArrayOutput() TemplatesTemplateListArrayOutput
	ToTemplatesTemplateListArrayOutputWithContext(context.Context) TemplatesTemplateListArrayOutput
}

type TemplatesTemplateListArray []TemplatesTemplateListInput

func (TemplatesTemplateListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplatesTemplateList)(nil)).Elem()
}

func (i TemplatesTemplateListArray) ToTemplatesTemplateListArrayOutput() TemplatesTemplateListArrayOutput {
	return i.ToTemplatesTemplateListArrayOutputWithContext(context.Background())
}

func (i TemplatesTemplateListArray) ToTemplatesTemplateListArrayOutputWithContext(ctx context.Context) TemplatesTemplateListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatesTemplateListArrayOutput)
}

type TemplatesTemplateListOutput struct{ *pulumi.OutputState }

func (TemplatesTemplateListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplatesTemplateList)(nil)).Elem()
}

func (o TemplatesTemplateListOutput) ToTemplatesTemplateListOutput() TemplatesTemplateListOutput {
	return o
}

func (o TemplatesTemplateListOutput) ToTemplatesTemplateListOutputWithContext(ctx context.Context) TemplatesTemplateListOutput {
	return o
}

func (o TemplatesTemplateListOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TemplatesTemplateList) []string { return v.Addresses }).(pulumi.StringArrayOutput)
}

func (o TemplatesTemplateListOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TemplatesTemplateList) string { return v.Id }).(pulumi.StringOutput)
}

func (o TemplatesTemplateListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TemplatesTemplateList) string { return v.Name }).(pulumi.StringOutput)
}

type TemplatesTemplateListArrayOutput struct{ *pulumi.OutputState }

func (TemplatesTemplateListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TemplatesTemplateList)(nil)).Elem()
}

func (o TemplatesTemplateListArrayOutput) ToTemplatesTemplateListArrayOutput() TemplatesTemplateListArrayOutput {
	return o
}

func (o TemplatesTemplateListArrayOutput) ToTemplatesTemplateListArrayOutputWithContext(ctx context.Context) TemplatesTemplateListArrayOutput {
	return o
}

func (o TemplatesTemplateListArrayOutput) Index(i pulumi.IntInput) TemplatesTemplateListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TemplatesTemplateList {
		return vs[0].([]TemplatesTemplateList)[vs[1].(int)]
	}).(TemplatesTemplateListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateGroupsGroupListInput)(nil)).Elem(), TemplateGroupsGroupListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateGroupsGroupListArrayInput)(nil)).Elem(), TemplateGroupsGroupListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplatesTemplateListInput)(nil)).Elem(), TemplatesTemplateListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplatesTemplateListArrayInput)(nil)).Elem(), TemplatesTemplateListArray{})
	pulumi.RegisterOutputType(TemplateGroupsGroupListOutput{})
	pulumi.RegisterOutputType(TemplateGroupsGroupListArrayOutput{})
	pulumi.RegisterOutputType(TemplatesTemplateListOutput{})
	pulumi.RegisterOutputType(TemplatesTemplateListArrayOutput{})
}
