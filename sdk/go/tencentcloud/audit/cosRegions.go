// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package audit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CosRegions(ctx *pulumi.Context, args *CosRegionsArgs, opts ...pulumi.InvokeOption) (*CosRegionsResult, error) {
	var rv CosRegionsResult
	err := ctx.Invoke("tencentcloud:Audit/cosRegions:CosRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking CosRegions.
type CosRegionsArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by CosRegions.
type CosRegionsResult struct {
	AuditCosRegionLists []CosRegionsAuditCosRegionList `pulumi:"auditCosRegionLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func CosRegionsOutput(ctx *pulumi.Context, args CosRegionsOutputArgs, opts ...pulumi.InvokeOption) CosRegionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CosRegionsResult, error) {
			args := v.(CosRegionsArgs)
			r, err := CosRegions(ctx, &args, opts...)
			var s CosRegionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CosRegionsResultOutput)
}

// A collection of arguments for invoking CosRegions.
type CosRegionsOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (CosRegionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CosRegionsArgs)(nil)).Elem()
}

// A collection of values returned by CosRegions.
type CosRegionsResultOutput struct{ *pulumi.OutputState }

func (CosRegionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosRegionsResult)(nil)).Elem()
}

func (o CosRegionsResultOutput) ToCosRegionsResultOutput() CosRegionsResultOutput {
	return o
}

func (o CosRegionsResultOutput) ToCosRegionsResultOutputWithContext(ctx context.Context) CosRegionsResultOutput {
	return o
}

func (o CosRegionsResultOutput) AuditCosRegionLists() CosRegionsAuditCosRegionListArrayOutput {
	return o.ApplyT(func(v CosRegionsResult) []CosRegionsAuditCosRegionList { return v.AuditCosRegionLists }).(CosRegionsAuditCosRegionListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o CosRegionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CosRegionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o CosRegionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CosRegionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CosRegionsResultOutput{})
}
