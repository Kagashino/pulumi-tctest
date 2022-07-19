// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstancesInstanceList struct {
	ChargeType       string                          `pulumi:"chargeType"`
	CreateTime       string                          `pulumi:"createTime"`
	Ip               string                          `pulumi:"ip"`
	MemSize          int                             `pulumi:"memSize"`
	Name             string                          `pulumi:"name"`
	NodeInfos        []InstancesInstanceListNodeInfo `pulumi:"nodeInfos"`
	Port             int                             `pulumi:"port"`
	ProjectId        int                             `pulumi:"projectId"`
	RedisId          string                          `pulumi:"redisId"`
	RedisReplicasNum int                             `pulumi:"redisReplicasNum"`
	RedisShardNum    int                             `pulumi:"redisShardNum"`
	Status           string                          `pulumi:"status"`
	SubnetId         string                          `pulumi:"subnetId"`
	Tags             map[string]interface{}          `pulumi:"tags"`
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type   string `pulumi:"type"`
	TypeId int    `pulumi:"typeId"`
	VpcId  string `pulumi:"vpcId"`
	Zone   string `pulumi:"zone"`
}

// InstancesInstanceListInput is an input type that accepts InstancesInstanceListArgs and InstancesInstanceListOutput values.
// You can construct a concrete instance of `InstancesInstanceListInput` via:
//
//          InstancesInstanceListArgs{...}
type InstancesInstanceListInput interface {
	pulumi.Input

	ToInstancesInstanceListOutput() InstancesInstanceListOutput
	ToInstancesInstanceListOutputWithContext(context.Context) InstancesInstanceListOutput
}

type InstancesInstanceListArgs struct {
	ChargeType       pulumi.StringInput                      `pulumi:"chargeType"`
	CreateTime       pulumi.StringInput                      `pulumi:"createTime"`
	Ip               pulumi.StringInput                      `pulumi:"ip"`
	MemSize          pulumi.IntInput                         `pulumi:"memSize"`
	Name             pulumi.StringInput                      `pulumi:"name"`
	NodeInfos        InstancesInstanceListNodeInfoArrayInput `pulumi:"nodeInfos"`
	Port             pulumi.IntInput                         `pulumi:"port"`
	ProjectId        pulumi.IntInput                         `pulumi:"projectId"`
	RedisId          pulumi.StringInput                      `pulumi:"redisId"`
	RedisReplicasNum pulumi.IntInput                         `pulumi:"redisReplicasNum"`
	RedisShardNum    pulumi.IntInput                         `pulumi:"redisShardNum"`
	Status           pulumi.StringInput                      `pulumi:"status"`
	SubnetId         pulumi.StringInput                      `pulumi:"subnetId"`
	Tags             pulumi.MapInput                         `pulumi:"tags"`
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type   pulumi.StringInput `pulumi:"type"`
	TypeId pulumi.IntInput    `pulumi:"typeId"`
	VpcId  pulumi.StringInput `pulumi:"vpcId"`
	Zone   pulumi.StringInput `pulumi:"zone"`
}

func (InstancesInstanceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceList)(nil)).Elem()
}

func (i InstancesInstanceListArgs) ToInstancesInstanceListOutput() InstancesInstanceListOutput {
	return i.ToInstancesInstanceListOutputWithContext(context.Background())
}

func (i InstancesInstanceListArgs) ToInstancesInstanceListOutputWithContext(ctx context.Context) InstancesInstanceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListOutput)
}

// InstancesInstanceListArrayInput is an input type that accepts InstancesInstanceListArray and InstancesInstanceListArrayOutput values.
// You can construct a concrete instance of `InstancesInstanceListArrayInput` via:
//
//          InstancesInstanceListArray{ InstancesInstanceListArgs{...} }
type InstancesInstanceListArrayInput interface {
	pulumi.Input

	ToInstancesInstanceListArrayOutput() InstancesInstanceListArrayOutput
	ToInstancesInstanceListArrayOutputWithContext(context.Context) InstancesInstanceListArrayOutput
}

type InstancesInstanceListArray []InstancesInstanceListInput

func (InstancesInstanceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceList)(nil)).Elem()
}

func (i InstancesInstanceListArray) ToInstancesInstanceListArrayOutput() InstancesInstanceListArrayOutput {
	return i.ToInstancesInstanceListArrayOutputWithContext(context.Background())
}

func (i InstancesInstanceListArray) ToInstancesInstanceListArrayOutputWithContext(ctx context.Context) InstancesInstanceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListArrayOutput)
}

type InstancesInstanceListOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceList)(nil)).Elem()
}

func (o InstancesInstanceListOutput) ToInstancesInstanceListOutput() InstancesInstanceListOutput {
	return o
}

func (o InstancesInstanceListOutput) ToInstancesInstanceListOutputWithContext(ctx context.Context) InstancesInstanceListOutput {
	return o
}

func (o InstancesInstanceListOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.ChargeType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Ip }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) MemSize() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.MemSize }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Name }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) NodeInfos() InstancesInstanceListNodeInfoArrayOutput {
	return o.ApplyT(func(v InstancesInstanceList) []InstancesInstanceListNodeInfo { return v.NodeInfos }).(InstancesInstanceListNodeInfoArrayOutput)
}

func (o InstancesInstanceListOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.Port }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) RedisId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.RedisId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) RedisReplicasNum() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.RedisReplicasNum }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) RedisShardNum() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.RedisShardNum }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Status }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v InstancesInstanceList) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
func (o InstancesInstanceListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Type }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) TypeId() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.TypeId }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.VpcId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Zone }).(pulumi.StringOutput)
}

type InstancesInstanceListArrayOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceList)(nil)).Elem()
}

func (o InstancesInstanceListArrayOutput) ToInstancesInstanceListArrayOutput() InstancesInstanceListArrayOutput {
	return o
}

func (o InstancesInstanceListArrayOutput) ToInstancesInstanceListArrayOutputWithContext(ctx context.Context) InstancesInstanceListArrayOutput {
	return o
}

func (o InstancesInstanceListArrayOutput) Index(i pulumi.IntInput) InstancesInstanceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesInstanceList {
		return vs[0].([]InstancesInstanceList)[vs[1].(int)]
	}).(InstancesInstanceListOutput)
}

type InstancesInstanceListNodeInfo struct {
	Id     int  `pulumi:"id"`
	Master bool `pulumi:"master"`
	ZoneId int  `pulumi:"zoneId"`
}

// InstancesInstanceListNodeInfoInput is an input type that accepts InstancesInstanceListNodeInfoArgs and InstancesInstanceListNodeInfoOutput values.
// You can construct a concrete instance of `InstancesInstanceListNodeInfoInput` via:
//
//          InstancesInstanceListNodeInfoArgs{...}
type InstancesInstanceListNodeInfoInput interface {
	pulumi.Input

	ToInstancesInstanceListNodeInfoOutput() InstancesInstanceListNodeInfoOutput
	ToInstancesInstanceListNodeInfoOutputWithContext(context.Context) InstancesInstanceListNodeInfoOutput
}

type InstancesInstanceListNodeInfoArgs struct {
	Id     pulumi.IntInput  `pulumi:"id"`
	Master pulumi.BoolInput `pulumi:"master"`
	ZoneId pulumi.IntInput  `pulumi:"zoneId"`
}

func (InstancesInstanceListNodeInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListNodeInfo)(nil)).Elem()
}

func (i InstancesInstanceListNodeInfoArgs) ToInstancesInstanceListNodeInfoOutput() InstancesInstanceListNodeInfoOutput {
	return i.ToInstancesInstanceListNodeInfoOutputWithContext(context.Background())
}

func (i InstancesInstanceListNodeInfoArgs) ToInstancesInstanceListNodeInfoOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListNodeInfoOutput)
}

// InstancesInstanceListNodeInfoArrayInput is an input type that accepts InstancesInstanceListNodeInfoArray and InstancesInstanceListNodeInfoArrayOutput values.
// You can construct a concrete instance of `InstancesInstanceListNodeInfoArrayInput` via:
//
//          InstancesInstanceListNodeInfoArray{ InstancesInstanceListNodeInfoArgs{...} }
type InstancesInstanceListNodeInfoArrayInput interface {
	pulumi.Input

	ToInstancesInstanceListNodeInfoArrayOutput() InstancesInstanceListNodeInfoArrayOutput
	ToInstancesInstanceListNodeInfoArrayOutputWithContext(context.Context) InstancesInstanceListNodeInfoArrayOutput
}

type InstancesInstanceListNodeInfoArray []InstancesInstanceListNodeInfoInput

func (InstancesInstanceListNodeInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListNodeInfo)(nil)).Elem()
}

func (i InstancesInstanceListNodeInfoArray) ToInstancesInstanceListNodeInfoArrayOutput() InstancesInstanceListNodeInfoArrayOutput {
	return i.ToInstancesInstanceListNodeInfoArrayOutputWithContext(context.Background())
}

func (i InstancesInstanceListNodeInfoArray) ToInstancesInstanceListNodeInfoArrayOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListNodeInfoArrayOutput)
}

type InstancesInstanceListNodeInfoOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListNodeInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListNodeInfo)(nil)).Elem()
}

func (o InstancesInstanceListNodeInfoOutput) ToInstancesInstanceListNodeInfoOutput() InstancesInstanceListNodeInfoOutput {
	return o
}

func (o InstancesInstanceListNodeInfoOutput) ToInstancesInstanceListNodeInfoOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoOutput {
	return o
}

func (o InstancesInstanceListNodeInfoOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfo) int { return v.Id }).(pulumi.IntOutput)
}

func (o InstancesInstanceListNodeInfoOutput) Master() pulumi.BoolOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfo) bool { return v.Master }).(pulumi.BoolOutput)
}

func (o InstancesInstanceListNodeInfoOutput) ZoneId() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfo) int { return v.ZoneId }).(pulumi.IntOutput)
}

type InstancesInstanceListNodeInfoArrayOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListNodeInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListNodeInfo)(nil)).Elem()
}

func (o InstancesInstanceListNodeInfoArrayOutput) ToInstancesInstanceListNodeInfoArrayOutput() InstancesInstanceListNodeInfoArrayOutput {
	return o
}

func (o InstancesInstanceListNodeInfoArrayOutput) ToInstancesInstanceListNodeInfoArrayOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoArrayOutput {
	return o
}

func (o InstancesInstanceListNodeInfoArrayOutput) Index(i pulumi.IntInput) InstancesInstanceListNodeInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesInstanceListNodeInfo {
		return vs[0].([]InstancesInstanceListNodeInfo)[vs[1].(int)]
	}).(InstancesInstanceListNodeInfoOutput)
}

type ZoneConfigList struct {
	// Deprecated: It has been deprecated from version 1.26.0. Use `shard_memories` instead.
	MemSizes          []int `pulumi:"memSizes"`
	RedisReplicasNums []int `pulumi:"redisReplicasNums"`
	RedisShardNums    []int `pulumi:"redisShardNums"`
	ShardMemories     []int `pulumi:"shardMemories"`
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type    string `pulumi:"type"`
	TypeId  int    `pulumi:"typeId"`
	Version string `pulumi:"version"`
	Zone    string `pulumi:"zone"`
}

// ZoneConfigListInput is an input type that accepts ZoneConfigListArgs and ZoneConfigListOutput values.
// You can construct a concrete instance of `ZoneConfigListInput` via:
//
//          ZoneConfigListArgs{...}
type ZoneConfigListInput interface {
	pulumi.Input

	ToZoneConfigListOutput() ZoneConfigListOutput
	ToZoneConfigListOutputWithContext(context.Context) ZoneConfigListOutput
}

type ZoneConfigListArgs struct {
	// Deprecated: It has been deprecated from version 1.26.0. Use `shard_memories` instead.
	MemSizes          pulumi.IntArrayInput `pulumi:"memSizes"`
	RedisReplicasNums pulumi.IntArrayInput `pulumi:"redisReplicasNums"`
	RedisShardNums    pulumi.IntArrayInput `pulumi:"redisShardNums"`
	ShardMemories     pulumi.IntArrayInput `pulumi:"shardMemories"`
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type    pulumi.StringInput `pulumi:"type"`
	TypeId  pulumi.IntInput    `pulumi:"typeId"`
	Version pulumi.StringInput `pulumi:"version"`
	Zone    pulumi.StringInput `pulumi:"zone"`
}

func (ZoneConfigListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneConfigList)(nil)).Elem()
}

func (i ZoneConfigListArgs) ToZoneConfigListOutput() ZoneConfigListOutput {
	return i.ToZoneConfigListOutputWithContext(context.Background())
}

func (i ZoneConfigListArgs) ToZoneConfigListOutputWithContext(ctx context.Context) ZoneConfigListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneConfigListOutput)
}

// ZoneConfigListArrayInput is an input type that accepts ZoneConfigListArray and ZoneConfigListArrayOutput values.
// You can construct a concrete instance of `ZoneConfigListArrayInput` via:
//
//          ZoneConfigListArray{ ZoneConfigListArgs{...} }
type ZoneConfigListArrayInput interface {
	pulumi.Input

	ToZoneConfigListArrayOutput() ZoneConfigListArrayOutput
	ToZoneConfigListArrayOutputWithContext(context.Context) ZoneConfigListArrayOutput
}

type ZoneConfigListArray []ZoneConfigListInput

func (ZoneConfigListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneConfigList)(nil)).Elem()
}

func (i ZoneConfigListArray) ToZoneConfigListArrayOutput() ZoneConfigListArrayOutput {
	return i.ToZoneConfigListArrayOutputWithContext(context.Background())
}

func (i ZoneConfigListArray) ToZoneConfigListArrayOutputWithContext(ctx context.Context) ZoneConfigListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneConfigListArrayOutput)
}

type ZoneConfigListOutput struct{ *pulumi.OutputState }

func (ZoneConfigListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneConfigList)(nil)).Elem()
}

func (o ZoneConfigListOutput) ToZoneConfigListOutput() ZoneConfigListOutput {
	return o
}

func (o ZoneConfigListOutput) ToZoneConfigListOutputWithContext(ctx context.Context) ZoneConfigListOutput {
	return o
}

// Deprecated: It has been deprecated from version 1.26.0. Use `shard_memories` instead.
func (o ZoneConfigListOutput) MemSizes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ZoneConfigList) []int { return v.MemSizes }).(pulumi.IntArrayOutput)
}

func (o ZoneConfigListOutput) RedisReplicasNums() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ZoneConfigList) []int { return v.RedisReplicasNums }).(pulumi.IntArrayOutput)
}

func (o ZoneConfigListOutput) RedisShardNums() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ZoneConfigList) []int { return v.RedisShardNums }).(pulumi.IntArrayOutput)
}

func (o ZoneConfigListOutput) ShardMemories() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ZoneConfigList) []int { return v.ShardMemories }).(pulumi.IntArrayOutput)
}

// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
func (o ZoneConfigListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneConfigList) string { return v.Type }).(pulumi.StringOutput)
}

func (o ZoneConfigListOutput) TypeId() pulumi.IntOutput {
	return o.ApplyT(func(v ZoneConfigList) int { return v.TypeId }).(pulumi.IntOutput)
}

func (o ZoneConfigListOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneConfigList) string { return v.Version }).(pulumi.StringOutput)
}

func (o ZoneConfigListOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneConfigList) string { return v.Zone }).(pulumi.StringOutput)
}

type ZoneConfigListArrayOutput struct{ *pulumi.OutputState }

func (ZoneConfigListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneConfigList)(nil)).Elem()
}

func (o ZoneConfigListArrayOutput) ToZoneConfigListArrayOutput() ZoneConfigListArrayOutput {
	return o
}

func (o ZoneConfigListArrayOutput) ToZoneConfigListArrayOutputWithContext(ctx context.Context) ZoneConfigListArrayOutput {
	return o
}

func (o ZoneConfigListArrayOutput) Index(i pulumi.IntInput) ZoneConfigListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneConfigList {
		return vs[0].([]ZoneConfigList)[vs[1].(int)]
	}).(ZoneConfigListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListInput)(nil)).Elem(), InstancesInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListArrayInput)(nil)).Elem(), InstancesInstanceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListNodeInfoInput)(nil)).Elem(), InstancesInstanceListNodeInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListNodeInfoArrayInput)(nil)).Elem(), InstancesInstanceListNodeInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneConfigListInput)(nil)).Elem(), ZoneConfigListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneConfigListArrayInput)(nil)).Elem(), ZoneConfigListArray{})
	pulumi.RegisterOutputType(InstancesInstanceListOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListArrayOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListNodeInfoOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListNodeInfoArrayOutput{})
	pulumi.RegisterOutputType(ZoneConfigListOutput{})
	pulumi.RegisterOutputType(ZoneConfigListArrayOutput{})
}
