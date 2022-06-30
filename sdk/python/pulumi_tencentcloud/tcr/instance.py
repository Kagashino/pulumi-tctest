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
                 instance_type: pulumi.Input[str],
                 delete_bucket: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_public_operation: Optional[pulumi.Input[bool]] = None,
                 security_policies: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] instance_type: TCR types. Valid values are: `standard`, `basic`, `premium`.
        :param pulumi.Input[bool] delete_bucket: Indicate to delete the COS bucket which is auto-created with the instance or not.
        :param pulumi.Input[str] name: Name of the TCR instance.
        :param pulumi.Input[bool] open_public_operation: Control public network access.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]] security_policies: Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
        :param pulumi.Input[Mapping[str, Any]] tags: The available tags within this TCR instance.
        """
        pulumi.set(__self__, "instance_type", instance_type)
        if delete_bucket is not None:
            pulumi.set(__self__, "delete_bucket", delete_bucket)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if open_public_operation is not None:
            pulumi.set(__self__, "open_public_operation", open_public_operation)
        if security_policies is not None:
            pulumi.set(__self__, "security_policies", security_policies)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        TCR types. Valid values are: `standard`, `basic`, `premium`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="deleteBucket")
    def delete_bucket(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate to delete the COS bucket which is auto-created with the instance or not.
        """
        return pulumi.get(self, "delete_bucket")

    @delete_bucket.setter
    def delete_bucket(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_bucket", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the TCR instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openPublicOperation")
    def open_public_operation(self) -> Optional[pulumi.Input[bool]]:
        """
        Control public network access.
        """
        return pulumi.get(self, "open_public_operation")

    @open_public_operation.setter
    def open_public_operation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "open_public_operation", value)

    @property
    @pulumi.getter(name="securityPolicies")
    def security_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]]]:
        """
        Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
        """
        return pulumi.get(self, "security_policies")

    @security_policies.setter
    def security_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]]]):
        pulumi.set(self, "security_policies", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The available tags within this TCR instance.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 delete_bucket: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 internal_end_point: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_public_operation: Optional[pulumi.Input[bool]] = None,
                 public_domain: Optional[pulumi.Input[str]] = None,
                 public_status: Optional[pulumi.Input[str]] = None,
                 security_policies: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[bool] delete_bucket: Indicate to delete the COS bucket which is auto-created with the instance or not.
        :param pulumi.Input[str] instance_type: TCR types. Valid values are: `standard`, `basic`, `premium`.
        :param pulumi.Input[str] internal_end_point: Internal address for access of the TCR instance.
        :param pulumi.Input[str] name: Name of the TCR instance.
        :param pulumi.Input[bool] open_public_operation: Control public network access.
        :param pulumi.Input[str] public_domain: Public address for access of the TCR instance.
        :param pulumi.Input[str] public_status: Status of the TCR instance public network access.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]] security_policies: Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
        :param pulumi.Input[str] status: Status of the TCR instance.
        :param pulumi.Input[Mapping[str, Any]] tags: The available tags within this TCR instance.
        """
        if delete_bucket is not None:
            pulumi.set(__self__, "delete_bucket", delete_bucket)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if internal_end_point is not None:
            pulumi.set(__self__, "internal_end_point", internal_end_point)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if open_public_operation is not None:
            pulumi.set(__self__, "open_public_operation", open_public_operation)
        if public_domain is not None:
            pulumi.set(__self__, "public_domain", public_domain)
        if public_status is not None:
            pulumi.set(__self__, "public_status", public_status)
        if security_policies is not None:
            pulumi.set(__self__, "security_policies", security_policies)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="deleteBucket")
    def delete_bucket(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate to delete the COS bucket which is auto-created with the instance or not.
        """
        return pulumi.get(self, "delete_bucket")

    @delete_bucket.setter
    def delete_bucket(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_bucket", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        TCR types. Valid values are: `standard`, `basic`, `premium`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="internalEndPoint")
    def internal_end_point(self) -> Optional[pulumi.Input[str]]:
        """
        Internal address for access of the TCR instance.
        """
        return pulumi.get(self, "internal_end_point")

    @internal_end_point.setter
    def internal_end_point(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_end_point", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the TCR instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openPublicOperation")
    def open_public_operation(self) -> Optional[pulumi.Input[bool]]:
        """
        Control public network access.
        """
        return pulumi.get(self, "open_public_operation")

    @open_public_operation.setter
    def open_public_operation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "open_public_operation", value)

    @property
    @pulumi.getter(name="publicDomain")
    def public_domain(self) -> Optional[pulumi.Input[str]]:
        """
        Public address for access of the TCR instance.
        """
        return pulumi.get(self, "public_domain")

    @public_domain.setter
    def public_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_domain", value)

    @property
    @pulumi.getter(name="publicStatus")
    def public_status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the TCR instance public network access.
        """
        return pulumi.get(self, "public_status")

    @public_status.setter
    def public_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_status", value)

    @property
    @pulumi.getter(name="securityPolicies")
    def security_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]]]:
        """
        Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
        """
        return pulumi.get(self, "security_policies")

    @security_policies.setter
    def security_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityPolicyArgs']]]]):
        pulumi.set(self, "security_policies", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the TCR instance.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The available tags within this TCR instance.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_bucket: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_public_operation: Optional[pulumi.Input[bool]] = None,
                 security_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityPolicyArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Create a Instance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_bucket: Indicate to delete the COS bucket which is auto-created with the instance or not.
        :param pulumi.Input[str] instance_type: TCR types. Valid values are: `standard`, `basic`, `premium`.
        :param pulumi.Input[str] name: Name of the TCR instance.
        :param pulumi.Input[bool] open_public_operation: Control public network access.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityPolicyArgs']]]] security_policies: Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
        :param pulumi.Input[Mapping[str, Any]] tags: The available tags within this TCR instance.
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
                 delete_bucket: Optional[pulumi.Input[bool]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 open_public_operation: Optional[pulumi.Input[bool]] = None,
                 security_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityPolicyArgs']]]]] = None,
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
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["delete_bucket"] = delete_bucket
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["name"] = name
            __props__.__dict__["open_public_operation"] = open_public_operation
            __props__.__dict__["security_policies"] = security_policies
            __props__.__dict__["tags"] = tags
            __props__.__dict__["internal_end_point"] = None
            __props__.__dict__["public_domain"] = None
            __props__.__dict__["public_status"] = None
            __props__.__dict__["status"] = None
        super(Instance, __self__).__init__(
            'tencentcloud:Tcr/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delete_bucket: Optional[pulumi.Input[bool]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            internal_end_point: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            open_public_operation: Optional[pulumi.Input[bool]] = None,
            public_domain: Optional[pulumi.Input[str]] = None,
            public_status: Optional[pulumi.Input[str]] = None,
            security_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityPolicyArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_bucket: Indicate to delete the COS bucket which is auto-created with the instance or not.
        :param pulumi.Input[str] instance_type: TCR types. Valid values are: `standard`, `basic`, `premium`.
        :param pulumi.Input[str] internal_end_point: Internal address for access of the TCR instance.
        :param pulumi.Input[str] name: Name of the TCR instance.
        :param pulumi.Input[bool] open_public_operation: Control public network access.
        :param pulumi.Input[str] public_domain: Public address for access of the TCR instance.
        :param pulumi.Input[str] public_status: Status of the TCR instance public network access.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityPolicyArgs']]]] security_policies: Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
        :param pulumi.Input[str] status: Status of the TCR instance.
        :param pulumi.Input[Mapping[str, Any]] tags: The available tags within this TCR instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["delete_bucket"] = delete_bucket
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["internal_end_point"] = internal_end_point
        __props__.__dict__["name"] = name
        __props__.__dict__["open_public_operation"] = open_public_operation
        __props__.__dict__["public_domain"] = public_domain
        __props__.__dict__["public_status"] = public_status
        __props__.__dict__["security_policies"] = security_policies
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteBucket")
    def delete_bucket(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate to delete the COS bucket which is auto-created with the instance or not.
        """
        return pulumi.get(self, "delete_bucket")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        TCR types. Valid values are: `standard`, `basic`, `premium`.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="internalEndPoint")
    def internal_end_point(self) -> pulumi.Output[str]:
        """
        Internal address for access of the TCR instance.
        """
        return pulumi.get(self, "internal_end_point")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the TCR instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openPublicOperation")
    def open_public_operation(self) -> pulumi.Output[Optional[bool]]:
        """
        Control public network access.
        """
        return pulumi.get(self, "open_public_operation")

    @property
    @pulumi.getter(name="publicDomain")
    def public_domain(self) -> pulumi.Output[str]:
        """
        Public address for access of the TCR instance.
        """
        return pulumi.get(self, "public_domain")

    @property
    @pulumi.getter(name="publicStatus")
    def public_status(self) -> pulumi.Output[str]:
        """
        Status of the TCR instance public network access.
        """
        return pulumi.get(self, "public_status")

    @property
    @pulumi.getter(name="securityPolicies")
    def security_policies(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceSecurityPolicy']]]:
        """
        Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
        """
        return pulumi.get(self, "security_policies")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the TCR instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The available tags within this TCR instance.
        """
        return pulumi.get(self, "tags")

