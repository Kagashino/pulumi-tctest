# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CustomerGatewayArgs', 'CustomerGateway']

@pulumi.input_type
class CustomerGatewayArgs:
    def __init__(__self__, *,
                 public_ip_address: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a CustomerGateway resource.
        :param pulumi.Input[str] public_ip_address: Public IP of the customer gateway.
        :param pulumi.Input[str] name: Name of the customer gateway. The length of character is limited to 1-60.
        :param pulumi.Input[Mapping[str, Any]] tags: A list of tags used to associate different resources.
        """
        pulumi.set(__self__, "public_ip_address", public_ip_address)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> pulumi.Input[str]:
        """
        Public IP of the customer gateway.
        """
        return pulumi.get(self, "public_ip_address")

    @public_ip_address.setter
    def public_ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_ip_address", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the customer gateway. The length of character is limited to 1-60.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A list of tags used to associate different resources.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _CustomerGatewayState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_ip_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering CustomerGateway resources.
        :param pulumi.Input[str] create_time: Create time of the customer gateway.
        :param pulumi.Input[str] name: Name of the customer gateway. The length of character is limited to 1-60.
        :param pulumi.Input[str] public_ip_address: Public IP of the customer gateway.
        :param pulumi.Input[Mapping[str, Any]] tags: A list of tags used to associate different resources.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_ip_address is not None:
            pulumi.set(__self__, "public_ip_address", public_ip_address)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the customer gateway.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the customer gateway. The length of character is limited to 1-60.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        Public IP of the customer gateway.
        """
        return pulumi.get(self, "public_ip_address")

    @public_ip_address.setter
    def public_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_ip_address", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A list of tags used to associate different resources.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class CustomerGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_ip_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Create a CustomerGateway resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the customer gateway. The length of character is limited to 1-60.
        :param pulumi.Input[str] public_ip_address: Public IP of the customer gateway.
        :param pulumi.Input[Mapping[str, Any]] tags: A list of tags used to associate different resources.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomerGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CustomerGateway resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CustomerGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomerGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_ip_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
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
            __props__ = CustomerGatewayArgs.__new__(CustomerGatewayArgs)

            __props__.__dict__["name"] = name
            if public_ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'public_ip_address'")
            __props__.__dict__["public_ip_address"] = public_ip_address
            __props__.__dict__["tags"] = tags
            __props__.__dict__["create_time"] = None
        super(CustomerGateway, __self__).__init__(
            'tencentcloud:Vpn/customerGateway:CustomerGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public_ip_address: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'CustomerGateway':
        """
        Get an existing CustomerGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Create time of the customer gateway.
        :param pulumi.Input[str] name: Name of the customer gateway. The length of character is limited to 1-60.
        :param pulumi.Input[str] public_ip_address: Public IP of the customer gateway.
        :param pulumi.Input[Mapping[str, Any]] tags: A list of tags used to associate different resources.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomerGatewayState.__new__(_CustomerGatewayState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["name"] = name
        __props__.__dict__["public_ip_address"] = public_ip_address
        __props__.__dict__["tags"] = tags
        return CustomerGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the customer gateway.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the customer gateway. The length of character is limited to 1-60.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> pulumi.Output[str]:
        """
        Public IP of the customer gateway.
        """
        return pulumi.get(self, "public_ip_address")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A list of tags used to associate different resources.
        """
        return pulumi.get(self, "tags")

