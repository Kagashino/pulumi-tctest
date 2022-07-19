// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func VipEipAttachments(ctx *pulumi.Context, args *VipEipAttachmentsArgs, opts ...pulumi.InvokeOption) (*VipEipAttachmentsResult, error) {
	var rv VipEipAttachmentsResult
	err := ctx.Invoke("tctest:Ha/vipEipAttachments:VipEipAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking VipEipAttachments.
type VipEipAttachmentsArgs struct {
	AddressIp        *string `pulumi:"addressIp"`
	HavipId          string  `pulumi:"havipId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by VipEipAttachments.
type VipEipAttachmentsResult struct {
	AddressIp               *string                                   `pulumi:"addressIp"`
	HaVipEipAttachmentLists []VipEipAttachmentsHaVipEipAttachmentList `pulumi:"haVipEipAttachmentLists"`
	HavipId                 string                                    `pulumi:"havipId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func VipEipAttachmentsOutput(ctx *pulumi.Context, args VipEipAttachmentsOutputArgs, opts ...pulumi.InvokeOption) VipEipAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VipEipAttachmentsResult, error) {
			args := v.(VipEipAttachmentsArgs)
			r, err := VipEipAttachments(ctx, &args, opts...)
			var s VipEipAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VipEipAttachmentsResultOutput)
}

// A collection of arguments for invoking VipEipAttachments.
type VipEipAttachmentsOutputArgs struct {
	AddressIp        pulumi.StringPtrInput `pulumi:"addressIp"`
	HavipId          pulumi.StringInput    `pulumi:"havipId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (VipEipAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VipEipAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by VipEipAttachments.
type VipEipAttachmentsResultOutput struct{ *pulumi.OutputState }

func (VipEipAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VipEipAttachmentsResult)(nil)).Elem()
}

func (o VipEipAttachmentsResultOutput) ToVipEipAttachmentsResultOutput() VipEipAttachmentsResultOutput {
	return o
}

func (o VipEipAttachmentsResultOutput) ToVipEipAttachmentsResultOutputWithContext(ctx context.Context) VipEipAttachmentsResultOutput {
	return o
}

func (o VipEipAttachmentsResultOutput) AddressIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VipEipAttachmentsResult) *string { return v.AddressIp }).(pulumi.StringPtrOutput)
}

func (o VipEipAttachmentsResultOutput) HaVipEipAttachmentLists() VipEipAttachmentsHaVipEipAttachmentListArrayOutput {
	return o.ApplyT(func(v VipEipAttachmentsResult) []VipEipAttachmentsHaVipEipAttachmentList {
		return v.HaVipEipAttachmentLists
	}).(VipEipAttachmentsHaVipEipAttachmentListArrayOutput)
}

func (o VipEipAttachmentsResultOutput) HavipId() pulumi.StringOutput {
	return o.ApplyT(func(v VipEipAttachmentsResult) string { return v.HavipId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o VipEipAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VipEipAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VipEipAttachmentsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VipEipAttachmentsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VipEipAttachmentsResultOutput{})
}
