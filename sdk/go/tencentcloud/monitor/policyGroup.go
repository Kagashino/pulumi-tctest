// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyGroup struct {
	pulumi.CustomResourceState

	// A list binding objects(list only those in the `provider.region`). Each element contains the following attributes:
	BindingObjects PolicyGroupBindingObjectArrayOutput `pulumi:"bindingObjects"`
	// A list of threshold rules. Each element contains the following attributes:
	Conditions PolicyGroupConditionArrayOutput `pulumi:"conditions"`
	// A list of dimensions for this policy group.
	DimensionGroups pulumi.StringArrayOutput `pulumi:"dimensionGroups"`
	// A list of event rules. Each element contains the following attributes:
	EventConditions PolicyGroupEventConditionArrayOutput `pulumi:"eventConditions"`
	// Policy group name, length should between 1 and 20.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
	// alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
	IsUnionRule pulumi.IntPtrOutput `pulumi:"isUnionRule"`
	// Recently edited user uin.
	LastEditUin pulumi.StringOutput `pulumi:"lastEditUin"`
	// Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
	// `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
	PolicyViewName pulumi.StringOutput `pulumi:"policyViewName"`
	// The project id to which the policy group belongs, default is `0`.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// A list of receivers. Each element contains the following attributes:
	Receivers PolicyGroupReceiverArrayOutput `pulumi:"receivers"`
	// Policy group's remark information.
	Remark pulumi.StringOutput `pulumi:"remark"`
	// Support regions this policy group.
	SupportRegions pulumi.StringArrayOutput `pulumi:"supportRegions"`
	// The policy group update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewPolicyGroup registers a new resource with the given unique name, arguments, and options.
func NewPolicyGroup(ctx *pulumi.Context,
	name string, args *PolicyGroupArgs, opts ...pulumi.ResourceOption) (*PolicyGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.PolicyViewName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyViewName'")
	}
	if args.Remark == nil {
		return nil, errors.New("invalid value for required argument 'Remark'")
	}
	var resource PolicyGroup
	err := ctx.RegisterResource("tencentcloud:Monitor/policyGroup:PolicyGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyGroup gets an existing PolicyGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyGroupState, opts ...pulumi.ResourceOption) (*PolicyGroup, error) {
	var resource PolicyGroup
	err := ctx.ReadResource("tencentcloud:Monitor/policyGroup:PolicyGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyGroup resources.
type policyGroupState struct {
	// A list binding objects(list only those in the `provider.region`). Each element contains the following attributes:
	BindingObjects []PolicyGroupBindingObject `pulumi:"bindingObjects"`
	// A list of threshold rules. Each element contains the following attributes:
	Conditions []PolicyGroupCondition `pulumi:"conditions"`
	// A list of dimensions for this policy group.
	DimensionGroups []string `pulumi:"dimensionGroups"`
	// A list of event rules. Each element contains the following attributes:
	EventConditions []PolicyGroupEventCondition `pulumi:"eventConditions"`
	// Policy group name, length should between 1 and 20.
	GroupName *string `pulumi:"groupName"`
	// The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
	// alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
	IsUnionRule *int `pulumi:"isUnionRule"`
	// Recently edited user uin.
	LastEditUin *string `pulumi:"lastEditUin"`
	// Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
	// `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
	PolicyViewName *string `pulumi:"policyViewName"`
	// The project id to which the policy group belongs, default is `0`.
	ProjectId *int `pulumi:"projectId"`
	// A list of receivers. Each element contains the following attributes:
	Receivers []PolicyGroupReceiver `pulumi:"receivers"`
	// Policy group's remark information.
	Remark *string `pulumi:"remark"`
	// Support regions this policy group.
	SupportRegions []string `pulumi:"supportRegions"`
	// The policy group update time.
	UpdateTime *string `pulumi:"updateTime"`
}

type PolicyGroupState struct {
	// A list binding objects(list only those in the `provider.region`). Each element contains the following attributes:
	BindingObjects PolicyGroupBindingObjectArrayInput
	// A list of threshold rules. Each element contains the following attributes:
	Conditions PolicyGroupConditionArrayInput
	// A list of dimensions for this policy group.
	DimensionGroups pulumi.StringArrayInput
	// A list of event rules. Each element contains the following attributes:
	EventConditions PolicyGroupEventConditionArrayInput
	// Policy group name, length should between 1 and 20.
	GroupName pulumi.StringPtrInput
	// The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
	// alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
	IsUnionRule pulumi.IntPtrInput
	// Recently edited user uin.
	LastEditUin pulumi.StringPtrInput
	// Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
	// `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
	PolicyViewName pulumi.StringPtrInput
	// The project id to which the policy group belongs, default is `0`.
	ProjectId pulumi.IntPtrInput
	// A list of receivers. Each element contains the following attributes:
	Receivers PolicyGroupReceiverArrayInput
	// Policy group's remark information.
	Remark pulumi.StringPtrInput
	// Support regions this policy group.
	SupportRegions pulumi.StringArrayInput
	// The policy group update time.
	UpdateTime pulumi.StringPtrInput
}

func (PolicyGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyGroupState)(nil)).Elem()
}

type policyGroupArgs struct {
	// A list of threshold rules. Each element contains the following attributes:
	Conditions []PolicyGroupCondition `pulumi:"conditions"`
	// A list of event rules. Each element contains the following attributes:
	EventConditions []PolicyGroupEventCondition `pulumi:"eventConditions"`
	// Policy group name, length should between 1 and 20.
	GroupName string `pulumi:"groupName"`
	// The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
	// alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
	IsUnionRule *int `pulumi:"isUnionRule"`
	// Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
	// `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
	PolicyViewName string `pulumi:"policyViewName"`
	// The project id to which the policy group belongs, default is `0`.
	ProjectId *int `pulumi:"projectId"`
	// Policy group's remark information.
	Remark string `pulumi:"remark"`
}

// The set of arguments for constructing a PolicyGroup resource.
type PolicyGroupArgs struct {
	// A list of threshold rules. Each element contains the following attributes:
	Conditions PolicyGroupConditionArrayInput
	// A list of event rules. Each element contains the following attributes:
	EventConditions PolicyGroupEventConditionArrayInput
	// Policy group name, length should between 1 and 20.
	GroupName pulumi.StringInput
	// The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
	// alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
	IsUnionRule pulumi.IntPtrInput
	// Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
	// `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
	PolicyViewName pulumi.StringInput
	// The project id to which the policy group belongs, default is `0`.
	ProjectId pulumi.IntPtrInput
	// Policy group's remark information.
	Remark pulumi.StringInput
}

func (PolicyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyGroupArgs)(nil)).Elem()
}

type PolicyGroupInput interface {
	pulumi.Input

	ToPolicyGroupOutput() PolicyGroupOutput
	ToPolicyGroupOutputWithContext(ctx context.Context) PolicyGroupOutput
}

func (*PolicyGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyGroup)(nil)).Elem()
}

func (i *PolicyGroup) ToPolicyGroupOutput() PolicyGroupOutput {
	return i.ToPolicyGroupOutputWithContext(context.Background())
}

func (i *PolicyGroup) ToPolicyGroupOutputWithContext(ctx context.Context) PolicyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyGroupOutput)
}

// PolicyGroupArrayInput is an input type that accepts PolicyGroupArray and PolicyGroupArrayOutput values.
// You can construct a concrete instance of `PolicyGroupArrayInput` via:
//
//          PolicyGroupArray{ PolicyGroupArgs{...} }
type PolicyGroupArrayInput interface {
	pulumi.Input

	ToPolicyGroupArrayOutput() PolicyGroupArrayOutput
	ToPolicyGroupArrayOutputWithContext(context.Context) PolicyGroupArrayOutput
}

type PolicyGroupArray []PolicyGroupInput

func (PolicyGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyGroup)(nil)).Elem()
}

func (i PolicyGroupArray) ToPolicyGroupArrayOutput() PolicyGroupArrayOutput {
	return i.ToPolicyGroupArrayOutputWithContext(context.Background())
}

func (i PolicyGroupArray) ToPolicyGroupArrayOutputWithContext(ctx context.Context) PolicyGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyGroupArrayOutput)
}

// PolicyGroupMapInput is an input type that accepts PolicyGroupMap and PolicyGroupMapOutput values.
// You can construct a concrete instance of `PolicyGroupMapInput` via:
//
//          PolicyGroupMap{ "key": PolicyGroupArgs{...} }
type PolicyGroupMapInput interface {
	pulumi.Input

	ToPolicyGroupMapOutput() PolicyGroupMapOutput
	ToPolicyGroupMapOutputWithContext(context.Context) PolicyGroupMapOutput
}

type PolicyGroupMap map[string]PolicyGroupInput

func (PolicyGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyGroup)(nil)).Elem()
}

func (i PolicyGroupMap) ToPolicyGroupMapOutput() PolicyGroupMapOutput {
	return i.ToPolicyGroupMapOutputWithContext(context.Background())
}

func (i PolicyGroupMap) ToPolicyGroupMapOutputWithContext(ctx context.Context) PolicyGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyGroupMapOutput)
}

type PolicyGroupOutput struct{ *pulumi.OutputState }

func (PolicyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyGroup)(nil)).Elem()
}

func (o PolicyGroupOutput) ToPolicyGroupOutput() PolicyGroupOutput {
	return o
}

func (o PolicyGroupOutput) ToPolicyGroupOutputWithContext(ctx context.Context) PolicyGroupOutput {
	return o
}

// A list binding objects(list only those in the `provider.region`). Each element contains the following attributes:
func (o PolicyGroupOutput) BindingObjects() PolicyGroupBindingObjectArrayOutput {
	return o.ApplyT(func(v *PolicyGroup) PolicyGroupBindingObjectArrayOutput { return v.BindingObjects }).(PolicyGroupBindingObjectArrayOutput)
}

// A list of threshold rules. Each element contains the following attributes:
func (o PolicyGroupOutput) Conditions() PolicyGroupConditionArrayOutput {
	return o.ApplyT(func(v *PolicyGroup) PolicyGroupConditionArrayOutput { return v.Conditions }).(PolicyGroupConditionArrayOutput)
}

// A list of dimensions for this policy group.
func (o PolicyGroupOutput) DimensionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.StringArrayOutput { return v.DimensionGroups }).(pulumi.StringArrayOutput)
}

// A list of event rules. Each element contains the following attributes:
func (o PolicyGroupOutput) EventConditions() PolicyGroupEventConditionArrayOutput {
	return o.ApplyT(func(v *PolicyGroup) PolicyGroupEventConditionArrayOutput { return v.EventConditions }).(PolicyGroupEventConditionArrayOutput)
}

// Policy group name, length should between 1 and 20.
func (o PolicyGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
// alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
func (o PolicyGroupOutput) IsUnionRule() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.IntPtrOutput { return v.IsUnionRule }).(pulumi.IntPtrOutput)
}

// Recently edited user uin.
func (o PolicyGroupOutput) LastEditUin() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.StringOutput { return v.LastEditUin }).(pulumi.StringOutput)
}

// Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
// `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
func (o PolicyGroupOutput) PolicyViewName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.StringOutput { return v.PolicyViewName }).(pulumi.StringOutput)
}

// The project id to which the policy group belongs, default is `0`.
func (o PolicyGroupOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// A list of receivers. Each element contains the following attributes:
func (o PolicyGroupOutput) Receivers() PolicyGroupReceiverArrayOutput {
	return o.ApplyT(func(v *PolicyGroup) PolicyGroupReceiverArrayOutput { return v.Receivers }).(PolicyGroupReceiverArrayOutput)
}

// Policy group's remark information.
func (o PolicyGroupOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.StringOutput { return v.Remark }).(pulumi.StringOutput)
}

// Support regions this policy group.
func (o PolicyGroupOutput) SupportRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.StringArrayOutput { return v.SupportRegions }).(pulumi.StringArrayOutput)
}

// The policy group update time.
func (o PolicyGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type PolicyGroupArrayOutput struct{ *pulumi.OutputState }

func (PolicyGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyGroup)(nil)).Elem()
}

func (o PolicyGroupArrayOutput) ToPolicyGroupArrayOutput() PolicyGroupArrayOutput {
	return o
}

func (o PolicyGroupArrayOutput) ToPolicyGroupArrayOutputWithContext(ctx context.Context) PolicyGroupArrayOutput {
	return o
}

func (o PolicyGroupArrayOutput) Index(i pulumi.IntInput) PolicyGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyGroup {
		return vs[0].([]*PolicyGroup)[vs[1].(int)]
	}).(PolicyGroupOutput)
}

type PolicyGroupMapOutput struct{ *pulumi.OutputState }

func (PolicyGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyGroup)(nil)).Elem()
}

func (o PolicyGroupMapOutput) ToPolicyGroupMapOutput() PolicyGroupMapOutput {
	return o
}

func (o PolicyGroupMapOutput) ToPolicyGroupMapOutputWithContext(ctx context.Context) PolicyGroupMapOutput {
	return o
}

func (o PolicyGroupMapOutput) MapIndex(k pulumi.StringInput) PolicyGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyGroup {
		return vs[0].(map[string]*PolicyGroup)[vs[1].(string)]
	}).(PolicyGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyGroupInput)(nil)).Elem(), &PolicyGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyGroupArrayInput)(nil)).Elem(), PolicyGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyGroupMapInput)(nil)).Elem(), PolicyGroupMap{})
	pulumi.RegisterOutputType(PolicyGroupOutput{})
	pulumi.RegisterOutputType(PolicyGroupArrayOutput{})
	pulumi.RegisterOutputType(PolicyGroupMapOutput{})
}
