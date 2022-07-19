# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AssociationArgs', 'Association']

@pulumi.input_type
class AssociationArgs:
    def __init__(__self__, *,
                 eip_id: pulumi.Input[str],
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Association resource.
        :param pulumi.Input[str] eip_id: The ID of EIP.
        :param pulumi.Input[str] instance_id: The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
               `private_ip fields`.
        :param pulumi.Input[str] network_interface_id: Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
        :param pulumi.Input[str] private_ip: Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
        """
        pulumi.set(__self__, "eip_id", eip_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if private_ip is not None:
            pulumi.set(__self__, "private_ip", private_ip)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> pulumi.Input[str]:
        """
        The ID of EIP.
        """
        return pulumi.get(self, "eip_id")

    @eip_id.setter
    def eip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "eip_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
        `private_ip fields`.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip", value)


@pulumi.input_type
class _AssociationState:
    def __init__(__self__, *,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Association resources.
        :param pulumi.Input[str] eip_id: The ID of EIP.
        :param pulumi.Input[str] instance_id: The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
               `private_ip fields`.
        :param pulumi.Input[str] network_interface_id: Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
        :param pulumi.Input[str] private_ip: Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
        """
        if eip_id is not None:
            pulumi.set(__self__, "eip_id", eip_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if private_ip is not None:
            pulumi.set(__self__, "private_ip", private_ip)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of EIP.
        """
        return pulumi.get(self, "eip_id")

    @eip_id.setter
    def eip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
        `private_ip fields`.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip", value)


class Association(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Association resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip_id: The ID of EIP.
        :param pulumi.Input[str] instance_id: The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
               `private_ip fields`.
        :param pulumi.Input[str] network_interface_id: Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
        :param pulumi.Input[str] private_ip: Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Association resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
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
            __props__ = AssociationArgs.__new__(AssociationArgs)

            if eip_id is None and not opts.urn:
                raise TypeError("Missing required property 'eip_id'")
            __props__.__dict__["eip_id"] = eip_id
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["network_interface_id"] = network_interface_id
            __props__.__dict__["private_ip"] = private_ip
        super(Association, __self__).__init__(
            'tctest:Eip/association:Association',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eip_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            private_ip: Optional[pulumi.Input[str]] = None) -> 'Association':
        """
        Get an existing Association resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip_id: The ID of EIP.
        :param pulumi.Input[str] instance_id: The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
               `private_ip fields`.
        :param pulumi.Input[str] network_interface_id: Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
        :param pulumi.Input[str] private_ip: Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AssociationState.__new__(_AssociationState)

        __props__.__dict__["eip_id"] = eip_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["network_interface_id"] = network_interface_id
        __props__.__dict__["private_ip"] = private_ip
        return Association(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> pulumi.Output[str]:
        """
        The ID of EIP.
        """
        return pulumi.get(self, "eip_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The CVM or CLB instance id going to bind with the EIP. This field is conflict with `network_interface_id` and
        `private_ip fields`.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> pulumi.Output[str]:
        """
        Indicates an IP belongs to the `network_interface_id`. This field is conflict with `instance_id`.
        """
        return pulumi.get(self, "private_ip")

