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

__all__ = [
    'GroupsResult',
    'AwaitableGroupsResult',
    'groups',
    'groups_output',
]

@pulumi.output_type
class GroupsResult:
    """
    A collection of values returned by Groups.
    """
    def __init__(__self__, id=None, name=None, placement_group_id=None, placement_group_lists=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if placement_group_id and not isinstance(placement_group_id, str):
            raise TypeError("Expected argument 'placement_group_id' to be a str")
        pulumi.set(__self__, "placement_group_id", placement_group_id)
        if placement_group_lists and not isinstance(placement_group_lists, list):
            raise TypeError("Expected argument 'placement_group_lists' to be a list")
        pulumi.set(__self__, "placement_group_lists", placement_group_lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> Optional[str]:
        return pulumi.get(self, "placement_group_id")

    @property
    @pulumi.getter(name="placementGroupLists")
    def placement_group_lists(self) -> Sequence['outputs.GroupsPlacementGroupListResult']:
        return pulumi.get(self, "placement_group_lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGroupsResult(GroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GroupsResult(
            id=self.id,
            name=self.name,
            placement_group_id=self.placement_group_id,
            placement_group_lists=self.placement_group_lists,
            result_output_file=self.result_output_file)


def groups(name: Optional[str] = None,
           placement_group_id: Optional[str] = None,
           result_output_file: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['placementGroupId'] = placement_group_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Placement/groups:Groups', __args__, opts=opts, typ=GroupsResult).value

    return AwaitableGroupsResult(
        id=__ret__.id,
        name=__ret__.name,
        placement_group_id=__ret__.placement_group_id,
        placement_group_lists=__ret__.placement_group_lists,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(groups)
def groups_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                  placement_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
