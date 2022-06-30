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
    'DomainsResult',
    'AwaitableDomainsResult',
    'domains',
    'domains_output',
]

@pulumi.output_type
class DomainsResult:
    """
    A collection of values returned by Domains.
    """
    def __init__(__self__, domain=None, domain_lists=None, full_url_cache=None, https_switch=None, id=None, origin_pull_protocol=None, result_output_file=None, service_type=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domain_lists and not isinstance(domain_lists, list):
            raise TypeError("Expected argument 'domain_lists' to be a list")
        pulumi.set(__self__, "domain_lists", domain_lists)
        if full_url_cache and not isinstance(full_url_cache, bool):
            raise TypeError("Expected argument 'full_url_cache' to be a bool")
        pulumi.set(__self__, "full_url_cache", full_url_cache)
        if https_switch and not isinstance(https_switch, str):
            raise TypeError("Expected argument 'https_switch' to be a str")
        pulumi.set(__self__, "https_switch", https_switch)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if origin_pull_protocol and not isinstance(origin_pull_protocol, str):
            raise TypeError("Expected argument 'origin_pull_protocol' to be a str")
        pulumi.set(__self__, "origin_pull_protocol", origin_pull_protocol)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if service_type and not isinstance(service_type, str):
            raise TypeError("Expected argument 'service_type' to be a str")
        pulumi.set(__self__, "service_type", service_type)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainLists")
    def domain_lists(self) -> Sequence['outputs.DomainsDomainListResult']:
        return pulumi.get(self, "domain_lists")

    @property
    @pulumi.getter(name="fullUrlCache")
    def full_url_cache(self) -> Optional[bool]:
        return pulumi.get(self, "full_url_cache")

    @property
    @pulumi.getter(name="httpsSwitch")
    def https_switch(self) -> Optional[str]:
        return pulumi.get(self, "https_switch")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="originPullProtocol")
    def origin_pull_protocol(self) -> Optional[str]:
        return pulumi.get(self, "origin_pull_protocol")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[str]:
        return pulumi.get(self, "service_type")


class AwaitableDomainsResult(DomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return DomainsResult(
            domain=self.domain,
            domain_lists=self.domain_lists,
            full_url_cache=self.full_url_cache,
            https_switch=self.https_switch,
            id=self.id,
            origin_pull_protocol=self.origin_pull_protocol,
            result_output_file=self.result_output_file,
            service_type=self.service_type)


def domains(domain: Optional[str] = None,
            full_url_cache: Optional[bool] = None,
            https_switch: Optional[str] = None,
            origin_pull_protocol: Optional[str] = None,
            result_output_file: Optional[str] = None,
            service_type: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableDomainsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['fullUrlCache'] = full_url_cache
    __args__['httpsSwitch'] = https_switch
    __args__['originPullProtocol'] = origin_pull_protocol
    __args__['resultOutputFile'] = result_output_file
    __args__['serviceType'] = service_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cdn/domains:Domains', __args__, opts=opts, typ=DomainsResult).value

    return AwaitableDomainsResult(
        domain=__ret__.domain,
        domain_lists=__ret__.domain_lists,
        full_url_cache=__ret__.full_url_cache,
        https_switch=__ret__.https_switch,
        id=__ret__.id,
        origin_pull_protocol=__ret__.origin_pull_protocol,
        result_output_file=__ret__.result_output_file,
        service_type=__ret__.service_type)


@_utilities.lift_output_func(domains)
def domains_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                   full_url_cache: Optional[pulumi.Input[Optional[bool]]] = None,
                   https_switch: Optional[pulumi.Input[Optional[str]]] = None,
                   origin_pull_protocol: Optional[pulumi.Input[Optional[str]]] = None,
                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                   service_type: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[DomainsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
