// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tke

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScaleWorker struct {
	pulumi.CustomResourceState

	// ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Configurations of data disk.
	DataDisks ScaleWorkerDataDiskArrayOutput `pulumi:"dataDisks"`
	// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
	DesiredPodNum pulumi.IntPtrOutput `pulumi:"desiredPodNum"`
	// Docker graph path. Default is `/var/lib/docker`.
	DockerGraphPath pulumi.StringPtrOutput `pulumi:"dockerGraphPath"`
	// Custom parameter information related to the node.
	ExtraArgs pulumi.StringArrayOutput `pulumi:"extraArgs"`
	// Labels of kubernetes scale worker created nodes.
	Labels pulumi.MapOutput `pulumi:"labels"`
	// Mount target. Default is not mounting.
	MountTarget pulumi.StringPtrOutput `pulumi:"mountTarget"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrOutput `pulumi:"unschedulable"`
	// Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
	WorkerConfig ScaleWorkerWorkerConfigOutput `pulumi:"workerConfig"`
	// An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
	WorkerInstancesLists ScaleWorkerWorkerInstancesListArrayOutput `pulumi:"workerInstancesLists"`
}

// NewScaleWorker registers a new resource with the given unique name, arguments, and options.
func NewScaleWorker(ctx *pulumi.Context,
	name string, args *ScaleWorkerArgs, opts ...pulumi.ResourceOption) (*ScaleWorker, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.WorkerConfig == nil {
		return nil, errors.New("invalid value for required argument 'WorkerConfig'")
	}
	var resource ScaleWorker
	err := ctx.RegisterResource("tctest:Tke/scaleWorker:ScaleWorker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScaleWorker gets an existing ScaleWorker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScaleWorker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScaleWorkerState, opts ...pulumi.ResourceOption) (*ScaleWorker, error) {
	var resource ScaleWorker
	err := ctx.ReadResource("tctest:Tke/scaleWorker:ScaleWorker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScaleWorker resources.
type scaleWorkerState struct {
	// ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Configurations of data disk.
	DataDisks []ScaleWorkerDataDisk `pulumi:"dataDisks"`
	// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
	DesiredPodNum *int `pulumi:"desiredPodNum"`
	// Docker graph path. Default is `/var/lib/docker`.
	DockerGraphPath *string `pulumi:"dockerGraphPath"`
	// Custom parameter information related to the node.
	ExtraArgs []string `pulumi:"extraArgs"`
	// Labels of kubernetes scale worker created nodes.
	Labels map[string]interface{} `pulumi:"labels"`
	// Mount target. Default is not mounting.
	MountTarget *string `pulumi:"mountTarget"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable *int `pulumi:"unschedulable"`
	// Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
	WorkerConfig *ScaleWorkerWorkerConfig `pulumi:"workerConfig"`
	// An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
	WorkerInstancesLists []ScaleWorkerWorkerInstancesList `pulumi:"workerInstancesLists"`
}

type ScaleWorkerState struct {
	// ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// Configurations of data disk.
	DataDisks ScaleWorkerDataDiskArrayInput
	// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
	DesiredPodNum pulumi.IntPtrInput
	// Docker graph path. Default is `/var/lib/docker`.
	DockerGraphPath pulumi.StringPtrInput
	// Custom parameter information related to the node.
	ExtraArgs pulumi.StringArrayInput
	// Labels of kubernetes scale worker created nodes.
	Labels pulumi.MapInput
	// Mount target. Default is not mounting.
	MountTarget pulumi.StringPtrInput
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrInput
	// Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
	WorkerConfig ScaleWorkerWorkerConfigPtrInput
	// An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
	WorkerInstancesLists ScaleWorkerWorkerInstancesListArrayInput
}

func (ScaleWorkerState) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleWorkerState)(nil)).Elem()
}

type scaleWorkerArgs struct {
	// ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// Configurations of data disk.
	DataDisks []ScaleWorkerDataDisk `pulumi:"dataDisks"`
	// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
	DesiredPodNum *int `pulumi:"desiredPodNum"`
	// Docker graph path. Default is `/var/lib/docker`.
	DockerGraphPath *string `pulumi:"dockerGraphPath"`
	// Custom parameter information related to the node.
	ExtraArgs []string `pulumi:"extraArgs"`
	// Labels of kubernetes scale worker created nodes.
	Labels map[string]interface{} `pulumi:"labels"`
	// Mount target. Default is not mounting.
	MountTarget *string `pulumi:"mountTarget"`
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable *int `pulumi:"unschedulable"`
	// Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
	WorkerConfig ScaleWorkerWorkerConfig `pulumi:"workerConfig"`
}

// The set of arguments for constructing a ScaleWorker resource.
type ScaleWorkerArgs struct {
	// ID of the cluster.
	ClusterId pulumi.StringInput
	// Configurations of data disk.
	DataDisks ScaleWorkerDataDiskArrayInput
	// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
	DesiredPodNum pulumi.IntPtrInput
	// Docker graph path. Default is `/var/lib/docker`.
	DockerGraphPath pulumi.StringPtrInput
	// Custom parameter information related to the node.
	ExtraArgs pulumi.StringArrayInput
	// Labels of kubernetes scale worker created nodes.
	Labels pulumi.MapInput
	// Mount target. Default is not mounting.
	MountTarget pulumi.StringPtrInput
	// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
	Unschedulable pulumi.IntPtrInput
	// Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
	WorkerConfig ScaleWorkerWorkerConfigInput
}

func (ScaleWorkerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleWorkerArgs)(nil)).Elem()
}

type ScaleWorkerInput interface {
	pulumi.Input

	ToScaleWorkerOutput() ScaleWorkerOutput
	ToScaleWorkerOutputWithContext(ctx context.Context) ScaleWorkerOutput
}

func (*ScaleWorker) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleWorker)(nil)).Elem()
}

func (i *ScaleWorker) ToScaleWorkerOutput() ScaleWorkerOutput {
	return i.ToScaleWorkerOutputWithContext(context.Background())
}

func (i *ScaleWorker) ToScaleWorkerOutputWithContext(ctx context.Context) ScaleWorkerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleWorkerOutput)
}

// ScaleWorkerArrayInput is an input type that accepts ScaleWorkerArray and ScaleWorkerArrayOutput values.
// You can construct a concrete instance of `ScaleWorkerArrayInput` via:
//
//          ScaleWorkerArray{ ScaleWorkerArgs{...} }
type ScaleWorkerArrayInput interface {
	pulumi.Input

	ToScaleWorkerArrayOutput() ScaleWorkerArrayOutput
	ToScaleWorkerArrayOutputWithContext(context.Context) ScaleWorkerArrayOutput
}

type ScaleWorkerArray []ScaleWorkerInput

func (ScaleWorkerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScaleWorker)(nil)).Elem()
}

func (i ScaleWorkerArray) ToScaleWorkerArrayOutput() ScaleWorkerArrayOutput {
	return i.ToScaleWorkerArrayOutputWithContext(context.Background())
}

func (i ScaleWorkerArray) ToScaleWorkerArrayOutputWithContext(ctx context.Context) ScaleWorkerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleWorkerArrayOutput)
}

// ScaleWorkerMapInput is an input type that accepts ScaleWorkerMap and ScaleWorkerMapOutput values.
// You can construct a concrete instance of `ScaleWorkerMapInput` via:
//
//          ScaleWorkerMap{ "key": ScaleWorkerArgs{...} }
type ScaleWorkerMapInput interface {
	pulumi.Input

	ToScaleWorkerMapOutput() ScaleWorkerMapOutput
	ToScaleWorkerMapOutputWithContext(context.Context) ScaleWorkerMapOutput
}

type ScaleWorkerMap map[string]ScaleWorkerInput

func (ScaleWorkerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScaleWorker)(nil)).Elem()
}

func (i ScaleWorkerMap) ToScaleWorkerMapOutput() ScaleWorkerMapOutput {
	return i.ToScaleWorkerMapOutputWithContext(context.Background())
}

func (i ScaleWorkerMap) ToScaleWorkerMapOutputWithContext(ctx context.Context) ScaleWorkerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleWorkerMapOutput)
}

type ScaleWorkerOutput struct{ *pulumi.OutputState }

func (ScaleWorkerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleWorker)(nil)).Elem()
}

func (o ScaleWorkerOutput) ToScaleWorkerOutput() ScaleWorkerOutput {
	return o
}

func (o ScaleWorkerOutput) ToScaleWorkerOutputWithContext(ctx context.Context) ScaleWorkerOutput {
	return o
}

// ID of the cluster.
func (o ScaleWorkerOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScaleWorker) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Configurations of data disk.
func (o ScaleWorkerOutput) DataDisks() ScaleWorkerDataDiskArrayOutput {
	return o.ApplyT(func(v *ScaleWorker) ScaleWorkerDataDiskArrayOutput { return v.DataDisks }).(ScaleWorkerDataDiskArrayOutput)
}

// Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
func (o ScaleWorkerOutput) DesiredPodNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleWorker) pulumi.IntPtrOutput { return v.DesiredPodNum }).(pulumi.IntPtrOutput)
}

// Docker graph path. Default is `/var/lib/docker`.
func (o ScaleWorkerOutput) DockerGraphPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScaleWorker) pulumi.StringPtrOutput { return v.DockerGraphPath }).(pulumi.StringPtrOutput)
}

// Custom parameter information related to the node.
func (o ScaleWorkerOutput) ExtraArgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScaleWorker) pulumi.StringArrayOutput { return v.ExtraArgs }).(pulumi.StringArrayOutput)
}

// Labels of kubernetes scale worker created nodes.
func (o ScaleWorkerOutput) Labels() pulumi.MapOutput {
	return o.ApplyT(func(v *ScaleWorker) pulumi.MapOutput { return v.Labels }).(pulumi.MapOutput)
}

// Mount target. Default is not mounting.
func (o ScaleWorkerOutput) MountTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScaleWorker) pulumi.StringPtrOutput { return v.MountTarget }).(pulumi.StringPtrOutput)
}

// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
func (o ScaleWorkerOutput) Unschedulable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleWorker) pulumi.IntPtrOutput { return v.Unschedulable }).(pulumi.IntPtrOutput)
}

// Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
func (o ScaleWorkerOutput) WorkerConfig() ScaleWorkerWorkerConfigOutput {
	return o.ApplyT(func(v *ScaleWorker) ScaleWorkerWorkerConfigOutput { return v.WorkerConfig }).(ScaleWorkerWorkerConfigOutput)
}

// An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
func (o ScaleWorkerOutput) WorkerInstancesLists() ScaleWorkerWorkerInstancesListArrayOutput {
	return o.ApplyT(func(v *ScaleWorker) ScaleWorkerWorkerInstancesListArrayOutput { return v.WorkerInstancesLists }).(ScaleWorkerWorkerInstancesListArrayOutput)
}

type ScaleWorkerArrayOutput struct{ *pulumi.OutputState }

func (ScaleWorkerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScaleWorker)(nil)).Elem()
}

func (o ScaleWorkerArrayOutput) ToScaleWorkerArrayOutput() ScaleWorkerArrayOutput {
	return o
}

func (o ScaleWorkerArrayOutput) ToScaleWorkerArrayOutputWithContext(ctx context.Context) ScaleWorkerArrayOutput {
	return o
}

func (o ScaleWorkerArrayOutput) Index(i pulumi.IntInput) ScaleWorkerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScaleWorker {
		return vs[0].([]*ScaleWorker)[vs[1].(int)]
	}).(ScaleWorkerOutput)
}

type ScaleWorkerMapOutput struct{ *pulumi.OutputState }

func (ScaleWorkerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScaleWorker)(nil)).Elem()
}

func (o ScaleWorkerMapOutput) ToScaleWorkerMapOutput() ScaleWorkerMapOutput {
	return o
}

func (o ScaleWorkerMapOutput) ToScaleWorkerMapOutputWithContext(ctx context.Context) ScaleWorkerMapOutput {
	return o
}

func (o ScaleWorkerMapOutput) MapIndex(k pulumi.StringInput) ScaleWorkerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScaleWorker {
		return vs[0].(map[string]*ScaleWorker)[vs[1].(string)]
	}).(ScaleWorkerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleWorkerInput)(nil)).Elem(), &ScaleWorker{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleWorkerArrayInput)(nil)).Elem(), ScaleWorkerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleWorkerMapInput)(nil)).Elem(), ScaleWorkerMap{})
	pulumi.RegisterOutputType(ScaleWorkerOutput{})
	pulumi.RegisterOutputType(ScaleWorkerArrayOutput{})
	pulumi.RegisterOutputType(ScaleWorkerMapOutput{})
}
