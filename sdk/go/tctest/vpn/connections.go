// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Connections(ctx *pulumi.Context, args *ConnectionsArgs, opts ...pulumi.InvokeOption) (*ConnectionsResult, error) {
	var rv ConnectionsResult
	err := ctx.Invoke("tctest:Vpn/connections:Connections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Connections.
type ConnectionsArgs struct {
	CustomerGatewayId *string                `pulumi:"customerGatewayId"`
	Id                *string                `pulumi:"id"`
	Name              *string                `pulumi:"name"`
	ResultOutputFile  *string                `pulumi:"resultOutputFile"`
	Tags              map[string]interface{} `pulumi:"tags"`
	VpcId             *string                `pulumi:"vpcId"`
	VpnGatewayId      *string                `pulumi:"vpnGatewayId"`
}

// A collection of values returned by Connections.
type ConnectionsResult struct {
	ConnectionLists   []ConnectionsConnectionList `pulumi:"connectionLists"`
	CustomerGatewayId *string                     `pulumi:"customerGatewayId"`
	Id                *string                     `pulumi:"id"`
	Name              *string                     `pulumi:"name"`
	ResultOutputFile  *string                     `pulumi:"resultOutputFile"`
	Tags              map[string]interface{}      `pulumi:"tags"`
	VpcId             *string                     `pulumi:"vpcId"`
	VpnGatewayId      *string                     `pulumi:"vpnGatewayId"`
}

func ConnectionsOutput(ctx *pulumi.Context, args ConnectionsOutputArgs, opts ...pulumi.InvokeOption) ConnectionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ConnectionsResult, error) {
			args := v.(ConnectionsArgs)
			r, err := Connections(ctx, &args, opts...)
			var s ConnectionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ConnectionsResultOutput)
}

// A collection of arguments for invoking Connections.
type ConnectionsOutputArgs struct {
	CustomerGatewayId pulumi.StringPtrInput `pulumi:"customerGatewayId"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile  pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Tags              pulumi.MapInput       `pulumi:"tags"`
	VpcId             pulumi.StringPtrInput `pulumi:"vpcId"`
	VpnGatewayId      pulumi.StringPtrInput `pulumi:"vpnGatewayId"`
}

func (ConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionsArgs)(nil)).Elem()
}

// A collection of values returned by Connections.
type ConnectionsResultOutput struct{ *pulumi.OutputState }

func (ConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionsResult)(nil)).Elem()
}

func (o ConnectionsResultOutput) ToConnectionsResultOutput() ConnectionsResultOutput {
	return o
}

func (o ConnectionsResultOutput) ToConnectionsResultOutputWithContext(ctx context.Context) ConnectionsResultOutput {
	return o
}

func (o ConnectionsResultOutput) ConnectionLists() ConnectionsConnectionListArrayOutput {
	return o.ApplyT(func(v ConnectionsResult) []ConnectionsConnectionList { return v.ConnectionLists }).(ConnectionsConnectionListArrayOutput)
}

func (o ConnectionsResultOutput) CustomerGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.CustomerGatewayId }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v ConnectionsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o ConnectionsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) VpnGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.VpnGatewayId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionsResultOutput{})
}
