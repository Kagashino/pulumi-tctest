// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GatewayRoutes(ctx *pulumi.Context, args *GatewayRoutesArgs, opts ...pulumi.InvokeOption) (*GatewayRoutesResult, error) {
	var rv GatewayRoutesResult
	err := ctx.Invoke("tctest:Vpn/gatewayRoutes:GatewayRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GatewayRoutes.
type GatewayRoutesArgs struct {
	DestinationCidr  *string `pulumi:"destinationCidr"`
	InstanceId       *string `pulumi:"instanceId"`
	InstanceType     *string `pulumi:"instanceType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	VpnGatewayId     string  `pulumi:"vpnGatewayId"`
}

// A collection of values returned by GatewayRoutes.
type GatewayRoutesResult struct {
	DestinationCidr *string `pulumi:"destinationCidr"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string                             `pulumi:"id"`
	InstanceId           *string                            `pulumi:"instanceId"`
	InstanceType         *string                            `pulumi:"instanceType"`
	ResultOutputFile     *string                            `pulumi:"resultOutputFile"`
	VpnGatewayId         string                             `pulumi:"vpnGatewayId"`
	VpnGatewayRouteLists []GatewayRoutesVpnGatewayRouteList `pulumi:"vpnGatewayRouteLists"`
}

func GatewayRoutesOutput(ctx *pulumi.Context, args GatewayRoutesOutputArgs, opts ...pulumi.InvokeOption) GatewayRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GatewayRoutesResult, error) {
			args := v.(GatewayRoutesArgs)
			r, err := GatewayRoutes(ctx, &args, opts...)
			var s GatewayRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GatewayRoutesResultOutput)
}

// A collection of arguments for invoking GatewayRoutes.
type GatewayRoutesOutputArgs struct {
	DestinationCidr  pulumi.StringPtrInput `pulumi:"destinationCidr"`
	InstanceId       pulumi.StringPtrInput `pulumi:"instanceId"`
	InstanceType     pulumi.StringPtrInput `pulumi:"instanceType"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	VpnGatewayId     pulumi.StringInput    `pulumi:"vpnGatewayId"`
}

func (GatewayRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRoutesArgs)(nil)).Elem()
}

// A collection of values returned by GatewayRoutes.
type GatewayRoutesResultOutput struct{ *pulumi.OutputState }

func (GatewayRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRoutesResult)(nil)).Elem()
}

func (o GatewayRoutesResultOutput) ToGatewayRoutesResultOutput() GatewayRoutesResultOutput {
	return o
}

func (o GatewayRoutesResultOutput) ToGatewayRoutesResultOutputWithContext(ctx context.Context) GatewayRoutesResultOutput {
	return o
}

func (o GatewayRoutesResultOutput) DestinationCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayRoutesResult) *string { return v.DestinationCidr }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GatewayRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GatewayRoutesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayRoutesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GatewayRoutesResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayRoutesResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o GatewayRoutesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayRoutesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GatewayRoutesResultOutput) VpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRoutesResult) string { return v.VpnGatewayId }).(pulumi.StringOutput)
}

func (o GatewayRoutesResultOutput) VpnGatewayRouteLists() GatewayRoutesVpnGatewayRouteListArrayOutput {
	return o.ApplyT(func(v GatewayRoutesResult) []GatewayRoutesVpnGatewayRouteList { return v.VpnGatewayRouteLists }).(GatewayRoutesVpnGatewayRouteListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayRoutesResultOutput{})
}
