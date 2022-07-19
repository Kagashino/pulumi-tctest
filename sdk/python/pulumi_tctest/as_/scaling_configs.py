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
    'ScalingConfigsResult',
    'AwaitableScalingConfigsResult',
    'scaling_configs',
    'scaling_configs_output',
]

@pulumi.output_type
class ScalingConfigsResult:
    """
    A collection of values returned by ScalingConfigs.
    """
    def __init__(__self__, configuration_id=None, configuration_lists=None, configuration_name=None, id=None, result_output_file=None):
        if configuration_id and not isinstance(configuration_id, str):
            raise TypeError("Expected argument 'configuration_id' to be a str")
        pulumi.set(__self__, "configuration_id", configuration_id)
        if configuration_lists and not isinstance(configuration_lists, list):
            raise TypeError("Expected argument 'configuration_lists' to be a list")
        pulumi.set(__self__, "configuration_lists", configuration_lists)
        if configuration_name and not isinstance(configuration_name, str):
            raise TypeError("Expected argument 'configuration_name' to be a str")
        pulumi.set(__self__, "configuration_name", configuration_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="configurationId")
    def configuration_id(self) -> Optional[str]:
        return pulumi.get(self, "configuration_id")

    @property
    @pulumi.getter(name="configurationLists")
    def configuration_lists(self) -> Sequence['outputs.ScalingConfigsConfigurationListResult']:
        return pulumi.get(self, "configuration_lists")

    @property
    @pulumi.getter(name="configurationName")
    def configuration_name(self) -> Optional[str]:
        return pulumi.get(self, "configuration_name")

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


class AwaitableScalingConfigsResult(ScalingConfigsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalingConfigsResult(
            configuration_id=self.configuration_id,
            configuration_lists=self.configuration_lists,
            configuration_name=self.configuration_name,
            id=self.id,
            result_output_file=self.result_output_file)


def scaling_configs(configuration_id: Optional[str] = None,
                    configuration_name: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalingConfigsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['configurationId'] = configuration_id
    __args__['configurationName'] = configuration_name
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:As/scalingConfigs:ScalingConfigs', __args__, opts=opts, typ=ScalingConfigsResult).value

    return AwaitableScalingConfigsResult(
        configuration_id=__ret__.configuration_id,
        configuration_lists=__ret__.configuration_lists,
        configuration_name=__ret__.configuration_name,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(scaling_configs)
def scaling_configs_output(configuration_id: Optional[pulumi.Input[Optional[str]]] = None,
                           configuration_name: Optional[pulumi.Input[Optional[str]]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ScalingConfigsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...