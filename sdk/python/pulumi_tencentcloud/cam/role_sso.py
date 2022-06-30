# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RoleSSOArgs', 'RoleSSO']

@pulumi.input_type
class RoleSSOArgs:
    def __init__(__self__, *,
                 client_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 identity_key: pulumi.Input[str],
                 identity_url: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RoleSSO resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_ids: Client ids.
        :param pulumi.Input[str] identity_key: Sign the public key.
        :param pulumi.Input[str] identity_url: Identity provider URL.
        :param pulumi.Input[str] description: The description of resource.
        :param pulumi.Input[str] name: The name of resource.
        """
        pulumi.set(__self__, "client_ids", client_ids)
        pulumi.set(__self__, "identity_key", identity_key)
        pulumi.set(__self__, "identity_url", identity_url)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="clientIds")
    def client_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Client ids.
        """
        return pulumi.get(self, "client_ids")

    @client_ids.setter
    def client_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "client_ids", value)

    @property
    @pulumi.getter(name="identityKey")
    def identity_key(self) -> pulumi.Input[str]:
        """
        Sign the public key.
        """
        return pulumi.get(self, "identity_key")

    @identity_key.setter
    def identity_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_key", value)

    @property
    @pulumi.getter(name="identityUrl")
    def identity_url(self) -> pulumi.Input[str]:
        """
        Identity provider URL.
        """
        return pulumi.get(self, "identity_url")

    @identity_url.setter
    def identity_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_url", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RoleSSOState:
    def __init__(__self__, *,
                 client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_key: Optional[pulumi.Input[str]] = None,
                 identity_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RoleSSO resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_ids: Client ids.
        :param pulumi.Input[str] description: The description of resource.
        :param pulumi.Input[str] identity_key: Sign the public key.
        :param pulumi.Input[str] identity_url: Identity provider URL.
        :param pulumi.Input[str] name: The name of resource.
        """
        if client_ids is not None:
            pulumi.set(__self__, "client_ids", client_ids)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if identity_key is not None:
            pulumi.set(__self__, "identity_key", identity_key)
        if identity_url is not None:
            pulumi.set(__self__, "identity_url", identity_url)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="clientIds")
    def client_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Client ids.
        """
        return pulumi.get(self, "client_ids")

    @client_ids.setter
    def client_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "client_ids", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="identityKey")
    def identity_key(self) -> Optional[pulumi.Input[str]]:
        """
        Sign the public key.
        """
        return pulumi.get(self, "identity_key")

    @identity_key.setter
    def identity_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_key", value)

    @property
    @pulumi.getter(name="identityUrl")
    def identity_url(self) -> Optional[pulumi.Input[str]]:
        """
        Identity provider URL.
        """
        return pulumi.get(self, "identity_url")

    @identity_url.setter
    def identity_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class RoleSSO(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_key: Optional[pulumi.Input[str]] = None,
                 identity_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a RoleSSO resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_ids: Client ids.
        :param pulumi.Input[str] description: The description of resource.
        :param pulumi.Input[str] identity_key: Sign the public key.
        :param pulumi.Input[str] identity_url: Identity provider URL.
        :param pulumi.Input[str] name: The name of resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleSSOArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RoleSSO resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RoleSSOArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleSSOArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_key: Optional[pulumi.Input[str]] = None,
                 identity_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = RoleSSOArgs.__new__(RoleSSOArgs)

            if client_ids is None and not opts.urn:
                raise TypeError("Missing required property 'client_ids'")
            __props__.__dict__["client_ids"] = client_ids
            __props__.__dict__["description"] = description
            if identity_key is None and not opts.urn:
                raise TypeError("Missing required property 'identity_key'")
            __props__.__dict__["identity_key"] = identity_key
            if identity_url is None and not opts.urn:
                raise TypeError("Missing required property 'identity_url'")
            __props__.__dict__["identity_url"] = identity_url
            __props__.__dict__["name"] = name
        super(RoleSSO, __self__).__init__(
            'tencentcloud:Cam/roleSSO:RoleSSO',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            identity_key: Optional[pulumi.Input[str]] = None,
            identity_url: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'RoleSSO':
        """
        Get an existing RoleSSO resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] client_ids: Client ids.
        :param pulumi.Input[str] description: The description of resource.
        :param pulumi.Input[str] identity_key: Sign the public key.
        :param pulumi.Input[str] identity_url: Identity provider URL.
        :param pulumi.Input[str] name: The name of resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleSSOState.__new__(_RoleSSOState)

        __props__.__dict__["client_ids"] = client_ids
        __props__.__dict__["description"] = description
        __props__.__dict__["identity_key"] = identity_key
        __props__.__dict__["identity_url"] = identity_url
        __props__.__dict__["name"] = name
        return RoleSSO(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientIds")
    def client_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Client ids.
        """
        return pulumi.get(self, "client_ids")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="identityKey")
    def identity_key(self) -> pulumi.Output[str]:
        """
        Sign the public key.
        """
        return pulumi.get(self, "identity_key")

    @property
    @pulumi.getter(name="identityUrl")
    def identity_url(self) -> pulumi.Output[str]:
        """
        Identity provider URL.
        """
        return pulumi.get(self, "identity_url")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of resource.
        """
        return pulumi.get(self, "name")

