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
    def __init__(__self__, cluster_type=None, id=None, instance_id=None, instance_lists=None, instance_name_prefix=None, result_output_file=None, tags=None):
        if cluster_type and not isinstance(cluster_type, str):
            raise TypeError("Expected argument 'cluster_type' to be a str")
        pulumi.set(__self__, "cluster_type", cluster_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if instance_name_prefix and not isinstance(instance_name_prefix, str):
            raise TypeError("Expected argument 'instance_name_prefix' to be a str")
        pulumi.set(__self__, "instance_name_prefix", instance_name_prefix)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[str]:
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.InstancesInstanceListResult']:
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter(name="instanceNamePrefix")
    def instance_name_prefix(self) -> Optional[str]:
        return pulumi.get(self, "instance_name_prefix")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            cluster_type=self.cluster_type,
            id=self.id,
            instance_id=self.instance_id,
            instance_lists=self.instance_lists,
            instance_name_prefix=self.instance_name_prefix,
            result_output_file=self.result_output_file,
            tags=self.tags)


def instances(cluster_type: Optional[str] = None,
              instance_id: Optional[str] = None,
              instance_name_prefix: Optional[str] = None,
              result_output_file: Optional[str] = None,
              tags: Optional[Mapping[str, Any]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterType'] = cluster_type
    __args__['instanceId'] = instance_id
    __args__['instanceNamePrefix'] = instance_name_prefix
    __args__['resultOutputFile'] = result_output_file
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Mongodb/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        cluster_type=__ret__.cluster_type,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_lists=__ret__.instance_lists,
        instance_name_prefix=__ret__.instance_name_prefix,
        result_output_file=__ret__.result_output_file,
        tags=__ret__.tags)


@_utilities.lift_output_func(instances)
def instances_output(cluster_type: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_name_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
