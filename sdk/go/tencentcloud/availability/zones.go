// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package availability

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Zones(ctx *pulumi.Context, args *ZonesArgs, opts ...pulumi.InvokeOption) (*ZonesResult, error) {
	var rv ZonesResult
	err := ctx.Invoke("tencentcloud:Availability/zones:Zones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Zones.
type ZonesArgs struct {
	IncludeUnavailable *bool   `pulumi:"includeUnavailable"`
	Name               *string `pulumi:"name"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by Zones.
type ZonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string      `pulumi:"id"`
	IncludeUnavailable *bool       `pulumi:"includeUnavailable"`
	Name               *string     `pulumi:"name"`
	ResultOutputFile   *string     `pulumi:"resultOutputFile"`
	Zones              []ZonesZone `pulumi:"zones"`
}

func ZonesOutput(ctx *pulumi.Context, args ZonesOutputArgs, opts ...pulumi.InvokeOption) ZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ZonesResult, error) {
			args := v.(ZonesArgs)
			r, err := Zones(ctx, &args, opts...)
			var s ZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ZonesResultOutput)
}

// A collection of arguments for invoking Zones.
type ZonesOutputArgs struct {
	IncludeUnavailable pulumi.BoolPtrInput   `pulumi:"includeUnavailable"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile   pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (ZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonesArgs)(nil)).Elem()
}

// A collection of values returned by Zones.
type ZonesResultOutput struct{ *pulumi.OutputState }

func (ZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonesResult)(nil)).Elem()
}

func (o ZonesResultOutput) ToZonesResultOutput() ZonesResultOutput {
	return o
}

func (o ZonesResultOutput) ToZonesResultOutputWithContext(ctx context.Context) ZonesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o ZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ZonesResultOutput) IncludeUnavailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ZonesResult) *bool { return v.IncludeUnavailable }).(pulumi.BoolPtrOutput)
}

func (o ZonesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ZonesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ZonesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ZonesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o ZonesResultOutput) Zones() ZonesZoneArrayOutput {
	return o.ApplyT(func(v ZonesResult) []ZonesZone { return v.Zones }).(ZonesZoneArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ZonesResultOutput{})
}