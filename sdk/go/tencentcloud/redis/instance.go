// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Instance struct {
	pulumi.CustomResourceState

	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	AutoRenewFlag pulumi.IntPtrOutput `pulumi:"autoRenewFlag"`
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud
	// International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create
	// new with new charge type.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// The time when the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted
	// instead of staying recycle bin. Note: only works for `PREPAID` instance.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// IP address of an instance.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The memory volume of an available instance(in MB), please refer to
	// `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size
	// of the instance; when Redis is cluster type, it represents memory size of per sharding.
	MemSize pulumi.IntOutput `pulumi:"memSize"`
	// Instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	NoAuth pulumi.BoolPtrOutput `pulumi:"noAuth"`
	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password
	// empty.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or
	// the Redis instance will be recreated.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod pulumi.IntPtrOutput `pulumi:"prepaidPeriod"`
	// Specifies which project the instance should belong to.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// The number of instance copies. This is not required for standalone and master slave versions.
	RedisReplicasNum pulumi.IntPtrOutput `pulumi:"redisReplicasNum"`
	// The number of instance shard. This is not required for standalone and master slave versions.
	RedisShardNum pulumi.IntPtrOutput `pulumi:"redisShardNum"`
	// ID of replica nodes available zone. This is not required for standalone and master slave versions.
	ReplicaZoneIds pulumi.IntArrayOutput `pulumi:"replicaZoneIds"`
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica
	// read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed
	// to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the
	// recommended number of replicas >=2.
	ReplicasReadOnly pulumi.BoolPtrOutput `pulumi:"replicasReadOnly"`
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Current status of an instance, maybe: init, processing, online, isolate and todelete.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies which subnet the instance should belong to.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Instance tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Instance type. Available values:
	// `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`,
	// specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	//
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or
	// [document](https://intl.cloud.tencent.com/document/product/239/32069).
	TypeId pulumi.IntPtrOutput `pulumi:"typeId"`
	// ID of the vpc with which the instance is to be associated.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.MemSize == nil {
		return nil, errors.New("invalid value for required argument 'MemSize'")
	}
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Redis/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Redis/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud
	// International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create
	// new with new charge type.
	ChargeType *string `pulumi:"chargeType"`
	// The time when the instance was created.
	CreateTime *string `pulumi:"createTime"`
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted
	// instead of staying recycle bin. Note: only works for `PREPAID` instance.
	ForceDelete *bool `pulumi:"forceDelete"`
	// IP address of an instance.
	Ip *string `pulumi:"ip"`
	// The memory volume of an available instance(in MB), please refer to
	// `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size
	// of the instance; when Redis is cluster type, it represents memory size of per sharding.
	MemSize *int `pulumi:"memSize"`
	// Instance name.
	Name *string `pulumi:"name"`
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	NoAuth *bool `pulumi:"noAuth"`
	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password
	// empty.
	Password *string `pulumi:"password"`
	// The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or
	// the Redis instance will be recreated.
	Port *int `pulumi:"port"`
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// Specifies which project the instance should belong to.
	ProjectId *int `pulumi:"projectId"`
	// The number of instance copies. This is not required for standalone and master slave versions.
	RedisReplicasNum *int `pulumi:"redisReplicasNum"`
	// The number of instance shard. This is not required for standalone and master slave versions.
	RedisShardNum *int `pulumi:"redisShardNum"`
	// ID of replica nodes available zone. This is not required for standalone and master slave versions.
	ReplicaZoneIds []int `pulumi:"replicaZoneIds"`
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica
	// read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed
	// to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the
	// recommended number of replicas >=2.
	ReplicasReadOnly *bool `pulumi:"replicasReadOnly"`
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Current status of an instance, maybe: init, processing, online, isolate and todelete.
	Status *string `pulumi:"status"`
	// Specifies which subnet the instance should belong to.
	SubnetId *string `pulumi:"subnetId"`
	// Instance tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// Instance type. Available values:
	// `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`,
	// specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	//
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type *string `pulumi:"type"`
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or
	// [document](https://intl.cloud.tencent.com/document/product/239/32069).
	TypeId *int `pulumi:"typeId"`
	// ID of the vpc with which the instance is to be associated.
	VpcId *string `pulumi:"vpcId"`
}

type InstanceState struct {
	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	AutoRenewFlag pulumi.IntPtrInput
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	AvailabilityZone pulumi.StringPtrInput
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud
	// International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create
	// new with new charge type.
	ChargeType pulumi.StringPtrInput
	// The time when the instance was created.
	CreateTime pulumi.StringPtrInput
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted
	// instead of staying recycle bin. Note: only works for `PREPAID` instance.
	ForceDelete pulumi.BoolPtrInput
	// IP address of an instance.
	Ip pulumi.StringPtrInput
	// The memory volume of an available instance(in MB), please refer to
	// `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size
	// of the instance; when Redis is cluster type, it represents memory size of per sharding.
	MemSize pulumi.IntPtrInput
	// Instance name.
	Name pulumi.StringPtrInput
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	NoAuth pulumi.BoolPtrInput
	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password
	// empty.
	Password pulumi.StringPtrInput
	// The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or
	// the Redis instance will be recreated.
	Port pulumi.IntPtrInput
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod pulumi.IntPtrInput
	// Specifies which project the instance should belong to.
	ProjectId pulumi.IntPtrInput
	// The number of instance copies. This is not required for standalone and master slave versions.
	RedisReplicasNum pulumi.IntPtrInput
	// The number of instance shard. This is not required for standalone and master slave versions.
	RedisShardNum pulumi.IntPtrInput
	// ID of replica nodes available zone. This is not required for standalone and master slave versions.
	ReplicaZoneIds pulumi.IntArrayInput
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica
	// read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed
	// to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the
	// recommended number of replicas >=2.
	ReplicasReadOnly pulumi.BoolPtrInput
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	SecurityGroups pulumi.StringArrayInput
	// Current status of an instance, maybe: init, processing, online, isolate and todelete.
	Status pulumi.StringPtrInput
	// Specifies which subnet the instance should belong to.
	SubnetId pulumi.StringPtrInput
	// Instance tags.
	Tags pulumi.MapInput
	// Instance type. Available values:
	// `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`,
	// specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	//
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type pulumi.StringPtrInput
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or
	// [document](https://intl.cloud.tencent.com/document/product/239/32069).
	TypeId pulumi.IntPtrInput
	// ID of the vpc with which the instance is to be associated.
	VpcId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud
	// International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create
	// new with new charge type.
	ChargeType *string `pulumi:"chargeType"`
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted
	// instead of staying recycle bin. Note: only works for `PREPAID` instance.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The memory volume of an available instance(in MB), please refer to
	// `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size
	// of the instance; when Redis is cluster type, it represents memory size of per sharding.
	MemSize int `pulumi:"memSize"`
	// Instance name.
	Name *string `pulumi:"name"`
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	NoAuth *bool `pulumi:"noAuth"`
	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password
	// empty.
	Password *string `pulumi:"password"`
	// The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or
	// the Redis instance will be recreated.
	Port *int `pulumi:"port"`
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod *int `pulumi:"prepaidPeriod"`
	// Specifies which project the instance should belong to.
	ProjectId *int `pulumi:"projectId"`
	// The number of instance copies. This is not required for standalone and master slave versions.
	RedisReplicasNum *int `pulumi:"redisReplicasNum"`
	// The number of instance shard. This is not required for standalone and master slave versions.
	RedisShardNum *int `pulumi:"redisShardNum"`
	// ID of replica nodes available zone. This is not required for standalone and master slave versions.
	ReplicaZoneIds []int `pulumi:"replicaZoneIds"`
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica
	// read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed
	// to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the
	// recommended number of replicas >=2.
	ReplicasReadOnly *bool `pulumi:"replicasReadOnly"`
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Specifies which subnet the instance should belong to.
	SubnetId *string `pulumi:"subnetId"`
	// Instance tags.
	Tags map[string]interface{} `pulumi:"tags"`
	// Instance type. Available values:
	// `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`,
	// specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	//
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type *string `pulumi:"type"`
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or
	// [document](https://intl.cloud.tencent.com/document/product/239/32069).
	TypeId *int `pulumi:"typeId"`
	// ID of the vpc with which the instance is to be associated.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
	AutoRenewFlag pulumi.IntPtrInput
	// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
	AvailabilityZone pulumi.StringInput
	// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud
	// International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create
	// new with new charge type.
	ChargeType pulumi.StringPtrInput
	// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted
	// instead of staying recycle bin. Note: only works for `PREPAID` instance.
	ForceDelete pulumi.BoolPtrInput
	// The memory volume of an available instance(in MB), please refer to
	// `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size
	// of the instance; when Redis is cluster type, it represents memory size of per sharding.
	MemSize pulumi.IntInput
	// Instance name.
	Name pulumi.StringPtrInput
	// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
	NoAuth pulumi.BoolPtrInput
	// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password
	// empty.
	Password pulumi.StringPtrInput
	// The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or
	// the Redis instance will be recreated.
	Port pulumi.IntPtrInput
	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
	// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	PrepaidPeriod pulumi.IntPtrInput
	// Specifies which project the instance should belong to.
	ProjectId pulumi.IntPtrInput
	// The number of instance copies. This is not required for standalone and master slave versions.
	RedisReplicasNum pulumi.IntPtrInput
	// The number of instance shard. This is not required for standalone and master slave versions.
	RedisShardNum pulumi.IntPtrInput
	// ID of replica nodes available zone. This is not required for standalone and master slave versions.
	ReplicaZoneIds pulumi.IntArrayInput
	// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica
	// read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed
	// to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the
	// recommended number of replicas >=2.
	ReplicasReadOnly pulumi.BoolPtrInput
	// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
	SecurityGroups pulumi.StringArrayInput
	// Specifies which subnet the instance should belong to.
	SubnetId pulumi.StringPtrInput
	// Instance tags.
	Tags pulumi.MapInput
	// Instance type. Available values:
	// `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`,
	// specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
	//
	// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
	Type pulumi.StringPtrInput
	// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or
	// [document](https://intl.cloud.tencent.com/document/product/239/32069).
	TypeId pulumi.IntPtrInput
	// ID of the vpc with which the instance is to be associated.
	VpcId pulumi.StringPtrInput
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

// Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
func (o InstanceOutput) AutoRenewFlag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.AutoRenewFlag }).(pulumi.IntPtrOutput)
}

// The available zone ID of an instance to be created, please refer to `tencentcloud_redis_zone_config.list`.
func (o InstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud
// International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create
// new with new charge type.
func (o InstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// The time when the instance was created.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted
// instead of staying recycle bin. Note: only works for `PREPAID` instance.
func (o InstanceOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

// IP address of an instance.
func (o InstanceOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The memory volume of an available instance(in MB), please refer to
// `tencentcloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size
// of the instance; when Redis is cluster type, it represents memory size of per sharding.
func (o InstanceOutput) MemSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.MemSize }).(pulumi.IntOutput)
}

// Instance name.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
func (o InstanceOutput) NoAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.NoAuth }).(pulumi.BoolPtrOutput)
}

// Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password
// empty.
func (o InstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or
// the Redis instance will be recreated.
func (o InstanceOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
// Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
func (o InstanceOutput) PrepaidPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.PrepaidPeriod }).(pulumi.IntPtrOutput)
}

// Specifies which project the instance should belong to.
func (o InstanceOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// The number of instance copies. This is not required for standalone and master slave versions.
func (o InstanceOutput) RedisReplicasNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.RedisReplicasNum }).(pulumi.IntPtrOutput)
}

// The number of instance shard. This is not required for standalone and master slave versions.
func (o InstanceOutput) RedisShardNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.RedisShardNum }).(pulumi.IntPtrOutput)
}

// ID of replica nodes available zone. This is not required for standalone and master slave versions.
func (o InstanceOutput) ReplicaZoneIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntArrayOutput { return v.ReplicaZoneIds }).(pulumi.IntArrayOutput)
}

// Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica
// read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed
// to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the
// recommended number of replicas >=2.
func (o InstanceOutput) ReplicasReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.ReplicasReadOnly }).(pulumi.BoolPtrOutput)
}

// ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
func (o InstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Current status of an instance, maybe: init, processing, online, isolate and todelete.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies which subnet the instance should belong to.
func (o InstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Instance tags.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Instance type. Available values:
// `cluster_ckv`,`cluster_redis5.0`,`cluster_redis`,`master_slave_ckv`,`master_slave_redis4.0`,`master_slave_redis5.0`,`master_slave_redis`,`standalone_redis`,
// specific region support specific types, need to refer data `tencentcloud_redis_zone_config`.
//
// Deprecated: It has been deprecated from version 1.33.1. Please use 'type_id' instead.
func (o InstanceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Instance type. Available values reference data source `tencentcloud_redis_zone_config` or
// [document](https://intl.cloud.tencent.com/document/product/239/32069).
func (o InstanceOutput) TypeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.TypeId }).(pulumi.IntPtrOutput)
}

// ID of the vpc with which the instance is to be associated.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
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
