// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tke

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodePool struct {
	pulumi.CustomResourceState

	// Auto scaling config parameters.
	AutoScalingConfig NodePoolAutoScalingConfigOutput `pulumi:"autoScalingConfig"`
	// The auto scaling group ID.
	AutoScalingGroupId pulumi.StringOutput `pulumi:"autoScalingGroupId"`
	// ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Seconds of scaling group cool down. Default value is `300`.
	DefaultCooldown pulumi.IntOutput `pulumi:"defaultCooldown"`
	// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
	DeleteKeepInstance pulumi.BoolPtrOutput `pulumi:"deleteKeepInstance"`
	// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
	DesiredCapacity pulumi.IntOutput `pulumi:"desiredCapacity"`
	// Indicate whether to enable auto scaling or not.
	EnableAutoScale pulumi.BoolPtrOutput `pulumi:"enableAutoScale"`
	// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English,
	// numbers,'/','-', and does not allow beginning with ('/').
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The launch config ID.
	LaunchConfigId pulumi.StringOutput `pulumi:"launchConfigId"`
	// Maximum number of node.
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// Minimum number of node.
	MinSize pulumi.IntOutput `pulumi:"minSize"`
	// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
	MultiZoneSubnetPolicy pulumi.StringPtrOutput `pulumi:"multiZoneSubnetPolicy"`
	// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores,
	// separators (`-`) and decimal points.
	Name pulumi.StringOutput `pulumi:"name"`
	// Node config.
	NodeConfig NodePoolNodeConfigPtrOutput `pulumi:"nodeConfig"`
	// The total node count.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1
	// LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new
	// nodes, not including the existing nodes.
	NodeOs pulumi.StringPtrOutput `pulumi:"nodeOs"`
	// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter
	// will only affect new nodes, not including the existing nodes.
	NodeOsType pulumi.StringPtrOutput `pulumi:"nodeOsType"`
	// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
	RetryPolicy pulumi.StringPtrOutput `pulumi:"retryPolicy"`
	// Name of relative scaling group.
	ScalingGroupName pulumi.StringOutput `pulumi:"scalingGroupName"`
	// Project ID the scaling group belongs to.
	ScalingGroupProjectId pulumi.IntPtrOutput `pulumi:"scalingGroupProjectId"`
	// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances),
	// `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given
	// priority to the shutdown of the instance. If the number of instances is still lower than the expected number of
	// instances after the startup, the instance will be created, and the method of destroying the instance will still be used
	// for shrinking).
	ScalingMode pulumi.StringPtrOutput `pulumi:"scalingMode"`
	// Status of the node pool.
	Status pulumi.StringOutput `pulumi:"status"`
	// ID list of subnet, and for VPC it is required.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Taints of kubernetes node pool created nodes.
	Taints NodePoolTaintArrayOutput `pulumi:"taints"`
	// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
	TerminationPolicies pulumi.StringOutput `pulumi:"terminationPolicies"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrOutput `pulumi:"unschedulable"`
	// ID of VPC network.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// List of auto scaling group available zones, for Basic network it is required.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewNodePool registers a new resource with the given unique name, arguments, and options.
func NewNodePool(ctx *pulumi.Context,
	name string, args *NodePoolArgs, opts ...pulumi.ResourceOption) (*NodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoScalingConfig == nil {
		return nil, errors.New("invalid value for required argument 'AutoScalingConfig'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.MaxSize == nil {
		return nil, errors.New("invalid value for required argument 'MaxSize'")
	}
	if args.MinSize == nil {
		return nil, errors.New("invalid value for required argument 'MinSize'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource NodePool
	err := ctx.RegisterResource("tencentcloud:Tke/nodePool:NodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodePool gets an existing NodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodePoolState, opts ...pulumi.ResourceOption) (*NodePool, error) {
	var resource NodePool
	err := ctx.ReadResource("tencentcloud:Tke/nodePool:NodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodePool resources.
type nodePoolState struct {
	// Auto scaling config parameters.
	AutoScalingConfig *NodePoolAutoScalingConfig `pulumi:"autoScalingConfig"`
	// The auto scaling group ID.
	AutoScalingGroupId *string `pulumi:"autoScalingGroupId"`
	// ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Seconds of scaling group cool down. Default value is `300`.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
	DeleteKeepInstance *bool `pulumi:"deleteKeepInstance"`
	// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// Indicate whether to enable auto scaling or not.
	EnableAutoScale *bool `pulumi:"enableAutoScale"`
	// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English,
	// numbers,'/','-', and does not allow beginning with ('/').
	Labels map[string]interface{} `pulumi:"labels"`
	// The launch config ID.
	LaunchConfigId *string `pulumi:"launchConfigId"`
	// Maximum number of node.
	MaxSize *int `pulumi:"maxSize"`
	// Minimum number of node.
	MinSize *int `pulumi:"minSize"`
	// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
	MultiZoneSubnetPolicy *string `pulumi:"multiZoneSubnetPolicy"`
	// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores,
	// separators (`-`) and decimal points.
	Name *string `pulumi:"name"`
	// Node config.
	NodeConfig *NodePoolNodeConfig `pulumi:"nodeConfig"`
	// The total node count.
	NodeCount *int `pulumi:"nodeCount"`
	// Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1
	// LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new
	// nodes, not including the existing nodes.
	NodeOs *string `pulumi:"nodeOs"`
	// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter
	// will only affect new nodes, not including the existing nodes.
	NodeOsType *string `pulumi:"nodeOsType"`
	// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
	RetryPolicy *string `pulumi:"retryPolicy"`
	// Name of relative scaling group.
	ScalingGroupName *string `pulumi:"scalingGroupName"`
	// Project ID the scaling group belongs to.
	ScalingGroupProjectId *int `pulumi:"scalingGroupProjectId"`
	// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances),
	// `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given
	// priority to the shutdown of the instance. If the number of instances is still lower than the expected number of
	// instances after the startup, the instance will be created, and the method of destroying the instance will still be used
	// for shrinking).
	ScalingMode *string `pulumi:"scalingMode"`
	// Status of the node pool.
	Status *string `pulumi:"status"`
	// ID list of subnet, and for VPC it is required.
	SubnetIds []string `pulumi:"subnetIds"`
	// Taints of kubernetes node pool created nodes.
	Taints []NodePoolTaint `pulumi:"taints"`
	// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
	TerminationPolicies *string `pulumi:"terminationPolicies"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable *int `pulumi:"unschedulable"`
	// ID of VPC network.
	VpcId *string `pulumi:"vpcId"`
	// List of auto scaling group available zones, for Basic network it is required.
	Zones []string `pulumi:"zones"`
}

type NodePoolState struct {
	// Auto scaling config parameters.
	AutoScalingConfig NodePoolAutoScalingConfigPtrInput
	// The auto scaling group ID.
	AutoScalingGroupId pulumi.StringPtrInput
	// ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// Seconds of scaling group cool down. Default value is `300`.
	DefaultCooldown pulumi.IntPtrInput
	// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
	DeleteKeepInstance pulumi.BoolPtrInput
	// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
	DesiredCapacity pulumi.IntPtrInput
	// Indicate whether to enable auto scaling or not.
	EnableAutoScale pulumi.BoolPtrInput
	// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English,
	// numbers,'/','-', and does not allow beginning with ('/').
	Labels pulumi.MapInput
	// The launch config ID.
	LaunchConfigId pulumi.StringPtrInput
	// Maximum number of node.
	MaxSize pulumi.IntPtrInput
	// Minimum number of node.
	MinSize pulumi.IntPtrInput
	// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
	MultiZoneSubnetPolicy pulumi.StringPtrInput
	// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores,
	// separators (`-`) and decimal points.
	Name pulumi.StringPtrInput
	// Node config.
	NodeConfig NodePoolNodeConfigPtrInput
	// The total node count.
	NodeCount pulumi.IntPtrInput
	// Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1
	// LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new
	// nodes, not including the existing nodes.
	NodeOs pulumi.StringPtrInput
	// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter
	// will only affect new nodes, not including the existing nodes.
	NodeOsType pulumi.StringPtrInput
	// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
	RetryPolicy pulumi.StringPtrInput
	// Name of relative scaling group.
	ScalingGroupName pulumi.StringPtrInput
	// Project ID the scaling group belongs to.
	ScalingGroupProjectId pulumi.IntPtrInput
	// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances),
	// `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given
	// priority to the shutdown of the instance. If the number of instances is still lower than the expected number of
	// instances after the startup, the instance will be created, and the method of destroying the instance will still be used
	// for shrinking).
	ScalingMode pulumi.StringPtrInput
	// Status of the node pool.
	Status pulumi.StringPtrInput
	// ID list of subnet, and for VPC it is required.
	SubnetIds pulumi.StringArrayInput
	// Taints of kubernetes node pool created nodes.
	Taints NodePoolTaintArrayInput
	// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
	TerminationPolicies pulumi.StringPtrInput
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrInput
	// ID of VPC network.
	VpcId pulumi.StringPtrInput
	// List of auto scaling group available zones, for Basic network it is required.
	Zones pulumi.StringArrayInput
}

func (NodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolState)(nil)).Elem()
}

type nodePoolArgs struct {
	// Auto scaling config parameters.
	AutoScalingConfig NodePoolAutoScalingConfig `pulumi:"autoScalingConfig"`
	// ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// Seconds of scaling group cool down. Default value is `300`.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
	DeleteKeepInstance *bool `pulumi:"deleteKeepInstance"`
	// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// Indicate whether to enable auto scaling or not.
	EnableAutoScale *bool `pulumi:"enableAutoScale"`
	// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English,
	// numbers,'/','-', and does not allow beginning with ('/').
	Labels map[string]interface{} `pulumi:"labels"`
	// Maximum number of node.
	MaxSize int `pulumi:"maxSize"`
	// Minimum number of node.
	MinSize int `pulumi:"minSize"`
	// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
	MultiZoneSubnetPolicy *string `pulumi:"multiZoneSubnetPolicy"`
	// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores,
	// separators (`-`) and decimal points.
	Name *string `pulumi:"name"`
	// Node config.
	NodeConfig *NodePoolNodeConfig `pulumi:"nodeConfig"`
	// Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1
	// LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new
	// nodes, not including the existing nodes.
	NodeOs *string `pulumi:"nodeOs"`
	// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter
	// will only affect new nodes, not including the existing nodes.
	NodeOsType *string `pulumi:"nodeOsType"`
	// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
	RetryPolicy *string `pulumi:"retryPolicy"`
	// Name of relative scaling group.
	ScalingGroupName *string `pulumi:"scalingGroupName"`
	// Project ID the scaling group belongs to.
	ScalingGroupProjectId *int `pulumi:"scalingGroupProjectId"`
	// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances),
	// `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given
	// priority to the shutdown of the instance. If the number of instances is still lower than the expected number of
	// instances after the startup, the instance will be created, and the method of destroying the instance will still be used
	// for shrinking).
	ScalingMode *string `pulumi:"scalingMode"`
	// ID list of subnet, and for VPC it is required.
	SubnetIds []string `pulumi:"subnetIds"`
	// Taints of kubernetes node pool created nodes.
	Taints []NodePoolTaint `pulumi:"taints"`
	// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
	TerminationPolicies *string `pulumi:"terminationPolicies"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable *int `pulumi:"unschedulable"`
	// ID of VPC network.
	VpcId string `pulumi:"vpcId"`
	// List of auto scaling group available zones, for Basic network it is required.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a NodePool resource.
type NodePoolArgs struct {
	// Auto scaling config parameters.
	AutoScalingConfig NodePoolAutoScalingConfigInput
	// ID of the cluster.
	ClusterId pulumi.StringInput
	// Seconds of scaling group cool down. Default value is `300`.
	DefaultCooldown pulumi.IntPtrInput
	// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
	DeleteKeepInstance pulumi.BoolPtrInput
	// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
	DesiredCapacity pulumi.IntPtrInput
	// Indicate whether to enable auto scaling or not.
	EnableAutoScale pulumi.BoolPtrInput
	// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English,
	// numbers,'/','-', and does not allow beginning with ('/').
	Labels pulumi.MapInput
	// Maximum number of node.
	MaxSize pulumi.IntInput
	// Minimum number of node.
	MinSize pulumi.IntInput
	// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
	MultiZoneSubnetPolicy pulumi.StringPtrInput
	// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores,
	// separators (`-`) and decimal points.
	Name pulumi.StringPtrInput
	// Node config.
	NodeConfig NodePoolNodeConfigPtrInput
	// Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1
	// LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new
	// nodes, not including the existing nodes.
	NodeOs pulumi.StringPtrInput
	// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter
	// will only affect new nodes, not including the existing nodes.
	NodeOsType pulumi.StringPtrInput
	// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
	RetryPolicy pulumi.StringPtrInput
	// Name of relative scaling group.
	ScalingGroupName pulumi.StringPtrInput
	// Project ID the scaling group belongs to.
	ScalingGroupProjectId pulumi.IntPtrInput
	// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances),
	// `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given
	// priority to the shutdown of the instance. If the number of instances is still lower than the expected number of
	// instances after the startup, the instance will be created, and the method of destroying the instance will still be used
	// for shrinking).
	ScalingMode pulumi.StringPtrInput
	// ID list of subnet, and for VPC it is required.
	SubnetIds pulumi.StringArrayInput
	// Taints of kubernetes node pool created nodes.
	Taints NodePoolTaintArrayInput
	// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
	TerminationPolicies pulumi.StringPtrInput
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrInput
	// ID of VPC network.
	VpcId pulumi.StringInput
	// List of auto scaling group available zones, for Basic network it is required.
	Zones pulumi.StringArrayInput
}

func (NodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolArgs)(nil)).Elem()
}

type NodePoolInput interface {
	pulumi.Input

	ToNodePoolOutput() NodePoolOutput
	ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput
}

func (*NodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePool)(nil)).Elem()
}

func (i *NodePool) ToNodePoolOutput() NodePoolOutput {
	return i.ToNodePoolOutputWithContext(context.Background())
}

func (i *NodePool) ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolOutput)
}

// NodePoolArrayInput is an input type that accepts NodePoolArray and NodePoolArrayOutput values.
// You can construct a concrete instance of `NodePoolArrayInput` via:
//
//          NodePoolArray{ NodePoolArgs{...} }
type NodePoolArrayInput interface {
	pulumi.Input

	ToNodePoolArrayOutput() NodePoolArrayOutput
	ToNodePoolArrayOutputWithContext(context.Context) NodePoolArrayOutput
}

type NodePoolArray []NodePoolInput

func (NodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodePool)(nil)).Elem()
}

func (i NodePoolArray) ToNodePoolArrayOutput() NodePoolArrayOutput {
	return i.ToNodePoolArrayOutputWithContext(context.Background())
}

func (i NodePoolArray) ToNodePoolArrayOutputWithContext(ctx context.Context) NodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolArrayOutput)
}

// NodePoolMapInput is an input type that accepts NodePoolMap and NodePoolMapOutput values.
// You can construct a concrete instance of `NodePoolMapInput` via:
//
//          NodePoolMap{ "key": NodePoolArgs{...} }
type NodePoolMapInput interface {
	pulumi.Input

	ToNodePoolMapOutput() NodePoolMapOutput
	ToNodePoolMapOutputWithContext(context.Context) NodePoolMapOutput
}

type NodePoolMap map[string]NodePoolInput

func (NodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodePool)(nil)).Elem()
}

func (i NodePoolMap) ToNodePoolMapOutput() NodePoolMapOutput {
	return i.ToNodePoolMapOutputWithContext(context.Background())
}

func (i NodePoolMap) ToNodePoolMapOutputWithContext(ctx context.Context) NodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolMapOutput)
}

type NodePoolOutput struct{ *pulumi.OutputState }

func (NodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePool)(nil)).Elem()
}

func (o NodePoolOutput) ToNodePoolOutput() NodePoolOutput {
	return o
}

func (o NodePoolOutput) ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput {
	return o
}

// Auto scaling config parameters.
func (o NodePoolOutput) AutoScalingConfig() NodePoolAutoScalingConfigOutput {
	return o.ApplyT(func(v *NodePool) NodePoolAutoScalingConfigOutput { return v.AutoScalingConfig }).(NodePoolAutoScalingConfigOutput)
}

// The auto scaling group ID.
func (o NodePoolOutput) AutoScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.AutoScalingGroupId }).(pulumi.StringOutput)
}

// ID of the cluster.
func (o NodePoolOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Seconds of scaling group cool down. Default value is `300`.
func (o NodePoolOutput) DefaultCooldown() pulumi.IntOutput {
	return o.ApplyT(func(v *NodePool) pulumi.IntOutput { return v.DefaultCooldown }).(pulumi.IntOutput)
}

// Indicate to keep the CVM instance when delete the node pool. Default is `true`.
func (o NodePoolOutput) DeleteKeepInstance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.BoolPtrOutput { return v.DeleteKeepInstance }).(pulumi.BoolPtrOutput)
}

// Desired capacity ot the node. If `enable_auto_scale` is set `true`, this will be a computed parameter.
func (o NodePoolOutput) DesiredCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *NodePool) pulumi.IntOutput { return v.DesiredCapacity }).(pulumi.IntOutput)
}

// Indicate whether to enable auto scaling or not.
func (o NodePoolOutput) EnableAutoScale() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.BoolPtrOutput { return v.EnableAutoScale }).(pulumi.BoolPtrOutput)
}

// Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English,
// numbers,'/','-', and does not allow beginning with ('/').
func (o NodePoolOutput) Labels() pulumi.MapOutput {
	return o.ApplyT(func(v *NodePool) pulumi.MapOutput { return v.Labels }).(pulumi.MapOutput)
}

// The launch config ID.
func (o NodePoolOutput) LaunchConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.LaunchConfigId }).(pulumi.StringOutput)
}

// Maximum number of node.
func (o NodePoolOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *NodePool) pulumi.IntOutput { return v.MaxSize }).(pulumi.IntOutput)
}

// Minimum number of node.
func (o NodePoolOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v *NodePool) pulumi.IntOutput { return v.MinSize }).(pulumi.IntOutput)
}

// Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
func (o NodePoolOutput) MultiZoneSubnetPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringPtrOutput { return v.MultiZoneSubnetPolicy }).(pulumi.StringPtrOutput)
}

// Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores,
// separators (`-`) and decimal points.
func (o NodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Node config.
func (o NodePoolOutput) NodeConfig() NodePoolNodeConfigPtrOutput {
	return o.ApplyT(func(v *NodePool) NodePoolNodeConfigPtrOutput { return v.NodeConfig }).(NodePoolNodeConfigPtrOutput)
}

// The total node count.
func (o NodePoolOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NodePool) pulumi.IntOutput { return v.NodeCount }).(pulumi.IntOutput)
}

// Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1
// LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new
// nodes, not including the existing nodes.
func (o NodePoolOutput) NodeOs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringPtrOutput { return v.NodeOs }).(pulumi.StringPtrOutput)
}

// The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter
// will only affect new nodes, not including the existing nodes.
func (o NodePoolOutput) NodeOsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringPtrOutput { return v.NodeOsType }).(pulumi.StringPtrOutput)
}

// Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
func (o NodePoolOutput) RetryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringPtrOutput { return v.RetryPolicy }).(pulumi.StringPtrOutput)
}

// Name of relative scaling group.
func (o NodePoolOutput) ScalingGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.ScalingGroupName }).(pulumi.StringOutput)
}

// Project ID the scaling group belongs to.
func (o NodePoolOutput) ScalingGroupProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.IntPtrOutput { return v.ScalingGroupProjectId }).(pulumi.IntPtrOutput)
}

// Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances),
// `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given
// priority to the shutdown of the instance. If the number of instances is still lower than the expected number of
// instances after the startup, the instance will be created, and the method of destroying the instance will still be used
// for shrinking).
func (o NodePoolOutput) ScalingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringPtrOutput { return v.ScalingMode }).(pulumi.StringPtrOutput)
}

// Status of the node pool.
func (o NodePoolOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// ID list of subnet, and for VPC it is required.
func (o NodePoolOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Taints of kubernetes node pool created nodes.
func (o NodePoolOutput) Taints() NodePoolTaintArrayOutput {
	return o.ApplyT(func(v *NodePool) NodePoolTaintArrayOutput { return v.Taints }).(NodePoolTaintArrayOutput)
}

// Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
func (o NodePoolOutput) TerminationPolicies() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.TerminationPolicies }).(pulumi.StringOutput)
}

// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
func (o NodePoolOutput) Unschedulable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodePool) pulumi.IntPtrOutput { return v.Unschedulable }).(pulumi.IntPtrOutput)
}

// ID of VPC network.
func (o NodePoolOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// List of auto scaling group available zones, for Basic network it is required.
func (o NodePoolOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NodePool) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

type NodePoolArrayOutput struct{ *pulumi.OutputState }

func (NodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodePool)(nil)).Elem()
}

func (o NodePoolArrayOutput) ToNodePoolArrayOutput() NodePoolArrayOutput {
	return o
}

func (o NodePoolArrayOutput) ToNodePoolArrayOutputWithContext(ctx context.Context) NodePoolArrayOutput {
	return o
}

func (o NodePoolArrayOutput) Index(i pulumi.IntInput) NodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodePool {
		return vs[0].([]*NodePool)[vs[1].(int)]
	}).(NodePoolOutput)
}

type NodePoolMapOutput struct{ *pulumi.OutputState }

func (NodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodePool)(nil)).Elem()
}

func (o NodePoolMapOutput) ToNodePoolMapOutput() NodePoolMapOutput {
	return o
}

func (o NodePoolMapOutput) ToNodePoolMapOutputWithContext(ctx context.Context) NodePoolMapOutput {
	return o
}

func (o NodePoolMapOutput) MapIndex(k pulumi.StringInput) NodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodePool {
		return vs[0].(map[string]*NodePool)[vs[1].(string)]
	}).(NodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodePoolInput)(nil)).Elem(), &NodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodePoolArrayInput)(nil)).Elem(), NodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodePoolMapInput)(nil)).Elem(), NodePoolMap{})
	pulumi.RegisterOutputType(NodePoolOutput{})
	pulumi.RegisterOutputType(NodePoolArrayOutput{})
	pulumi.RegisterOutputType(NodePoolMapOutput{})
}
