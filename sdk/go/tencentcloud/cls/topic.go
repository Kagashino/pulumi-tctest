// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Topic struct {
	pulumi.CustomResourceState

	// Whether to enable automatic split. Default value: true.
	AutoSplit pulumi.BoolPtrOutput `pulumi:"autoSplit"`
	// Logset ID.
	LogsetId pulumi.StringOutput `pulumi:"logsetId"`
	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
	MaxSplitPartitions pulumi.IntPtrOutput `pulumi:"maxSplitPartitions"`
	// Number of log topic partitions. Default value: 1. Maximum value: 10.
	PartitionCount pulumi.IntPtrOutput `pulumi:"partitionCount"`
	// Lifecycle in days. Value range: 1~366. Default value: 30.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
	// passed in, please contact the customer service to add the log topic to the allowlist first..
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Log topic name.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogsetId == nil {
		return nil, errors.New("invalid value for required argument 'LogsetId'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	var resource Topic
	err := ctx.RegisterResource("tencentcloud:Cls/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("tencentcloud:Cls/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Whether to enable automatic split. Default value: true.
	AutoSplit *bool `pulumi:"autoSplit"`
	// Logset ID.
	LogsetId *string `pulumi:"logsetId"`
	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
	MaxSplitPartitions *int `pulumi:"maxSplitPartitions"`
	// Number of log topic partitions. Default value: 1. Maximum value: 10.
	PartitionCount *int `pulumi:"partitionCount"`
	// Lifecycle in days. Value range: 1~366. Default value: 30.
	Period *int `pulumi:"period"`
	// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
	// passed in, please contact the customer service to add the log topic to the allowlist first..
	StorageType *string `pulumi:"storageType"`
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags map[string]interface{} `pulumi:"tags"`
	// Log topic name.
	TopicName *string `pulumi:"topicName"`
}

type TopicState struct {
	// Whether to enable automatic split. Default value: true.
	AutoSplit pulumi.BoolPtrInput
	// Logset ID.
	LogsetId pulumi.StringPtrInput
	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
	MaxSplitPartitions pulumi.IntPtrInput
	// Number of log topic partitions. Default value: 1. Maximum value: 10.
	PartitionCount pulumi.IntPtrInput
	// Lifecycle in days. Value range: 1~366. Default value: 30.
	Period pulumi.IntPtrInput
	// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
	// passed in, please contact the customer service to add the log topic to the allowlist first..
	StorageType pulumi.StringPtrInput
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags pulumi.MapInput
	// Log topic name.
	TopicName pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// Whether to enable automatic split. Default value: true.
	AutoSplit *bool `pulumi:"autoSplit"`
	// Logset ID.
	LogsetId string `pulumi:"logsetId"`
	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
	MaxSplitPartitions *int `pulumi:"maxSplitPartitions"`
	// Number of log topic partitions. Default value: 1. Maximum value: 10.
	PartitionCount *int `pulumi:"partitionCount"`
	// Lifecycle in days. Value range: 1~366. Default value: 30.
	Period *int `pulumi:"period"`
	// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
	// passed in, please contact the customer service to add the log topic to the allowlist first..
	StorageType *string `pulumi:"storageType"`
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags map[string]interface{} `pulumi:"tags"`
	// Log topic name.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// Whether to enable automatic split. Default value: true.
	AutoSplit pulumi.BoolPtrInput
	// Logset ID.
	LogsetId pulumi.StringInput
	// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
	MaxSplitPartitions pulumi.IntPtrInput
	// Number of log topic partitions. Default value: 1. Maximum value: 10.
	PartitionCount pulumi.IntPtrInput
	// Lifecycle in days. Value range: 1~366. Default value: 30.
	Period pulumi.IntPtrInput
	// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
	// passed in, please contact the customer service to add the log topic to the allowlist first..
	StorageType pulumi.StringPtrInput
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags pulumi.MapInput
	// Log topic name.
	TopicName pulumi.StringInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//          TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//          TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

// Whether to enable automatic split. Default value: true.
func (o TopicOutput) AutoSplit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.AutoSplit }).(pulumi.BoolPtrOutput)
}

// Logset ID.
func (o TopicOutput) LogsetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.LogsetId }).(pulumi.StringOutput)
}

// Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.
func (o TopicOutput) MaxSplitPartitions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.MaxSplitPartitions }).(pulumi.IntPtrOutput)
}

// Number of log topic partitions. Default value: 1. Maximum value: 10.
func (o TopicOutput) PartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.PartitionCount }).(pulumi.IntPtrOutput)
}

// Lifecycle in days. Value range: 1~366. Default value: 30.
func (o TopicOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is
// passed in, please contact the customer service to add the log topic to the allowlist first..
func (o TopicOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.StorageType }).(pulumi.StringPtrOutput)
}

// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
func (o TopicOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Topic) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Log topic name.
func (o TopicOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].([]*Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].(map[string]*Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicArrayInput)(nil)).Elem(), TopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicMapInput)(nil)).Elem(), TopicMap{})
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
