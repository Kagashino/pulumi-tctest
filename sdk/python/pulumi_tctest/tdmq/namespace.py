# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NamespaceArgs', 'Namespace']

@pulumi.input_type
class NamespaceArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 environ_name: pulumi.Input[str],
                 msg_ttl: pulumi.Input[int],
                 remark: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Namespace resource.
        :param pulumi.Input[str] cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] environ_name: The name of namespace to be created.
        :param pulumi.Input[int] msg_ttl: The expiration time of unconsumed message.
        :param pulumi.Input[str] remark: Description of the namespace.
        :param pulumi.Input[Mapping[str, Any]] retention_policy: The Policy of message to retain.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "environ_name", environ_name)
        pulumi.set(__self__, "msg_ttl", msg_ttl)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The Dedicated Cluster Id.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="environName")
    def environ_name(self) -> pulumi.Input[str]:
        """
        The name of namespace to be created.
        """
        return pulumi.get(self, "environ_name")

    @environ_name.setter
    def environ_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "environ_name", value)

    @property
    @pulumi.getter(name="msgTtl")
    def msg_ttl(self) -> pulumi.Input[int]:
        """
        The expiration time of unconsumed message.
        """
        return pulumi.get(self, "msg_ttl")

    @msg_ttl.setter
    def msg_ttl(self, value: pulumi.Input[int]):
        pulumi.set(self, "msg_ttl", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the namespace.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The Policy of message to retain.
        """
        return pulumi.get(self, "retention_policy")

    @retention_policy.setter
    def retention_policy(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "retention_policy", value)


@pulumi.input_type
class _NamespaceState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 environ_name: Optional[pulumi.Input[str]] = None,
                 msg_ttl: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Namespace resources.
        :param pulumi.Input[str] cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] environ_name: The name of namespace to be created.
        :param pulumi.Input[int] msg_ttl: The expiration time of unconsumed message.
        :param pulumi.Input[str] remark: Description of the namespace.
        :param pulumi.Input[Mapping[str, Any]] retention_policy: The Policy of message to retain.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if environ_name is not None:
            pulumi.set(__self__, "environ_name", environ_name)
        if msg_ttl is not None:
            pulumi.set(__self__, "msg_ttl", msg_ttl)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Dedicated Cluster Id.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="environName")
    def environ_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of namespace to be created.
        """
        return pulumi.get(self, "environ_name")

    @environ_name.setter
    def environ_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environ_name", value)

    @property
    @pulumi.getter(name="msgTtl")
    def msg_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The expiration time of unconsumed message.
        """
        return pulumi.get(self, "msg_ttl")

    @msg_ttl.setter
    def msg_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "msg_ttl", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the namespace.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The Policy of message to retain.
        """
        return pulumi.get(self, "retention_policy")

    @retention_policy.setter
    def retention_policy(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "retention_policy", value)


class Namespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 environ_name: Optional[pulumi.Input[str]] = None,
                 msg_ttl: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Create a Namespace resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] environ_name: The name of namespace to be created.
        :param pulumi.Input[int] msg_ttl: The expiration time of unconsumed message.
        :param pulumi.Input[str] remark: Description of the namespace.
        :param pulumi.Input[Mapping[str, Any]] retention_policy: The Policy of message to retain.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Namespace resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 environ_name: Optional[pulumi.Input[str]] = None,
                 msg_ttl: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input[Mapping[str, Any]]] = None,
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
            __props__ = NamespaceArgs.__new__(NamespaceArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if environ_name is None and not opts.urn:
                raise TypeError("Missing required property 'environ_name'")
            __props__.__dict__["environ_name"] = environ_name
            if msg_ttl is None and not opts.urn:
                raise TypeError("Missing required property 'msg_ttl'")
            __props__.__dict__["msg_ttl"] = msg_ttl
            __props__.__dict__["remark"] = remark
            __props__.__dict__["retention_policy"] = retention_policy
        super(Namespace, __self__).__init__(
            'tctest:Tdmq/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            environ_name: Optional[pulumi.Input[str]] = None,
            msg_ttl: Optional[pulumi.Input[int]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            retention_policy: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Namespace':
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] environ_name: The name of namespace to be created.
        :param pulumi.Input[int] msg_ttl: The expiration time of unconsumed message.
        :param pulumi.Input[str] remark: Description of the namespace.
        :param pulumi.Input[Mapping[str, Any]] retention_policy: The Policy of message to retain.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NamespaceState.__new__(_NamespaceState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["environ_name"] = environ_name
        __props__.__dict__["msg_ttl"] = msg_ttl
        __props__.__dict__["remark"] = remark
        __props__.__dict__["retention_policy"] = retention_policy
        return Namespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The Dedicated Cluster Id.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="environName")
    def environ_name(self) -> pulumi.Output[str]:
        """
        The name of namespace to be created.
        """
        return pulumi.get(self, "environ_name")

    @property
    @pulumi.getter(name="msgTtl")
    def msg_ttl(self) -> pulumi.Output[int]:
        """
        The expiration time of unconsumed message.
        """
        return pulumi.get(self, "msg_ttl")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the namespace.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The Policy of message to retain.
        """
        return pulumi.get(self, "retention_policy")

