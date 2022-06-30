// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package image

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Instance struct {
	pulumi.CustomResourceState

	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
	// image.
	DataDiskIds pulumi.StringArrayOutput `pulumi:"dataDiskIds"`
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
	// will be made after shutdown.
	ForcePoweroff pulumi.BoolPtrOutput `pulumi:"forcePoweroff"`
	// Image Description.
	ImageDescription pulumi.StringPtrOutput `pulumi:"imageDescription"`
	// Image name.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// Cloud server instance ID.
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
	// passed in simultaneously with InstanceId.
	SnapshotIds pulumi.StringArrayOutput `pulumi:"snapshotIds"`
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
	// Syspre function.
	Sysprep pulumi.BoolPtrOutput `pulumi:"sysprep"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Image/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Image/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
	// image.
	DataDiskIds []string `pulumi:"dataDiskIds"`
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
	// will be made after shutdown.
	ForcePoweroff *bool `pulumi:"forcePoweroff"`
	// Image Description.
	ImageDescription *string `pulumi:"imageDescription"`
	// Image name.
	ImageName *string `pulumi:"imageName"`
	// Cloud server instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
	// passed in simultaneously with InstanceId.
	SnapshotIds []string `pulumi:"snapshotIds"`
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
	// Syspre function.
	Sysprep *bool `pulumi:"sysprep"`
}

type InstanceState struct {
	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
	// image.
	DataDiskIds pulumi.StringArrayInput
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
	// will be made after shutdown.
	ForcePoweroff pulumi.BoolPtrInput
	// Image Description.
	ImageDescription pulumi.StringPtrInput
	// Image name.
	ImageName pulumi.StringPtrInput
	// Cloud server instance ID.
	InstanceId pulumi.StringPtrInput
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
	// passed in simultaneously with InstanceId.
	SnapshotIds pulumi.StringArrayInput
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
	// Syspre function.
	Sysprep pulumi.BoolPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
	// image.
	DataDiskIds []string `pulumi:"dataDiskIds"`
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
	// will be made after shutdown.
	ForcePoweroff *bool `pulumi:"forcePoweroff"`
	// Image Description.
	ImageDescription *string `pulumi:"imageDescription"`
	// Image name.
	ImageName string `pulumi:"imageName"`
	// Cloud server instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
	// passed in simultaneously with InstanceId.
	SnapshotIds []string `pulumi:"snapshotIds"`
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
	// Syspre function.
	Sysprep *bool `pulumi:"sysprep"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
	// image.
	DataDiskIds pulumi.StringArrayInput
	// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
	// will be made after shutdown.
	ForcePoweroff pulumi.BoolPtrInput
	// Image Description.
	ImageDescription pulumi.StringPtrInput
	// Image name.
	ImageName pulumi.StringInput
	// Cloud server instance ID.
	InstanceId pulumi.StringPtrInput
	// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
	// passed in simultaneously with InstanceId.
	SnapshotIds pulumi.StringArrayInput
	// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
	// Syspre function.
	Sysprep pulumi.BoolPtrInput
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

// Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the
// image.
func (o InstanceOutput) DataDiskIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.DataDiskIds }).(pulumi.StringArrayOutput)
}

// Set whether to force shutdown during mirroring. The default value is `false`, when set to true, it means that the mirror
// will be made after shutdown.
func (o InstanceOutput) ForcePoweroff() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.ForcePoweroff }).(pulumi.BoolPtrOutput)
}

// Image Description.
func (o InstanceOutput) ImageDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ImageDescription }).(pulumi.StringPtrOutput)
}

// Image name.
func (o InstanceOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

// Cloud server instance ID.
func (o InstanceOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be
// passed in simultaneously with InstanceId.
func (o InstanceOutput) SnapshotIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SnapshotIds }).(pulumi.StringArrayOutput)
}

// Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the
// Syspre function.
func (o InstanceOutput) Sysprep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.Sysprep }).(pulumi.BoolPtrOutput)
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
