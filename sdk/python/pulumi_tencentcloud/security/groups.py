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
    def __init__(__self__, id=None, name=None, project_id=None, result_output_file=None, security_group_id=None, security_groups=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence['outputs.GroupsSecurityGroupResult']:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGroupsResult(GroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GroupsResult(
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            result_output_file=self.result_output_file,
            security_group_id=self.security_group_id,
            security_groups=self.security_groups,
            tags=self.tags)


def groups(name: Optional[str] = None,
           project_id: Optional[int] = None,
           result_output_file: Optional[str] = None,
           security_group_id: Optional[str] = None,
           tags: Optional[Mapping[str, Any]] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['resultOutputFile'] = result_output_file
    __args__['securityGroupId'] = security_group_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Security/groups:Groups', __args__, opts=opts, typ=GroupsResult).value

    return AwaitableGroupsResult(
        id=__ret__.id,
        name=__ret__.name,
        project_id=__ret__.project_id,
        result_output_file=__ret__.result_output_file,
        security_group_id=__ret__.security_group_id,
        security_groups=__ret__.security_groups,
        tags=__ret__.tags)


@_utilities.lift_output_func(groups)
def groups_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                  project_id: Optional[pulumi.Input[Optional[int]]] = None,
                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                  security_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                  tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
