// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tke

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ClusterLevels(ctx *pulumi.Context, args *ClusterLevelsArgs, opts ...pulumi.InvokeOption) (*ClusterLevelsResult, error) {
	var rv ClusterLevelsResult
	err := ctx.Invoke("tctest:Tke/clusterLevels:ClusterLevels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ClusterLevels.
type ClusterLevelsArgs struct {
	ClusterId        *string `pulumi:"clusterId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by ClusterLevels.
type ClusterLevelsResult struct {
	ClusterId *string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string              `pulumi:"id"`
	Lists            []ClusterLevelsList `pulumi:"lists"`
	ResultOutputFile *string             `pulumi:"resultOutputFile"`
}

func ClusterLevelsOutput(ctx *pulumi.Context, args ClusterLevelsOutputArgs, opts ...pulumi.InvokeOption) ClusterLevelsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ClusterLevelsResult, error) {
			args := v.(ClusterLevelsArgs)
			r, err := ClusterLevels(ctx, &args, opts...)
			var s ClusterLevelsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ClusterLevelsResultOutput)
}

// A collection of arguments for invoking ClusterLevels.
type ClusterLevelsOutputArgs struct {
	ClusterId        pulumi.StringPtrInput `pulumi:"clusterId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (ClusterLevelsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLevelsArgs)(nil)).Elem()
}

// A collection of values returned by ClusterLevels.
type ClusterLevelsResultOutput struct{ *pulumi.OutputState }

func (ClusterLevelsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLevelsResult)(nil)).Elem()
}

func (o ClusterLevelsResultOutput) ToClusterLevelsResultOutput() ClusterLevelsResultOutput {
	return o
}

func (o ClusterLevelsResultOutput) ToClusterLevelsResultOutputWithContext(ctx context.Context) ClusterLevelsResultOutput {
	return o
}

func (o ClusterLevelsResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterLevelsResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ClusterLevelsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterLevelsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ClusterLevelsResultOutput) Lists() ClusterLevelsListArrayOutput {
	return o.ApplyT(func(v ClusterLevelsResult) []ClusterLevelsList { return v.Lists }).(ClusterLevelsListArrayOutput)
}

func (o ClusterLevelsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterLevelsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterLevelsResultOutput{})
}
