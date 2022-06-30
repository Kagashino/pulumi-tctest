// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Snapshots(ctx *pulumi.Context, args *SnapshotsArgs, opts ...pulumi.InvokeOption) (*SnapshotsResult, error) {
	var rv SnapshotsResult
	err := ctx.Invoke("tencentcloud:Cbs/snapshots:Snapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Snapshots.
type SnapshotsArgs struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	ProjectId        *int    `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SnapshotId       *string `pulumi:"snapshotId"`
	SnapshotName     *string `pulumi:"snapshotName"`
	StorageId        *string `pulumi:"storageId"`
	StorageUsage     *string `pulumi:"storageUsage"`
}

// A collection of values returned by Snapshots.
type SnapshotsResult struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                  `pulumi:"id"`
	ProjectId        *int                    `pulumi:"projectId"`
	ResultOutputFile *string                 `pulumi:"resultOutputFile"`
	SnapshotId       *string                 `pulumi:"snapshotId"`
	SnapshotLists    []SnapshotsSnapshotList `pulumi:"snapshotLists"`
	SnapshotName     *string                 `pulumi:"snapshotName"`
	StorageId        *string                 `pulumi:"storageId"`
	StorageUsage     *string                 `pulumi:"storageUsage"`
}

func SnapshotsOutput(ctx *pulumi.Context, args SnapshotsOutputArgs, opts ...pulumi.InvokeOption) SnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SnapshotsResult, error) {
			args := v.(SnapshotsArgs)
			r, err := Snapshots(ctx, &args, opts...)
			var s SnapshotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SnapshotsResultOutput)
}

// A collection of arguments for invoking Snapshots.
type SnapshotsOutputArgs struct {
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	ProjectId        pulumi.IntPtrInput    `pulumi:"projectId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SnapshotId       pulumi.StringPtrInput `pulumi:"snapshotId"`
	SnapshotName     pulumi.StringPtrInput `pulumi:"snapshotName"`
	StorageId        pulumi.StringPtrInput `pulumi:"storageId"`
	StorageUsage     pulumi.StringPtrInput `pulumi:"storageUsage"`
}

func (SnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotsArgs)(nil)).Elem()
}

// A collection of values returned by Snapshots.
type SnapshotsResultOutput struct{ *pulumi.OutputState }

func (SnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotsResult)(nil)).Elem()
}

func (o SnapshotsResultOutput) ToSnapshotsResultOutput() SnapshotsResultOutput {
	return o
}

func (o SnapshotsResultOutput) ToSnapshotsResultOutputWithContext(ctx context.Context) SnapshotsResultOutput {
	return o
}

func (o SnapshotsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o SnapshotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SnapshotsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o SnapshotsResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SnapshotsResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o SnapshotsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o SnapshotsResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotsResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o SnapshotsResultOutput) SnapshotLists() SnapshotsSnapshotListArrayOutput {
	return o.ApplyT(func(v SnapshotsResult) []SnapshotsSnapshotList { return v.SnapshotLists }).(SnapshotsSnapshotListArrayOutput)
}

func (o SnapshotsResultOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotsResult) *string { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

func (o SnapshotsResultOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotsResult) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

func (o SnapshotsResultOutput) StorageUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotsResult) *string { return v.StorageUsage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotsResultOutput{})
}
