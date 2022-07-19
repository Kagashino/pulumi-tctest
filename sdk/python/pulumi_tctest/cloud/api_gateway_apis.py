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
    'APIGatewayAPIsResult',
    'AwaitableAPIGatewayAPIsResult',
    'api_gateway_apis',
    'api_gateway_apis_output',
]

@pulumi.output_type
class APIGatewayAPIsResult:
    """
    A collection of values returned by APIGatewayAPIs.
    """
    def __init__(__self__, api_id=None, api_name=None, id=None, lists=None, result_output_file=None, service_id=None):
        if api_id and not isinstance(api_id, str):
            raise TypeError("Expected argument 'api_id' to be a str")
        pulumi.set(__self__, "api_id", api_id)
        if api_name and not isinstance(api_name, str):
            raise TypeError("Expected argument 'api_name' to be a str")
        pulumi.set(__self__, "api_name", api_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> Optional[str]:
        return pulumi.get(self, "api_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.APIGatewayAPIsListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        return pulumi.get(self, "service_id")


class AwaitableAPIGatewayAPIsResult(APIGatewayAPIsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return APIGatewayAPIsResult(
            api_id=self.api_id,
            api_name=self.api_name,
            id=self.id,
            lists=self.lists,
            result_output_file=self.result_output_file,
            service_id=self.service_id)


def api_gateway_apis(api_id: Optional[str] = None,
                     api_name: Optional[str] = None,
                     result_output_file: Optional[str] = None,
                     service_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAPIGatewayAPIsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['apiName'] = api_name
    __args__['resultOutputFile'] = result_output_file
    __args__['serviceId'] = service_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Cloud/aPIGatewayAPIs:APIGatewayAPIs', __args__, opts=opts, typ=APIGatewayAPIsResult).value

    return AwaitableAPIGatewayAPIsResult(
        api_id=__ret__.api_id,
        api_name=__ret__.api_name,
        id=__ret__.id,
        lists=__ret__.lists,
        result_output_file=__ret__.result_output_file,
        service_id=__ret__.service_id)


@_utilities.lift_output_func(api_gateway_apis)
def api_gateway_apis_output(api_id: Optional[pulumi.Input[Optional[str]]] = None,
                            api_name: Optional[pulumi.Input[Optional[str]]] = None,
                            result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            service_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[APIGatewayAPIsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
