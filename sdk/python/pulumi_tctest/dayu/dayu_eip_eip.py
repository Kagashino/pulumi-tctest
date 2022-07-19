# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DayuEipEipArgs', 'DayuEipEip']

@pulumi.input_type
class DayuEipEipArgs:
    def __init__(__self__, *,
                 bind_resource_id: pulumi.Input[str],
                 bind_resource_region: pulumi.Input[str],
                 bind_resource_type: pulumi.Input[str],
                 eip: pulumi.Input[str],
                 resource_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DayuEipEip resource.
        :param pulumi.Input[str] bind_resource_id: Resource id to bind.
        :param pulumi.Input[str] bind_resource_region: Resource region to bind.
        :param pulumi.Input[str] bind_resource_type: Resource type to bind, value range [`clb`, `cvm`].
        :param pulumi.Input[str] eip: Eip of the resource.
        :param pulumi.Input[str] resource_id: ID of the resource.
        """
        pulumi.set(__self__, "bind_resource_id", bind_resource_id)
        pulumi.set(__self__, "bind_resource_region", bind_resource_region)
        pulumi.set(__self__, "bind_resource_type", bind_resource_type)
        pulumi.set(__self__, "eip", eip)
        pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="bindResourceId")
    def bind_resource_id(self) -> pulumi.Input[str]:
        """
        Resource id to bind.
        """
        return pulumi.get(self, "bind_resource_id")

    @bind_resource_id.setter
    def bind_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_resource_id", value)

    @property
    @pulumi.getter(name="bindResourceRegion")
    def bind_resource_region(self) -> pulumi.Input[str]:
        """
        Resource region to bind.
        """
        return pulumi.get(self, "bind_resource_region")

    @bind_resource_region.setter
    def bind_resource_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_resource_region", value)

    @property
    @pulumi.getter(name="bindResourceType")
    def bind_resource_type(self) -> pulumi.Input[str]:
        """
        Resource type to bind, value range [`clb`, `cvm`].
        """
        return pulumi.get(self, "bind_resource_type")

    @bind_resource_type.setter
    def bind_resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "bind_resource_type", value)

    @property
    @pulumi.getter
    def eip(self) -> pulumi.Input[str]:
        """
        Eip of the resource.
        """
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: pulumi.Input[str]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)


@pulumi.input_type
class _DayuEipEipState:
    def __init__(__self__, *,
                 bind_resource_id: Optional[pulumi.Input[str]] = None,
                 bind_resource_region: Optional[pulumi.Input[str]] = None,
                 bind_resource_type: Optional[pulumi.Input[str]] = None,
                 created_time: Optional[pulumi.Input[str]] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 eip_address_status: Optional[pulumi.Input[str]] = None,
                 eip_bound_rsc_eni: Optional[pulumi.Input[str]] = None,
                 eip_bound_rsc_ins: Optional[pulumi.Input[str]] = None,
                 eip_bound_rsc_vip: Optional[pulumi.Input[str]] = None,
                 expired_time: Optional[pulumi.Input[str]] = None,
                 modify_time: Optional[pulumi.Input[str]] = None,
                 protection_status: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DayuEipEip resources.
        :param pulumi.Input[str] bind_resource_id: Resource id to bind.
        :param pulumi.Input[str] bind_resource_region: Resource region to bind.
        :param pulumi.Input[str] bind_resource_type: Resource type to bind, value range [`clb`, `cvm`].
        :param pulumi.Input[str] created_time: Created time of the resource instance.
        :param pulumi.Input[str] eip: Eip of the resource.
        :param pulumi.Input[str] eip_address_status: Eip address status of the resource instance.
        :param pulumi.Input[str] eip_bound_rsc_eni: Eip bound rsc eni of the resource instance.
        :param pulumi.Input[str] eip_bound_rsc_ins: Eip bound rsc ins of the resource instance.
        :param pulumi.Input[str] eip_bound_rsc_vip: Eip bound rsc vip of the resource instance.
        :param pulumi.Input[str] expired_time: Expired time of the resource instance.
        :param pulumi.Input[str] modify_time: Modify time of the resource instance.
        :param pulumi.Input[str] protection_status: Protection status of the resource instance.
        :param pulumi.Input[str] resource_id: ID of the resource.
        :param pulumi.Input[str] resource_region: Region of the resource instance.
        """
        if bind_resource_id is not None:
            pulumi.set(__self__, "bind_resource_id", bind_resource_id)
        if bind_resource_region is not None:
            pulumi.set(__self__, "bind_resource_region", bind_resource_region)
        if bind_resource_type is not None:
            pulumi.set(__self__, "bind_resource_type", bind_resource_type)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if eip is not None:
            pulumi.set(__self__, "eip", eip)
        if eip_address_status is not None:
            pulumi.set(__self__, "eip_address_status", eip_address_status)
        if eip_bound_rsc_eni is not None:
            pulumi.set(__self__, "eip_bound_rsc_eni", eip_bound_rsc_eni)
        if eip_bound_rsc_ins is not None:
            pulumi.set(__self__, "eip_bound_rsc_ins", eip_bound_rsc_ins)
        if eip_bound_rsc_vip is not None:
            pulumi.set(__self__, "eip_bound_rsc_vip", eip_bound_rsc_vip)
        if expired_time is not None:
            pulumi.set(__self__, "expired_time", expired_time)
        if modify_time is not None:
            pulumi.set(__self__, "modify_time", modify_time)
        if protection_status is not None:
            pulumi.set(__self__, "protection_status", protection_status)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_region is not None:
            pulumi.set(__self__, "resource_region", resource_region)

    @property
    @pulumi.getter(name="bindResourceId")
    def bind_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource id to bind.
        """
        return pulumi.get(self, "bind_resource_id")

    @bind_resource_id.setter
    def bind_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_resource_id", value)

    @property
    @pulumi.getter(name="bindResourceRegion")
    def bind_resource_region(self) -> Optional[pulumi.Input[str]]:
        """
        Resource region to bind.
        """
        return pulumi.get(self, "bind_resource_region")

    @bind_resource_region.setter
    def bind_resource_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_resource_region", value)

    @property
    @pulumi.getter(name="bindResourceType")
    def bind_resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Resource type to bind, value range [`clb`, `cvm`].
        """
        return pulumi.get(self, "bind_resource_type")

    @bind_resource_type.setter
    def bind_resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_resource_type", value)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[pulumi.Input[str]]:
        """
        Created time of the resource instance.
        """
        return pulumi.get(self, "created_time")

    @created_time.setter
    def created_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_time", value)

    @property
    @pulumi.getter
    def eip(self) -> Optional[pulumi.Input[str]]:
        """
        Eip of the resource.
        """
        return pulumi.get(self, "eip")

    @eip.setter
    def eip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip", value)

    @property
    @pulumi.getter(name="eipAddressStatus")
    def eip_address_status(self) -> Optional[pulumi.Input[str]]:
        """
        Eip address status of the resource instance.
        """
        return pulumi.get(self, "eip_address_status")

    @eip_address_status.setter
    def eip_address_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_address_status", value)

    @property
    @pulumi.getter(name="eipBoundRscEni")
    def eip_bound_rsc_eni(self) -> Optional[pulumi.Input[str]]:
        """
        Eip bound rsc eni of the resource instance.
        """
        return pulumi.get(self, "eip_bound_rsc_eni")

    @eip_bound_rsc_eni.setter
    def eip_bound_rsc_eni(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_bound_rsc_eni", value)

    @property
    @pulumi.getter(name="eipBoundRscIns")
    def eip_bound_rsc_ins(self) -> Optional[pulumi.Input[str]]:
        """
        Eip bound rsc ins of the resource instance.
        """
        return pulumi.get(self, "eip_bound_rsc_ins")

    @eip_bound_rsc_ins.setter
    def eip_bound_rsc_ins(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_bound_rsc_ins", value)

    @property
    @pulumi.getter(name="eipBoundRscVip")
    def eip_bound_rsc_vip(self) -> Optional[pulumi.Input[str]]:
        """
        Eip bound rsc vip of the resource instance.
        """
        return pulumi.get(self, "eip_bound_rsc_vip")

    @eip_bound_rsc_vip.setter
    def eip_bound_rsc_vip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_bound_rsc_vip", value)

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> Optional[pulumi.Input[str]]:
        """
        Expired time of the resource instance.
        """
        return pulumi.get(self, "expired_time")

    @expired_time.setter
    def expired_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expired_time", value)

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> Optional[pulumi.Input[str]]:
        """
        Modify time of the resource instance.
        """
        return pulumi.get(self, "modify_time")

    @modify_time.setter
    def modify_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_time", value)

    @property
    @pulumi.getter(name="protectionStatus")
    def protection_status(self) -> Optional[pulumi.Input[str]]:
        """
        Protection status of the resource instance.
        """
        return pulumi.get(self, "protection_status")

    @protection_status.setter
    def protection_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protection_status", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceRegion")
    def resource_region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of the resource instance.
        """
        return pulumi.get(self, "resource_region")

    @resource_region.setter
    def resource_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_region", value)


class DayuEipEip(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bind_resource_id: Optional[pulumi.Input[str]] = None,
                 bind_resource_region: Optional[pulumi.Input[str]] = None,
                 bind_resource_type: Optional[pulumi.Input[str]] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DayuEipEip resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bind_resource_id: Resource id to bind.
        :param pulumi.Input[str] bind_resource_region: Resource region to bind.
        :param pulumi.Input[str] bind_resource_type: Resource type to bind, value range [`clb`, `cvm`].
        :param pulumi.Input[str] eip: Eip of the resource.
        :param pulumi.Input[str] resource_id: ID of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DayuEipEipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DayuEipEip resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DayuEipEipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DayuEipEipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bind_resource_id: Optional[pulumi.Input[str]] = None,
                 bind_resource_region: Optional[pulumi.Input[str]] = None,
                 bind_resource_type: Optional[pulumi.Input[str]] = None,
                 eip: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DayuEipEipArgs.__new__(DayuEipEipArgs)

            if bind_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'bind_resource_id'")
            __props__.__dict__["bind_resource_id"] = bind_resource_id
            if bind_resource_region is None and not opts.urn:
                raise TypeError("Missing required property 'bind_resource_region'")
            __props__.__dict__["bind_resource_region"] = bind_resource_region
            if bind_resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'bind_resource_type'")
            __props__.__dict__["bind_resource_type"] = bind_resource_type
            if eip is None and not opts.urn:
                raise TypeError("Missing required property 'eip'")
            __props__.__dict__["eip"] = eip
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            __props__.__dict__["created_time"] = None
            __props__.__dict__["eip_address_status"] = None
            __props__.__dict__["eip_bound_rsc_eni"] = None
            __props__.__dict__["eip_bound_rsc_ins"] = None
            __props__.__dict__["eip_bound_rsc_vip"] = None
            __props__.__dict__["expired_time"] = None
            __props__.__dict__["modify_time"] = None
            __props__.__dict__["protection_status"] = None
            __props__.__dict__["resource_region"] = None
        super(DayuEipEip, __self__).__init__(
            'tctest:Dayu/dayuEipEip:DayuEipEip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bind_resource_id: Optional[pulumi.Input[str]] = None,
            bind_resource_region: Optional[pulumi.Input[str]] = None,
            bind_resource_type: Optional[pulumi.Input[str]] = None,
            created_time: Optional[pulumi.Input[str]] = None,
            eip: Optional[pulumi.Input[str]] = None,
            eip_address_status: Optional[pulumi.Input[str]] = None,
            eip_bound_rsc_eni: Optional[pulumi.Input[str]] = None,
            eip_bound_rsc_ins: Optional[pulumi.Input[str]] = None,
            eip_bound_rsc_vip: Optional[pulumi.Input[str]] = None,
            expired_time: Optional[pulumi.Input[str]] = None,
            modify_time: Optional[pulumi.Input[str]] = None,
            protection_status: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_region: Optional[pulumi.Input[str]] = None) -> 'DayuEipEip':
        """
        Get an existing DayuEipEip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bind_resource_id: Resource id to bind.
        :param pulumi.Input[str] bind_resource_region: Resource region to bind.
        :param pulumi.Input[str] bind_resource_type: Resource type to bind, value range [`clb`, `cvm`].
        :param pulumi.Input[str] created_time: Created time of the resource instance.
        :param pulumi.Input[str] eip: Eip of the resource.
        :param pulumi.Input[str] eip_address_status: Eip address status of the resource instance.
        :param pulumi.Input[str] eip_bound_rsc_eni: Eip bound rsc eni of the resource instance.
        :param pulumi.Input[str] eip_bound_rsc_ins: Eip bound rsc ins of the resource instance.
        :param pulumi.Input[str] eip_bound_rsc_vip: Eip bound rsc vip of the resource instance.
        :param pulumi.Input[str] expired_time: Expired time of the resource instance.
        :param pulumi.Input[str] modify_time: Modify time of the resource instance.
        :param pulumi.Input[str] protection_status: Protection status of the resource instance.
        :param pulumi.Input[str] resource_id: ID of the resource.
        :param pulumi.Input[str] resource_region: Region of the resource instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DayuEipEipState.__new__(_DayuEipEipState)

        __props__.__dict__["bind_resource_id"] = bind_resource_id
        __props__.__dict__["bind_resource_region"] = bind_resource_region
        __props__.__dict__["bind_resource_type"] = bind_resource_type
        __props__.__dict__["created_time"] = created_time
        __props__.__dict__["eip"] = eip
        __props__.__dict__["eip_address_status"] = eip_address_status
        __props__.__dict__["eip_bound_rsc_eni"] = eip_bound_rsc_eni
        __props__.__dict__["eip_bound_rsc_ins"] = eip_bound_rsc_ins
        __props__.__dict__["eip_bound_rsc_vip"] = eip_bound_rsc_vip
        __props__.__dict__["expired_time"] = expired_time
        __props__.__dict__["modify_time"] = modify_time
        __props__.__dict__["protection_status"] = protection_status
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_region"] = resource_region
        return DayuEipEip(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bindResourceId")
    def bind_resource_id(self) -> pulumi.Output[str]:
        """
        Resource id to bind.
        """
        return pulumi.get(self, "bind_resource_id")

    @property
    @pulumi.getter(name="bindResourceRegion")
    def bind_resource_region(self) -> pulumi.Output[str]:
        """
        Resource region to bind.
        """
        return pulumi.get(self, "bind_resource_region")

    @property
    @pulumi.getter(name="bindResourceType")
    def bind_resource_type(self) -> pulumi.Output[str]:
        """
        Resource type to bind, value range [`clb`, `cvm`].
        """
        return pulumi.get(self, "bind_resource_type")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        Created time of the resource instance.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def eip(self) -> pulumi.Output[str]:
        """
        Eip of the resource.
        """
        return pulumi.get(self, "eip")

    @property
    @pulumi.getter(name="eipAddressStatus")
    def eip_address_status(self) -> pulumi.Output[str]:
        """
        Eip address status of the resource instance.
        """
        return pulumi.get(self, "eip_address_status")

    @property
    @pulumi.getter(name="eipBoundRscEni")
    def eip_bound_rsc_eni(self) -> pulumi.Output[str]:
        """
        Eip bound rsc eni of the resource instance.
        """
        return pulumi.get(self, "eip_bound_rsc_eni")

    @property
    @pulumi.getter(name="eipBoundRscIns")
    def eip_bound_rsc_ins(self) -> pulumi.Output[str]:
        """
        Eip bound rsc ins of the resource instance.
        """
        return pulumi.get(self, "eip_bound_rsc_ins")

    @property
    @pulumi.getter(name="eipBoundRscVip")
    def eip_bound_rsc_vip(self) -> pulumi.Output[str]:
        """
        Eip bound rsc vip of the resource instance.
        """
        return pulumi.get(self, "eip_bound_rsc_vip")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> pulumi.Output[str]:
        """
        Expired time of the resource instance.
        """
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> pulumi.Output[str]:
        """
        Modify time of the resource instance.
        """
        return pulumi.get(self, "modify_time")

    @property
    @pulumi.getter(name="protectionStatus")
    def protection_status(self) -> pulumi.Output[str]:
        """
        Protection status of the resource instance.
        """
        return pulumi.get(self, "protection_status")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceRegion")
    def resource_region(self) -> pulumi.Output[str]:
        """
        Region of the resource instance.
        """
        return pulumi.get(self, "resource_region")

