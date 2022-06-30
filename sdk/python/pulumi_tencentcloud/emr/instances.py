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
    'InstancesResult',
    'AwaitableInstancesResult',
    'instances',
    'instances_output',
]

@pulumi.output_type
class InstancesResult:
    """
    A collection of values returned by Instances.
    """
    def __init__(__self__, clusters=None, display_strategy=None, id=None, instance_ids=None, project_id=None, result_output_file=None):
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if display_strategy and not isinstance(display_strategy, str):
            raise TypeError("Expected argument 'display_strategy' to be a str")
        pulumi.set(__self__, "display_strategy", display_strategy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.InstancesClusterResult']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="displayStrategy")
    def display_strategy(self) -> str:
        return pulumi.get(self, "display_strategy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            clusters=self.clusters,
            display_strategy=self.display_strategy,
            id=self.id,
            instance_ids=self.instance_ids,
            project_id=self.project_id,
            result_output_file=self.result_output_file)


def instances(display_strategy: Optional[str] = None,
              instance_ids: Optional[Sequence[str]] = None,
              project_id: Optional[int] = None,
              result_output_file: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['displayStrategy'] = display_strategy
    __args__['instanceIds'] = instance_ids
    __args__['projectId'] = project_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Emr/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        clusters=__ret__.clusters,
        display_strategy=__ret__.display_strategy,
        id=__ret__.id,
        instance_ids=__ret__.instance_ids,
        project_id=__ret__.project_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(instances)
def instances_output(display_strategy: Optional[pulumi.Input[str]] = None,
                     instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     project_id: Optional[pulumi.Input[Optional[int]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
