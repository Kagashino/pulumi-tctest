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

__all__ = ['DdosPolicyV2Args', 'DdosPolicyV2']

@pulumi.input_type
class DdosPolicyV2Args:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]]] = None,
                 black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]]] = None,
                 business: Optional[pulumi.Input[str]] = None,
                 ddos_ai: Optional[pulumi.Input[str]] = None,
                 ddos_connect_limit: Optional[pulumi.Input['DdosPolicyV2DdosConnectLimitArgs']] = None,
                 ddos_geo_ip_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]] = None,
                 ddos_level: Optional[pulumi.Input[str]] = None,
                 ddos_speed_limit_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]]] = None,
                 ddos_threshold: Optional[pulumi.Input[int]] = None,
                 packet_filters: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]]] = None,
                 protocol_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]]] = None):
        """
        The set of arguments for constructing a DdosPolicyV2 resource.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]] acls: Port ACL policy for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]] black_white_ips: DDoS-protected IP blacklist and whitelist.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[str] ddos_ai: AI protection switch, take the value [`on`, `off`].
        :param pulumi.Input['DdosPolicyV2DdosConnectLimitArgs'] ddos_connect_limit: DDoS connection suppression options.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]] ddos_geo_ip_block_configs: DDoS-protected area block configuration.
        :param pulumi.Input[str] ddos_level: Protection class, value [`low`, `middle`, `high`].
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]] ddos_speed_limit_configs: Access speed limit configuration for DDoS protection.
        :param pulumi.Input[int] ddos_threshold: DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it
               means that the default value is adopted.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]] packet_filters: Feature filtering rules for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]] protocol_block_configs: Protocol block configuration for DDoS protection.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if black_white_ips is not None:
            pulumi.set(__self__, "black_white_ips", black_white_ips)
        if business is not None:
            pulumi.set(__self__, "business", business)
        if ddos_ai is not None:
            pulumi.set(__self__, "ddos_ai", ddos_ai)
        if ddos_connect_limit is not None:
            pulumi.set(__self__, "ddos_connect_limit", ddos_connect_limit)
        if ddos_geo_ip_block_configs is not None:
            pulumi.set(__self__, "ddos_geo_ip_block_configs", ddos_geo_ip_block_configs)
        if ddos_level is not None:
            pulumi.set(__self__, "ddos_level", ddos_level)
        if ddos_speed_limit_configs is not None:
            pulumi.set(__self__, "ddos_speed_limit_configs", ddos_speed_limit_configs)
        if ddos_threshold is not None:
            pulumi.set(__self__, "ddos_threshold", ddos_threshold)
        if packet_filters is not None:
            pulumi.set(__self__, "packet_filters", packet_filters)
        if protocol_block_configs is not None:
            pulumi.set(__self__, "protocol_block_configs", protocol_block_configs)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource instance.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]]]:
        """
        Port ACL policy for DDoS protection.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="blackWhiteIps")
    def black_white_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]]]:
        """
        DDoS-protected IP blacklist and whitelist.
        """
        return pulumi.get(self, "black_white_ips")

    @black_white_ips.setter
    def black_white_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]]]):
        pulumi.set(self, "black_white_ips", value)

    @property
    @pulumi.getter
    def business(self) -> Optional[pulumi.Input[str]]:
        """
        Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
        packet; net indicates anti-anti-ip pro version.
        """
        return pulumi.get(self, "business")

    @business.setter
    def business(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "business", value)

    @property
    @pulumi.getter(name="ddosAi")
    def ddos_ai(self) -> Optional[pulumi.Input[str]]:
        """
        AI protection switch, take the value [`on`, `off`].
        """
        return pulumi.get(self, "ddos_ai")

    @ddos_ai.setter
    def ddos_ai(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ddos_ai", value)

    @property
    @pulumi.getter(name="ddosConnectLimit")
    def ddos_connect_limit(self) -> Optional[pulumi.Input['DdosPolicyV2DdosConnectLimitArgs']]:
        """
        DDoS connection suppression options.
        """
        return pulumi.get(self, "ddos_connect_limit")

    @ddos_connect_limit.setter
    def ddos_connect_limit(self, value: Optional[pulumi.Input['DdosPolicyV2DdosConnectLimitArgs']]):
        pulumi.set(self, "ddos_connect_limit", value)

    @property
    @pulumi.getter(name="ddosGeoIpBlockConfigs")
    def ddos_geo_ip_block_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]]:
        """
        DDoS-protected area block configuration.
        """
        return pulumi.get(self, "ddos_geo_ip_block_configs")

    @ddos_geo_ip_block_configs.setter
    def ddos_geo_ip_block_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]]):
        pulumi.set(self, "ddos_geo_ip_block_configs", value)

    @property
    @pulumi.getter(name="ddosLevel")
    def ddos_level(self) -> Optional[pulumi.Input[str]]:
        """
        Protection class, value [`low`, `middle`, `high`].
        """
        return pulumi.get(self, "ddos_level")

    @ddos_level.setter
    def ddos_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ddos_level", value)

    @property
    @pulumi.getter(name="ddosSpeedLimitConfigs")
    def ddos_speed_limit_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]]]:
        """
        Access speed limit configuration for DDoS protection.
        """
        return pulumi.get(self, "ddos_speed_limit_configs")

    @ddos_speed_limit_configs.setter
    def ddos_speed_limit_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]]]):
        pulumi.set(self, "ddos_speed_limit_configs", value)

    @property
    @pulumi.getter(name="ddosThreshold")
    def ddos_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it
        means that the default value is adopted.
        """
        return pulumi.get(self, "ddos_threshold")

    @ddos_threshold.setter
    def ddos_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ddos_threshold", value)

    @property
    @pulumi.getter(name="packetFilters")
    def packet_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]]]:
        """
        Feature filtering rules for DDoS protection.
        """
        return pulumi.get(self, "packet_filters")

    @packet_filters.setter
    def packet_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]]]):
        pulumi.set(self, "packet_filters", value)

    @property
    @pulumi.getter(name="protocolBlockConfigs")
    def protocol_block_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]]]:
        """
        Protocol block configuration for DDoS protection.
        """
        return pulumi.get(self, "protocol_block_configs")

    @protocol_block_configs.setter
    def protocol_block_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]]]):
        pulumi.set(self, "protocol_block_configs", value)


@pulumi.input_type
class _DdosPolicyV2State:
    def __init__(__self__, *,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]]] = None,
                 black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]]] = None,
                 business: Optional[pulumi.Input[str]] = None,
                 ddos_ai: Optional[pulumi.Input[str]] = None,
                 ddos_connect_limit: Optional[pulumi.Input['DdosPolicyV2DdosConnectLimitArgs']] = None,
                 ddos_geo_ip_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]] = None,
                 ddos_level: Optional[pulumi.Input[str]] = None,
                 ddos_speed_limit_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]]] = None,
                 ddos_threshold: Optional[pulumi.Input[int]] = None,
                 packet_filters: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]]] = None,
                 protocol_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DdosPolicyV2 resources.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]] acls: Port ACL policy for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]] black_white_ips: DDoS-protected IP blacklist and whitelist.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[str] ddos_ai: AI protection switch, take the value [`on`, `off`].
        :param pulumi.Input['DdosPolicyV2DdosConnectLimitArgs'] ddos_connect_limit: DDoS connection suppression options.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]] ddos_geo_ip_block_configs: DDoS-protected area block configuration.
        :param pulumi.Input[str] ddos_level: Protection class, value [`low`, `middle`, `high`].
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]] ddos_speed_limit_configs: Access speed limit configuration for DDoS protection.
        :param pulumi.Input[int] ddos_threshold: DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it
               means that the default value is adopted.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]] packet_filters: Feature filtering rules for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]] protocol_block_configs: Protocol block configuration for DDoS protection.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        """
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if black_white_ips is not None:
            pulumi.set(__self__, "black_white_ips", black_white_ips)
        if business is not None:
            pulumi.set(__self__, "business", business)
        if ddos_ai is not None:
            pulumi.set(__self__, "ddos_ai", ddos_ai)
        if ddos_connect_limit is not None:
            pulumi.set(__self__, "ddos_connect_limit", ddos_connect_limit)
        if ddos_geo_ip_block_configs is not None:
            pulumi.set(__self__, "ddos_geo_ip_block_configs", ddos_geo_ip_block_configs)
        if ddos_level is not None:
            pulumi.set(__self__, "ddos_level", ddos_level)
        if ddos_speed_limit_configs is not None:
            pulumi.set(__self__, "ddos_speed_limit_configs", ddos_speed_limit_configs)
        if ddos_threshold is not None:
            pulumi.set(__self__, "ddos_threshold", ddos_threshold)
        if packet_filters is not None:
            pulumi.set(__self__, "packet_filters", packet_filters)
        if protocol_block_configs is not None:
            pulumi.set(__self__, "protocol_block_configs", protocol_block_configs)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]]]:
        """
        Port ACL policy for DDoS protection.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2AclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="blackWhiteIps")
    def black_white_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]]]:
        """
        DDoS-protected IP blacklist and whitelist.
        """
        return pulumi.get(self, "black_white_ips")

    @black_white_ips.setter
    def black_white_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2BlackWhiteIpArgs']]]]):
        pulumi.set(self, "black_white_ips", value)

    @property
    @pulumi.getter
    def business(self) -> Optional[pulumi.Input[str]]:
        """
        Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
        packet; net indicates anti-anti-ip pro version.
        """
        return pulumi.get(self, "business")

    @business.setter
    def business(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "business", value)

    @property
    @pulumi.getter(name="ddosAi")
    def ddos_ai(self) -> Optional[pulumi.Input[str]]:
        """
        AI protection switch, take the value [`on`, `off`].
        """
        return pulumi.get(self, "ddos_ai")

    @ddos_ai.setter
    def ddos_ai(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ddos_ai", value)

    @property
    @pulumi.getter(name="ddosConnectLimit")
    def ddos_connect_limit(self) -> Optional[pulumi.Input['DdosPolicyV2DdosConnectLimitArgs']]:
        """
        DDoS connection suppression options.
        """
        return pulumi.get(self, "ddos_connect_limit")

    @ddos_connect_limit.setter
    def ddos_connect_limit(self, value: Optional[pulumi.Input['DdosPolicyV2DdosConnectLimitArgs']]):
        pulumi.set(self, "ddos_connect_limit", value)

    @property
    @pulumi.getter(name="ddosGeoIpBlockConfigs")
    def ddos_geo_ip_block_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]]:
        """
        DDoS-protected area block configuration.
        """
        return pulumi.get(self, "ddos_geo_ip_block_configs")

    @ddos_geo_ip_block_configs.setter
    def ddos_geo_ip_block_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]]):
        pulumi.set(self, "ddos_geo_ip_block_configs", value)

    @property
    @pulumi.getter(name="ddosLevel")
    def ddos_level(self) -> Optional[pulumi.Input[str]]:
        """
        Protection class, value [`low`, `middle`, `high`].
        """
        return pulumi.get(self, "ddos_level")

    @ddos_level.setter
    def ddos_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ddos_level", value)

    @property
    @pulumi.getter(name="ddosSpeedLimitConfigs")
    def ddos_speed_limit_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]]]:
        """
        Access speed limit configuration for DDoS protection.
        """
        return pulumi.get(self, "ddos_speed_limit_configs")

    @ddos_speed_limit_configs.setter
    def ddos_speed_limit_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2DdosSpeedLimitConfigArgs']]]]):
        pulumi.set(self, "ddos_speed_limit_configs", value)

    @property
    @pulumi.getter(name="ddosThreshold")
    def ddos_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it
        means that the default value is adopted.
        """
        return pulumi.get(self, "ddos_threshold")

    @ddos_threshold.setter
    def ddos_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ddos_threshold", value)

    @property
    @pulumi.getter(name="packetFilters")
    def packet_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]]]:
        """
        Feature filtering rules for DDoS protection.
        """
        return pulumi.get(self, "packet_filters")

    @packet_filters.setter
    def packet_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2PacketFilterArgs']]]]):
        pulumi.set(self, "packet_filters", value)

    @property
    @pulumi.getter(name="protocolBlockConfigs")
    def protocol_block_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]]]:
        """
        Protocol block configuration for DDoS protection.
        """
        return pulumi.get(self, "protocol_block_configs")

    @protocol_block_configs.setter
    def protocol_block_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DdosPolicyV2ProtocolBlockConfigArgs']]]]):
        pulumi.set(self, "protocol_block_configs", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource instance.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)


class DdosPolicyV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2AclArgs']]]]] = None,
                 black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2BlackWhiteIpArgs']]]]] = None,
                 business: Optional[pulumi.Input[str]] = None,
                 ddos_ai: Optional[pulumi.Input[str]] = None,
                 ddos_connect_limit: Optional[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosConnectLimitArgs']]] = None,
                 ddos_geo_ip_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]]] = None,
                 ddos_level: Optional[pulumi.Input[str]] = None,
                 ddos_speed_limit_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosSpeedLimitConfigArgs']]]]] = None,
                 ddos_threshold: Optional[pulumi.Input[int]] = None,
                 packet_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2PacketFilterArgs']]]]] = None,
                 protocol_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2ProtocolBlockConfigArgs']]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DdosPolicyV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2AclArgs']]]] acls: Port ACL policy for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2BlackWhiteIpArgs']]]] black_white_ips: DDoS-protected IP blacklist and whitelist.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[str] ddos_ai: AI protection switch, take the value [`on`, `off`].
        :param pulumi.Input[pulumi.InputType['DdosPolicyV2DdosConnectLimitArgs']] ddos_connect_limit: DDoS connection suppression options.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]] ddos_geo_ip_block_configs: DDoS-protected area block configuration.
        :param pulumi.Input[str] ddos_level: Protection class, value [`low`, `middle`, `high`].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosSpeedLimitConfigArgs']]]] ddos_speed_limit_configs: Access speed limit configuration for DDoS protection.
        :param pulumi.Input[int] ddos_threshold: DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it
               means that the default value is adopted.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2PacketFilterArgs']]]] packet_filters: Feature filtering rules for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2ProtocolBlockConfigArgs']]]] protocol_block_configs: Protocol block configuration for DDoS protection.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DdosPolicyV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DdosPolicyV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DdosPolicyV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DdosPolicyV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2AclArgs']]]]] = None,
                 black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2BlackWhiteIpArgs']]]]] = None,
                 business: Optional[pulumi.Input[str]] = None,
                 ddos_ai: Optional[pulumi.Input[str]] = None,
                 ddos_connect_limit: Optional[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosConnectLimitArgs']]] = None,
                 ddos_geo_ip_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]]] = None,
                 ddos_level: Optional[pulumi.Input[str]] = None,
                 ddos_speed_limit_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosSpeedLimitConfigArgs']]]]] = None,
                 ddos_threshold: Optional[pulumi.Input[int]] = None,
                 packet_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2PacketFilterArgs']]]]] = None,
                 protocol_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2ProtocolBlockConfigArgs']]]]] = None,
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
            __props__ = DdosPolicyV2Args.__new__(DdosPolicyV2Args)

            __props__.__dict__["acls"] = acls
            __props__.__dict__["black_white_ips"] = black_white_ips
            __props__.__dict__["business"] = business
            __props__.__dict__["ddos_ai"] = ddos_ai
            __props__.__dict__["ddos_connect_limit"] = ddos_connect_limit
            __props__.__dict__["ddos_geo_ip_block_configs"] = ddos_geo_ip_block_configs
            __props__.__dict__["ddos_level"] = ddos_level
            __props__.__dict__["ddos_speed_limit_configs"] = ddos_speed_limit_configs
            __props__.__dict__["ddos_threshold"] = ddos_threshold
            __props__.__dict__["packet_filters"] = packet_filters
            __props__.__dict__["protocol_block_configs"] = protocol_block_configs
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
        super(DdosPolicyV2, __self__).__init__(
            'tctest:Dayu/ddosPolicyV2:DdosPolicyV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2AclArgs']]]]] = None,
            black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2BlackWhiteIpArgs']]]]] = None,
            business: Optional[pulumi.Input[str]] = None,
            ddos_ai: Optional[pulumi.Input[str]] = None,
            ddos_connect_limit: Optional[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosConnectLimitArgs']]] = None,
            ddos_geo_ip_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]]] = None,
            ddos_level: Optional[pulumi.Input[str]] = None,
            ddos_speed_limit_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosSpeedLimitConfigArgs']]]]] = None,
            ddos_threshold: Optional[pulumi.Input[int]] = None,
            packet_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2PacketFilterArgs']]]]] = None,
            protocol_block_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2ProtocolBlockConfigArgs']]]]] = None,
            resource_id: Optional[pulumi.Input[str]] = None) -> 'DdosPolicyV2':
        """
        Get an existing DdosPolicyV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2AclArgs']]]] acls: Port ACL policy for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2BlackWhiteIpArgs']]]] black_white_ips: DDoS-protected IP blacklist and whitelist.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[str] ddos_ai: AI protection switch, take the value [`on`, `off`].
        :param pulumi.Input[pulumi.InputType['DdosPolicyV2DdosConnectLimitArgs']] ddos_connect_limit: DDoS connection suppression options.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosGeoIpBlockConfigArgs']]]] ddos_geo_ip_block_configs: DDoS-protected area block configuration.
        :param pulumi.Input[str] ddos_level: Protection class, value [`low`, `middle`, `high`].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2DdosSpeedLimitConfigArgs']]]] ddos_speed_limit_configs: Access speed limit configuration for DDoS protection.
        :param pulumi.Input[int] ddos_threshold: DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it
               means that the default value is adopted.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2PacketFilterArgs']]]] packet_filters: Feature filtering rules for DDoS protection.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DdosPolicyV2ProtocolBlockConfigArgs']]]] protocol_block_configs: Protocol block configuration for DDoS protection.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DdosPolicyV2State.__new__(_DdosPolicyV2State)

        __props__.__dict__["acls"] = acls
        __props__.__dict__["black_white_ips"] = black_white_ips
        __props__.__dict__["business"] = business
        __props__.__dict__["ddos_ai"] = ddos_ai
        __props__.__dict__["ddos_connect_limit"] = ddos_connect_limit
        __props__.__dict__["ddos_geo_ip_block_configs"] = ddos_geo_ip_block_configs
        __props__.__dict__["ddos_level"] = ddos_level
        __props__.__dict__["ddos_speed_limit_configs"] = ddos_speed_limit_configs
        __props__.__dict__["ddos_threshold"] = ddos_threshold
        __props__.__dict__["packet_filters"] = packet_filters
        __props__.__dict__["protocol_block_configs"] = protocol_block_configs
        __props__.__dict__["resource_id"] = resource_id
        return DdosPolicyV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acls(self) -> pulumi.Output[Optional[Sequence['outputs.DdosPolicyV2Acl']]]:
        """
        Port ACL policy for DDoS protection.
        """
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter(name="blackWhiteIps")
    def black_white_ips(self) -> pulumi.Output[Optional[Sequence['outputs.DdosPolicyV2BlackWhiteIp']]]:
        """
        DDoS-protected IP blacklist and whitelist.
        """
        return pulumi.get(self, "black_white_ips")

    @property
    @pulumi.getter
    def business(self) -> pulumi.Output[Optional[str]]:
        """
        Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
        packet; net indicates anti-anti-ip pro version.
        """
        return pulumi.get(self, "business")

    @property
    @pulumi.getter(name="ddosAi")
    def ddos_ai(self) -> pulumi.Output[Optional[str]]:
        """
        AI protection switch, take the value [`on`, `off`].
        """
        return pulumi.get(self, "ddos_ai")

    @property
    @pulumi.getter(name="ddosConnectLimit")
    def ddos_connect_limit(self) -> pulumi.Output[Optional['outputs.DdosPolicyV2DdosConnectLimit']]:
        """
        DDoS connection suppression options.
        """
        return pulumi.get(self, "ddos_connect_limit")

    @property
    @pulumi.getter(name="ddosGeoIpBlockConfigs")
    def ddos_geo_ip_block_configs(self) -> pulumi.Output[Optional[Sequence['outputs.DdosPolicyV2DdosGeoIpBlockConfig']]]:
        """
        DDoS-protected area block configuration.
        """
        return pulumi.get(self, "ddos_geo_ip_block_configs")

    @property
    @pulumi.getter(name="ddosLevel")
    def ddos_level(self) -> pulumi.Output[Optional[str]]:
        """
        Protection class, value [`low`, `middle`, `high`].
        """
        return pulumi.get(self, "ddos_level")

    @property
    @pulumi.getter(name="ddosSpeedLimitConfigs")
    def ddos_speed_limit_configs(self) -> pulumi.Output[Optional[Sequence['outputs.DdosPolicyV2DdosSpeedLimitConfig']]]:
        """
        Access speed limit configuration for DDoS protection.
        """
        return pulumi.get(self, "ddos_speed_limit_configs")

    @property
    @pulumi.getter(name="ddosThreshold")
    def ddos_threshold(self) -> pulumi.Output[Optional[int]]:
        """
        DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it
        means that the default value is adopted.
        """
        return pulumi.get(self, "ddos_threshold")

    @property
    @pulumi.getter(name="packetFilters")
    def packet_filters(self) -> pulumi.Output[Optional[Sequence['outputs.DdosPolicyV2PacketFilter']]]:
        """
        Feature filtering rules for DDoS protection.
        """
        return pulumi.get(self, "packet_filters")

    @property
    @pulumi.getter(name="protocolBlockConfigs")
    def protocol_block_configs(self) -> pulumi.Output[Optional[Sequence['outputs.DdosPolicyV2ProtocolBlockConfig']]]:
        """
        Protocol block configuration for DDoS protection.
        """
        return pulumi.get(self, "protocol_block_configs")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource instance.
        """
        return pulumi.get(self, "resource_id")

