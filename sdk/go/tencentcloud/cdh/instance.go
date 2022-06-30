// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Instance struct {
	pulumi.CustomResourceState

	// The available zone for the CDH instance.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// Create time of the instance.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Id of CVM instances that have been created on the CDH instance.
	CvmInstanceIds pulumi.StringArrayOutput `pulumi:"cvmInstanceIds"`
	// Expired time of the instance.
	ExpiredTime pulumi.StringOutput `pulumi:"expiredTime"`
	// The name of the CDH instance. The max length of host_name is 60.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// An information list of host resource. Each element contains the following attributes:
	HostResources InstanceHostResourceArrayOutput `pulumi:"hostResources"`
	// State of the CDH instance.
	HostState pulumi.StringOutput `pulumi:"hostState"`
	// The type of the CDH instance.
	HostType pulumi.StringPtrOutput `pulumi:"hostType"`
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod pulumi.IntPtrOutput `pulumi:"prepaidPeriod"`
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
	// `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
	// neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
	// specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
	// balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	PrepaidRenewFlag pulumi.StringOutput `pulumi:"prepaidRenewFlag"`
	// The project the instance belongs to, default to 0.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Cdh/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("tencentcloud:Cdh/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The available zone for the CDH instance.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	ChargeType *string `pulumi:"chargeType"`
	// Create time of the instance.
	CreateTime *string `pulumi:"createTime"`
	// Id of CVM instances that have been created on the CDH instance.
	CvmInstanceIds []string `pulumi:"cvmInstanceIds"`
	// Expired time of the instance.
	ExpiredTime *string `pulumi:"expiredTime"`
	// The name of the CDH instance. The max length of host_name is 60.
	HostName *string `pulumi:"hostName"`
	// An information list of host resource. Each element contains the following attributes:
	HostResources []InstanceHostResource `pulumi:"hostResources"`
	// State of the CDH instance.
	HostState *string `pulumi:"hostState"`
	// The type of the CDH instance.
	HostType *string `pulumi:"hostType"`
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
	// `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
	// neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
	// specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
	// balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	PrepaidRenewFlag *string `pulumi:"prepaidRenewFlag"`
	// The project the instance belongs to, default to 0.
	ProjectId *int `pulumi:"projectId"`
}

type InstanceState struct {
	// The available zone for the CDH instance.
	AvailabilityZone pulumi.StringPtrInput
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	ChargeType pulumi.StringPtrInput
	// Create time of the instance.
	CreateTime pulumi.StringPtrInput
	// Id of CVM instances that have been created on the CDH instance.
	CvmInstanceIds pulumi.StringArrayInput
	// Expired time of the instance.
	ExpiredTime pulumi.StringPtrInput
	// The name of the CDH instance. The max length of host_name is 60.
	HostName pulumi.StringPtrInput
	// An information list of host resource. Each element contains the following attributes:
	HostResources InstanceHostResourceArrayInput
	// State of the CDH instance.
	HostState pulumi.StringPtrInput
	// The type of the CDH instance.
	HostType pulumi.StringPtrInput
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod pulumi.IntPtrInput
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
	// `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
	// neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
	// specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
	// balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	PrepaidRenewFlag pulumi.StringPtrInput
	// The project the instance belongs to, default to 0.
	ProjectId pulumi.IntPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The available zone for the CDH instance.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	ChargeType *string `pulumi:"chargeType"`
	// The name of the CDH instance. The max length of host_name is 60.
	HostName *string `pulumi:"hostName"`
	// The type of the CDH instance.
	HostType *string `pulumi:"hostType"`
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
	// `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
	// neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
	// specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
	// balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	PrepaidRenewFlag *string `pulumi:"prepaidRenewFlag"`
	// The project the instance belongs to, default to 0.
	ProjectId *int `pulumi:"projectId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The available zone for the CDH instance.
	AvailabilityZone pulumi.StringInput
	// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
	ChargeType pulumi.StringPtrInput
	// The name of the CDH instance. The max length of host_name is 60.
	HostName pulumi.StringPtrInput
	// The type of the CDH instance.
	HostType pulumi.StringPtrInput
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod pulumi.IntPtrInput
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
	// `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
	// neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
	// specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
	// balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
	PrepaidRenewFlag pulumi.StringPtrInput
	// The project the instance belongs to, default to 0.
	ProjectId pulumi.IntPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The available zone for the CDH instance.
func (o InstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
func (o InstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// Create time of the instance.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Id of CVM instances that have been created on the CDH instance.
func (o InstanceOutput) CvmInstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.CvmInstanceIds }).(pulumi.StringArrayOutput)
}

// Expired time of the instance.
func (o InstanceOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ExpiredTime }).(pulumi.StringOutput)
}

// The name of the CDH instance. The max length of host_name is 60.
func (o InstanceOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// An information list of host resource. Each element contains the following attributes:
func (o InstanceOutput) HostResources() InstanceHostResourceArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceHostResourceArrayOutput { return v.HostResources }).(InstanceHostResourceArrayOutput)
}

// State of the CDH instance.
func (o InstanceOutput) HostState() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.HostState }).(pulumi.StringOutput)
}

// The type of the CDH instance.
func (o InstanceOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.HostType }).(pulumi.StringPtrOutput)
}

// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
func (o InstanceOutput) PrepaidPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.PrepaidPeriod }).(pulumi.IntPtrOutput)
}

// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
// `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
// neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
// specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
// balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
func (o InstanceOutput) PrepaidRenewFlag() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PrepaidRenewFlag }).(pulumi.StringOutput)
}

// The project the instance belongs to, default to 0.
func (o InstanceOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
