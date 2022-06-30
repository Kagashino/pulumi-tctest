// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceRoleAttachment struct {
	pulumi.CustomResourceState

	// The id of tdmq cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Creation time of resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The name of tdmq namespace.
	EnvironId pulumi.StringOutput `pulumi:"environId"`
	// The permissions of tdmq role.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// The name of tdmq role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
}

// NewNamespaceRoleAttachment registers a new resource with the given unique name, arguments, and options.
func NewNamespaceRoleAttachment(ctx *pulumi.Context,
	name string, args *NamespaceRoleAttachmentArgs, opts ...pulumi.ResourceOption) (*NamespaceRoleAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.EnvironId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironId'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	var resource NamespaceRoleAttachment
	err := ctx.RegisterResource("tencentcloud:Tdmq/namespaceRoleAttachment:NamespaceRoleAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceRoleAttachment gets an existing NamespaceRoleAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceRoleAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceRoleAttachmentState, opts ...pulumi.ResourceOption) (*NamespaceRoleAttachment, error) {
	var resource NamespaceRoleAttachment
	err := ctx.ReadResource("tencentcloud:Tdmq/namespaceRoleAttachment:NamespaceRoleAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceRoleAttachment resources.
type namespaceRoleAttachmentState struct {
	// The id of tdmq cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Creation time of resource.
	CreateTime *string `pulumi:"createTime"`
	// The name of tdmq namespace.
	EnvironId *string `pulumi:"environId"`
	// The permissions of tdmq role.
	Permissions []string `pulumi:"permissions"`
	// The name of tdmq role.
	RoleName *string `pulumi:"roleName"`
}

type NamespaceRoleAttachmentState struct {
	// The id of tdmq cluster.
	ClusterId pulumi.StringPtrInput
	// Creation time of resource.
	CreateTime pulumi.StringPtrInput
	// The name of tdmq namespace.
	EnvironId pulumi.StringPtrInput
	// The permissions of tdmq role.
	Permissions pulumi.StringArrayInput
	// The name of tdmq role.
	RoleName pulumi.StringPtrInput
}

func (NamespaceRoleAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceRoleAttachmentState)(nil)).Elem()
}

type namespaceRoleAttachmentArgs struct {
	// The id of tdmq cluster.
	ClusterId string `pulumi:"clusterId"`
	// The name of tdmq namespace.
	EnvironId string `pulumi:"environId"`
	// The permissions of tdmq role.
	Permissions []string `pulumi:"permissions"`
	// The name of tdmq role.
	RoleName string `pulumi:"roleName"`
}

// The set of arguments for constructing a NamespaceRoleAttachment resource.
type NamespaceRoleAttachmentArgs struct {
	// The id of tdmq cluster.
	ClusterId pulumi.StringInput
	// The name of tdmq namespace.
	EnvironId pulumi.StringInput
	// The permissions of tdmq role.
	Permissions pulumi.StringArrayInput
	// The name of tdmq role.
	RoleName pulumi.StringInput
}

func (NamespaceRoleAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceRoleAttachmentArgs)(nil)).Elem()
}

type NamespaceRoleAttachmentInput interface {
	pulumi.Input

	ToNamespaceRoleAttachmentOutput() NamespaceRoleAttachmentOutput
	ToNamespaceRoleAttachmentOutputWithContext(ctx context.Context) NamespaceRoleAttachmentOutput
}

func (*NamespaceRoleAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceRoleAttachment)(nil)).Elem()
}

func (i *NamespaceRoleAttachment) ToNamespaceRoleAttachmentOutput() NamespaceRoleAttachmentOutput {
	return i.ToNamespaceRoleAttachmentOutputWithContext(context.Background())
}

func (i *NamespaceRoleAttachment) ToNamespaceRoleAttachmentOutputWithContext(ctx context.Context) NamespaceRoleAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceRoleAttachmentOutput)
}

// NamespaceRoleAttachmentArrayInput is an input type that accepts NamespaceRoleAttachmentArray and NamespaceRoleAttachmentArrayOutput values.
// You can construct a concrete instance of `NamespaceRoleAttachmentArrayInput` via:
//
//          NamespaceRoleAttachmentArray{ NamespaceRoleAttachmentArgs{...} }
type NamespaceRoleAttachmentArrayInput interface {
	pulumi.Input

	ToNamespaceRoleAttachmentArrayOutput() NamespaceRoleAttachmentArrayOutput
	ToNamespaceRoleAttachmentArrayOutputWithContext(context.Context) NamespaceRoleAttachmentArrayOutput
}

type NamespaceRoleAttachmentArray []NamespaceRoleAttachmentInput

func (NamespaceRoleAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceRoleAttachment)(nil)).Elem()
}

func (i NamespaceRoleAttachmentArray) ToNamespaceRoleAttachmentArrayOutput() NamespaceRoleAttachmentArrayOutput {
	return i.ToNamespaceRoleAttachmentArrayOutputWithContext(context.Background())
}

func (i NamespaceRoleAttachmentArray) ToNamespaceRoleAttachmentArrayOutputWithContext(ctx context.Context) NamespaceRoleAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceRoleAttachmentArrayOutput)
}

// NamespaceRoleAttachmentMapInput is an input type that accepts NamespaceRoleAttachmentMap and NamespaceRoleAttachmentMapOutput values.
// You can construct a concrete instance of `NamespaceRoleAttachmentMapInput` via:
//
//          NamespaceRoleAttachmentMap{ "key": NamespaceRoleAttachmentArgs{...} }
type NamespaceRoleAttachmentMapInput interface {
	pulumi.Input

	ToNamespaceRoleAttachmentMapOutput() NamespaceRoleAttachmentMapOutput
	ToNamespaceRoleAttachmentMapOutputWithContext(context.Context) NamespaceRoleAttachmentMapOutput
}

type NamespaceRoleAttachmentMap map[string]NamespaceRoleAttachmentInput

func (NamespaceRoleAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceRoleAttachment)(nil)).Elem()
}

func (i NamespaceRoleAttachmentMap) ToNamespaceRoleAttachmentMapOutput() NamespaceRoleAttachmentMapOutput {
	return i.ToNamespaceRoleAttachmentMapOutputWithContext(context.Background())
}

func (i NamespaceRoleAttachmentMap) ToNamespaceRoleAttachmentMapOutputWithContext(ctx context.Context) NamespaceRoleAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceRoleAttachmentMapOutput)
}

type NamespaceRoleAttachmentOutput struct{ *pulumi.OutputState }

func (NamespaceRoleAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceRoleAttachment)(nil)).Elem()
}

func (o NamespaceRoleAttachmentOutput) ToNamespaceRoleAttachmentOutput() NamespaceRoleAttachmentOutput {
	return o
}

func (o NamespaceRoleAttachmentOutput) ToNamespaceRoleAttachmentOutputWithContext(ctx context.Context) NamespaceRoleAttachmentOutput {
	return o
}

// The id of tdmq cluster.
func (o NamespaceRoleAttachmentOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceRoleAttachment) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation time of resource.
func (o NamespaceRoleAttachmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceRoleAttachment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The name of tdmq namespace.
func (o NamespaceRoleAttachmentOutput) EnvironId() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceRoleAttachment) pulumi.StringOutput { return v.EnvironId }).(pulumi.StringOutput)
}

// The permissions of tdmq role.
func (o NamespaceRoleAttachmentOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceRoleAttachment) pulumi.StringArrayOutput { return v.Permissions }).(pulumi.StringArrayOutput)
}

// The name of tdmq role.
func (o NamespaceRoleAttachmentOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceRoleAttachment) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

type NamespaceRoleAttachmentArrayOutput struct{ *pulumi.OutputState }

func (NamespaceRoleAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceRoleAttachment)(nil)).Elem()
}

func (o NamespaceRoleAttachmentArrayOutput) ToNamespaceRoleAttachmentArrayOutput() NamespaceRoleAttachmentArrayOutput {
	return o
}

func (o NamespaceRoleAttachmentArrayOutput) ToNamespaceRoleAttachmentArrayOutputWithContext(ctx context.Context) NamespaceRoleAttachmentArrayOutput {
	return o
}

func (o NamespaceRoleAttachmentArrayOutput) Index(i pulumi.IntInput) NamespaceRoleAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NamespaceRoleAttachment {
		return vs[0].([]*NamespaceRoleAttachment)[vs[1].(int)]
	}).(NamespaceRoleAttachmentOutput)
}

type NamespaceRoleAttachmentMapOutput struct{ *pulumi.OutputState }

func (NamespaceRoleAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceRoleAttachment)(nil)).Elem()
}

func (o NamespaceRoleAttachmentMapOutput) ToNamespaceRoleAttachmentMapOutput() NamespaceRoleAttachmentMapOutput {
	return o
}

func (o NamespaceRoleAttachmentMapOutput) ToNamespaceRoleAttachmentMapOutputWithContext(ctx context.Context) NamespaceRoleAttachmentMapOutput {
	return o
}

func (o NamespaceRoleAttachmentMapOutput) MapIndex(k pulumi.StringInput) NamespaceRoleAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NamespaceRoleAttachment {
		return vs[0].(map[string]*NamespaceRoleAttachment)[vs[1].(string)]
	}).(NamespaceRoleAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceRoleAttachmentInput)(nil)).Elem(), &NamespaceRoleAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceRoleAttachmentArrayInput)(nil)).Elem(), NamespaceRoleAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceRoleAttachmentMapInput)(nil)).Elem(), NamespaceRoleAttachmentMap{})
	pulumi.RegisterOutputType(NamespaceRoleAttachmentOutput{})
	pulumi.RegisterOutputType(NamespaceRoleAttachmentArrayOutput{})
	pulumi.RegisterOutputType(NamespaceRoleAttachmentMapOutput{})
}
