// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SAMLProvider struct {
	pulumi.CustomResourceState

	// The create time of the CAM SAML provider.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the CAM SAML provider.
	Description pulumi.StringOutput `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringOutput `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the CAM SAML provider.
	ProviderArn pulumi.StringOutput `pulumi:"providerArn"`
	// The last update time of the CAM SAML provider.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSAMLProvider registers a new resource with the given unique name, arguments, and options.
func NewSAMLProvider(ctx *pulumi.Context,
	name string, args *SAMLProviderArgs, opts ...pulumi.ResourceOption) (*SAMLProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.MetaData == nil {
		return nil, errors.New("invalid value for required argument 'MetaData'")
	}
	var resource SAMLProvider
	err := ctx.RegisterResource("tencentcloud:Cam/sAMLProvider:SAMLProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSAMLProvider gets an existing SAMLProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSAMLProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAMLProviderState, opts ...pulumi.ResourceOption) (*SAMLProvider, error) {
	var resource SAMLProvider
	err := ctx.ReadResource("tencentcloud:Cam/sAMLProvider:SAMLProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SAMLProvider resources.
type samlproviderState struct {
	// The create time of the CAM SAML provider.
	CreateTime *string `pulumi:"createTime"`
	// The description of the CAM SAML provider.
	Description *string `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData *string `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name *string `pulumi:"name"`
	// The ARN of the CAM SAML provider.
	ProviderArn *string `pulumi:"providerArn"`
	// The last update time of the CAM SAML provider.
	UpdateTime *string `pulumi:"updateTime"`
}

type SAMLProviderState struct {
	// The create time of the CAM SAML provider.
	CreateTime pulumi.StringPtrInput
	// The description of the CAM SAML provider.
	Description pulumi.StringPtrInput
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringPtrInput
	// Name of CAM SAML provider.
	Name pulumi.StringPtrInput
	// The ARN of the CAM SAML provider.
	ProviderArn pulumi.StringPtrInput
	// The last update time of the CAM SAML provider.
	UpdateTime pulumi.StringPtrInput
}

func (SAMLProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlproviderState)(nil)).Elem()
}

type samlproviderArgs struct {
	// The description of the CAM SAML provider.
	Description string `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData string `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SAMLProvider resource.
type SAMLProviderArgs struct {
	// The description of the CAM SAML provider.
	Description pulumi.StringInput
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringInput
	// Name of CAM SAML provider.
	Name pulumi.StringPtrInput
}

func (SAMLProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlproviderArgs)(nil)).Elem()
}

type SAMLProviderInput interface {
	pulumi.Input

	ToSAMLProviderOutput() SAMLProviderOutput
	ToSAMLProviderOutputWithContext(ctx context.Context) SAMLProviderOutput
}

func (*SAMLProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SAMLProvider)(nil)).Elem()
}

func (i *SAMLProvider) ToSAMLProviderOutput() SAMLProviderOutput {
	return i.ToSAMLProviderOutputWithContext(context.Background())
}

func (i *SAMLProvider) ToSAMLProviderOutputWithContext(ctx context.Context) SAMLProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAMLProviderOutput)
}

// SAMLProviderArrayInput is an input type that accepts SAMLProviderArray and SAMLProviderArrayOutput values.
// You can construct a concrete instance of `SAMLProviderArrayInput` via:
//
//          SAMLProviderArray{ SAMLProviderArgs{...} }
type SAMLProviderArrayInput interface {
	pulumi.Input

	ToSAMLProviderArrayOutput() SAMLProviderArrayOutput
	ToSAMLProviderArrayOutputWithContext(context.Context) SAMLProviderArrayOutput
}

type SAMLProviderArray []SAMLProviderInput

func (SAMLProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SAMLProvider)(nil)).Elem()
}

func (i SAMLProviderArray) ToSAMLProviderArrayOutput() SAMLProviderArrayOutput {
	return i.ToSAMLProviderArrayOutputWithContext(context.Background())
}

func (i SAMLProviderArray) ToSAMLProviderArrayOutputWithContext(ctx context.Context) SAMLProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAMLProviderArrayOutput)
}

// SAMLProviderMapInput is an input type that accepts SAMLProviderMap and SAMLProviderMapOutput values.
// You can construct a concrete instance of `SAMLProviderMapInput` via:
//
//          SAMLProviderMap{ "key": SAMLProviderArgs{...} }
type SAMLProviderMapInput interface {
	pulumi.Input

	ToSAMLProviderMapOutput() SAMLProviderMapOutput
	ToSAMLProviderMapOutputWithContext(context.Context) SAMLProviderMapOutput
}

type SAMLProviderMap map[string]SAMLProviderInput

func (SAMLProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SAMLProvider)(nil)).Elem()
}

func (i SAMLProviderMap) ToSAMLProviderMapOutput() SAMLProviderMapOutput {
	return i.ToSAMLProviderMapOutputWithContext(context.Background())
}

func (i SAMLProviderMap) ToSAMLProviderMapOutputWithContext(ctx context.Context) SAMLProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAMLProviderMapOutput)
}

type SAMLProviderOutput struct{ *pulumi.OutputState }

func (SAMLProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAMLProvider)(nil)).Elem()
}

func (o SAMLProviderOutput) ToSAMLProviderOutput() SAMLProviderOutput {
	return o
}

func (o SAMLProviderOutput) ToSAMLProviderOutputWithContext(ctx context.Context) SAMLProviderOutput {
	return o
}

// The create time of the CAM SAML provider.
func (o SAMLProviderOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SAMLProvider) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the CAM SAML provider.
func (o SAMLProviderOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SAMLProvider) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The meta data document of the CAM SAML provider.
func (o SAMLProviderOutput) MetaData() pulumi.StringOutput {
	return o.ApplyT(func(v *SAMLProvider) pulumi.StringOutput { return v.MetaData }).(pulumi.StringOutput)
}

// Name of CAM SAML provider.
func (o SAMLProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAMLProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the CAM SAML provider.
func (o SAMLProviderOutput) ProviderArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SAMLProvider) pulumi.StringOutput { return v.ProviderArn }).(pulumi.StringOutput)
}

// The last update time of the CAM SAML provider.
func (o SAMLProviderOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SAMLProvider) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type SAMLProviderArrayOutput struct{ *pulumi.OutputState }

func (SAMLProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SAMLProvider)(nil)).Elem()
}

func (o SAMLProviderArrayOutput) ToSAMLProviderArrayOutput() SAMLProviderArrayOutput {
	return o
}

func (o SAMLProviderArrayOutput) ToSAMLProviderArrayOutputWithContext(ctx context.Context) SAMLProviderArrayOutput {
	return o
}

func (o SAMLProviderArrayOutput) Index(i pulumi.IntInput) SAMLProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SAMLProvider {
		return vs[0].([]*SAMLProvider)[vs[1].(int)]
	}).(SAMLProviderOutput)
}

type SAMLProviderMapOutput struct{ *pulumi.OutputState }

func (SAMLProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SAMLProvider)(nil)).Elem()
}

func (o SAMLProviderMapOutput) ToSAMLProviderMapOutput() SAMLProviderMapOutput {
	return o
}

func (o SAMLProviderMapOutput) ToSAMLProviderMapOutputWithContext(ctx context.Context) SAMLProviderMapOutput {
	return o
}

func (o SAMLProviderMapOutput) MapIndex(k pulumi.StringInput) SAMLProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SAMLProvider {
		return vs[0].(map[string]*SAMLProvider)[vs[1].(string)]
	}).(SAMLProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SAMLProviderInput)(nil)).Elem(), &SAMLProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*SAMLProviderArrayInput)(nil)).Elem(), SAMLProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SAMLProviderMapInput)(nil)).Elem(), SAMLProviderMap{})
	pulumi.RegisterOutputType(SAMLProviderOutput{})
	pulumi.RegisterOutputType(SAMLProviderArrayOutput{})
	pulumi.RegisterOutputType(SAMLProviderMapOutput{})
}
