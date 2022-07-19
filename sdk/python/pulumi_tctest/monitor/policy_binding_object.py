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

__all__ = ['PolicyBindingObjectArgs', 'PolicyBindingObject']

@pulumi.input_type
class PolicyBindingObjectArgs:
    def __init__(__self__, *,
                 dimensions: pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]],
                 policy_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a PolicyBindingObject resource.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]] dimensions: A list objects. Each element contains the following attributes:
        :param pulumi.Input[str] policy_id: Alarm policy ID for binding objects.
        """
        pulumi.set(__self__, "dimensions", dimensions)
        pulumi.set(__self__, "policy_id", policy_id)

    @property
    @pulumi.getter
    def dimensions(self) -> pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]]:
        """
        A list objects. Each element contains the following attributes:
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]]):
        pulumi.set(self, "dimensions", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        Alarm policy ID for binding objects.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)


@pulumi.input_type
class _PolicyBindingObjectState:
    def __init__(__self__, *,
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyBindingObject resources.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]] dimensions: A list objects. Each element contains the following attributes:
        :param pulumi.Input[str] policy_id: Alarm policy ID for binding objects.
        """
        if dimensions is not None:
            pulumi.set(__self__, "dimensions", dimensions)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]]]:
        """
        A list objects. Each element contains the following attributes:
        """
        return pulumi.get(self, "dimensions")

    @dimensions.setter
    def dimensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyBindingObjectDimensionArgs']]]]):
        pulumi.set(self, "dimensions", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Alarm policy ID for binding objects.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)


class PolicyBindingObject(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyBindingObjectDimensionArgs']]]]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PolicyBindingObject resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyBindingObjectDimensionArgs']]]] dimensions: A list objects. Each element contains the following attributes:
        :param pulumi.Input[str] policy_id: Alarm policy ID for binding objects.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyBindingObjectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PolicyBindingObject resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PolicyBindingObjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyBindingObjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dimensions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyBindingObjectDimensionArgs']]]]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = PolicyBindingObjectArgs.__new__(PolicyBindingObjectArgs)

            if dimensions is None and not opts.urn:
                raise TypeError("Missing required property 'dimensions'")
            __props__.__dict__["dimensions"] = dimensions
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
        super(PolicyBindingObject, __self__).__init__(
            'tctest:Monitor/policyBindingObject:PolicyBindingObject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dimensions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyBindingObjectDimensionArgs']]]]] = None,
            policy_id: Optional[pulumi.Input[str]] = None) -> 'PolicyBindingObject':
        """
        Get an existing PolicyBindingObject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyBindingObjectDimensionArgs']]]] dimensions: A list objects. Each element contains the following attributes:
        :param pulumi.Input[str] policy_id: Alarm policy ID for binding objects.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyBindingObjectState.__new__(_PolicyBindingObjectState)

        __props__.__dict__["dimensions"] = dimensions
        __props__.__dict__["policy_id"] = policy_id
        return PolicyBindingObject(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dimensions(self) -> pulumi.Output[Sequence['outputs.PolicyBindingObjectDimension']]:
        """
        A list objects. Each element contains the following attributes:
        """
        return pulumi.get(self, "dimensions")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        Alarm policy ID for binding objects.
        """
        return pulumi.get(self, "policy_id")

