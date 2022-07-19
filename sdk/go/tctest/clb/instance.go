// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Instance struct {
	pulumi.CustomResourceState

	// IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
	AddressIpVersion pulumi.StringOutput `pulumi:"addressIpVersion"`
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	BandwidthPackageId pulumi.StringPtrOutput `pulumi:"bandwidthPackageId"`
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	ClbName pulumi.StringOutput `pulumi:"clbName"`
	// The virtual service address table of the CLB.
	ClbVips pulumi.StringArrayOutput `pulumi:"clbVips"`
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	InternetBandwidthMaxOut pulumi.IntOutput `pulumi:"internetBandwidthMaxOut"`
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
	// `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType pulumi.StringOutput `pulumi:"internetChargeType"`
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
	// backend instance security group.
	LoadBalancerPassToTarget pulumi.BoolPtrOutput `pulumi:"loadBalancerPassToTarget"`
	// The id of log set.
	LogSetId pulumi.StringPtrOutput `pulumi:"logSetId"`
	// The id of log topic.
	LogTopicId pulumi.StringPtrOutput `pulumi:"logTopicId"`
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	MasterZoneId pulumi.StringPtrOutput `pulumi:"masterZoneId"`
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// ID of the project within the CLB instance, `0` - Default Project.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
	// traffic when the master is down.
	SlaveZoneId pulumi.StringPtrOutput `pulumi:"slaveZoneId"`
	// Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// The available tags within this CLB.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoRegion pulumi.StringOutput `pulumi:"targetRegionInfoRegion"`
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoVpcId pulumi.StringOutput `pulumi:"targetRegionInfoVpcId"`
	// Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China
	// Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing
	// (BANDWIDTH_PACKAGE).
	VipIsp pulumi.StringOutput `pulumi:"vipIsp"`
	// VPC ID of the CLB.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Available zone id, only applicable to open CLB.
	ZoneId pulumi.StringPtrOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClbName == nil {
		return nil, errors.New("invalid value for required argument 'ClbName'")
	}
	if args.NetworkType == nil {
		return nil, errors.New("invalid value for required argument 'NetworkType'")
	}
	var resource Instance
	err := ctx.RegisterResource("tctest:Clb/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tctest:Clb/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	ClbName *string `pulumi:"clbName"`
	// The virtual service address table of the CLB.
	ClbVips []string `pulumi:"clbVips"`
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	InternetBandwidthMaxOut *int `pulumi:"internetBandwidthMaxOut"`
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
	// `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
	// backend instance security group.
	LoadBalancerPassToTarget *bool `pulumi:"loadBalancerPassToTarget"`
	// The id of log set.
	LogSetId *string `pulumi:"logSetId"`
	// The id of log topic.
	LogTopicId *string `pulumi:"logTopicId"`
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	MasterZoneId *string `pulumi:"masterZoneId"`
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	NetworkType *string `pulumi:"networkType"`
	// ID of the project within the CLB instance, `0` - Default Project.
	ProjectId *int `pulumi:"projectId"`
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
	// traffic when the master is down.
	SlaveZoneId *string `pulumi:"slaveZoneId"`
	// Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
	SubnetId *string `pulumi:"subnetId"`
	// The available tags within this CLB.
	Tags map[string]interface{} `pulumi:"tags"`
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoRegion *string `pulumi:"targetRegionInfoRegion"`
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoVpcId *string `pulumi:"targetRegionInfoVpcId"`
	// Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China
	// Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing
	// (BANDWIDTH_PACKAGE).
	VipIsp *string `pulumi:"vipIsp"`
	// VPC ID of the CLB.
	VpcId *string `pulumi:"vpcId"`
	// Available zone id, only applicable to open CLB.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
	AddressIpVersion pulumi.StringPtrInput
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	BandwidthPackageId pulumi.StringPtrInput
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	ClbName pulumi.StringPtrInput
	// The virtual service address table of the CLB.
	ClbVips pulumi.StringArrayInput
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	InternetBandwidthMaxOut pulumi.IntPtrInput
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
	// `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType pulumi.StringPtrInput
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
	// backend instance security group.
	LoadBalancerPassToTarget pulumi.BoolPtrInput
	// The id of log set.
	LogSetId pulumi.StringPtrInput
	// The id of log topic.
	LogTopicId pulumi.StringPtrInput
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	MasterZoneId pulumi.StringPtrInput
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	NetworkType pulumi.StringPtrInput
	// ID of the project within the CLB instance, `0` - Default Project.
	ProjectId pulumi.IntPtrInput
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	SecurityGroups pulumi.StringArrayInput
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
	// traffic when the master is down.
	SlaveZoneId pulumi.StringPtrInput
	// Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
	SubnetId pulumi.StringPtrInput
	// The available tags within this CLB.
	Tags pulumi.MapInput
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoRegion pulumi.StringPtrInput
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoVpcId pulumi.StringPtrInput
	// Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China
	// Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing
	// (BANDWIDTH_PACKAGE).
	VipIsp pulumi.StringPtrInput
	// VPC ID of the CLB.
	VpcId pulumi.StringPtrInput
	// Available zone id, only applicable to open CLB.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	ClbName string `pulumi:"clbName"`
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	InternetBandwidthMaxOut *int `pulumi:"internetBandwidthMaxOut"`
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
	// `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
	// backend instance security group.
	LoadBalancerPassToTarget *bool `pulumi:"loadBalancerPassToTarget"`
	// The id of log set.
	LogSetId *string `pulumi:"logSetId"`
	// The id of log topic.
	LogTopicId *string `pulumi:"logTopicId"`
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	MasterZoneId *string `pulumi:"masterZoneId"`
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	NetworkType string `pulumi:"networkType"`
	// ID of the project within the CLB instance, `0` - Default Project.
	ProjectId *int `pulumi:"projectId"`
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
	// traffic when the master is down.
	SlaveZoneId *string `pulumi:"slaveZoneId"`
	// Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
	SubnetId *string `pulumi:"subnetId"`
	// The available tags within this CLB.
	Tags map[string]interface{} `pulumi:"tags"`
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoRegion *string `pulumi:"targetRegionInfoRegion"`
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoVpcId *string `pulumi:"targetRegionInfoVpcId"`
	// VPC ID of the CLB.
	VpcId *string `pulumi:"vpcId"`
	// Available zone id, only applicable to open CLB.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
	AddressIpVersion pulumi.StringPtrInput
	// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
	BandwidthPackageId pulumi.StringPtrInput
	// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
	ClbName pulumi.StringInput
	// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
	InternetBandwidthMaxOut pulumi.IntPtrInput
	// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
	// `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType pulumi.StringPtrInput
	// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
	// backend instance security group.
	LoadBalancerPassToTarget pulumi.BoolPtrInput
	// The id of log set.
	LogSetId pulumi.StringPtrInput
	// The id of log topic.
	LogTopicId pulumi.StringPtrInput
	// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
	MasterZoneId pulumi.StringPtrInput
	// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
	NetworkType pulumi.StringInput
	// ID of the project within the CLB instance, `0` - Default Project.
	ProjectId pulumi.IntPtrInput
	// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
	SecurityGroups pulumi.StringArrayInput
	// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
	// traffic when the master is down.
	SlaveZoneId pulumi.StringPtrInput
	// Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
	SubnetId pulumi.StringPtrInput
	// The available tags within this CLB.
	Tags pulumi.MapInput
	// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoRegion pulumi.StringPtrInput
	// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
	TargetRegionInfoVpcId pulumi.StringPtrInput
	// VPC ID of the CLB.
	VpcId pulumi.StringPtrInput
	// Available zone id, only applicable to open CLB.
	ZoneId pulumi.StringPtrInput
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

// IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
func (o InstanceOutput) AddressIpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AddressIpVersion }).(pulumi.StringOutput)
}

// Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
func (o InstanceOutput) BandwidthPackageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.BandwidthPackageId }).(pulumi.StringPtrOutput)
}

// Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
func (o InstanceOutput) ClbName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ClbName }).(pulumi.StringOutput)
}

// The virtual service address table of the CLB.
func (o InstanceOutput) ClbVips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.ClbVips }).(pulumi.StringArrayOutput)
}

// Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
func (o InstanceOutput) InternetBandwidthMaxOut() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.InternetBandwidthMaxOut }).(pulumi.IntOutput)
}

// Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
// `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
func (o InstanceOutput) InternetChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InternetChargeType }).(pulumi.StringOutput)
}

// Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
// backend instance security group.
func (o InstanceOutput) LoadBalancerPassToTarget() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.LoadBalancerPassToTarget }).(pulumi.BoolPtrOutput)
}

// The id of log set.
func (o InstanceOutput) LogSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.LogSetId }).(pulumi.StringPtrOutput)
}

// The id of log topic.
func (o InstanceOutput) LogTopicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.LogTopicId }).(pulumi.StringPtrOutput)
}

// Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
func (o InstanceOutput) MasterZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.MasterZoneId }).(pulumi.StringPtrOutput)
}

// Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
func (o InstanceOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// ID of the project within the CLB instance, `0` - Default Project.
func (o InstanceOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
func (o InstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
// traffic when the master is down.
func (o InstanceOutput) SlaveZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SlaveZoneId }).(pulumi.StringPtrOutput)
}

// Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
func (o InstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The available tags within this CLB.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
func (o InstanceOutput) TargetRegionInfoRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TargetRegionInfoRegion }).(pulumi.StringOutput)
}

// Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
func (o InstanceOutput) TargetRegionInfoVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TargetRegionInfoVpcId }).(pulumi.StringOutput)
}

// Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China
// Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing
// (BANDWIDTH_PACKAGE).
func (o InstanceOutput) VipIsp() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VipIsp }).(pulumi.StringOutput)
}

// VPC ID of the CLB.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Available zone id, only applicable to open CLB.
func (o InstanceOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ZoneId }).(pulumi.StringPtrOutput)
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
