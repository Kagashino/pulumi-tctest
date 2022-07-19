// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ccn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func BandwidthLimits(ctx *pulumi.Context, args *BandwidthLimitsArgs, opts ...pulumi.InvokeOption) (*BandwidthLimitsResult, error) {
	var rv BandwidthLimitsResult
	err := ctx.Invoke("tctest:Ccn/bandwidthLimits:BandwidthLimits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking BandwidthLimits.
type BandwidthLimitsArgs struct {
	CcnId            string  `pulumi:"ccnId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by BandwidthLimits.
type BandwidthLimitsResult struct {
	CcnId string `pulumi:"ccnId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                 `pulumi:"id"`
	Limits           []BandwidthLimitsLimit `pulumi:"limits"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
}

func BandwidthLimitsOutput(ctx *pulumi.Context, args BandwidthLimitsOutputArgs, opts ...pulumi.InvokeOption) BandwidthLimitsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BandwidthLimitsResult, error) {
			args := v.(BandwidthLimitsArgs)
			r, err := BandwidthLimits(ctx, &args, opts...)
			var s BandwidthLimitsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(BandwidthLimitsResultOutput)
}

// A collection of arguments for invoking BandwidthLimits.
type BandwidthLimitsOutputArgs struct {
	CcnId            pulumi.StringInput    `pulumi:"ccnId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (BandwidthLimitsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthLimitsArgs)(nil)).Elem()
}

// A collection of values returned by BandwidthLimits.
type BandwidthLimitsResultOutput struct{ *pulumi.OutputState }

func (BandwidthLimitsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthLimitsResult)(nil)).Elem()
}

func (o BandwidthLimitsResultOutput) ToBandwidthLimitsResultOutput() BandwidthLimitsResultOutput {
	return o
}

func (o BandwidthLimitsResultOutput) ToBandwidthLimitsResultOutputWithContext(ctx context.Context) BandwidthLimitsResultOutput {
	return o
}

func (o BandwidthLimitsResultOutput) CcnId() pulumi.StringOutput {
	return o.ApplyT(func(v BandwidthLimitsResult) string { return v.CcnId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o BandwidthLimitsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v BandwidthLimitsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o BandwidthLimitsResultOutput) Limits() BandwidthLimitsLimitArrayOutput {
	return o.ApplyT(func(v BandwidthLimitsResult) []BandwidthLimitsLimit { return v.Limits }).(BandwidthLimitsLimitArrayOutput)
}

func (o BandwidthLimitsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthLimitsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BandwidthLimitsResultOutput{})
}
