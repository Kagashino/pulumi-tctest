// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instances

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Instances(ctx *pulumi.Context, args *InstancesArgs, opts ...pulumi.InvokeOption) (*InstancesResult, error) {
	var rv InstancesResult
	err := ctx.Invoke("tencentcloud:Instances/instances:Instances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Instances.
type InstancesArgs struct {
	AvailabilityZone *string                `pulumi:"availabilityZone"`
	InstanceId       *string                `pulumi:"instanceId"`
	InstanceName     *string                `pulumi:"instanceName"`
	ProjectId        *int                   `pulumi:"projectId"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
	SubnetId         *string                `pulumi:"subnetId"`
	Tags             map[string]interface{} `pulumi:"tags"`
	VpcId            *string                `pulumi:"vpcId"`
}

// A collection of values returned by Instances.
type InstancesResult struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                  `pulumi:"id"`
	InstanceId       *string                 `pulumi:"instanceId"`
	InstanceLists    []InstancesInstanceList `pulumi:"instanceLists"`
	InstanceName     *string                 `pulumi:"instanceName"`
	ProjectId        *int                    `pulumi:"projectId"`
	ResultOutputFile *string                 `pulumi:"resultOutputFile"`
	SubnetId         *string                 `pulumi:"subnetId"`
	Tags             map[string]interface{}  `pulumi:"tags"`
	VpcId            *string                 `pulumi:"vpcId"`
}

func InstancesOutput(ctx *pulumi.Context, args InstancesOutputArgs, opts ...pulumi.InvokeOption) InstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (InstancesResult, error) {
			args := v.(InstancesArgs)
			r, err := Instances(ctx, &args, opts...)
			var s InstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(InstancesResultOutput)
}

// A collection of arguments for invoking Instances.
type InstancesOutputArgs struct {
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	InstanceId       pulumi.StringPtrInput `pulumi:"instanceId"`
	InstanceName     pulumi.StringPtrInput `pulumi:"instanceName"`
	ProjectId        pulumi.IntPtrInput    `pulumi:"projectId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SubnetId         pulumi.StringPtrInput `pulumi:"subnetId"`
	Tags             pulumi.MapInput       `pulumi:"tags"`
	VpcId            pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (InstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesArgs)(nil)).Elem()
}

// A collection of values returned by Instances.
type InstancesResultOutput struct{ *pulumi.OutputState }

func (InstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesResult)(nil)).Elem()
}

func (o InstancesResultOutput) ToInstancesResultOutput() InstancesResultOutput {
	return o
}

func (o InstancesResultOutput) ToInstancesResultOutputWithContext(ctx context.Context) InstancesResultOutput {
	return o
}

func (o InstancesResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o InstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o InstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) InstanceLists() InstancesInstanceListArrayOutput {
	return o.ApplyT(func(v InstancesResult) []InstancesInstanceList { return v.InstanceLists }).(InstancesInstanceListArrayOutput)
}

func (o InstancesResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstancesResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o InstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v InstancesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o InstancesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(InstancesResultOutput{})
}
