# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReservedInstanceArgs', 'ReservedInstance']

@pulumi.input_type
class ReservedInstanceArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[str],
                 instance_count: pulumi.Input[int],
                 reserved_instance_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReservedInstance resource.
        :param pulumi.Input[str] config_id: Configuration ID of the reserved instance.
        :param pulumi.Input[int] instance_count: Number of reserved instances to be purchased.
        :param pulumi.Input[str] reserved_instance_name: Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
               Up to 60 characters (including pattern strings) are supported.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "instance_count", instance_count)
        if reserved_instance_name is not None:
            pulumi.set(__self__, "reserved_instance_name", reserved_instance_name)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[str]:
        """
        Configuration ID of the reserved instance.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Input[int]:
        """
        Number of reserved instances to be purchased.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="reservedInstanceName")
    def reserved_instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
        Up to 60 characters (including pattern strings) are supported.
        """
        return pulumi.get(self, "reserved_instance_name")

    @reserved_instance_name.setter
    def reserved_instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reserved_instance_name", value)


@pulumi.input_type
class _ReservedInstanceState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 reserved_instance_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReservedInstance resources.
        :param pulumi.Input[str] config_id: Configuration ID of the reserved instance.
        :param pulumi.Input[str] end_time: Expiry time of the RI.
        :param pulumi.Input[int] instance_count: Number of reserved instances to be purchased.
        :param pulumi.Input[str] reserved_instance_name: Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
               Up to 60 characters (including pattern strings) are supported.
        :param pulumi.Input[str] start_time: Start time of the RI.
        :param pulumi.Input[str] status: Status of the RI at the time of purchase.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if instance_count is not None:
            pulumi.set(__self__, "instance_count", instance_count)
        if reserved_instance_name is not None:
            pulumi.set(__self__, "reserved_instance_name", reserved_instance_name)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[str]]:
        """
        Configuration ID of the reserved instance.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        Expiry time of the RI.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of reserved instances to be purchased.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="reservedInstanceName")
    def reserved_instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
        Up to 60 characters (including pattern strings) are supported.
        """
        return pulumi.get(self, "reserved_instance_name")

    @reserved_instance_name.setter
    def reserved_instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reserved_instance_name", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Start time of the RI.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the RI at the time of purchase.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ReservedInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 reserved_instance_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ReservedInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: Configuration ID of the reserved instance.
        :param pulumi.Input[int] instance_count: Number of reserved instances to be purchased.
        :param pulumi.Input[str] reserved_instance_name: Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
               Up to 60 characters (including pattern strings) are supported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReservedInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ReservedInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ReservedInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReservedInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 reserved_instance_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReservedInstanceArgs.__new__(ReservedInstanceArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if instance_count is None and not opts.urn:
                raise TypeError("Missing required property 'instance_count'")
            __props__.__dict__["instance_count"] = instance_count
            __props__.__dict__["reserved_instance_name"] = reserved_instance_name
            __props__.__dict__["end_time"] = None
            __props__.__dict__["start_time"] = None
            __props__.__dict__["status"] = None
        super(ReservedInstance, __self__).__init__(
            'tctest:Cvm/reservedInstance:ReservedInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[str]] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            instance_count: Optional[pulumi.Input[int]] = None,
            reserved_instance_name: Optional[pulumi.Input[str]] = None,
            start_time: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ReservedInstance':
        """
        Get an existing ReservedInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: Configuration ID of the reserved instance.
        :param pulumi.Input[str] end_time: Expiry time of the RI.
        :param pulumi.Input[int] instance_count: Number of reserved instances to be purchased.
        :param pulumi.Input[str] reserved_instance_name: Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
               Up to 60 characters (including pattern strings) are supported.
        :param pulumi.Input[str] start_time: Start time of the RI.
        :param pulumi.Input[str] status: Status of the RI at the time of purchase.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReservedInstanceState.__new__(_ReservedInstanceState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["instance_count"] = instance_count
        __props__.__dict__["reserved_instance_name"] = reserved_instance_name
        __props__.__dict__["start_time"] = start_time
        __props__.__dict__["status"] = status
        return ReservedInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[str]:
        """
        Configuration ID of the reserved instance.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        Expiry time of the RI.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Output[int]:
        """
        Number of reserved instances to be purchased.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="reservedInstanceName")
    def reserved_instance_name(self) -> pulumi.Output[Optional[str]]:
        """
        Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
        Up to 60 characters (including pattern strings) are supported.
        """
        return pulumi.get(self, "reserved_instance_name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        Start time of the RI.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the RI at the time of purchase.
        """
        return pulumi.get(self, "status")

