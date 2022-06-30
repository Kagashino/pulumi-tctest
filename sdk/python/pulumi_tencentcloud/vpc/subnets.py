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
    'SubnetsResult',
    'AwaitableSubnetsResult',
    'subnets',
    'subnets_output',
]

@pulumi.output_type
class SubnetsResult:
    """
    A collection of values returned by Subnets.
    """
    def __init__(__self__, availability_zone=None, cidr_block=None, id=None, instance_lists=None, is_default=None, is_remote_vpc_snat=None, name=None, result_output_file=None, subnet_id=None, tag_key=None, tags=None, vpc_id=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        pulumi.set(__self__, "cidr_block", cidr_block)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if is_remote_vpc_snat and not isinstance(is_remote_vpc_snat, bool):
            raise TypeError("Expected argument 'is_remote_vpc_snat' to be a bool")
        pulumi.set(__self__, "is_remote_vpc_snat", is_remote_vpc_snat)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tag_key and not isinstance(tag_key, str):
            raise TypeError("Expected argument 'tag_key' to be a str")
        pulumi.set(__self__, "tag_key", tag_key)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[str]:
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.SubnetsInstanceListResult']:
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[bool]:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="isRemoteVpcSnat")
    def is_remote_vpc_snat(self) -> Optional[bool]:
        return pulumi.get(self, "is_remote_vpc_snat")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="tagKey")
    def tag_key(self) -> Optional[str]:
        return pulumi.get(self, "tag_key")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableSubnetsResult(SubnetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SubnetsResult(
            availability_zone=self.availability_zone,
            cidr_block=self.cidr_block,
            id=self.id,
            instance_lists=self.instance_lists,
            is_default=self.is_default,
            is_remote_vpc_snat=self.is_remote_vpc_snat,
            name=self.name,
            result_output_file=self.result_output_file,
            subnet_id=self.subnet_id,
            tag_key=self.tag_key,
            tags=self.tags,
            vpc_id=self.vpc_id)


def subnets(availability_zone: Optional[str] = None,
            cidr_block: Optional[str] = None,
            is_default: Optional[bool] = None,
            is_remote_vpc_snat: Optional[bool] = None,
            name: Optional[str] = None,
            result_output_file: Optional[str] = None,
            subnet_id: Optional[str] = None,
            tag_key: Optional[str] = None,
            tags: Optional[Mapping[str, Any]] = None,
            vpc_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSubnetsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['cidrBlock'] = cidr_block
    __args__['isDefault'] = is_default
    __args__['isRemoteVpcSnat'] = is_remote_vpc_snat
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['subnetId'] = subnet_id
    __args__['tagKey'] = tag_key
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpc/subnets:Subnets', __args__, opts=opts, typ=SubnetsResult).value

    return AwaitableSubnetsResult(
        availability_zone=__ret__.availability_zone,
        cidr_block=__ret__.cidr_block,
        id=__ret__.id,
        instance_lists=__ret__.instance_lists,
        is_default=__ret__.is_default,
        is_remote_vpc_snat=__ret__.is_remote_vpc_snat,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file,
        subnet_id=__ret__.subnet_id,
        tag_key=__ret__.tag_key,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(subnets)
def subnets_output(availability_zone: Optional[pulumi.Input[Optional[str]]] = None,
                   cidr_block: Optional[pulumi.Input[Optional[str]]] = None,
                   is_default: Optional[pulumi.Input[Optional[bool]]] = None,
                   is_remote_vpc_snat: Optional[pulumi.Input[Optional[bool]]] = None,
                   name: Optional[pulumi.Input[Optional[str]]] = None,
                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                   subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                   tag_key: Optional[pulumi.Input[Optional[str]]] = None,
                   tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                   vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SubnetsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
