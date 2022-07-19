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
from ._inputs import *

__all__ = [
    'ProductEventResult',
    'AwaitableProductEventResult',
    'product_event',
    'product_event_output',
]

@pulumi.output_type
class ProductEventResult:
    """
    A collection of values returned by ProductEvent.
    """
    def __init__(__self__, dimensions=None, end_time=None, event_names=None, id=None, instance_ids=None, is_alarm_config=None, lists=None, product_names=None, project_ids=None, region_lists=None, result_output_file=None, start_time=None, statuses=None, types=None):
        if dimensions and not isinstance(dimensions, list):
            raise TypeError("Expected argument 'dimensions' to be a list")
        pulumi.set(__self__, "dimensions", dimensions)
        if end_time and not isinstance(end_time, int):
            raise TypeError("Expected argument 'end_time' to be a int")
        pulumi.set(__self__, "end_time", end_time)
        if event_names and not isinstance(event_names, list):
            raise TypeError("Expected argument 'event_names' to be a list")
        pulumi.set(__self__, "event_names", event_names)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if is_alarm_config and not isinstance(is_alarm_config, int):
            raise TypeError("Expected argument 'is_alarm_config' to be a int")
        pulumi.set(__self__, "is_alarm_config", is_alarm_config)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if product_names and not isinstance(product_names, list):
            raise TypeError("Expected argument 'product_names' to be a list")
        pulumi.set(__self__, "product_names", product_names)
        if project_ids and not isinstance(project_ids, list):
            raise TypeError("Expected argument 'project_ids' to be a list")
        pulumi.set(__self__, "project_ids", project_ids)
        if region_lists and not isinstance(region_lists, list):
            raise TypeError("Expected argument 'region_lists' to be a list")
        pulumi.set(__self__, "region_lists", region_lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_time and not isinstance(start_time, int):
            raise TypeError("Expected argument 'start_time' to be a int")
        pulumi.set(__self__, "start_time", start_time)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if types and not isinstance(types, list):
            raise TypeError("Expected argument 'types' to be a list")
        pulumi.set(__self__, "types", types)

    @property
    @pulumi.getter
    def dimensions(self) -> Optional[Sequence['outputs.ProductEventDimensionResult']]:
        return pulumi.get(self, "dimensions")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[int]:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="eventNames")
    def event_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "event_names")

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
    @pulumi.getter(name="isAlarmConfig")
    def is_alarm_config(self) -> Optional[int]:
        return pulumi.get(self, "is_alarm_config")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.ProductEventListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="productNames")
    def product_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "product_names")

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "project_ids")

    @property
    @pulumi.getter(name="regionLists")
    def region_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "region_lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[int]:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter
    def types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "types")


class AwaitableProductEventResult(ProductEventResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ProductEventResult(
            dimensions=self.dimensions,
            end_time=self.end_time,
            event_names=self.event_names,
            id=self.id,
            instance_ids=self.instance_ids,
            is_alarm_config=self.is_alarm_config,
            lists=self.lists,
            product_names=self.product_names,
            project_ids=self.project_ids,
            region_lists=self.region_lists,
            result_output_file=self.result_output_file,
            start_time=self.start_time,
            statuses=self.statuses,
            types=self.types)


def product_event(dimensions: Optional[Sequence[pulumi.InputType['ProductEventDimensionArgs']]] = None,
                  end_time: Optional[int] = None,
                  event_names: Optional[Sequence[str]] = None,
                  instance_ids: Optional[Sequence[str]] = None,
                  is_alarm_config: Optional[int] = None,
                  product_names: Optional[Sequence[str]] = None,
                  project_ids: Optional[Sequence[str]] = None,
                  region_lists: Optional[Sequence[str]] = None,
                  result_output_file: Optional[str] = None,
                  start_time: Optional[int] = None,
                  statuses: Optional[Sequence[str]] = None,
                  types: Optional[Sequence[str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableProductEventResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['dimensions'] = dimensions
    __args__['endTime'] = end_time
    __args__['eventNames'] = event_names
    __args__['instanceIds'] = instance_ids
    __args__['isAlarmConfig'] = is_alarm_config
    __args__['productNames'] = product_names
    __args__['projectIds'] = project_ids
    __args__['regionLists'] = region_lists
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    __args__['statuses'] = statuses
    __args__['types'] = types
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Monitor/productEvent:ProductEvent', __args__, opts=opts, typ=ProductEventResult).value

    return AwaitableProductEventResult(
        dimensions=__ret__.dimensions,
        end_time=__ret__.end_time,
        event_names=__ret__.event_names,
        id=__ret__.id,
        instance_ids=__ret__.instance_ids,
        is_alarm_config=__ret__.is_alarm_config,
        lists=__ret__.lists,
        product_names=__ret__.product_names,
        project_ids=__ret__.project_ids,
        region_lists=__ret__.region_lists,
        result_output_file=__ret__.result_output_file,
        start_time=__ret__.start_time,
        statuses=__ret__.statuses,
        types=__ret__.types)


@_utilities.lift_output_func(product_event)
def product_event_output(dimensions: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['ProductEventDimensionArgs']]]]] = None,
                         end_time: Optional[pulumi.Input[Optional[int]]] = None,
                         event_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         instance_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         is_alarm_config: Optional[pulumi.Input[Optional[int]]] = None,
                         product_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         project_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         region_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         start_time: Optional[pulumi.Input[Optional[int]]] = None,
                         statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ProductEventResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
