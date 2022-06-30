# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnapshotPolicyAttachmentArgs', 'SnapshotPolicyAttachment']

@pulumi.input_type
class SnapshotPolicyAttachmentArgs:
    def __init__(__self__, *,
                 snapshot_policy_id: pulumi.Input[str],
                 storage_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a SnapshotPolicyAttachment resource.
        :param pulumi.Input[str] snapshot_policy_id: ID of CBS snapshot policy.
        :param pulumi.Input[str] storage_id: ID of CBS.
        """
        pulumi.set(__self__, "snapshot_policy_id", snapshot_policy_id)
        pulumi.set(__self__, "storage_id", storage_id)

    @property
    @pulumi.getter(name="snapshotPolicyId")
    def snapshot_policy_id(self) -> pulumi.Input[str]:
        """
        ID of CBS snapshot policy.
        """
        return pulumi.get(self, "snapshot_policy_id")

    @snapshot_policy_id.setter
    def snapshot_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_policy_id", value)

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> pulumi.Input[str]:
        """
        ID of CBS.
        """
        return pulumi.get(self, "storage_id")

    @storage_id.setter
    def storage_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_id", value)


@pulumi.input_type
class _SnapshotPolicyAttachmentState:
    def __init__(__self__, *,
                 snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnapshotPolicyAttachment resources.
        :param pulumi.Input[str] snapshot_policy_id: ID of CBS snapshot policy.
        :param pulumi.Input[str] storage_id: ID of CBS.
        """
        if snapshot_policy_id is not None:
            pulumi.set(__self__, "snapshot_policy_id", snapshot_policy_id)
        if storage_id is not None:
            pulumi.set(__self__, "storage_id", storage_id)

    @property
    @pulumi.getter(name="snapshotPolicyId")
    def snapshot_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of CBS snapshot policy.
        """
        return pulumi.get(self, "snapshot_policy_id")

    @snapshot_policy_id.setter
    def snapshot_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_policy_id", value)

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of CBS.
        """
        return pulumi.get(self, "storage_id")

    @storage_id.setter
    def storage_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_id", value)


class SnapshotPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SnapshotPolicyAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] snapshot_policy_id: ID of CBS snapshot policy.
        :param pulumi.Input[str] storage_id: ID of CBS.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SnapshotPolicyAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SnapshotPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = SnapshotPolicyAttachmentArgs.__new__(SnapshotPolicyAttachmentArgs)

            if snapshot_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_policy_id'")
            __props__.__dict__["snapshot_policy_id"] = snapshot_policy_id
            if storage_id is None and not opts.urn:
                raise TypeError("Missing required property 'storage_id'")
            __props__.__dict__["storage_id"] = storage_id
        super(SnapshotPolicyAttachment, __self__).__init__(
            'tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            snapshot_policy_id: Optional[pulumi.Input[str]] = None,
            storage_id: Optional[pulumi.Input[str]] = None) -> 'SnapshotPolicyAttachment':
        """
        Get an existing SnapshotPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] snapshot_policy_id: ID of CBS snapshot policy.
        :param pulumi.Input[str] storage_id: ID of CBS.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotPolicyAttachmentState.__new__(_SnapshotPolicyAttachmentState)

        __props__.__dict__["snapshot_policy_id"] = snapshot_policy_id
        __props__.__dict__["storage_id"] = storage_id
        return SnapshotPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="snapshotPolicyId")
    def snapshot_policy_id(self) -> pulumi.Output[str]:
        """
        ID of CBS snapshot policy.
        """
        return pulumi.get(self, "snapshot_policy_id")

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> pulumi.Output[str]:
        """
        ID of CBS.
        """
        return pulumi.get(self, "storage_id")

