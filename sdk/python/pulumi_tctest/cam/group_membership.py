# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupMembershipArgs', 'GroupMembership']

@pulumi.input_type
class GroupMembershipArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[str],
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a GroupMembership resource.
        :param pulumi.Input[str] group_id: ID of CAM group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: ID set of the CAM group members.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: User name set as ID of the CAM group members.
        """
        pulumi.set(__self__, "group_id", group_id)
        if user_ids is not None:
            warnings.warn("""It has been deprecated from version 1.59.5. Use `user_names` instead.""", DeprecationWarning)
            pulumi.log.warn("""user_ids is deprecated: It has been deprecated from version 1.59.5. Use `user_names` instead.""")
        if user_ids is not None:
            pulumi.set(__self__, "user_ids", user_ids)
        if user_names is not None:
            pulumi.set(__self__, "user_names", user_names)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        ID of CAM group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ID set of the CAM group members.
        """
        return pulumi.get(self, "user_ids")

    @user_ids.setter
    def user_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_ids", value)

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User name set as ID of the CAM group members.
        """
        return pulumi.get(self, "user_names")

    @user_names.setter
    def user_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_names", value)


@pulumi.input_type
class _GroupMembershipState:
    def __init__(__self__, *,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering GroupMembership resources.
        :param pulumi.Input[str] group_id: ID of CAM group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: ID set of the CAM group members.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: User name set as ID of the CAM group members.
        """
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if user_ids is not None:
            warnings.warn("""It has been deprecated from version 1.59.5. Use `user_names` instead.""", DeprecationWarning)
            pulumi.log.warn("""user_ids is deprecated: It has been deprecated from version 1.59.5. Use `user_names` instead.""")
        if user_ids is not None:
            pulumi.set(__self__, "user_ids", user_ids)
        if user_names is not None:
            pulumi.set(__self__, "user_names", user_names)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of CAM group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ID set of the CAM group members.
        """
        return pulumi.get(self, "user_ids")

    @user_ids.setter
    def user_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_ids", value)

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User name set as ID of the CAM group members.
        """
        return pulumi.get(self, "user_names")

    @user_names.setter
    def user_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_names", value)


class GroupMembership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a GroupMembership resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: ID of CAM group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: ID set of the CAM group members.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: User name set as ID of the CAM group members.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a GroupMembership resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param GroupMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = GroupMembershipArgs.__new__(GroupMembershipArgs)

            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if user_ids is not None and not opts.urn:
                warnings.warn("""It has been deprecated from version 1.59.5. Use `user_names` instead.""", DeprecationWarning)
                pulumi.log.warn("""user_ids is deprecated: It has been deprecated from version 1.59.5. Use `user_names` instead.""")
            __props__.__dict__["user_ids"] = user_ids
            __props__.__dict__["user_names"] = user_names
        super(GroupMembership, __self__).__init__(
            'tctest:Cam/groupMembership:GroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'GroupMembership':
        """
        Get an existing GroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: ID of CAM group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: ID set of the CAM group members.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_names: User name set as ID of the CAM group members.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupMembershipState.__new__(_GroupMembershipState)

        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["user_ids"] = user_ids
        __props__.__dict__["user_names"] = user_names
        return GroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        ID of CAM group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        ID set of the CAM group members.
        """
        return pulumi.get(self, "user_ids")

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        User name set as ID of the CAM group members.
        """
        return pulumi.get(self, "user_names")

