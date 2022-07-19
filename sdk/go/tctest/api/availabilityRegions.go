// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package api

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func AvailabilityRegions(ctx *pulumi.Context, args *AvailabilityRegionsArgs, opts ...pulumi.InvokeOption) (*AvailabilityRegionsResult, error) {
	var rv AvailabilityRegionsResult
	err := ctx.Invoke("tctest:Api/availabilityRegions:AvailabilityRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking AvailabilityRegions.
type AvailabilityRegionsArgs struct {
	IncludeUnavailable *bool   `pulumi:"includeUnavailable"`
	Name               *string `pulumi:"name"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by AvailabilityRegions.
type AvailabilityRegionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string                      `pulumi:"id"`
	IncludeUnavailable *bool                       `pulumi:"includeUnavailable"`
	Name               *string                     `pulumi:"name"`
	Regions            []AvailabilityRegionsRegion `pulumi:"regions"`
	ResultOutputFile   *string                     `pulumi:"resultOutputFile"`
}

func AvailabilityRegionsOutput(ctx *pulumi.Context, args AvailabilityRegionsOutputArgs, opts ...pulumi.InvokeOption) AvailabilityRegionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AvailabilityRegionsResult, error) {
			args := v.(AvailabilityRegionsArgs)
			r, err := AvailabilityRegions(ctx, &args, opts...)
			var s AvailabilityRegionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AvailabilityRegionsResultOutput)
}

// A collection of arguments for invoking AvailabilityRegions.
type AvailabilityRegionsOutputArgs struct {
	IncludeUnavailable pulumi.BoolPtrInput   `pulumi:"includeUnavailable"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile   pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (AvailabilityRegionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityRegionsArgs)(nil)).Elem()
}

// A collection of values returned by AvailabilityRegions.
type AvailabilityRegionsResultOutput struct{ *pulumi.OutputState }

func (AvailabilityRegionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityRegionsResult)(nil)).Elem()
}

func (o AvailabilityRegionsResultOutput) ToAvailabilityRegionsResultOutput() AvailabilityRegionsResultOutput {
	return o
}

func (o AvailabilityRegionsResultOutput) ToAvailabilityRegionsResultOutputWithContext(ctx context.Context) AvailabilityRegionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o AvailabilityRegionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilityRegionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o AvailabilityRegionsResultOutput) IncludeUnavailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AvailabilityRegionsResult) *bool { return v.IncludeUnavailable }).(pulumi.BoolPtrOutput)
}

func (o AvailabilityRegionsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityRegionsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AvailabilityRegionsResultOutput) Regions() AvailabilityRegionsRegionArrayOutput {
	return o.ApplyT(func(v AvailabilityRegionsResult) []AvailabilityRegionsRegion { return v.Regions }).(AvailabilityRegionsRegionArrayOutput)
}

func (o AvailabilityRegionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityRegionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilityRegionsResultOutput{})
}