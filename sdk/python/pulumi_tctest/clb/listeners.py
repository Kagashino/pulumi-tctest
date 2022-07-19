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
    'ListenersResult',
    'AwaitableListenersResult',
    'listeners',
    'listeners_output',
]

@pulumi.output_type
class ListenersResult:
    """
    A collection of values returned by Listeners.
    """
    def __init__(__self__, clb_id=None, id=None, listener_id=None, listener_lists=None, port=None, protocol=None, result_output_file=None):
        if clb_id and not isinstance(clb_id, str):
            raise TypeError("Expected argument 'clb_id' to be a str")
        pulumi.set(__self__, "clb_id", clb_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if listener_lists and not isinstance(listener_lists, list):
            raise TypeError("Expected argument 'listener_lists' to be a list")
        pulumi.set(__self__, "listener_lists", listener_lists)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clbId")
    def clb_id(self) -> str:
        return pulumi.get(self, "clb_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[str]:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="listenerLists")
    def listener_lists(self) -> Sequence['outputs.ListenersListenerListResult']:
        return pulumi.get(self, "listener_lists")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableListenersResult(ListenersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListenersResult(
            clb_id=self.clb_id,
            id=self.id,
            listener_id=self.listener_id,
            listener_lists=self.listener_lists,
            port=self.port,
            protocol=self.protocol,
            result_output_file=self.result_output_file)


def listeners(clb_id: Optional[str] = None,
              listener_id: Optional[str] = None,
              port: Optional[int] = None,
              protocol: Optional[str] = None,
              result_output_file: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListenersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clbId'] = clb_id
    __args__['listenerId'] = listener_id
    __args__['port'] = port
    __args__['protocol'] = protocol
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Clb/listeners:Listeners', __args__, opts=opts, typ=ListenersResult).value

    return AwaitableListenersResult(
        clb_id=__ret__.clb_id,
        id=__ret__.id,
        listener_id=__ret__.listener_id,
        listener_lists=__ret__.listener_lists,
        port=__ret__.port,
        protocol=__ret__.protocol,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(listeners)
def listeners_output(clb_id: Optional[pulumi.Input[str]] = None,
                     listener_id: Optional[pulumi.Input[Optional[str]]] = None,
                     port: Optional[pulumi.Input[Optional[int]]] = None,
                     protocol: Optional[pulumi.Input[Optional[str]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListenersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
