// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VipEipAttachmentsHaVipEipAttachmentList struct {
	AddressIp string `pulumi:"addressIp"`
	HavipId   string `pulumi:"havipId"`
}

// VipEipAttachmentsHaVipEipAttachmentListInput is an input type that accepts VipEipAttachmentsHaVipEipAttachmentListArgs and VipEipAttachmentsHaVipEipAttachmentListOutput values.
// You can construct a concrete instance of `VipEipAttachmentsHaVipEipAttachmentListInput` via:
//
//          VipEipAttachmentsHaVipEipAttachmentListArgs{...}
type VipEipAttachmentsHaVipEipAttachmentListInput interface {
	pulumi.Input

	ToVipEipAttachmentsHaVipEipAttachmentListOutput() VipEipAttachmentsHaVipEipAttachmentListOutput
	ToVipEipAttachmentsHaVipEipAttachmentListOutputWithContext(context.Context) VipEipAttachmentsHaVipEipAttachmentListOutput
}

type VipEipAttachmentsHaVipEipAttachmentListArgs struct {
	AddressIp pulumi.StringInput `pulumi:"addressIp"`
	HavipId   pulumi.StringInput `pulumi:"havipId"`
}

func (VipEipAttachmentsHaVipEipAttachmentListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VipEipAttachmentsHaVipEipAttachmentList)(nil)).Elem()
}

func (i VipEipAttachmentsHaVipEipAttachmentListArgs) ToVipEipAttachmentsHaVipEipAttachmentListOutput() VipEipAttachmentsHaVipEipAttachmentListOutput {
	return i.ToVipEipAttachmentsHaVipEipAttachmentListOutputWithContext(context.Background())
}

func (i VipEipAttachmentsHaVipEipAttachmentListArgs) ToVipEipAttachmentsHaVipEipAttachmentListOutputWithContext(ctx context.Context) VipEipAttachmentsHaVipEipAttachmentListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipEipAttachmentsHaVipEipAttachmentListOutput)
}

// VipEipAttachmentsHaVipEipAttachmentListArrayInput is an input type that accepts VipEipAttachmentsHaVipEipAttachmentListArray and VipEipAttachmentsHaVipEipAttachmentListArrayOutput values.
// You can construct a concrete instance of `VipEipAttachmentsHaVipEipAttachmentListArrayInput` via:
//
//          VipEipAttachmentsHaVipEipAttachmentListArray{ VipEipAttachmentsHaVipEipAttachmentListArgs{...} }
type VipEipAttachmentsHaVipEipAttachmentListArrayInput interface {
	pulumi.Input

	ToVipEipAttachmentsHaVipEipAttachmentListArrayOutput() VipEipAttachmentsHaVipEipAttachmentListArrayOutput
	ToVipEipAttachmentsHaVipEipAttachmentListArrayOutputWithContext(context.Context) VipEipAttachmentsHaVipEipAttachmentListArrayOutput
}

type VipEipAttachmentsHaVipEipAttachmentListArray []VipEipAttachmentsHaVipEipAttachmentListInput

func (VipEipAttachmentsHaVipEipAttachmentListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VipEipAttachmentsHaVipEipAttachmentList)(nil)).Elem()
}

func (i VipEipAttachmentsHaVipEipAttachmentListArray) ToVipEipAttachmentsHaVipEipAttachmentListArrayOutput() VipEipAttachmentsHaVipEipAttachmentListArrayOutput {
	return i.ToVipEipAttachmentsHaVipEipAttachmentListArrayOutputWithContext(context.Background())
}

func (i VipEipAttachmentsHaVipEipAttachmentListArray) ToVipEipAttachmentsHaVipEipAttachmentListArrayOutputWithContext(ctx context.Context) VipEipAttachmentsHaVipEipAttachmentListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipEipAttachmentsHaVipEipAttachmentListArrayOutput)
}

type VipEipAttachmentsHaVipEipAttachmentListOutput struct{ *pulumi.OutputState }

func (VipEipAttachmentsHaVipEipAttachmentListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VipEipAttachmentsHaVipEipAttachmentList)(nil)).Elem()
}

func (o VipEipAttachmentsHaVipEipAttachmentListOutput) ToVipEipAttachmentsHaVipEipAttachmentListOutput() VipEipAttachmentsHaVipEipAttachmentListOutput {
	return o
}

func (o VipEipAttachmentsHaVipEipAttachmentListOutput) ToVipEipAttachmentsHaVipEipAttachmentListOutputWithContext(ctx context.Context) VipEipAttachmentsHaVipEipAttachmentListOutput {
	return o
}

func (o VipEipAttachmentsHaVipEipAttachmentListOutput) AddressIp() pulumi.StringOutput {
	return o.ApplyT(func(v VipEipAttachmentsHaVipEipAttachmentList) string { return v.AddressIp }).(pulumi.StringOutput)
}

func (o VipEipAttachmentsHaVipEipAttachmentListOutput) HavipId() pulumi.StringOutput {
	return o.ApplyT(func(v VipEipAttachmentsHaVipEipAttachmentList) string { return v.HavipId }).(pulumi.StringOutput)
}

type VipEipAttachmentsHaVipEipAttachmentListArrayOutput struct{ *pulumi.OutputState }

func (VipEipAttachmentsHaVipEipAttachmentListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VipEipAttachmentsHaVipEipAttachmentList)(nil)).Elem()
}

func (o VipEipAttachmentsHaVipEipAttachmentListArrayOutput) ToVipEipAttachmentsHaVipEipAttachmentListArrayOutput() VipEipAttachmentsHaVipEipAttachmentListArrayOutput {
	return o
}

func (o VipEipAttachmentsHaVipEipAttachmentListArrayOutput) ToVipEipAttachmentsHaVipEipAttachmentListArrayOutputWithContext(ctx context.Context) VipEipAttachmentsHaVipEipAttachmentListArrayOutput {
	return o
}

func (o VipEipAttachmentsHaVipEipAttachmentListArrayOutput) Index(i pulumi.IntInput) VipEipAttachmentsHaVipEipAttachmentListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VipEipAttachmentsHaVipEipAttachmentList {
		return vs[0].([]VipEipAttachmentsHaVipEipAttachmentList)[vs[1].(int)]
	}).(VipEipAttachmentsHaVipEipAttachmentListOutput)
}

type VipsHaVipList struct {
	AddressIp          string `pulumi:"addressIp"`
	CreateTime         string `pulumi:"createTime"`
	Id                 string `pulumi:"id"`
	InstanceId         string `pulumi:"instanceId"`
	Name               string `pulumi:"name"`
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	State              string `pulumi:"state"`
	SubnetId           string `pulumi:"subnetId"`
	Vip                string `pulumi:"vip"`
	VpcId              string `pulumi:"vpcId"`
}

// VipsHaVipListInput is an input type that accepts VipsHaVipListArgs and VipsHaVipListOutput values.
// You can construct a concrete instance of `VipsHaVipListInput` via:
//
//          VipsHaVipListArgs{...}
type VipsHaVipListInput interface {
	pulumi.Input

	ToVipsHaVipListOutput() VipsHaVipListOutput
	ToVipsHaVipListOutputWithContext(context.Context) VipsHaVipListOutput
}

type VipsHaVipListArgs struct {
	AddressIp          pulumi.StringInput `pulumi:"addressIp"`
	CreateTime         pulumi.StringInput `pulumi:"createTime"`
	Id                 pulumi.StringInput `pulumi:"id"`
	InstanceId         pulumi.StringInput `pulumi:"instanceId"`
	Name               pulumi.StringInput `pulumi:"name"`
	NetworkInterfaceId pulumi.StringInput `pulumi:"networkInterfaceId"`
	State              pulumi.StringInput `pulumi:"state"`
	SubnetId           pulumi.StringInput `pulumi:"subnetId"`
	Vip                pulumi.StringInput `pulumi:"vip"`
	VpcId              pulumi.StringInput `pulumi:"vpcId"`
}

func (VipsHaVipListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VipsHaVipList)(nil)).Elem()
}

func (i VipsHaVipListArgs) ToVipsHaVipListOutput() VipsHaVipListOutput {
	return i.ToVipsHaVipListOutputWithContext(context.Background())
}

func (i VipsHaVipListArgs) ToVipsHaVipListOutputWithContext(ctx context.Context) VipsHaVipListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipsHaVipListOutput)
}

// VipsHaVipListArrayInput is an input type that accepts VipsHaVipListArray and VipsHaVipListArrayOutput values.
// You can construct a concrete instance of `VipsHaVipListArrayInput` via:
//
//          VipsHaVipListArray{ VipsHaVipListArgs{...} }
type VipsHaVipListArrayInput interface {
	pulumi.Input

	ToVipsHaVipListArrayOutput() VipsHaVipListArrayOutput
	ToVipsHaVipListArrayOutputWithContext(context.Context) VipsHaVipListArrayOutput
}

type VipsHaVipListArray []VipsHaVipListInput

func (VipsHaVipListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VipsHaVipList)(nil)).Elem()
}

func (i VipsHaVipListArray) ToVipsHaVipListArrayOutput() VipsHaVipListArrayOutput {
	return i.ToVipsHaVipListArrayOutputWithContext(context.Background())
}

func (i VipsHaVipListArray) ToVipsHaVipListArrayOutputWithContext(ctx context.Context) VipsHaVipListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VipsHaVipListArrayOutput)
}

type VipsHaVipListOutput struct{ *pulumi.OutputState }

func (VipsHaVipListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VipsHaVipList)(nil)).Elem()
}

func (o VipsHaVipListOutput) ToVipsHaVipListOutput() VipsHaVipListOutput {
	return o
}

func (o VipsHaVipListOutput) ToVipsHaVipListOutputWithContext(ctx context.Context) VipsHaVipListOutput {
	return o
}

func (o VipsHaVipListOutput) AddressIp() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.AddressIp }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.Id }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.Name }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.State }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) Vip() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.Vip }).(pulumi.StringOutput)
}

func (o VipsHaVipListOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v VipsHaVipList) string { return v.VpcId }).(pulumi.StringOutput)
}

type VipsHaVipListArrayOutput struct{ *pulumi.OutputState }

func (VipsHaVipListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VipsHaVipList)(nil)).Elem()
}

func (o VipsHaVipListArrayOutput) ToVipsHaVipListArrayOutput() VipsHaVipListArrayOutput {
	return o
}

func (o VipsHaVipListArrayOutput) ToVipsHaVipListArrayOutputWithContext(ctx context.Context) VipsHaVipListArrayOutput {
	return o
}

func (o VipsHaVipListArrayOutput) Index(i pulumi.IntInput) VipsHaVipListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VipsHaVipList {
		return vs[0].([]VipsHaVipList)[vs[1].(int)]
	}).(VipsHaVipListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VipEipAttachmentsHaVipEipAttachmentListInput)(nil)).Elem(), VipEipAttachmentsHaVipEipAttachmentListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VipEipAttachmentsHaVipEipAttachmentListArrayInput)(nil)).Elem(), VipEipAttachmentsHaVipEipAttachmentListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VipsHaVipListInput)(nil)).Elem(), VipsHaVipListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VipsHaVipListArrayInput)(nil)).Elem(), VipsHaVipListArray{})
	pulumi.RegisterOutputType(VipEipAttachmentsHaVipEipAttachmentListOutput{})
	pulumi.RegisterOutputType(VipEipAttachmentsHaVipEipAttachmentListArrayOutput{})
	pulumi.RegisterOutputType(VipsHaVipListOutput{})
	pulumi.RegisterOutputType(VipsHaVipListArrayOutput{})
}
