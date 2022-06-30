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
    'IdlsResult',
    'AwaitableIdlsResult',
    'idls',
    'idls_output',
]

@pulumi.output_type
class IdlsResult:
    """
    A collection of values returned by Idls.
    """
    def __init__(__self__, cluster_id=None, id=None, lists=None, result_output_file=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.IdlsListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableIdlsResult(IdlsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IdlsResult(
            cluster_id=self.cluster_id,
            id=self.id,
            lists=self.lists,
            result_output_file=self.result_output_file)


def idls(cluster_id: Optional[str] = None,
         result_output_file: Optional[str] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIdlsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tcaplus/idls:Idls', __args__, opts=opts, typ=IdlsResult).value

    return AwaitableIdlsResult(
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        lists=__ret__.lists,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(idls)
def idls_output(cluster_id: Optional[pulumi.Input[str]] = None,
                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IdlsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
