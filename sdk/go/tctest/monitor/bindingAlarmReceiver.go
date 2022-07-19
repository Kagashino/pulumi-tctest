// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BindingAlarmReceiver struct {
	pulumi.CustomResourceState

	// Policy group ID for binding receivers.
	GroupId pulumi.IntOutput `pulumi:"groupId"`
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the
	// following attributes:
	Receivers BindingAlarmReceiverReceiversPtrOutput `pulumi:"receivers"`
}

// NewBindingAlarmReceiver registers a new resource with the given unique name, arguments, and options.
func NewBindingAlarmReceiver(ctx *pulumi.Context,
	name string, args *BindingAlarmReceiverArgs, opts ...pulumi.ResourceOption) (*BindingAlarmReceiver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	var resource BindingAlarmReceiver
	err := ctx.RegisterResource("tctest:Monitor/bindingAlarmReceiver:BindingAlarmReceiver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBindingAlarmReceiver gets an existing BindingAlarmReceiver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBindingAlarmReceiver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingAlarmReceiverState, opts ...pulumi.ResourceOption) (*BindingAlarmReceiver, error) {
	var resource BindingAlarmReceiver
	err := ctx.ReadResource("tctest:Monitor/bindingAlarmReceiver:BindingAlarmReceiver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BindingAlarmReceiver resources.
type bindingAlarmReceiverState struct {
	// Policy group ID for binding receivers.
	GroupId *int `pulumi:"groupId"`
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the
	// following attributes:
	Receivers *BindingAlarmReceiverReceivers `pulumi:"receivers"`
}

type BindingAlarmReceiverState struct {
	// Policy group ID for binding receivers.
	GroupId pulumi.IntPtrInput
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the
	// following attributes:
	Receivers BindingAlarmReceiverReceiversPtrInput
}

func (BindingAlarmReceiverState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingAlarmReceiverState)(nil)).Elem()
}

type bindingAlarmReceiverArgs struct {
	// Policy group ID for binding receivers.
	GroupId int `pulumi:"groupId"`
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the
	// following attributes:
	Receivers *BindingAlarmReceiverReceivers `pulumi:"receivers"`
}

// The set of arguments for constructing a BindingAlarmReceiver resource.
type BindingAlarmReceiverArgs struct {
	// Policy group ID for binding receivers.
	GroupId pulumi.IntInput
	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the
	// following attributes:
	Receivers BindingAlarmReceiverReceiversPtrInput
}

func (BindingAlarmReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingAlarmReceiverArgs)(nil)).Elem()
}

type BindingAlarmReceiverInput interface {
	pulumi.Input

	ToBindingAlarmReceiverOutput() BindingAlarmReceiverOutput
	ToBindingAlarmReceiverOutputWithContext(ctx context.Context) BindingAlarmReceiverOutput
}

func (*BindingAlarmReceiver) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingAlarmReceiver)(nil)).Elem()
}

func (i *BindingAlarmReceiver) ToBindingAlarmReceiverOutput() BindingAlarmReceiverOutput {
	return i.ToBindingAlarmReceiverOutputWithContext(context.Background())
}

func (i *BindingAlarmReceiver) ToBindingAlarmReceiverOutputWithContext(ctx context.Context) BindingAlarmReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingAlarmReceiverOutput)
}

// BindingAlarmReceiverArrayInput is an input type that accepts BindingAlarmReceiverArray and BindingAlarmReceiverArrayOutput values.
// You can construct a concrete instance of `BindingAlarmReceiverArrayInput` via:
//
//          BindingAlarmReceiverArray{ BindingAlarmReceiverArgs{...} }
type BindingAlarmReceiverArrayInput interface {
	pulumi.Input

	ToBindingAlarmReceiverArrayOutput() BindingAlarmReceiverArrayOutput
	ToBindingAlarmReceiverArrayOutputWithContext(context.Context) BindingAlarmReceiverArrayOutput
}

type BindingAlarmReceiverArray []BindingAlarmReceiverInput

func (BindingAlarmReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindingAlarmReceiver)(nil)).Elem()
}

func (i BindingAlarmReceiverArray) ToBindingAlarmReceiverArrayOutput() BindingAlarmReceiverArrayOutput {
	return i.ToBindingAlarmReceiverArrayOutputWithContext(context.Background())
}

func (i BindingAlarmReceiverArray) ToBindingAlarmReceiverArrayOutputWithContext(ctx context.Context) BindingAlarmReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingAlarmReceiverArrayOutput)
}

// BindingAlarmReceiverMapInput is an input type that accepts BindingAlarmReceiverMap and BindingAlarmReceiverMapOutput values.
// You can construct a concrete instance of `BindingAlarmReceiverMapInput` via:
//
//          BindingAlarmReceiverMap{ "key": BindingAlarmReceiverArgs{...} }
type BindingAlarmReceiverMapInput interface {
	pulumi.Input

	ToBindingAlarmReceiverMapOutput() BindingAlarmReceiverMapOutput
	ToBindingAlarmReceiverMapOutputWithContext(context.Context) BindingAlarmReceiverMapOutput
}

type BindingAlarmReceiverMap map[string]BindingAlarmReceiverInput

func (BindingAlarmReceiverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindingAlarmReceiver)(nil)).Elem()
}

func (i BindingAlarmReceiverMap) ToBindingAlarmReceiverMapOutput() BindingAlarmReceiverMapOutput {
	return i.ToBindingAlarmReceiverMapOutputWithContext(context.Background())
}

func (i BindingAlarmReceiverMap) ToBindingAlarmReceiverMapOutputWithContext(ctx context.Context) BindingAlarmReceiverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingAlarmReceiverMapOutput)
}

type BindingAlarmReceiverOutput struct{ *pulumi.OutputState }

func (BindingAlarmReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingAlarmReceiver)(nil)).Elem()
}

func (o BindingAlarmReceiverOutput) ToBindingAlarmReceiverOutput() BindingAlarmReceiverOutput {
	return o
}

func (o BindingAlarmReceiverOutput) ToBindingAlarmReceiverOutputWithContext(ctx context.Context) BindingAlarmReceiverOutput {
	return o
}

// Policy group ID for binding receivers.
func (o BindingAlarmReceiverOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *BindingAlarmReceiver) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the
// following attributes:
func (o BindingAlarmReceiverOutput) Receivers() BindingAlarmReceiverReceiversPtrOutput {
	return o.ApplyT(func(v *BindingAlarmReceiver) BindingAlarmReceiverReceiversPtrOutput { return v.Receivers }).(BindingAlarmReceiverReceiversPtrOutput)
}

type BindingAlarmReceiverArrayOutput struct{ *pulumi.OutputState }

func (BindingAlarmReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindingAlarmReceiver)(nil)).Elem()
}

func (o BindingAlarmReceiverArrayOutput) ToBindingAlarmReceiverArrayOutput() BindingAlarmReceiverArrayOutput {
	return o
}

func (o BindingAlarmReceiverArrayOutput) ToBindingAlarmReceiverArrayOutputWithContext(ctx context.Context) BindingAlarmReceiverArrayOutput {
	return o
}

func (o BindingAlarmReceiverArrayOutput) Index(i pulumi.IntInput) BindingAlarmReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BindingAlarmReceiver {
		return vs[0].([]*BindingAlarmReceiver)[vs[1].(int)]
	}).(BindingAlarmReceiverOutput)
}

type BindingAlarmReceiverMapOutput struct{ *pulumi.OutputState }

func (BindingAlarmReceiverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindingAlarmReceiver)(nil)).Elem()
}

func (o BindingAlarmReceiverMapOutput) ToBindingAlarmReceiverMapOutput() BindingAlarmReceiverMapOutput {
	return o
}

func (o BindingAlarmReceiverMapOutput) ToBindingAlarmReceiverMapOutputWithContext(ctx context.Context) BindingAlarmReceiverMapOutput {
	return o
}

func (o BindingAlarmReceiverMapOutput) MapIndex(k pulumi.StringInput) BindingAlarmReceiverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BindingAlarmReceiver {
		return vs[0].(map[string]*BindingAlarmReceiver)[vs[1].(string)]
	}).(BindingAlarmReceiverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindingAlarmReceiverInput)(nil)).Elem(), &BindingAlarmReceiver{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingAlarmReceiverArrayInput)(nil)).Elem(), BindingAlarmReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingAlarmReceiverMapInput)(nil)).Elem(), BindingAlarmReceiverMap{})
	pulumi.RegisterOutputType(BindingAlarmReceiverOutput{})
	pulumi.RegisterOutputType(BindingAlarmReceiverArrayOutput{})
	pulumi.RegisterOutputType(BindingAlarmReceiverMapOutput{})
}
