// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type APIKeyAttachment struct {
	pulumi.CustomResourceState

	// ID of API key.
	ApiKeyId pulumi.StringOutput `pulumi:"apiKeyId"`
	// ID of the usage plan.
	UsagePlanId pulumi.StringOutput `pulumi:"usagePlanId"`
}

// NewAPIKeyAttachment registers a new resource with the given unique name, arguments, and options.
func NewAPIKeyAttachment(ctx *pulumi.Context,
	name string, args *APIKeyAttachmentArgs, opts ...pulumi.ResourceOption) (*APIKeyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKeyId == nil {
		return nil, errors.New("invalid value for required argument 'ApiKeyId'")
	}
	if args.UsagePlanId == nil {
		return nil, errors.New("invalid value for required argument 'UsagePlanId'")
	}
	var resource APIKeyAttachment
	err := ctx.RegisterResource("tctest:APIGateway/aPIKeyAttachment:APIKeyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIKeyAttachment gets an existing APIKeyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIKeyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIKeyAttachmentState, opts ...pulumi.ResourceOption) (*APIKeyAttachment, error) {
	var resource APIKeyAttachment
	err := ctx.ReadResource("tctest:APIGateway/aPIKeyAttachment:APIKeyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIKeyAttachment resources.
type apikeyAttachmentState struct {
	// ID of API key.
	ApiKeyId *string `pulumi:"apiKeyId"`
	// ID of the usage plan.
	UsagePlanId *string `pulumi:"usagePlanId"`
}

type APIKeyAttachmentState struct {
	// ID of API key.
	ApiKeyId pulumi.StringPtrInput
	// ID of the usage plan.
	UsagePlanId pulumi.StringPtrInput
}

func (APIKeyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeyAttachmentState)(nil)).Elem()
}

type apikeyAttachmentArgs struct {
	// ID of API key.
	ApiKeyId string `pulumi:"apiKeyId"`
	// ID of the usage plan.
	UsagePlanId string `pulumi:"usagePlanId"`
}

// The set of arguments for constructing a APIKeyAttachment resource.
type APIKeyAttachmentArgs struct {
	// ID of API key.
	ApiKeyId pulumi.StringInput
	// ID of the usage plan.
	UsagePlanId pulumi.StringInput
}

func (APIKeyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeyAttachmentArgs)(nil)).Elem()
}

type APIKeyAttachmentInput interface {
	pulumi.Input

	ToAPIKeyAttachmentOutput() APIKeyAttachmentOutput
	ToAPIKeyAttachmentOutputWithContext(ctx context.Context) APIKeyAttachmentOutput
}

func (*APIKeyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**APIKeyAttachment)(nil)).Elem()
}

func (i *APIKeyAttachment) ToAPIKeyAttachmentOutput() APIKeyAttachmentOutput {
	return i.ToAPIKeyAttachmentOutputWithContext(context.Background())
}

func (i *APIKeyAttachment) ToAPIKeyAttachmentOutputWithContext(ctx context.Context) APIKeyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyAttachmentOutput)
}

// APIKeyAttachmentArrayInput is an input type that accepts APIKeyAttachmentArray and APIKeyAttachmentArrayOutput values.
// You can construct a concrete instance of `APIKeyAttachmentArrayInput` via:
//
//          APIKeyAttachmentArray{ APIKeyAttachmentArgs{...} }
type APIKeyAttachmentArrayInput interface {
	pulumi.Input

	ToAPIKeyAttachmentArrayOutput() APIKeyAttachmentArrayOutput
	ToAPIKeyAttachmentArrayOutputWithContext(context.Context) APIKeyAttachmentArrayOutput
}

type APIKeyAttachmentArray []APIKeyAttachmentInput

func (APIKeyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIKeyAttachment)(nil)).Elem()
}

func (i APIKeyAttachmentArray) ToAPIKeyAttachmentArrayOutput() APIKeyAttachmentArrayOutput {
	return i.ToAPIKeyAttachmentArrayOutputWithContext(context.Background())
}

func (i APIKeyAttachmentArray) ToAPIKeyAttachmentArrayOutputWithContext(ctx context.Context) APIKeyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyAttachmentArrayOutput)
}

// APIKeyAttachmentMapInput is an input type that accepts APIKeyAttachmentMap and APIKeyAttachmentMapOutput values.
// You can construct a concrete instance of `APIKeyAttachmentMapInput` via:
//
//          APIKeyAttachmentMap{ "key": APIKeyAttachmentArgs{...} }
type APIKeyAttachmentMapInput interface {
	pulumi.Input

	ToAPIKeyAttachmentMapOutput() APIKeyAttachmentMapOutput
	ToAPIKeyAttachmentMapOutputWithContext(context.Context) APIKeyAttachmentMapOutput
}

type APIKeyAttachmentMap map[string]APIKeyAttachmentInput

func (APIKeyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIKeyAttachment)(nil)).Elem()
}

func (i APIKeyAttachmentMap) ToAPIKeyAttachmentMapOutput() APIKeyAttachmentMapOutput {
	return i.ToAPIKeyAttachmentMapOutputWithContext(context.Background())
}

func (i APIKeyAttachmentMap) ToAPIKeyAttachmentMapOutputWithContext(ctx context.Context) APIKeyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyAttachmentMapOutput)
}

type APIKeyAttachmentOutput struct{ *pulumi.OutputState }

func (APIKeyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIKeyAttachment)(nil)).Elem()
}

func (o APIKeyAttachmentOutput) ToAPIKeyAttachmentOutput() APIKeyAttachmentOutput {
	return o
}

func (o APIKeyAttachmentOutput) ToAPIKeyAttachmentOutputWithContext(ctx context.Context) APIKeyAttachmentOutput {
	return o
}

// ID of API key.
func (o APIKeyAttachmentOutput) ApiKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKeyAttachment) pulumi.StringOutput { return v.ApiKeyId }).(pulumi.StringOutput)
}

// ID of the usage plan.
func (o APIKeyAttachmentOutput) UsagePlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKeyAttachment) pulumi.StringOutput { return v.UsagePlanId }).(pulumi.StringOutput)
}

type APIKeyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (APIKeyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIKeyAttachment)(nil)).Elem()
}

func (o APIKeyAttachmentArrayOutput) ToAPIKeyAttachmentArrayOutput() APIKeyAttachmentArrayOutput {
	return o
}

func (o APIKeyAttachmentArrayOutput) ToAPIKeyAttachmentArrayOutputWithContext(ctx context.Context) APIKeyAttachmentArrayOutput {
	return o
}

func (o APIKeyAttachmentArrayOutput) Index(i pulumi.IntInput) APIKeyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *APIKeyAttachment {
		return vs[0].([]*APIKeyAttachment)[vs[1].(int)]
	}).(APIKeyAttachmentOutput)
}

type APIKeyAttachmentMapOutput struct{ *pulumi.OutputState }

func (APIKeyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIKeyAttachment)(nil)).Elem()
}

func (o APIKeyAttachmentMapOutput) ToAPIKeyAttachmentMapOutput() APIKeyAttachmentMapOutput {
	return o
}

func (o APIKeyAttachmentMapOutput) ToAPIKeyAttachmentMapOutputWithContext(ctx context.Context) APIKeyAttachmentMapOutput {
	return o
}

func (o APIKeyAttachmentMapOutput) MapIndex(k pulumi.StringInput) APIKeyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *APIKeyAttachment {
		return vs[0].(map[string]*APIKeyAttachment)[vs[1].(string)]
	}).(APIKeyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*APIKeyAttachmentInput)(nil)).Elem(), &APIKeyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIKeyAttachmentArrayInput)(nil)).Elem(), APIKeyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIKeyAttachmentMapInput)(nil)).Elem(), APIKeyAttachmentMap{})
	pulumi.RegisterOutputType(APIKeyAttachmentOutput{})
	pulumi.RegisterOutputType(APIKeyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(APIKeyAttachmentMapOutput{})
}
