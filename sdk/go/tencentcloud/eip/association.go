// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Association struct {
	pulumi.CustomResourceState

	// The ID of EIP.
	EipId pulumi.StringOutput `pulumi:"eipId"`
	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
	// `private_ip fields`.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
}

// NewAssociation registers a new resource with the given unique name, arguments, and options.
func NewAssociation(ctx *pulumi.Context,
	name string, args *AssociationArgs, opts ...pulumi.ResourceOption) (*Association, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EipId == nil {
		return nil, errors.New("invalid value for required argument 'EipId'")
	}
	var resource Association
	err := ctx.RegisterResource("tencentcloud:Eip/association:Association", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssociation gets an existing Association resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssociationState, opts ...pulumi.ResourceOption) (*Association, error) {
	var resource Association
	err := ctx.ReadResource("tencentcloud:Eip/association:Association", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Association resources.
type associationState struct {
	// The ID of EIP.
	EipId *string `pulumi:"eipId"`
	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
	// `private_ip fields`.
	InstanceId *string `pulumi:"instanceId"`
	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	PrivateIp *string `pulumi:"privateIp"`
}

type AssociationState struct {
	// The ID of EIP.
	EipId pulumi.StringPtrInput
	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
	// `private_ip fields`.
	InstanceId pulumi.StringPtrInput
	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	NetworkInterfaceId pulumi.StringPtrInput
	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	PrivateIp pulumi.StringPtrInput
}

func (AssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*associationState)(nil)).Elem()
}

type associationArgs struct {
	// The ID of EIP.
	EipId string `pulumi:"eipId"`
	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
	// `private_ip fields`.
	InstanceId *string `pulumi:"instanceId"`
	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	PrivateIp *string `pulumi:"privateIp"`
}

// The set of arguments for constructing a Association resource.
type AssociationArgs struct {
	// The ID of EIP.
	EipId pulumi.StringInput
	// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
	// `private_ip fields`.
	InstanceId pulumi.StringPtrInput
	// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
	NetworkInterfaceId pulumi.StringPtrInput
	// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
	PrivateIp pulumi.StringPtrInput
}

func (AssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*associationArgs)(nil)).Elem()
}

type AssociationInput interface {
	pulumi.Input

	ToAssociationOutput() AssociationOutput
	ToAssociationOutputWithContext(ctx context.Context) AssociationOutput
}

func (*Association) ElementType() reflect.Type {
	return reflect.TypeOf((**Association)(nil)).Elem()
}

func (i *Association) ToAssociationOutput() AssociationOutput {
	return i.ToAssociationOutputWithContext(context.Background())
}

func (i *Association) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationOutput)
}

// AssociationArrayInput is an input type that accepts AssociationArray and AssociationArrayOutput values.
// You can construct a concrete instance of `AssociationArrayInput` via:
//
//          AssociationArray{ AssociationArgs{...} }
type AssociationArrayInput interface {
	pulumi.Input

	ToAssociationArrayOutput() AssociationArrayOutput
	ToAssociationArrayOutputWithContext(context.Context) AssociationArrayOutput
}

type AssociationArray []AssociationInput

func (AssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Association)(nil)).Elem()
}

func (i AssociationArray) ToAssociationArrayOutput() AssociationArrayOutput {
	return i.ToAssociationArrayOutputWithContext(context.Background())
}

func (i AssociationArray) ToAssociationArrayOutputWithContext(ctx context.Context) AssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationArrayOutput)
}

// AssociationMapInput is an input type that accepts AssociationMap and AssociationMapOutput values.
// You can construct a concrete instance of `AssociationMapInput` via:
//
//          AssociationMap{ "key": AssociationArgs{...} }
type AssociationMapInput interface {
	pulumi.Input

	ToAssociationMapOutput() AssociationMapOutput
	ToAssociationMapOutputWithContext(context.Context) AssociationMapOutput
}

type AssociationMap map[string]AssociationInput

func (AssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Association)(nil)).Elem()
}

func (i AssociationMap) ToAssociationMapOutput() AssociationMapOutput {
	return i.ToAssociationMapOutputWithContext(context.Background())
}

func (i AssociationMap) ToAssociationMapOutputWithContext(ctx context.Context) AssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationMapOutput)
}

type AssociationOutput struct{ *pulumi.OutputState }

func (AssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Association)(nil)).Elem()
}

func (o AssociationOutput) ToAssociationOutput() AssociationOutput {
	return o
}

func (o AssociationOutput) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return o
}

// The ID of EIP.
func (o AssociationOutput) EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *Association) pulumi.StringOutput { return v.EipId }).(pulumi.StringOutput)
}

// The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
// `private_ip fields`.
func (o AssociationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Association) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
func (o AssociationOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Association) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
func (o AssociationOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Association) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

type AssociationArrayOutput struct{ *pulumi.OutputState }

func (AssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Association)(nil)).Elem()
}

func (o AssociationArrayOutput) ToAssociationArrayOutput() AssociationArrayOutput {
	return o
}

func (o AssociationArrayOutput) ToAssociationArrayOutputWithContext(ctx context.Context) AssociationArrayOutput {
	return o
}

func (o AssociationArrayOutput) Index(i pulumi.IntInput) AssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Association {
		return vs[0].([]*Association)[vs[1].(int)]
	}).(AssociationOutput)
}

type AssociationMapOutput struct{ *pulumi.OutputState }

func (AssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Association)(nil)).Elem()
}

func (o AssociationMapOutput) ToAssociationMapOutput() AssociationMapOutput {
	return o
}

func (o AssociationMapOutput) ToAssociationMapOutputWithContext(ctx context.Context) AssociationMapOutput {
	return o
}

func (o AssociationMapOutput) MapIndex(k pulumi.StringInput) AssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Association {
		return vs[0].(map[string]*Association)[vs[1].(string)]
	}).(AssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssociationInput)(nil)).Elem(), &Association{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssociationArrayInput)(nil)).Elem(), AssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssociationMapInput)(nil)).Elem(), AssociationMap{})
	pulumi.RegisterOutputType(AssociationOutput{})
	pulumi.RegisterOutputType(AssociationArrayOutput{})
	pulumi.RegisterOutputType(AssociationMapOutput{})
}
