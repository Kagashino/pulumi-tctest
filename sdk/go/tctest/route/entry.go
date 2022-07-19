// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Entry struct {
	pulumi.CustomResourceState

	// The RouteEntry's target network segment.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	NextHub pulumi.StringOutput `pulumi:"nextHub"`
	// The next hop type. Valid values:
	// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
	// and `instance`. `instance` points to CVM Instance.
	NextType pulumi.StringOutput `pulumi:"nextType"`
	// The ID of the route table.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewEntry registers a new resource with the given unique name, arguments, and options.
func NewEntry(ctx *pulumi.Context,
	name string, args *EntryArgs, opts ...pulumi.ResourceOption) (*Entry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.NextHub == nil {
		return nil, errors.New("invalid value for required argument 'NextHub'")
	}
	if args.NextType == nil {
		return nil, errors.New("invalid value for required argument 'NextType'")
	}
	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Entry
	err := ctx.RegisterResource("tctest:Route/entry:Entry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntry gets an existing Entry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryState, opts ...pulumi.ResourceOption) (*Entry, error) {
	var resource Entry
	err := ctx.ReadResource("tctest:Route/entry:Entry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Entry resources.
type entryState struct {
	// The RouteEntry's target network segment.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	NextHub *string `pulumi:"nextHub"`
	// The next hop type. Valid values:
	// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
	// and `instance`. `instance` points to CVM Instance.
	NextType *string `pulumi:"nextType"`
	// The ID of the route table.
	RouteTableId *string `pulumi:"routeTableId"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type EntryState struct {
	// The RouteEntry's target network segment.
	CidrBlock pulumi.StringPtrInput
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	NextHub pulumi.StringPtrInput
	// The next hop type. Valid values:
	// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
	// and `instance`. `instance` points to CVM Instance.
	NextType pulumi.StringPtrInput
	// The ID of the route table.
	RouteTableId pulumi.StringPtrInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (EntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryState)(nil)).Elem()
}

type entryArgs struct {
	// The RouteEntry's target network segment.
	CidrBlock string `pulumi:"cidrBlock"`
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	NextHub string `pulumi:"nextHub"`
	// The next hop type. Valid values:
	// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
	// and `instance`. `instance` points to CVM Instance.
	NextType string `pulumi:"nextType"`
	// The ID of the route table.
	RouteTableId string `pulumi:"routeTableId"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Entry resource.
type EntryArgs struct {
	// The RouteEntry's target network segment.
	CidrBlock pulumi.StringInput
	// The route entry's next hub. CVM instance ID or VPC router interface ID.
	NextHub pulumi.StringInput
	// The next hop type. Valid values:
	// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
	// and `instance`. `instance` points to CVM Instance.
	NextType pulumi.StringInput
	// The ID of the route table.
	RouteTableId pulumi.StringInput
	// The VPC ID.
	VpcId pulumi.StringInput
}

func (EntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryArgs)(nil)).Elem()
}

type EntryInput interface {
	pulumi.Input

	ToEntryOutput() EntryOutput
	ToEntryOutputWithContext(ctx context.Context) EntryOutput
}

func (*Entry) ElementType() reflect.Type {
	return reflect.TypeOf((**Entry)(nil)).Elem()
}

func (i *Entry) ToEntryOutput() EntryOutput {
	return i.ToEntryOutputWithContext(context.Background())
}

func (i *Entry) ToEntryOutputWithContext(ctx context.Context) EntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryOutput)
}

// EntryArrayInput is an input type that accepts EntryArray and EntryArrayOutput values.
// You can construct a concrete instance of `EntryArrayInput` via:
//
//          EntryArray{ EntryArgs{...} }
type EntryArrayInput interface {
	pulumi.Input

	ToEntryArrayOutput() EntryArrayOutput
	ToEntryArrayOutputWithContext(context.Context) EntryArrayOutput
}

type EntryArray []EntryInput

func (EntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Entry)(nil)).Elem()
}

func (i EntryArray) ToEntryArrayOutput() EntryArrayOutput {
	return i.ToEntryArrayOutputWithContext(context.Background())
}

func (i EntryArray) ToEntryArrayOutputWithContext(ctx context.Context) EntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryArrayOutput)
}

// EntryMapInput is an input type that accepts EntryMap and EntryMapOutput values.
// You can construct a concrete instance of `EntryMapInput` via:
//
//          EntryMap{ "key": EntryArgs{...} }
type EntryMapInput interface {
	pulumi.Input

	ToEntryMapOutput() EntryMapOutput
	ToEntryMapOutputWithContext(context.Context) EntryMapOutput
}

type EntryMap map[string]EntryInput

func (EntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Entry)(nil)).Elem()
}

func (i EntryMap) ToEntryMapOutput() EntryMapOutput {
	return i.ToEntryMapOutputWithContext(context.Background())
}

func (i EntryMap) ToEntryMapOutputWithContext(ctx context.Context) EntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryMapOutput)
}

type EntryOutput struct{ *pulumi.OutputState }

func (EntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Entry)(nil)).Elem()
}

func (o EntryOutput) ToEntryOutput() EntryOutput {
	return o
}

func (o EntryOutput) ToEntryOutputWithContext(ctx context.Context) EntryOutput {
	return o
}

// The RouteEntry's target network segment.
func (o EntryOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The route entry's next hub. CVM instance ID or VPC router interface ID.
func (o EntryOutput) NextHub() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.NextHub }).(pulumi.StringOutput)
}

// The next hop type. Valid values:
// `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway`
// and `instance`. `instance` points to CVM Instance.
func (o EntryOutput) NextType() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.NextType }).(pulumi.StringOutput)
}

// The ID of the route table.
func (o EntryOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

// The VPC ID.
func (o EntryOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Entry) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type EntryArrayOutput struct{ *pulumi.OutputState }

func (EntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Entry)(nil)).Elem()
}

func (o EntryArrayOutput) ToEntryArrayOutput() EntryArrayOutput {
	return o
}

func (o EntryArrayOutput) ToEntryArrayOutputWithContext(ctx context.Context) EntryArrayOutput {
	return o
}

func (o EntryArrayOutput) Index(i pulumi.IntInput) EntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Entry {
		return vs[0].([]*Entry)[vs[1].(int)]
	}).(EntryOutput)
}

type EntryMapOutput struct{ *pulumi.OutputState }

func (EntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Entry)(nil)).Elem()
}

func (o EntryMapOutput) ToEntryMapOutput() EntryMapOutput {
	return o
}

func (o EntryMapOutput) ToEntryMapOutputWithContext(ctx context.Context) EntryMapOutput {
	return o
}

func (o EntryMapOutput) MapIndex(k pulumi.StringInput) EntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Entry {
		return vs[0].(map[string]*Entry)[vs[1].(string)]
	}).(EntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryInput)(nil)).Elem(), &Entry{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryArrayInput)(nil)).Elem(), EntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryMapInput)(nil)).Elem(), EntryMap{})
	pulumi.RegisterOutputType(EntryOutput{})
	pulumi.RegisterOutputType(EntryArrayOutput{})
	pulumi.RegisterOutputType(EntryMapOutput{})
}
