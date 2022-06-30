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
    'PairsResult',
    'AwaitablePairsResult',
    'pairs',
    'pairs_output',
]

@pulumi.output_type
class PairsResult:
    """
    A collection of values returned by Pairs.
    """
    def __init__(__self__, id=None, key_id=None, key_name=None, key_pair_lists=None, project_id=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        pulumi.set(__self__, "key_id", key_id)
        if key_name and not isinstance(key_name, str):
            raise TypeError("Expected argument 'key_name' to be a str")
        pulumi.set(__self__, "key_name", key_name)
        if key_pair_lists and not isinstance(key_pair_lists, list):
            raise TypeError("Expected argument 'key_pair_lists' to be a list")
        pulumi.set(__self__, "key_pair_lists", key_pair_lists)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
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
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[str]:
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[str]:
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="keyPairLists")
    def key_pair_lists(self) -> Sequence['outputs.PairsKeyPairListResult']:
        return pulumi.get(self, "key_pair_lists")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitablePairsResult(PairsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return PairsResult(
            id=self.id,
            key_id=self.key_id,
            key_name=self.key_name,
            key_pair_lists=self.key_pair_lists,
            project_id=self.project_id,
            result_output_file=self.result_output_file)


def pairs(key_id: Optional[str] = None,
          key_name: Optional[str] = None,
          project_id: Optional[int] = None,
          result_output_file: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitablePairsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['keyId'] = key_id
    __args__['keyName'] = key_name
    __args__['projectId'] = project_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Key/pairs:Pairs', __args__, opts=opts, typ=PairsResult).value

    return AwaitablePairsResult(
        id=__ret__.id,
        key_id=__ret__.key_id,
        key_name=__ret__.key_name,
        key_pair_lists=__ret__.key_pair_lists,
        project_id=__ret__.project_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(pairs)
def pairs_output(key_id: Optional[pulumi.Input[Optional[str]]] = None,
                 key_name: Optional[pulumi.Input[Optional[str]]] = None,
                 project_id: Optional[pulumi.Input[Optional[int]]] = None,
                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[PairsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
