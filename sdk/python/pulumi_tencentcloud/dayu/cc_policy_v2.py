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

__all__ = ['CCPolicyV2Args', 'CCPolicyV2']

@pulumi.input_type
class CCPolicyV2Args:
    def __init__(__self__, *,
                 business: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 cc_black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]]] = None,
                 cc_geo_ip_policys: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]]] = None,
                 cc_precision_policys: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]]] = None,
                 cc_precision_req_limits: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]]] = None,
                 thresholds: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]]] = None):
        """
        The set of arguments for constructing a CCPolicyV2 resource.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]] cc_black_white_ips: Blacklist and whitelist.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]] cc_geo_ip_policys: Details of the CC region blocking policy list.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]] cc_precision_policys: CC Precision Protection List.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]] cc_precision_req_limits: CC frequency throttling policy.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]] thresholds: List of protection threshold configurations.
        """
        pulumi.set(__self__, "business", business)
        pulumi.set(__self__, "resource_id", resource_id)
        if cc_black_white_ips is not None:
            pulumi.set(__self__, "cc_black_white_ips", cc_black_white_ips)
        if cc_geo_ip_policys is not None:
            pulumi.set(__self__, "cc_geo_ip_policys", cc_geo_ip_policys)
        if cc_precision_policys is not None:
            pulumi.set(__self__, "cc_precision_policys", cc_precision_policys)
        if cc_precision_req_limits is not None:
            pulumi.set(__self__, "cc_precision_req_limits", cc_precision_req_limits)
        if thresholds is not None:
            pulumi.set(__self__, "thresholds", thresholds)

    @property
    @pulumi.getter
    def business(self) -> pulumi.Input[str]:
        """
        Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
        packet; net indicates anti-anti-ip pro version.
        """
        return pulumi.get(self, "business")

    @business.setter
    def business(self, value: pulumi.Input[str]):
        pulumi.set(self, "business", value)

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
    @pulumi.getter(name="ccBlackWhiteIps")
    def cc_black_white_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]]]:
        """
        Blacklist and whitelist.
        """
        return pulumi.get(self, "cc_black_white_ips")

    @cc_black_white_ips.setter
    def cc_black_white_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]]]):
        pulumi.set(self, "cc_black_white_ips", value)

    @property
    @pulumi.getter(name="ccGeoIpPolicys")
    def cc_geo_ip_policys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]]]:
        """
        Details of the CC region blocking policy list.
        """
        return pulumi.get(self, "cc_geo_ip_policys")

    @cc_geo_ip_policys.setter
    def cc_geo_ip_policys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]]]):
        pulumi.set(self, "cc_geo_ip_policys", value)

    @property
    @pulumi.getter(name="ccPrecisionPolicys")
    def cc_precision_policys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]]]:
        """
        CC Precision Protection List.
        """
        return pulumi.get(self, "cc_precision_policys")

    @cc_precision_policys.setter
    def cc_precision_policys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]]]):
        pulumi.set(self, "cc_precision_policys", value)

    @property
    @pulumi.getter(name="ccPrecisionReqLimits")
    def cc_precision_req_limits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]]]:
        """
        CC frequency throttling policy.
        """
        return pulumi.get(self, "cc_precision_req_limits")

    @cc_precision_req_limits.setter
    def cc_precision_req_limits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]]]):
        pulumi.set(self, "cc_precision_req_limits", value)

    @property
    @pulumi.getter
    def thresholds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]]]:
        """
        List of protection threshold configurations.
        """
        return pulumi.get(self, "thresholds")

    @thresholds.setter
    def thresholds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]]]):
        pulumi.set(self, "thresholds", value)


@pulumi.input_type
class _CCPolicyV2State:
    def __init__(__self__, *,
                 business: Optional[pulumi.Input[str]] = None,
                 cc_black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]]] = None,
                 cc_geo_ip_policys: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]]] = None,
                 cc_precision_policys: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]]] = None,
                 cc_precision_req_limits: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 thresholds: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]]] = None):
        """
        Input properties used for looking up and filtering CCPolicyV2 resources.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]] cc_black_white_ips: Blacklist and whitelist.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]] cc_geo_ip_policys: Details of the CC region blocking policy list.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]] cc_precision_policys: CC Precision Protection List.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]] cc_precision_req_limits: CC frequency throttling policy.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        :param pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]] thresholds: List of protection threshold configurations.
        """
        if business is not None:
            pulumi.set(__self__, "business", business)
        if cc_black_white_ips is not None:
            pulumi.set(__self__, "cc_black_white_ips", cc_black_white_ips)
        if cc_geo_ip_policys is not None:
            pulumi.set(__self__, "cc_geo_ip_policys", cc_geo_ip_policys)
        if cc_precision_policys is not None:
            pulumi.set(__self__, "cc_precision_policys", cc_precision_policys)
        if cc_precision_req_limits is not None:
            pulumi.set(__self__, "cc_precision_req_limits", cc_precision_req_limits)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if thresholds is not None:
            pulumi.set(__self__, "thresholds", thresholds)

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
    @pulumi.getter(name="ccBlackWhiteIps")
    def cc_black_white_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]]]:
        """
        Blacklist and whitelist.
        """
        return pulumi.get(self, "cc_black_white_ips")

    @cc_black_white_ips.setter
    def cc_black_white_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcBlackWhiteIpArgs']]]]):
        pulumi.set(self, "cc_black_white_ips", value)

    @property
    @pulumi.getter(name="ccGeoIpPolicys")
    def cc_geo_ip_policys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]]]:
        """
        Details of the CC region blocking policy list.
        """
        return pulumi.get(self, "cc_geo_ip_policys")

    @cc_geo_ip_policys.setter
    def cc_geo_ip_policys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcGeoIpPolicyArgs']]]]):
        pulumi.set(self, "cc_geo_ip_policys", value)

    @property
    @pulumi.getter(name="ccPrecisionPolicys")
    def cc_precision_policys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]]]:
        """
        CC Precision Protection List.
        """
        return pulumi.get(self, "cc_precision_policys")

    @cc_precision_policys.setter
    def cc_precision_policys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionPolicyArgs']]]]):
        pulumi.set(self, "cc_precision_policys", value)

    @property
    @pulumi.getter(name="ccPrecisionReqLimits")
    def cc_precision_req_limits(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]]]:
        """
        CC frequency throttling policy.
        """
        return pulumi.get(self, "cc_precision_req_limits")

    @cc_precision_req_limits.setter
    def cc_precision_req_limits(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2CcPrecisionReqLimitArgs']]]]):
        pulumi.set(self, "cc_precision_req_limits", value)

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

    @property
    @pulumi.getter
    def thresholds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]]]:
        """
        List of protection threshold configurations.
        """
        return pulumi.get(self, "thresholds")

    @thresholds.setter
    def thresholds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCPolicyV2ThresholdArgs']]]]):
        pulumi.set(self, "thresholds", value)


class CCPolicyV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 business: Optional[pulumi.Input[str]] = None,
                 cc_black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcBlackWhiteIpArgs']]]]] = None,
                 cc_geo_ip_policys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcGeoIpPolicyArgs']]]]] = None,
                 cc_precision_policys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionPolicyArgs']]]]] = None,
                 cc_precision_req_limits: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionReqLimitArgs']]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 thresholds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2ThresholdArgs']]]]] = None,
                 __props__=None):
        """
        Create a CCPolicyV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcBlackWhiteIpArgs']]]] cc_black_white_ips: Blacklist and whitelist.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcGeoIpPolicyArgs']]]] cc_geo_ip_policys: Details of the CC region blocking policy list.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionPolicyArgs']]]] cc_precision_policys: CC Precision Protection List.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionReqLimitArgs']]]] cc_precision_req_limits: CC frequency throttling policy.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2ThresholdArgs']]]] thresholds: List of protection threshold configurations.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CCPolicyV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CCPolicyV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CCPolicyV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CCPolicyV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 business: Optional[pulumi.Input[str]] = None,
                 cc_black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcBlackWhiteIpArgs']]]]] = None,
                 cc_geo_ip_policys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcGeoIpPolicyArgs']]]]] = None,
                 cc_precision_policys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionPolicyArgs']]]]] = None,
                 cc_precision_req_limits: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionReqLimitArgs']]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 thresholds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2ThresholdArgs']]]]] = None,
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
            __props__ = CCPolicyV2Args.__new__(CCPolicyV2Args)

            if business is None and not opts.urn:
                raise TypeError("Missing required property 'business'")
            __props__.__dict__["business"] = business
            __props__.__dict__["cc_black_white_ips"] = cc_black_white_ips
            __props__.__dict__["cc_geo_ip_policys"] = cc_geo_ip_policys
            __props__.__dict__["cc_precision_policys"] = cc_precision_policys
            __props__.__dict__["cc_precision_req_limits"] = cc_precision_req_limits
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            __props__.__dict__["thresholds"] = thresholds
        super(CCPolicyV2, __self__).__init__(
            'tencentcloud:Dayu/cCPolicyV2:CCPolicyV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            business: Optional[pulumi.Input[str]] = None,
            cc_black_white_ips: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcBlackWhiteIpArgs']]]]] = None,
            cc_geo_ip_policys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcGeoIpPolicyArgs']]]]] = None,
            cc_precision_policys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionPolicyArgs']]]]] = None,
            cc_precision_req_limits: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionReqLimitArgs']]]]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            thresholds: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2ThresholdArgs']]]]] = None) -> 'CCPolicyV2':
        """
        Get an existing CCPolicyV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] business: Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
               packet; net indicates anti-anti-ip pro version.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcBlackWhiteIpArgs']]]] cc_black_white_ips: Blacklist and whitelist.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcGeoIpPolicyArgs']]]] cc_geo_ip_policys: Details of the CC region blocking policy list.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionPolicyArgs']]]] cc_precision_policys: CC Precision Protection List.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2CcPrecisionReqLimitArgs']]]] cc_precision_req_limits: CC frequency throttling policy.
        :param pulumi.Input[str] resource_id: The ID of the resource instance.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCPolicyV2ThresholdArgs']]]] thresholds: List of protection threshold configurations.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CCPolicyV2State.__new__(_CCPolicyV2State)

        __props__.__dict__["business"] = business
        __props__.__dict__["cc_black_white_ips"] = cc_black_white_ips
        __props__.__dict__["cc_geo_ip_policys"] = cc_geo_ip_policys
        __props__.__dict__["cc_precision_policys"] = cc_precision_policys
        __props__.__dict__["cc_precision_req_limits"] = cc_precision_req_limits
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["thresholds"] = thresholds
        return CCPolicyV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def business(self) -> pulumi.Output[str]:
        """
        Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared
        packet; net indicates anti-anti-ip pro version.
        """
        return pulumi.get(self, "business")

    @property
    @pulumi.getter(name="ccBlackWhiteIps")
    def cc_black_white_ips(self) -> pulumi.Output[Optional[Sequence['outputs.CCPolicyV2CcBlackWhiteIp']]]:
        """
        Blacklist and whitelist.
        """
        return pulumi.get(self, "cc_black_white_ips")

    @property
    @pulumi.getter(name="ccGeoIpPolicys")
    def cc_geo_ip_policys(self) -> pulumi.Output[Optional[Sequence['outputs.CCPolicyV2CcGeoIpPolicy']]]:
        """
        Details of the CC region blocking policy list.
        """
        return pulumi.get(self, "cc_geo_ip_policys")

    @property
    @pulumi.getter(name="ccPrecisionPolicys")
    def cc_precision_policys(self) -> pulumi.Output[Optional[Sequence['outputs.CCPolicyV2CcPrecisionPolicy']]]:
        """
        CC Precision Protection List.
        """
        return pulumi.get(self, "cc_precision_policys")

    @property
    @pulumi.getter(name="ccPrecisionReqLimits")
    def cc_precision_req_limits(self) -> pulumi.Output[Optional[Sequence['outputs.CCPolicyV2CcPrecisionReqLimit']]]:
        """
        CC frequency throttling policy.
        """
        return pulumi.get(self, "cc_precision_req_limits")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource instance.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def thresholds(self) -> pulumi.Output[Optional[Sequence['outputs.CCPolicyV2Threshold']]]:
        """
        List of protection threshold configurations.
        """
        return pulumi.get(self, "thresholds")
