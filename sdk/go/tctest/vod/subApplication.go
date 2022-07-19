// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubApplication struct {
	pulumi.CustomResourceState

	// The time when the sub application was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Sub application description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and
	// must be unique under a user.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sub appliaction status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSubApplication registers a new resource with the given unique name, arguments, and options.
func NewSubApplication(ctx *pulumi.Context,
	name string, args *SubApplicationArgs, opts ...pulumi.ResourceOption) (*SubApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource SubApplication
	err := ctx.RegisterResource("tctest:Vod/subApplication:SubApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubApplication gets an existing SubApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubApplicationState, opts ...pulumi.ResourceOption) (*SubApplication, error) {
	var resource SubApplication
	err := ctx.ReadResource("tctest:Vod/subApplication:SubApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubApplication resources.
type subApplicationState struct {
	// The time when the sub application was created.
	CreateTime *string `pulumi:"createTime"`
	// Sub application description.
	Description *string `pulumi:"description"`
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and
	// must be unique under a user.
	Name *string `pulumi:"name"`
	// Sub appliaction status.
	Status *string `pulumi:"status"`
}

type SubApplicationState struct {
	// The time when the sub application was created.
	CreateTime pulumi.StringPtrInput
	// Sub application description.
	Description pulumi.StringPtrInput
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and
	// must be unique under a user.
	Name pulumi.StringPtrInput
	// Sub appliaction status.
	Status pulumi.StringPtrInput
}

func (SubApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subApplicationState)(nil)).Elem()
}

type subApplicationArgs struct {
	// Sub application description.
	Description *string `pulumi:"description"`
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and
	// must be unique under a user.
	Name *string `pulumi:"name"`
	// Sub appliaction status.
	Status string `pulumi:"status"`
}

// The set of arguments for constructing a SubApplication resource.
type SubApplicationArgs struct {
	// Sub application description.
	Description pulumi.StringPtrInput
	// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and
	// must be unique under a user.
	Name pulumi.StringPtrInput
	// Sub appliaction status.
	Status pulumi.StringInput
}

func (SubApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subApplicationArgs)(nil)).Elem()
}

type SubApplicationInput interface {
	pulumi.Input

	ToSubApplicationOutput() SubApplicationOutput
	ToSubApplicationOutputWithContext(ctx context.Context) SubApplicationOutput
}

func (*SubApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**SubApplication)(nil)).Elem()
}

func (i *SubApplication) ToSubApplicationOutput() SubApplicationOutput {
	return i.ToSubApplicationOutputWithContext(context.Background())
}

func (i *SubApplication) ToSubApplicationOutputWithContext(ctx context.Context) SubApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubApplicationOutput)
}

// SubApplicationArrayInput is an input type that accepts SubApplicationArray and SubApplicationArrayOutput values.
// You can construct a concrete instance of `SubApplicationArrayInput` via:
//
//          SubApplicationArray{ SubApplicationArgs{...} }
type SubApplicationArrayInput interface {
	pulumi.Input

	ToSubApplicationArrayOutput() SubApplicationArrayOutput
	ToSubApplicationArrayOutputWithContext(context.Context) SubApplicationArrayOutput
}

type SubApplicationArray []SubApplicationInput

func (SubApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubApplication)(nil)).Elem()
}

func (i SubApplicationArray) ToSubApplicationArrayOutput() SubApplicationArrayOutput {
	return i.ToSubApplicationArrayOutputWithContext(context.Background())
}

func (i SubApplicationArray) ToSubApplicationArrayOutputWithContext(ctx context.Context) SubApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubApplicationArrayOutput)
}

// SubApplicationMapInput is an input type that accepts SubApplicationMap and SubApplicationMapOutput values.
// You can construct a concrete instance of `SubApplicationMapInput` via:
//
//          SubApplicationMap{ "key": SubApplicationArgs{...} }
type SubApplicationMapInput interface {
	pulumi.Input

	ToSubApplicationMapOutput() SubApplicationMapOutput
	ToSubApplicationMapOutputWithContext(context.Context) SubApplicationMapOutput
}

type SubApplicationMap map[string]SubApplicationInput

func (SubApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubApplication)(nil)).Elem()
}

func (i SubApplicationMap) ToSubApplicationMapOutput() SubApplicationMapOutput {
	return i.ToSubApplicationMapOutputWithContext(context.Background())
}

func (i SubApplicationMap) ToSubApplicationMapOutputWithContext(ctx context.Context) SubApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubApplicationMapOutput)
}

type SubApplicationOutput struct{ *pulumi.OutputState }

func (SubApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubApplication)(nil)).Elem()
}

func (o SubApplicationOutput) ToSubApplicationOutput() SubApplicationOutput {
	return o
}

func (o SubApplicationOutput) ToSubApplicationOutputWithContext(ctx context.Context) SubApplicationOutput {
	return o
}

// The time when the sub application was created.
func (o SubApplicationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SubApplication) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Sub application description.
func (o SubApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubApplication) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and
// must be unique under a user.
func (o SubApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubApplication) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sub appliaction status.
func (o SubApplicationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SubApplication) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type SubApplicationArrayOutput struct{ *pulumi.OutputState }

func (SubApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubApplication)(nil)).Elem()
}

func (o SubApplicationArrayOutput) ToSubApplicationArrayOutput() SubApplicationArrayOutput {
	return o
}

func (o SubApplicationArrayOutput) ToSubApplicationArrayOutputWithContext(ctx context.Context) SubApplicationArrayOutput {
	return o
}

func (o SubApplicationArrayOutput) Index(i pulumi.IntInput) SubApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubApplication {
		return vs[0].([]*SubApplication)[vs[1].(int)]
	}).(SubApplicationOutput)
}

type SubApplicationMapOutput struct{ *pulumi.OutputState }

func (SubApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubApplication)(nil)).Elem()
}

func (o SubApplicationMapOutput) ToSubApplicationMapOutput() SubApplicationMapOutput {
	return o
}

func (o SubApplicationMapOutput) ToSubApplicationMapOutputWithContext(ctx context.Context) SubApplicationMapOutput {
	return o
}

func (o SubApplicationMapOutput) MapIndex(k pulumi.StringInput) SubApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubApplication {
		return vs[0].(map[string]*SubApplication)[vs[1].(string)]
	}).(SubApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubApplicationInput)(nil)).Elem(), &SubApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubApplicationArrayInput)(nil)).Elem(), SubApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubApplicationMapInput)(nil)).Elem(), SubApplicationMap{})
	pulumi.RegisterOutputType(SubApplicationOutput{})
	pulumi.RegisterOutputType(SubApplicationArrayOutput{})
	pulumi.RegisterOutputType(SubApplicationMapOutput{})
}
