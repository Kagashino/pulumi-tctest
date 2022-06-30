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
    'ReadonlyGroupsResult',
    'AwaitableReadonlyGroupsResult',
    'readonly_groups',
    'readonly_groups_output',
]

@pulumi.output_type
class ReadonlyGroupsResult:
    """
    A collection of values returned by ReadonlyGroups.
    """
    def __init__(__self__, id=None, lists=None, master_instance_id=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if master_instance_id and not isinstance(master_instance_id, str):
            raise TypeError("Expected argument 'master_instance_id' to be a str")
        pulumi.set(__self__, "master_instance_id", master_instance_id)
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
    def lists(self) -> Sequence['outputs.ReadonlyGroupsListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="masterInstanceId")
    def master_instance_id(self) -> Optional[str]:
        return pulumi.get(self, "master_instance_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableReadonlyGroupsResult(ReadonlyGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ReadonlyGroupsResult(
            id=self.id,
            lists=self.lists,
            master_instance_id=self.master_instance_id,
            result_output_file=self.result_output_file)


def readonly_groups(master_instance_id: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableReadonlyGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['masterInstanceId'] = master_instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/readonlyGroups:ReadonlyGroups', __args__, opts=opts, typ=ReadonlyGroupsResult).value

    return AwaitableReadonlyGroupsResult(
        id=__ret__.id,
        lists=__ret__.lists,
        master_instance_id=__ret__.master_instance_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(readonly_groups)
def readonly_groups_output(master_instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ReadonlyGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
