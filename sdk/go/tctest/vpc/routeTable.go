// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouteTable struct {
	pulumi.CustomResourceState

	// Creation time of the routing table.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicates whether it is the default routing table.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The name of routing table.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID list of the routing entries.
	RouteEntryIds pulumi.StringArrayOutput `pulumi:"routeEntryIds"`
	// ID list of the subnets associated with this route table.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// The tags of routing table.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// ID of VPC to which the route table should be associated.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource RouteTable
	err := ctx.RegisterResource("tctest:Vpc/routeTable:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("tctest:Vpc/routeTable:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
	// Creation time of the routing table.
	CreateTime *string `pulumi:"createTime"`
	// Indicates whether it is the default routing table.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of routing table.
	Name *string `pulumi:"name"`
	// ID list of the routing entries.
	RouteEntryIds []string `pulumi:"routeEntryIds"`
	// ID list of the subnets associated with this route table.
	SubnetIds []string `pulumi:"subnetIds"`
	// The tags of routing table.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of VPC to which the route table should be associated.
	VpcId *string `pulumi:"vpcId"`
}

type RouteTableState struct {
	// Creation time of the routing table.
	CreateTime pulumi.StringPtrInput
	// Indicates whether it is the default routing table.
	IsDefault pulumi.BoolPtrInput
	// The name of routing table.
	Name pulumi.StringPtrInput
	// ID list of the routing entries.
	RouteEntryIds pulumi.StringArrayInput
	// ID list of the subnets associated with this route table.
	SubnetIds pulumi.StringArrayInput
	// The tags of routing table.
	Tags pulumi.MapInput
	// ID of VPC to which the route table should be associated.
	VpcId pulumi.StringPtrInput
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	// The name of routing table.
	Name *string `pulumi:"name"`
	// The tags of routing table.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of VPC to which the route table should be associated.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// The name of routing table.
	Name pulumi.StringPtrInput
	// The tags of routing table.
	Tags pulumi.MapInput
	// ID of VPC to which the route table should be associated.
	VpcId pulumi.StringInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

// RouteTableArrayInput is an input type that accepts RouteTableArray and RouteTableArrayOutput values.
// You can construct a concrete instance of `RouteTableArrayInput` via:
//
//          RouteTableArray{ RouteTableArgs{...} }
type RouteTableArrayInput interface {
	pulumi.Input

	ToRouteTableArrayOutput() RouteTableArrayOutput
	ToRouteTableArrayOutputWithContext(context.Context) RouteTableArrayOutput
}

type RouteTableArray []RouteTableInput

func (RouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (i RouteTableArray) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return i.ToRouteTableArrayOutputWithContext(context.Background())
}

func (i RouteTableArray) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableArrayOutput)
}

// RouteTableMapInput is an input type that accepts RouteTableMap and RouteTableMapOutput values.
// You can construct a concrete instance of `RouteTableMapInput` via:
//
//          RouteTableMap{ "key": RouteTableArgs{...} }
type RouteTableMapInput interface {
	pulumi.Input

	ToRouteTableMapOutput() RouteTableMapOutput
	ToRouteTableMapOutputWithContext(context.Context) RouteTableMapOutput
}

type RouteTableMap map[string]RouteTableInput

func (RouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (i RouteTableMap) ToRouteTableMapOutput() RouteTableMapOutput {
	return i.ToRouteTableMapOutputWithContext(context.Background())
}

func (i RouteTableMap) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableMapOutput)
}

type RouteTableOutput struct{ *pulumi.OutputState }

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

// Creation time of the routing table.
func (o RouteTableOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether it is the default routing table.
func (o RouteTableOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The name of routing table.
func (o RouteTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID list of the routing entries.
func (o RouteTableOutput) RouteEntryIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringArrayOutput { return v.RouteEntryIds }).(pulumi.StringArrayOutput)
}

// ID list of the subnets associated with this route table.
func (o RouteTableOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The tags of routing table.
func (o RouteTableOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// ID of VPC to which the route table should be associated.
func (o RouteTableOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type RouteTableArrayOutput struct{ *pulumi.OutputState }

func (RouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) Index(i pulumi.IntInput) RouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].([]*RouteTable)[vs[1].(int)]
	}).(RouteTableOutput)
}

type RouteTableMapOutput struct{ *pulumi.OutputState }

func (RouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (o RouteTableMapOutput) ToRouteTableMapOutput() RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) MapIndex(k pulumi.StringInput) RouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].(map[string]*RouteTable)[vs[1].(string)]
	}).(RouteTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableInput)(nil)).Elem(), &RouteTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableArrayInput)(nil)).Elem(), RouteTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableMapInput)(nil)).Elem(), RouteTableMap{})
	pulumi.RegisterOutputType(RouteTableOutput{})
	pulumi.RegisterOutputType(RouteTableArrayOutput{})
	pulumi.RegisterOutputType(RouteTableMapOutput{})
}
