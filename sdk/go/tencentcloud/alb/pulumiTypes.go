// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerAttachmentBackend struct {
	InstanceId string `pulumi:"instanceId"`
	Port       int    `pulumi:"port"`
	Weight     *int   `pulumi:"weight"`
}

// ServerAttachmentBackendInput is an input type that accepts ServerAttachmentBackendArgs and ServerAttachmentBackendOutput values.
// You can construct a concrete instance of `ServerAttachmentBackendInput` via:
//
//          ServerAttachmentBackendArgs{...}
type ServerAttachmentBackendInput interface {
	pulumi.Input

	ToServerAttachmentBackendOutput() ServerAttachmentBackendOutput
	ToServerAttachmentBackendOutputWithContext(context.Context) ServerAttachmentBackendOutput
}

type ServerAttachmentBackendArgs struct {
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	Port       pulumi.IntInput    `pulumi:"port"`
	Weight     pulumi.IntPtrInput `pulumi:"weight"`
}

func (ServerAttachmentBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAttachmentBackend)(nil)).Elem()
}

func (i ServerAttachmentBackendArgs) ToServerAttachmentBackendOutput() ServerAttachmentBackendOutput {
	return i.ToServerAttachmentBackendOutputWithContext(context.Background())
}

func (i ServerAttachmentBackendArgs) ToServerAttachmentBackendOutputWithContext(ctx context.Context) ServerAttachmentBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAttachmentBackendOutput)
}

// ServerAttachmentBackendArrayInput is an input type that accepts ServerAttachmentBackendArray and ServerAttachmentBackendArrayOutput values.
// You can construct a concrete instance of `ServerAttachmentBackendArrayInput` via:
//
//          ServerAttachmentBackendArray{ ServerAttachmentBackendArgs{...} }
type ServerAttachmentBackendArrayInput interface {
	pulumi.Input

	ToServerAttachmentBackendArrayOutput() ServerAttachmentBackendArrayOutput
	ToServerAttachmentBackendArrayOutputWithContext(context.Context) ServerAttachmentBackendArrayOutput
}

type ServerAttachmentBackendArray []ServerAttachmentBackendInput

func (ServerAttachmentBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerAttachmentBackend)(nil)).Elem()
}

func (i ServerAttachmentBackendArray) ToServerAttachmentBackendArrayOutput() ServerAttachmentBackendArrayOutput {
	return i.ToServerAttachmentBackendArrayOutputWithContext(context.Background())
}

func (i ServerAttachmentBackendArray) ToServerAttachmentBackendArrayOutputWithContext(ctx context.Context) ServerAttachmentBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAttachmentBackendArrayOutput)
}

type ServerAttachmentBackendOutput struct{ *pulumi.OutputState }

func (ServerAttachmentBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAttachmentBackend)(nil)).Elem()
}

func (o ServerAttachmentBackendOutput) ToServerAttachmentBackendOutput() ServerAttachmentBackendOutput {
	return o
}

func (o ServerAttachmentBackendOutput) ToServerAttachmentBackendOutputWithContext(ctx context.Context) ServerAttachmentBackendOutput {
	return o
}

func (o ServerAttachmentBackendOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerAttachmentBackend) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o ServerAttachmentBackendOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ServerAttachmentBackend) int { return v.Port }).(pulumi.IntOutput)
}

func (o ServerAttachmentBackendOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServerAttachmentBackend) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type ServerAttachmentBackendArrayOutput struct{ *pulumi.OutputState }

func (ServerAttachmentBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerAttachmentBackend)(nil)).Elem()
}

func (o ServerAttachmentBackendArrayOutput) ToServerAttachmentBackendArrayOutput() ServerAttachmentBackendArrayOutput {
	return o
}

func (o ServerAttachmentBackendArrayOutput) ToServerAttachmentBackendArrayOutputWithContext(ctx context.Context) ServerAttachmentBackendArrayOutput {
	return o
}

func (o ServerAttachmentBackendArrayOutput) Index(i pulumi.IntInput) ServerAttachmentBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerAttachmentBackend {
		return vs[0].([]ServerAttachmentBackend)[vs[1].(int)]
	}).(ServerAttachmentBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerAttachmentBackendInput)(nil)).Elem(), ServerAttachmentBackendArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerAttachmentBackendArrayInput)(nil)).Elem(), ServerAttachmentBackendArray{})
	pulumi.RegisterOutputType(ServerAttachmentBackendOutput{})
	pulumi.RegisterOutputType(ServerAttachmentBackendArrayOutput{})
}
