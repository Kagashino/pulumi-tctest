// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineGroup struct {
	pulumi.CustomResourceState

	// Whether to enable automatic update for the machine group.
	AutoUpdate pulumi.BoolPtrOutput `pulumi:"autoUpdate"`
	// Machine group name, which must be unique.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Type of the machine group to be created.
	MachineGroupType MachineGroupMachineGroupTypeOutput `pulumi:"machineGroupType"`
	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled,
	// the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log
	// topics will be created, which will not incur fees.
	ServiceLogging pulumi.BoolPtrOutput `pulumi:"serviceLogging"`
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime pulumi.StringPtrOutput `pulumi:"updateEndTime"`
	// pdate start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime pulumi.StringPtrOutput `pulumi:"updateStartTime"`
}

// NewMachineGroup registers a new resource with the given unique name, arguments, and options.
func NewMachineGroup(ctx *pulumi.Context,
	name string, args *MachineGroupArgs, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.MachineGroupType == nil {
		return nil, errors.New("invalid value for required argument 'MachineGroupType'")
	}
	var resource MachineGroup
	err := ctx.RegisterResource("tctest:Cls/machineGroup:MachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineGroup gets an existing MachineGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineGroupState, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	var resource MachineGroup
	err := ctx.ReadResource("tctest:Cls/machineGroup:MachineGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineGroup resources.
type machineGroupState struct {
	// Whether to enable automatic update for the machine group.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// Machine group name, which must be unique.
	GroupName *string `pulumi:"groupName"`
	// Type of the machine group to be created.
	MachineGroupType *MachineGroupMachineGroupType `pulumi:"machineGroupType"`
	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled,
	// the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log
	// topics will be created, which will not incur fees.
	ServiceLogging *bool `pulumi:"serviceLogging"`
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags map[string]interface{} `pulumi:"tags"`
	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `pulumi:"updateEndTime"`
	// pdate start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `pulumi:"updateStartTime"`
}

type MachineGroupState struct {
	// Whether to enable automatic update for the machine group.
	AutoUpdate pulumi.BoolPtrInput
	// Machine group name, which must be unique.
	GroupName pulumi.StringPtrInput
	// Type of the machine group to be created.
	MachineGroupType MachineGroupMachineGroupTypePtrInput
	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled,
	// the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log
	// topics will be created, which will not incur fees.
	ServiceLogging pulumi.BoolPtrInput
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags pulumi.MapInput
	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime pulumi.StringPtrInput
	// pdate start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime pulumi.StringPtrInput
}

func (MachineGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupState)(nil)).Elem()
}

type machineGroupArgs struct {
	// Whether to enable automatic update for the machine group.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// Machine group name, which must be unique.
	GroupName string `pulumi:"groupName"`
	// Type of the machine group to be created.
	MachineGroupType MachineGroupMachineGroupType `pulumi:"machineGroupType"`
	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled,
	// the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log
	// topics will be created, which will not incur fees.
	ServiceLogging *bool `pulumi:"serviceLogging"`
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags map[string]interface{} `pulumi:"tags"`
	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime *string `pulumi:"updateEndTime"`
	// pdate start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime *string `pulumi:"updateStartTime"`
}

// The set of arguments for constructing a MachineGroup resource.
type MachineGroupArgs struct {
	// Whether to enable automatic update for the machine group.
	AutoUpdate pulumi.BoolPtrInput
	// Machine group name, which must be unique.
	GroupName pulumi.StringInput
	// Type of the machine group to be created.
	MachineGroupType MachineGroupMachineGroupTypeInput
	// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled,
	// the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log
	// topics will be created, which will not incur fees.
	ServiceLogging pulumi.BoolPtrInput
	// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
	Tags pulumi.MapInput
	// Update end time. We recommend you update LogListener during off-peak hours.
	UpdateEndTime pulumi.StringPtrInput
	// pdate start time. We recommend you update LogListener during off-peak hours.
	UpdateStartTime pulumi.StringPtrInput
}

func (MachineGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupArgs)(nil)).Elem()
}

type MachineGroupInput interface {
	pulumi.Input

	ToMachineGroupOutput() MachineGroupOutput
	ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput
}

func (*MachineGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (i *MachineGroup) ToMachineGroupOutput() MachineGroupOutput {
	return i.ToMachineGroupOutputWithContext(context.Background())
}

func (i *MachineGroup) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupOutput)
}

// MachineGroupArrayInput is an input type that accepts MachineGroupArray and MachineGroupArrayOutput values.
// You can construct a concrete instance of `MachineGroupArrayInput` via:
//
//          MachineGroupArray{ MachineGroupArgs{...} }
type MachineGroupArrayInput interface {
	pulumi.Input

	ToMachineGroupArrayOutput() MachineGroupArrayOutput
	ToMachineGroupArrayOutputWithContext(context.Context) MachineGroupArrayOutput
}

type MachineGroupArray []MachineGroupInput

func (MachineGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineGroup)(nil)).Elem()
}

func (i MachineGroupArray) ToMachineGroupArrayOutput() MachineGroupArrayOutput {
	return i.ToMachineGroupArrayOutputWithContext(context.Background())
}

func (i MachineGroupArray) ToMachineGroupArrayOutputWithContext(ctx context.Context) MachineGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupArrayOutput)
}

// MachineGroupMapInput is an input type that accepts MachineGroupMap and MachineGroupMapOutput values.
// You can construct a concrete instance of `MachineGroupMapInput` via:
//
//          MachineGroupMap{ "key": MachineGroupArgs{...} }
type MachineGroupMapInput interface {
	pulumi.Input

	ToMachineGroupMapOutput() MachineGroupMapOutput
	ToMachineGroupMapOutputWithContext(context.Context) MachineGroupMapOutput
}

type MachineGroupMap map[string]MachineGroupInput

func (MachineGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineGroup)(nil)).Elem()
}

func (i MachineGroupMap) ToMachineGroupMapOutput() MachineGroupMapOutput {
	return i.ToMachineGroupMapOutputWithContext(context.Background())
}

func (i MachineGroupMap) ToMachineGroupMapOutputWithContext(ctx context.Context) MachineGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupMapOutput)
}

type MachineGroupOutput struct{ *pulumi.OutputState }

func (MachineGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (o MachineGroupOutput) ToMachineGroupOutput() MachineGroupOutput {
	return o
}

func (o MachineGroupOutput) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return o
}

// Whether to enable automatic update for the machine group.
func (o MachineGroupOutput) AutoUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.BoolPtrOutput { return v.AutoUpdate }).(pulumi.BoolPtrOutput)
}

// Machine group name, which must be unique.
func (o MachineGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// Type of the machine group to be created.
func (o MachineGroupOutput) MachineGroupType() MachineGroupMachineGroupTypeOutput {
	return o.ApplyT(func(v *MachineGroup) MachineGroupMachineGroupTypeOutput { return v.MachineGroupType }).(MachineGroupMachineGroupTypeOutput)
}

// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled,
// the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log
// topics will be created, which will not incur fees.
func (o MachineGroupOutput) ServiceLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.BoolPtrOutput { return v.ServiceLogging }).(pulumi.BoolPtrOutput)
}

// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
func (o MachineGroupOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Update end time. We recommend you update LogListener during off-peak hours.
func (o MachineGroupOutput) UpdateEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.UpdateEndTime }).(pulumi.StringPtrOutput)
}

// pdate start time. We recommend you update LogListener during off-peak hours.
func (o MachineGroupOutput) UpdateStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.UpdateStartTime }).(pulumi.StringPtrOutput)
}

type MachineGroupArrayOutput struct{ *pulumi.OutputState }

func (MachineGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineGroup)(nil)).Elem()
}

func (o MachineGroupArrayOutput) ToMachineGroupArrayOutput() MachineGroupArrayOutput {
	return o
}

func (o MachineGroupArrayOutput) ToMachineGroupArrayOutputWithContext(ctx context.Context) MachineGroupArrayOutput {
	return o
}

func (o MachineGroupArrayOutput) Index(i pulumi.IntInput) MachineGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineGroup {
		return vs[0].([]*MachineGroup)[vs[1].(int)]
	}).(MachineGroupOutput)
}

type MachineGroupMapOutput struct{ *pulumi.OutputState }

func (MachineGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineGroup)(nil)).Elem()
}

func (o MachineGroupMapOutput) ToMachineGroupMapOutput() MachineGroupMapOutput {
	return o
}

func (o MachineGroupMapOutput) ToMachineGroupMapOutputWithContext(ctx context.Context) MachineGroupMapOutput {
	return o
}

func (o MachineGroupMapOutput) MapIndex(k pulumi.StringInput) MachineGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineGroup {
		return vs[0].(map[string]*MachineGroup)[vs[1].(string)]
	}).(MachineGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineGroupInput)(nil)).Elem(), &MachineGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineGroupArrayInput)(nil)).Elem(), MachineGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineGroupMapInput)(nil)).Elem(), MachineGroupMap{})
	pulumi.RegisterOutputType(MachineGroupOutput{})
	pulumi.RegisterOutputType(MachineGroupArrayOutput{})
	pulumi.RegisterOutputType(MachineGroupMapOutput{})
}
