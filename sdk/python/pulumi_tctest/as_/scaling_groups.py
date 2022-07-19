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
    'ScalingGroupsResult',
    'AwaitableScalingGroupsResult',
    'scaling_groups',
    'scaling_groups_output',
]

@pulumi.output_type
class ScalingGroupsResult:
    """
    A collection of values returned by ScalingGroups.
    """
    def __init__(__self__, configuration_id=None, id=None, result_output_file=None, scaling_group_id=None, scaling_group_lists=None, scaling_group_name=None, tags=None):
        if configuration_id and not isinstance(configuration_id, str):
            raise TypeError("Expected argument 'configuration_id' to be a str")
        pulumi.set(__self__, "configuration_id", configuration_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if scaling_group_id and not isinstance(scaling_group_id, str):
            raise TypeError("Expected argument 'scaling_group_id' to be a str")
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        if scaling_group_lists and not isinstance(scaling_group_lists, list):
            raise TypeError("Expected argument 'scaling_group_lists' to be a list")
        pulumi.set(__self__, "scaling_group_lists", scaling_group_lists)
        if scaling_group_name and not isinstance(scaling_group_name, str):
            raise TypeError("Expected argument 'scaling_group_name' to be a str")
        pulumi.set(__self__, "scaling_group_name", scaling_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="configurationId")
    def configuration_id(self) -> Optional[str]:
        return pulumi.get(self, "configuration_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[str]:
        return pulumi.get(self, "scaling_group_id")

    @property
    @pulumi.getter(name="scalingGroupLists")
    def scaling_group_lists(self) -> Sequence['outputs.ScalingGroupsScalingGroupListResult']:
        return pulumi.get(self, "scaling_group_lists")

    @property
    @pulumi.getter(name="scalingGroupName")
    def scaling_group_name(self) -> Optional[str]:
        return pulumi.get(self, "scaling_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableScalingGroupsResult(ScalingGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalingGroupsResult(
            configuration_id=self.configuration_id,
            id=self.id,
            result_output_file=self.result_output_file,
            scaling_group_id=self.scaling_group_id,
            scaling_group_lists=self.scaling_group_lists,
            scaling_group_name=self.scaling_group_name,
            tags=self.tags)


def scaling_groups(configuration_id: Optional[str] = None,
                   result_output_file: Optional[str] = None,
                   scaling_group_id: Optional[str] = None,
                   scaling_group_name: Optional[str] = None,
                   tags: Optional[Mapping[str, Any]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalingGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['configurationId'] = configuration_id
    __args__['resultOutputFile'] = result_output_file
    __args__['scalingGroupId'] = scaling_group_id
    __args__['scalingGroupName'] = scaling_group_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:As/scalingGroups:ScalingGroups', __args__, opts=opts, typ=ScalingGroupsResult).value

    return AwaitableScalingGroupsResult(
        configuration_id=__ret__.configuration_id,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        scaling_group_id=__ret__.scaling_group_id,
        scaling_group_lists=__ret__.scaling_group_lists,
        scaling_group_name=__ret__.scaling_group_name,
        tags=__ret__.tags)


@_utilities.lift_output_func(scaling_groups)
def scaling_groups_output(configuration_id: Optional[pulumi.Input[Optional[str]]] = None,
                          result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          scaling_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                          scaling_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                          tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ScalingGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
