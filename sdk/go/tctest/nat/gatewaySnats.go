// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GatewaySnats(ctx *pulumi.Context, args *GatewaySnatsArgs, opts ...pulumi.InvokeOption) (*GatewaySnatsResult, error) {
	var rv GatewaySnatsResult
	err := ctx.Invoke("tctest:Nat/gatewaySnats:GatewaySnats", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GatewaySnats.
type GatewaySnatsArgs struct {
	Description      *string  `pulumi:"description"`
	InstanceId       *string  `pulumi:"instanceId"`
	NatGatewayId     string   `pulumi:"natGatewayId"`
	PublicIpAddrs    []string `pulumi:"publicIpAddrs"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	SubnetId         *string  `pulumi:"subnetId"`
}

// A collection of values returned by GatewaySnats.
type GatewaySnatsResult struct {
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                 `pulumi:"id"`
	InstanceId       *string                `pulumi:"instanceId"`
	NatGatewayId     string                 `pulumi:"natGatewayId"`
	PublicIpAddrs    []string               `pulumi:"publicIpAddrs"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
	SnatLists        []GatewaySnatsSnatList `pulumi:"snatLists"`
	SubnetId         *string                `pulumi:"subnetId"`
}

func GatewaySnatsOutput(ctx *pulumi.Context, args GatewaySnatsOutputArgs, opts ...pulumi.InvokeOption) GatewaySnatsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GatewaySnatsResult, error) {
			args := v.(GatewaySnatsArgs)
			r, err := GatewaySnats(ctx, &args, opts...)
			var s GatewaySnatsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GatewaySnatsResultOutput)
}

// A collection of arguments for invoking GatewaySnats.
type GatewaySnatsOutputArgs struct {
	Description      pulumi.StringPtrInput   `pulumi:"description"`
	InstanceId       pulumi.StringPtrInput   `pulumi:"instanceId"`
	NatGatewayId     pulumi.StringInput      `pulumi:"natGatewayId"`
	PublicIpAddrs    pulumi.StringArrayInput `pulumi:"publicIpAddrs"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
	SubnetId         pulumi.StringPtrInput   `pulumi:"subnetId"`
}

func (GatewaySnatsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewaySnatsArgs)(nil)).Elem()
}

// A collection of values returned by GatewaySnats.
type GatewaySnatsResultOutput struct{ *pulumi.OutputState }

func (GatewaySnatsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewaySnatsResult)(nil)).Elem()
}

func (o GatewaySnatsResultOutput) ToGatewaySnatsResultOutput() GatewaySnatsResultOutput {
	return o
}

func (o GatewaySnatsResultOutput) ToGatewaySnatsResultOutputWithContext(ctx context.Context) GatewaySnatsResultOutput {
	return o
}

func (o GatewaySnatsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaySnatsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GatewaySnatsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GatewaySnatsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GatewaySnatsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaySnatsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GatewaySnatsResultOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GatewaySnatsResult) string { return v.NatGatewayId }).(pulumi.StringOutput)
}

func (o GatewaySnatsResultOutput) PublicIpAddrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewaySnatsResult) []string { return v.PublicIpAddrs }).(pulumi.StringArrayOutput)
}

func (o GatewaySnatsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaySnatsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GatewaySnatsResultOutput) SnatLists() GatewaySnatsSnatListArrayOutput {
	return o.ApplyT(func(v GatewaySnatsResult) []GatewaySnatsSnatList { return v.SnatLists }).(GatewaySnatsSnatListArrayOutput)
}

func (o GatewaySnatsResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaySnatsResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewaySnatsResultOutput{})
}
