# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GatewayInstanceArgs', 'GatewayInstance']

@pulumi.input_type
class GatewayInstanceArgs:
    def __init__(__self__, *,
                 network_instance_id: pulumi.Input[str],
                 network_type: pulumi.Input[str],
                 gateway_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GatewayInstance resource.
        :param pulumi.Input[str] network_instance_id: If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the
               available value is CCN instance ID.
        :param pulumi.Input[str] network_type: Type of associated network. Valid value: `VPC` and `CCN`.
        :param pulumi.Input[str] gateway_type: Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC
               can create two DCGs, the one is NAT type and the other is non-NAT type.
        :param pulumi.Input[str] name: Name of the DCG.
        """
        pulumi.set(__self__, "network_instance_id", network_instance_id)
        pulumi.set(__self__, "network_type", network_type)
        if gateway_type is not None:
            pulumi.set(__self__, "gateway_type", gateway_type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="networkInstanceId")
    def network_instance_id(self) -> pulumi.Input[str]:
        """
        If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the
        available value is CCN instance ID.
        """
        return pulumi.get(self, "network_instance_id")

    @network_instance_id.setter
    def network_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_instance_id", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Input[str]:
        """
        Type of associated network. Valid value: `VPC` and `CCN`.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="gatewayType")
    def gateway_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC
        can create two DCGs, the one is NAT type and the other is non-NAT type.
        """
        return pulumi.get(self, "gateway_type")

    @gateway_type.setter
    def gateway_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the DCG.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GatewayInstanceState:
    def __init__(__self__, *,
                 cnn_route_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 enable_bgp: Optional[pulumi.Input[bool]] = None,
                 gateway_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_instance_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewayInstance resources.
        :param pulumi.Input[str] cnn_route_type: Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP
               enabled.
        :param pulumi.Input[str] create_time: Creation time of resource.
        :param pulumi.Input[bool] enable_bgp: Indicates whether the BGP is enabled.
        :param pulumi.Input[str] gateway_type: Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC
               can create two DCGs, the one is NAT type and the other is non-NAT type.
        :param pulumi.Input[str] name: Name of the DCG.
        :param pulumi.Input[str] network_instance_id: If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the
               available value is CCN instance ID.
        :param pulumi.Input[str] network_type: Type of associated network. Valid value: `VPC` and `CCN`.
        """
        if cnn_route_type is not None:
            pulumi.set(__self__, "cnn_route_type", cnn_route_type)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if enable_bgp is not None:
            pulumi.set(__self__, "enable_bgp", enable_bgp)
        if gateway_type is not None:
            pulumi.set(__self__, "gateway_type", gateway_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_instance_id is not None:
            pulumi.set(__self__, "network_instance_id", network_instance_id)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)

    @property
    @pulumi.getter(name="cnnRouteType")
    def cnn_route_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP
        enabled.
        """
        return pulumi.get(self, "cnn_route_type")

    @cnn_route_type.setter
    def cnn_route_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cnn_route_type", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of resource.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="enableBgp")
    def enable_bgp(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the BGP is enabled.
        """
        return pulumi.get(self, "enable_bgp")

    @enable_bgp.setter
    def enable_bgp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_bgp", value)

    @property
    @pulumi.getter(name="gatewayType")
    def gateway_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC
        can create two DCGs, the one is NAT type and the other is non-NAT type.
        """
        return pulumi.get(self, "gateway_type")

    @gateway_type.setter
    def gateway_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the DCG.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkInstanceId")
    def network_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the
        available value is CCN instance ID.
        """
        return pulumi.get(self, "network_instance_id")

    @network_instance_id.setter
    def network_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_instance_id", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of associated network. Valid value: `VPC` and `CCN`.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)


class GatewayInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_instance_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a GatewayInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_type: Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC
               can create two DCGs, the one is NAT type and the other is non-NAT type.
        :param pulumi.Input[str] name: Name of the DCG.
        :param pulumi.Input[str] network_instance_id: If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the
               available value is CCN instance ID.
        :param pulumi.Input[str] network_type: Type of associated network. Valid value: `VPC` and `CCN`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a GatewayInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GatewayInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_instance_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = GatewayInstanceArgs.__new__(GatewayInstanceArgs)

            __props__.__dict__["gateway_type"] = gateway_type
            __props__.__dict__["name"] = name
            if network_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_instance_id'")
            __props__.__dict__["network_instance_id"] = network_instance_id
            if network_type is None and not opts.urn:
                raise TypeError("Missing required property 'network_type'")
            __props__.__dict__["network_type"] = network_type
            __props__.__dict__["cnn_route_type"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["enable_bgp"] = None
        super(GatewayInstance, __self__).__init__(
            'tencentcloud:Dc/gatewayInstance:GatewayInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cnn_route_type: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            enable_bgp: Optional[pulumi.Input[bool]] = None,
            gateway_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_instance_id: Optional[pulumi.Input[str]] = None,
            network_type: Optional[pulumi.Input[str]] = None) -> 'GatewayInstance':
        """
        Get an existing GatewayInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cnn_route_type: Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP
               enabled.
        :param pulumi.Input[str] create_time: Creation time of resource.
        :param pulumi.Input[bool] enable_bgp: Indicates whether the BGP is enabled.
        :param pulumi.Input[str] gateway_type: Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC
               can create two DCGs, the one is NAT type and the other is non-NAT type.
        :param pulumi.Input[str] name: Name of the DCG.
        :param pulumi.Input[str] network_instance_id: If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the
               available value is CCN instance ID.
        :param pulumi.Input[str] network_type: Type of associated network. Valid value: `VPC` and `CCN`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayInstanceState.__new__(_GatewayInstanceState)

        __props__.__dict__["cnn_route_type"] = cnn_route_type
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["enable_bgp"] = enable_bgp
        __props__.__dict__["gateway_type"] = gateway_type
        __props__.__dict__["name"] = name
        __props__.__dict__["network_instance_id"] = network_instance_id
        __props__.__dict__["network_type"] = network_type
        return GatewayInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cnnRouteType")
    def cnn_route_type(self) -> pulumi.Output[str]:
        """
        Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP
        enabled.
        """
        return pulumi.get(self, "cnn_route_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="enableBgp")
    def enable_bgp(self) -> pulumi.Output[bool]:
        """
        Indicates whether the BGP is enabled.
        """
        return pulumi.get(self, "enable_bgp")

    @property
    @pulumi.getter(name="gatewayType")
    def gateway_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC
        can create two DCGs, the one is NAT type and the other is non-NAT type.
        """
        return pulumi.get(self, "gateway_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the DCG.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInstanceId")
    def network_instance_id(self) -> pulumi.Output[str]:
        """
        If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the
        available value is CCN instance ID.
        """
        return pulumi.get(self, "network_instance_id")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[str]:
        """
        Type of associated network. Valid value: `VPC` and `CCN`.
        """
        return pulumi.get(self, "network_type")

