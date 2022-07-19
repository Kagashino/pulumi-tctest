// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayCcnRouteInstance struct {
	pulumi.CustomResourceState

	// As path list of the BGP.
	AsPaths pulumi.StringArrayOutput `pulumi:"asPaths"`
	// A network address segment of IDC.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// ID of the DCG.
	DcgId pulumi.StringOutput `pulumi:"dcgId"`
}

// NewGatewayCcnRouteInstance registers a new resource with the given unique name, arguments, and options.
func NewGatewayCcnRouteInstance(ctx *pulumi.Context,
	name string, args *GatewayCcnRouteInstanceArgs, opts ...pulumi.ResourceOption) (*GatewayCcnRouteInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.DcgId == nil {
		return nil, errors.New("invalid value for required argument 'DcgId'")
	}
	var resource GatewayCcnRouteInstance
	err := ctx.RegisterResource("tctest:Dc/gatewayCcnRouteInstance:GatewayCcnRouteInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayCcnRouteInstance gets an existing GatewayCcnRouteInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayCcnRouteInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayCcnRouteInstanceState, opts ...pulumi.ResourceOption) (*GatewayCcnRouteInstance, error) {
	var resource GatewayCcnRouteInstance
	err := ctx.ReadResource("tctest:Dc/gatewayCcnRouteInstance:GatewayCcnRouteInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayCcnRouteInstance resources.
type gatewayCcnRouteInstanceState struct {
	// As path list of the BGP.
	AsPaths []string `pulumi:"asPaths"`
	// A network address segment of IDC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// ID of the DCG.
	DcgId *string `pulumi:"dcgId"`
}

type GatewayCcnRouteInstanceState struct {
	// As path list of the BGP.
	AsPaths pulumi.StringArrayInput
	// A network address segment of IDC.
	CidrBlock pulumi.StringPtrInput
	// ID of the DCG.
	DcgId pulumi.StringPtrInput
}

func (GatewayCcnRouteInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCcnRouteInstanceState)(nil)).Elem()
}

type gatewayCcnRouteInstanceArgs struct {
	// A network address segment of IDC.
	CidrBlock string `pulumi:"cidrBlock"`
	// ID of the DCG.
	DcgId string `pulumi:"dcgId"`
}

// The set of arguments for constructing a GatewayCcnRouteInstance resource.
type GatewayCcnRouteInstanceArgs struct {
	// A network address segment of IDC.
	CidrBlock pulumi.StringInput
	// ID of the DCG.
	DcgId pulumi.StringInput
}

func (GatewayCcnRouteInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCcnRouteInstanceArgs)(nil)).Elem()
}

type GatewayCcnRouteInstanceInput interface {
	pulumi.Input

	ToGatewayCcnRouteInstanceOutput() GatewayCcnRouteInstanceOutput
	ToGatewayCcnRouteInstanceOutputWithContext(ctx context.Context) GatewayCcnRouteInstanceOutput
}

func (*GatewayCcnRouteInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCcnRouteInstance)(nil)).Elem()
}

func (i *GatewayCcnRouteInstance) ToGatewayCcnRouteInstanceOutput() GatewayCcnRouteInstanceOutput {
	return i.ToGatewayCcnRouteInstanceOutputWithContext(context.Background())
}

func (i *GatewayCcnRouteInstance) ToGatewayCcnRouteInstanceOutputWithContext(ctx context.Context) GatewayCcnRouteInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCcnRouteInstanceOutput)
}

// GatewayCcnRouteInstanceArrayInput is an input type that accepts GatewayCcnRouteInstanceArray and GatewayCcnRouteInstanceArrayOutput values.
// You can construct a concrete instance of `GatewayCcnRouteInstanceArrayInput` via:
//
//          GatewayCcnRouteInstanceArray{ GatewayCcnRouteInstanceArgs{...} }
type GatewayCcnRouteInstanceArrayInput interface {
	pulumi.Input

	ToGatewayCcnRouteInstanceArrayOutput() GatewayCcnRouteInstanceArrayOutput
	ToGatewayCcnRouteInstanceArrayOutputWithContext(context.Context) GatewayCcnRouteInstanceArrayOutput
}

type GatewayCcnRouteInstanceArray []GatewayCcnRouteInstanceInput

func (GatewayCcnRouteInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayCcnRouteInstance)(nil)).Elem()
}

func (i GatewayCcnRouteInstanceArray) ToGatewayCcnRouteInstanceArrayOutput() GatewayCcnRouteInstanceArrayOutput {
	return i.ToGatewayCcnRouteInstanceArrayOutputWithContext(context.Background())
}

func (i GatewayCcnRouteInstanceArray) ToGatewayCcnRouteInstanceArrayOutputWithContext(ctx context.Context) GatewayCcnRouteInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCcnRouteInstanceArrayOutput)
}

// GatewayCcnRouteInstanceMapInput is an input type that accepts GatewayCcnRouteInstanceMap and GatewayCcnRouteInstanceMapOutput values.
// You can construct a concrete instance of `GatewayCcnRouteInstanceMapInput` via:
//
//          GatewayCcnRouteInstanceMap{ "key": GatewayCcnRouteInstanceArgs{...} }
type GatewayCcnRouteInstanceMapInput interface {
	pulumi.Input

	ToGatewayCcnRouteInstanceMapOutput() GatewayCcnRouteInstanceMapOutput
	ToGatewayCcnRouteInstanceMapOutputWithContext(context.Context) GatewayCcnRouteInstanceMapOutput
}

type GatewayCcnRouteInstanceMap map[string]GatewayCcnRouteInstanceInput

func (GatewayCcnRouteInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayCcnRouteInstance)(nil)).Elem()
}

func (i GatewayCcnRouteInstanceMap) ToGatewayCcnRouteInstanceMapOutput() GatewayCcnRouteInstanceMapOutput {
	return i.ToGatewayCcnRouteInstanceMapOutputWithContext(context.Background())
}

func (i GatewayCcnRouteInstanceMap) ToGatewayCcnRouteInstanceMapOutputWithContext(ctx context.Context) GatewayCcnRouteInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCcnRouteInstanceMapOutput)
}

type GatewayCcnRouteInstanceOutput struct{ *pulumi.OutputState }

func (GatewayCcnRouteInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCcnRouteInstance)(nil)).Elem()
}

func (o GatewayCcnRouteInstanceOutput) ToGatewayCcnRouteInstanceOutput() GatewayCcnRouteInstanceOutput {
	return o
}

func (o GatewayCcnRouteInstanceOutput) ToGatewayCcnRouteInstanceOutputWithContext(ctx context.Context) GatewayCcnRouteInstanceOutput {
	return o
}

// As path list of the BGP.
func (o GatewayCcnRouteInstanceOutput) AsPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayCcnRouteInstance) pulumi.StringArrayOutput { return v.AsPaths }).(pulumi.StringArrayOutput)
}

// A network address segment of IDC.
func (o GatewayCcnRouteInstanceOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCcnRouteInstance) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// ID of the DCG.
func (o GatewayCcnRouteInstanceOutput) DcgId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCcnRouteInstance) pulumi.StringOutput { return v.DcgId }).(pulumi.StringOutput)
}

type GatewayCcnRouteInstanceArrayOutput struct{ *pulumi.OutputState }

func (GatewayCcnRouteInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayCcnRouteInstance)(nil)).Elem()
}

func (o GatewayCcnRouteInstanceArrayOutput) ToGatewayCcnRouteInstanceArrayOutput() GatewayCcnRouteInstanceArrayOutput {
	return o
}

func (o GatewayCcnRouteInstanceArrayOutput) ToGatewayCcnRouteInstanceArrayOutputWithContext(ctx context.Context) GatewayCcnRouteInstanceArrayOutput {
	return o
}

func (o GatewayCcnRouteInstanceArrayOutput) Index(i pulumi.IntInput) GatewayCcnRouteInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayCcnRouteInstance {
		return vs[0].([]*GatewayCcnRouteInstance)[vs[1].(int)]
	}).(GatewayCcnRouteInstanceOutput)
}

type GatewayCcnRouteInstanceMapOutput struct{ *pulumi.OutputState }

func (GatewayCcnRouteInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayCcnRouteInstance)(nil)).Elem()
}

func (o GatewayCcnRouteInstanceMapOutput) ToGatewayCcnRouteInstanceMapOutput() GatewayCcnRouteInstanceMapOutput {
	return o
}

func (o GatewayCcnRouteInstanceMapOutput) ToGatewayCcnRouteInstanceMapOutputWithContext(ctx context.Context) GatewayCcnRouteInstanceMapOutput {
	return o
}

func (o GatewayCcnRouteInstanceMapOutput) MapIndex(k pulumi.StringInput) GatewayCcnRouteInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayCcnRouteInstance {
		return vs[0].(map[string]*GatewayCcnRouteInstance)[vs[1].(string)]
	}).(GatewayCcnRouteInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCcnRouteInstanceInput)(nil)).Elem(), &GatewayCcnRouteInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCcnRouteInstanceArrayInput)(nil)).Elem(), GatewayCcnRouteInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayCcnRouteInstanceMapInput)(nil)).Elem(), GatewayCcnRouteInstanceMap{})
	pulumi.RegisterOutputType(GatewayCcnRouteInstanceOutput{})
	pulumi.RegisterOutputType(GatewayCcnRouteInstanceArrayOutput{})
	pulumi.RegisterOutputType(GatewayCcnRouteInstanceMapOutput{})
}
