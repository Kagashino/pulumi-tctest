# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PairArgs', 'Pair']

@pulumi.input_type
class PairArgs:
    def __init__(__self__, *,
                 key_name: pulumi.Input[str],
                 public_key: pulumi.Input[str],
                 project_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Pair resource.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one TencentCloud account.
        :param pulumi.Input[str] public_key: You can import an existing public key and using TencentCloud key pair to manage it.
        :param pulumi.Input[int] project_id: Specifys to which project the key pair belongs.
        """
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "public_key", public_key)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Input[str]:
        """
        The key pair's name. It is the only in one TencentCloud account.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Input[str]:
        """
        You can import an existing public key and using TencentCloud key pair to manage it.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Specifys to which project the key pair belongs.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _PairState:
    def __init__(__self__, *,
                 key_name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Pair resources.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one TencentCloud account.
        :param pulumi.Input[int] project_id: Specifys to which project the key pair belongs.
        :param pulumi.Input[str] public_key: You can import an existing public key and using TencentCloud key pair to manage it.
        """
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The key pair's name. It is the only in one TencentCloud account.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Specifys to which project the key pair belongs.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        You can import an existing public key and using TencentCloud key pair to manage it.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)


class Pair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Pair resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one TencentCloud account.
        :param pulumi.Input[int] project_id: Specifys to which project the key pair belongs.
        :param pulumi.Input[str] public_key: You can import an existing public key and using TencentCloud key pair to manage it.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PairArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Pair resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
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
            __props__ = PairArgs.__new__(PairArgs)

            if key_name is None and not opts.urn:
                raise TypeError("Missing required property 'key_name'")
            __props__.__dict__["key_name"] = key_name
            __props__.__dict__["project_id"] = project_id
            if public_key is None and not opts.urn:
                raise TypeError("Missing required property 'public_key'")
            __props__.__dict__["public_key"] = public_key
        super(Pair, __self__).__init__(
            'tctest:Key/pair:Pair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            public_key: Optional[pulumi.Input[str]] = None) -> 'Pair':
        """
        Get an existing Pair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_name: The key pair's name. It is the only in one TencentCloud account.
        :param pulumi.Input[int] project_id: Specifys to which project the key pair belongs.
        :param pulumi.Input[str] public_key: You can import an existing public key and using TencentCloud key pair to manage it.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PairState.__new__(_PairState)

        __props__.__dict__["key_name"] = key_name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["public_key"] = public_key
        return Pair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[str]:
        """
        The key pair's name. It is the only in one TencentCloud account.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[int]]:
        """
        Specifys to which project the key pair belongs.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        You can import an existing public key and using TencentCloud key pair to manage it.
        """
        return pulumi.get(self, "public_key")

