# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[str],
                 charge_type: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 prepaid_period: Optional[pulumi.Input[int]] = None,
                 prepaid_renew_flag: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] availability_zone: The available zone for the CDH instance.
        :param pulumi.Input[str] charge_type: The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
        :param pulumi.Input[str] host_name: The name of the CDH instance. The max length of host_name is 60.
        :param pulumi.Input[str] host_type: The type of the CDH instance.
        :param pulumi.Input[int] prepaid_period: The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
               Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        :param pulumi.Input[str] prepaid_renew_flag: Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
               `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
               neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
               specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
               balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
        :param pulumi.Input[int] project_id: The project the instance belongs to, default to 0.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if host_name is not None:
            pulumi.set(__self__, "host_name", host_name)
        if host_type is not None:
            pulumi.set(__self__, "host_type", host_type)
        if prepaid_period is not None:
            pulumi.set(__self__, "prepaid_period", prepaid_period)
        if prepaid_renew_flag is not None:
            pulumi.set(__self__, "prepaid_renew_flag", prepaid_renew_flag)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        """
        The available zone for the CDH instance.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the CDH instance. The max length of host_name is 60.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the CDH instance.
        """
        return pulumi.get(self, "host_type")

    @host_type.setter
    def host_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_type", value)

    @property
    @pulumi.getter(name="prepaidPeriod")
    def prepaid_period(self) -> Optional[pulumi.Input[int]]:
        """
        The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
        Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        """
        return pulumi.get(self, "prepaid_period")

    @prepaid_period.setter
    def prepaid_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prepaid_period", value)

    @property
    @pulumi.getter(name="prepaidRenewFlag")
    def prepaid_renew_flag(self) -> Optional[pulumi.Input[str]]:
        """
        Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
        `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
        neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
        specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
        balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
        """
        return pulumi.get(self, "prepaid_renew_flag")

    @prepaid_renew_flag.setter
    def prepaid_renew_flag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepaid_renew_flag", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        The project the instance belongs to, default to 0.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 cvm_instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 expired_time: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 host_resources: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceHostResourceArgs']]]] = None,
                 host_state: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 prepaid_period: Optional[pulumi.Input[int]] = None,
                 prepaid_renew_flag: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[str] availability_zone: The available zone for the CDH instance.
        :param pulumi.Input[str] charge_type: The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
        :param pulumi.Input[str] create_time: Create time of the instance.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cvm_instance_ids: Id of CVM instances that have been created on the CDH instance.
        :param pulumi.Input[str] expired_time: Expired time of the instance.
        :param pulumi.Input[str] host_name: The name of the CDH instance. The max length of host_name is 60.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceHostResourceArgs']]] host_resources: An information list of host resource. Each element contains the following attributes:
        :param pulumi.Input[str] host_state: State of the CDH instance.
        :param pulumi.Input[str] host_type: The type of the CDH instance.
        :param pulumi.Input[int] prepaid_period: The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
               Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        :param pulumi.Input[str] prepaid_renew_flag: Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
               `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
               neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
               specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
               balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
        :param pulumi.Input[int] project_id: The project the instance belongs to, default to 0.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if cvm_instance_ids is not None:
            pulumi.set(__self__, "cvm_instance_ids", cvm_instance_ids)
        if expired_time is not None:
            pulumi.set(__self__, "expired_time", expired_time)
        if host_name is not None:
            pulumi.set(__self__, "host_name", host_name)
        if host_resources is not None:
            pulumi.set(__self__, "host_resources", host_resources)
        if host_state is not None:
            pulumi.set(__self__, "host_state", host_state)
        if host_type is not None:
            pulumi.set(__self__, "host_type", host_type)
        if prepaid_period is not None:
            pulumi.set(__self__, "prepaid_period", prepaid_period)
        if prepaid_renew_flag is not None:
            pulumi.set(__self__, "prepaid_renew_flag", prepaid_renew_flag)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The available zone for the CDH instance.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the instance.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="cvmInstanceIds")
    def cvm_instance_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Id of CVM instances that have been created on the CDH instance.
        """
        return pulumi.get(self, "cvm_instance_ids")

    @cvm_instance_ids.setter
    def cvm_instance_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cvm_instance_ids", value)

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> Optional[pulumi.Input[str]]:
        """
        Expired time of the instance.
        """
        return pulumi.get(self, "expired_time")

    @expired_time.setter
    def expired_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expired_time", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the CDH instance. The max length of host_name is 60.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="hostResources")
    def host_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceHostResourceArgs']]]]:
        """
        An information list of host resource. Each element contains the following attributes:
        """
        return pulumi.get(self, "host_resources")

    @host_resources.setter
    def host_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceHostResourceArgs']]]]):
        pulumi.set(self, "host_resources", value)

    @property
    @pulumi.getter(name="hostState")
    def host_state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the CDH instance.
        """
        return pulumi.get(self, "host_state")

    @host_state.setter
    def host_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_state", value)

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the CDH instance.
        """
        return pulumi.get(self, "host_type")

    @host_type.setter
    def host_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_type", value)

    @property
    @pulumi.getter(name="prepaidPeriod")
    def prepaid_period(self) -> Optional[pulumi.Input[int]]:
        """
        The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
        Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        """
        return pulumi.get(self, "prepaid_period")

    @prepaid_period.setter
    def prepaid_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prepaid_period", value)

    @property
    @pulumi.getter(name="prepaidRenewFlag")
    def prepaid_renew_flag(self) -> Optional[pulumi.Input[str]]:
        """
        Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
        `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
        neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
        specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
        balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
        """
        return pulumi.get(self, "prepaid_renew_flag")

    @prepaid_renew_flag.setter
    def prepaid_renew_flag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prepaid_renew_flag", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        The project the instance belongs to, default to 0.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 prepaid_period: Optional[pulumi.Input[int]] = None,
                 prepaid_renew_flag: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Instance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The available zone for the CDH instance.
        :param pulumi.Input[str] charge_type: The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
        :param pulumi.Input[str] host_name: The name of the CDH instance. The max length of host_name is 60.
        :param pulumi.Input[str] host_type: The type of the CDH instance.
        :param pulumi.Input[int] prepaid_period: The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
               Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        :param pulumi.Input[str] prepaid_renew_flag: Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
               `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
               neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
               specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
               balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
        :param pulumi.Input[int] project_id: The project the instance belongs to, default to 0.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Instance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 host_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 prepaid_period: Optional[pulumi.Input[int]] = None,
                 prepaid_renew_flag: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        else:
            opts = copy.copy(opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceArgs.__new__(InstanceArgs)

            if availability_zone is None and not opts.urn:
                raise TypeError("Missing required property 'availability_zone'")
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["charge_type"] = charge_type
            __props__.__dict__["host_name"] = host_name
            __props__.__dict__["host_type"] = host_type
            __props__.__dict__["prepaid_period"] = prepaid_period
            __props__.__dict__["prepaid_renew_flag"] = prepaid_renew_flag
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["cvm_instance_ids"] = None
            __props__.__dict__["expired_time"] = None
            __props__.__dict__["host_resources"] = None
            __props__.__dict__["host_state"] = None
        super(Instance, __self__).__init__(
            'tencentcloud:Cdh/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            charge_type: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            cvm_instance_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            expired_time: Optional[pulumi.Input[str]] = None,
            host_name: Optional[pulumi.Input[str]] = None,
            host_resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceHostResourceArgs']]]]] = None,
            host_state: Optional[pulumi.Input[str]] = None,
            host_type: Optional[pulumi.Input[str]] = None,
            prepaid_period: Optional[pulumi.Input[int]] = None,
            prepaid_renew_flag: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The available zone for the CDH instance.
        :param pulumi.Input[str] charge_type: The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
        :param pulumi.Input[str] create_time: Create time of the instance.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cvm_instance_ids: Id of CVM instances that have been created on the CDH instance.
        :param pulumi.Input[str] expired_time: Expired time of the instance.
        :param pulumi.Input[str] host_name: The name of the CDH instance. The max length of host_name is 60.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceHostResourceArgs']]]] host_resources: An information list of host resource. Each element contains the following attributes:
        :param pulumi.Input[str] host_state: State of the CDH instance.
        :param pulumi.Input[str] host_type: The type of the CDH instance.
        :param pulumi.Input[int] prepaid_period: The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
               Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        :param pulumi.Input[str] prepaid_renew_flag: Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
               `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
               neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
               specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
               balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
        :param pulumi.Input[int] project_id: The project the instance belongs to, default to 0.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["charge_type"] = charge_type
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["cvm_instance_ids"] = cvm_instance_ids
        __props__.__dict__["expired_time"] = expired_time
        __props__.__dict__["host_name"] = host_name
        __props__.__dict__["host_resources"] = host_resources
        __props__.__dict__["host_state"] = host_state
        __props__.__dict__["host_type"] = host_type
        __props__.__dict__["prepaid_period"] = prepaid_period
        __props__.__dict__["prepaid_renew_flag"] = prepaid_renew_flag
        __props__.__dict__["project_id"] = project_id
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The available zone for the CDH instance.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        The charge type of instance. Valid values are `PREPAID`. The default is `PREPAID`.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the instance.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="cvmInstanceIds")
    def cvm_instance_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Id of CVM instances that have been created on the CDH instance.
        """
        return pulumi.get(self, "cvm_instance_ids")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> pulumi.Output[str]:
        """
        Expired time of the instance.
        """
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Output[str]:
        """
        The name of the CDH instance. The max length of host_name is 60.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="hostResources")
    def host_resources(self) -> pulumi.Output[Sequence['outputs.InstanceHostResource']]:
        """
        An information list of host resource. Each element contains the following attributes:
        """
        return pulumi.get(self, "host_resources")

    @property
    @pulumi.getter(name="hostState")
    def host_state(self) -> pulumi.Output[str]:
        """
        State of the CDH instance.
        """
        return pulumi.get(self, "host_state")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the CDH instance.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter(name="prepaidPeriod")
    def prepaid_period(self) -> pulumi.Output[Optional[int]]:
        """
        The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`.
        Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        """
        return pulumi.get(self, "prepaid_period")

    @property
    @pulumi.getter(name="prepaidRenewFlag")
    def prepaid_renew_flag(self) -> pulumi.Output[str]:
        """
        Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically,
        `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`:
        neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is
        specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account
        balance is sufficient. NOTE: it only works when charge_type is set to `PREPAID`.
        """
        return pulumi.get(self, "prepaid_renew_flag")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[int]]:
        """
        The project the instance belongs to, default to 0.
        """
        return pulumi.get(self, "project_id")

