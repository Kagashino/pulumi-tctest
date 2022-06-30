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

__all__ = ['CCHttpPolicyArgs', 'CCHttpPolicy']

@pulumi.input_type
class CCHttpPolicyArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 resource_type: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rule_lists: Optional[pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]]] = None,
                 smode: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a CCHttpPolicy resource.
        :param pulumi.Input[str] resource_id: ID of the resource that the CC self-define http policy works for.
        :param pulumi.Input[str] resource_type: Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
               `net`.
        :param pulumi.Input[str] action: Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
        :param pulumi.Input[int] frequency: Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
        :param pulumi.Input[str] ip: Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
               set one.
        :param pulumi.Input[str] name: Name of the CC self-define http policy. Length should between 1 and 20.
        :param pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]] rule_lists: Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
        :param pulumi.Input[str] smode: Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
               one.
        :param pulumi.Input[bool] switch: Indicate the CC self-define http policy takes effect or not.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if frequency is not None:
            pulumi.set(__self__, "frequency", frequency)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rule_lists is not None:
            pulumi.set(__self__, "rule_lists", rule_lists)
        if smode is not None:
            pulumi.set(__self__, "smode", smode)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        ID of the resource that the CC self-define http policy works for.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
        `net`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def frequency(self) -> Optional[pulumi.Input[int]]:
        """
        Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
        set one.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CC self-define http policy. Length should between 1 and 20.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ruleLists")
    def rule_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]]]:
        """
        Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
        """
        return pulumi.get(self, "rule_lists")

    @rule_lists.setter
    def rule_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]]]):
        pulumi.set(self, "rule_lists", value)

    @property
    @pulumi.getter
    def smode(self) -> Optional[pulumi.Input[str]]:
        """
        Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
        one.
        """
        return pulumi.get(self, "smode")

    @smode.setter
    def smode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "smode", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate the CC self-define http policy takes effect or not.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "switch", value)


@pulumi.input_type
class _CCHttpPolicyState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 rule_lists: Optional[pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]]] = None,
                 smode: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering CCHttpPolicy resources.
        :param pulumi.Input[str] action: Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
        :param pulumi.Input[str] create_time: Create time of the CC self-define http policy.
        :param pulumi.Input[int] frequency: Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
        :param pulumi.Input[str] ip: Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
               set one.
        :param pulumi.Input[str] name: Name of the CC self-define http policy. Length should between 1 and 20.
        :param pulumi.Input[str] policy_id: Id of the CC self-define http policy.
        :param pulumi.Input[str] resource_id: ID of the resource that the CC self-define http policy works for.
        :param pulumi.Input[str] resource_type: Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
               `net`.
        :param pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]] rule_lists: Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
        :param pulumi.Input[str] smode: Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
               one.
        :param pulumi.Input[bool] switch: Indicate the CC self-define http policy takes effect or not.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if frequency is not None:
            pulumi.set(__self__, "frequency", frequency)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if rule_lists is not None:
            pulumi.set(__self__, "rule_lists", rule_lists)
        if smode is not None:
            pulumi.set(__self__, "smode", smode)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the CC self-define http policy.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def frequency(self) -> Optional[pulumi.Input[int]]:
        """
        Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
        """
        return pulumi.get(self, "frequency")

    @frequency.setter
    def frequency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "frequency", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
        set one.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CC self-define http policy. Length should between 1 and 20.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of the CC self-define http policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the resource that the CC self-define http policy works for.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
        `net`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="ruleLists")
    def rule_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]]]:
        """
        Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
        """
        return pulumi.get(self, "rule_lists")

    @rule_lists.setter
    def rule_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CCHttpPolicyRuleListArgs']]]]):
        pulumi.set(self, "rule_lists", value)

    @property
    @pulumi.getter
    def smode(self) -> Optional[pulumi.Input[str]]:
        """
        Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
        one.
        """
        return pulumi.get(self, "smode")

    @smode.setter
    def smode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "smode", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate the CC self-define http policy takes effect or not.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "switch", value)


class CCHttpPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 rule_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCHttpPolicyRuleListArgs']]]]] = None,
                 smode: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a CCHttpPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
        :param pulumi.Input[int] frequency: Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
        :param pulumi.Input[str] ip: Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
               set one.
        :param pulumi.Input[str] name: Name of the CC self-define http policy. Length should between 1 and 20.
        :param pulumi.Input[str] resource_id: ID of the resource that the CC self-define http policy works for.
        :param pulumi.Input[str] resource_type: Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
               `net`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCHttpPolicyRuleListArgs']]]] rule_lists: Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
        :param pulumi.Input[str] smode: Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
               one.
        :param pulumi.Input[bool] switch: Indicate the CC self-define http policy takes effect or not.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CCHttpPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CCHttpPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CCHttpPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CCHttpPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 frequency: Optional[pulumi.Input[int]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 rule_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCHttpPolicyRuleListArgs']]]]] = None,
                 smode: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[bool]] = None,
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
            __props__ = CCHttpPolicyArgs.__new__(CCHttpPolicyArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["frequency"] = frequency
            __props__.__dict__["ip"] = ip
            __props__.__dict__["name"] = name
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["rule_lists"] = rule_lists
            __props__.__dict__["smode"] = smode
            __props__.__dict__["switch"] = switch
            __props__.__dict__["create_time"] = None
            __props__.__dict__["policy_id"] = None
        super(CCHttpPolicy, __self__).__init__(
            'tencentcloud:Dayu/cCHttpPolicy:CCHttpPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            frequency: Optional[pulumi.Input[int]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            rule_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCHttpPolicyRuleListArgs']]]]] = None,
            smode: Optional[pulumi.Input[str]] = None,
            switch: Optional[pulumi.Input[bool]] = None) -> 'CCHttpPolicy':
        """
        Get an existing CCHttpPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
        :param pulumi.Input[str] create_time: Create time of the CC self-define http policy.
        :param pulumi.Input[int] frequency: Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
        :param pulumi.Input[str] ip: Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
               set one.
        :param pulumi.Input[str] name: Name of the CC self-define http policy. Length should between 1 and 20.
        :param pulumi.Input[str] policy_id: Id of the CC self-define http policy.
        :param pulumi.Input[str] resource_id: ID of the resource that the CC self-define http policy works for.
        :param pulumi.Input[str] resource_type: Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
               `net`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CCHttpPolicyRuleListArgs']]]] rule_lists: Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
        :param pulumi.Input[str] smode: Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
               one.
        :param pulumi.Input[bool] switch: Indicate the CC self-define http policy takes effect or not.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CCHttpPolicyState.__new__(_CCHttpPolicyState)

        __props__.__dict__["action"] = action
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["frequency"] = frequency
        __props__.__dict__["ip"] = ip
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["rule_lists"] = rule_lists
        __props__.__dict__["smode"] = smode
        __props__.__dict__["switch"] = switch
        return CCHttpPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Action mode, only valid when `smode` is `matching`. Valid values are `alg` and `drop`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the CC self-define http policy.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Output[int]:
        """
        Max frequency per minute, only valid when `smode` is `speedlimit`, the valid value ranges from 1 to 10000.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        Ip of the CC self-define http policy, only valid when `resource_type` is `bgp-multip`. The num of list items can only be
        set one.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the CC self-define http policy. Length should between 1 and 20.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        Id of the CC self-define http policy.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        ID of the resource that the CC self-define http policy works for.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        Type of the resource that the CC self-define http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and
        `net`.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="ruleLists")
    def rule_lists(self) -> pulumi.Output[Optional[Sequence['outputs.CCHttpPolicyRuleList']]]:
        """
        Rule list of the CC self-define http policy, only valid when `smode` is `matching`.
        """
        return pulumi.get(self, "rule_lists")

    @property
    @pulumi.getter
    def smode(self) -> pulumi.Output[Optional[str]]:
        """
        Match mode, and valid values are `matching`, `speedlimit`. Note: the speed limit type CC self-define policy can only set
        one.
        """
        return pulumi.get(self, "smode")

    @property
    @pulumi.getter
    def switch(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate the CC self-define http policy takes effect or not.
        """
        return pulumi.get(self, "switch")
