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
    'XlogsResult',
    'AwaitableXlogsResult',
    'xlogs',
    'xlogs_output',
]

@pulumi.output_type
class XlogsResult:
    """
    A collection of values returned by Xlogs.
    """
    def __init__(__self__, end_time=None, id=None, instance_id=None, lists=None, result_output_file=None, start_time=None):
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.XlogsListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[str]:
        return pulumi.get(self, "start_time")


class AwaitableXlogsResult(XlogsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return XlogsResult(
            end_time=self.end_time,
            id=self.id,
            instance_id=self.instance_id,
            lists=self.lists,
            result_output_file=self.result_output_file,
            start_time=self.start_time)


def xlogs(end_time: Optional[str] = None,
          instance_id: Optional[str] = None,
          result_output_file: Optional[str] = None,
          start_time: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableXlogsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['endTime'] = end_time
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Postgresql/xlogs:Xlogs', __args__, opts=opts, typ=XlogsResult).value

    return AwaitableXlogsResult(
        end_time=__ret__.end_time,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        lists=__ret__.lists,
        result_output_file=__ret__.result_output_file,
        start_time=__ret__.start_time)


@_utilities.lift_output_func(xlogs)
def xlogs_output(end_time: Optional[pulumi.Input[Optional[str]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 start_time: Optional[pulumi.Input[Optional[str]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[XlogsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
