# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcAttachmentArgs', 'VpcAttachment']

@pulumi.input_type
class VpcAttachmentArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 enable_public_domain_dns: Optional[pulumi.Input[bool]] = None,
                 enable_vpc_domain_dns: Optional[pulumi.Input[bool]] = None,
                 region_id: Optional[pulumi.Input[int]] = None,
                 region_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcAttachment resource.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[str] subnet_id: ID of subnet.
        :param pulumi.Input[str] vpc_id: ID of VPC.
        :param pulumi.Input[bool] enable_public_domain_dns: Whether to enable public domain dns. Default value is `false`.
        :param pulumi.Input[bool] enable_vpc_domain_dns: Whether to enable vpc domain dns. Default value is `false`.
        :param pulumi.Input[int] region_id: ID of region. Conflict with region_name, can not be set at the same time.
        :param pulumi.Input[str] region_name: Name of region. Conflict with region_id, can not be set at the same time.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if enable_public_domain_dns is not None:
            pulumi.set(__self__, "enable_public_domain_dns", enable_public_domain_dns)
        if enable_vpc_domain_dns is not None:
            pulumi.set(__self__, "enable_vpc_domain_dns", enable_vpc_domain_dns)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        ID of subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        ID of VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="enablePublicDomainDns")
    def enable_public_domain_dns(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable public domain dns. Default value is `false`.
        """
        return pulumi.get(self, "enable_public_domain_dns")

    @enable_public_domain_dns.setter
    def enable_public_domain_dns(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_public_domain_dns", value)

    @property
    @pulumi.getter(name="enableVpcDomainDns")
    def enable_vpc_domain_dns(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable vpc domain dns. Default value is `false`.
        """
        return pulumi.get(self, "enable_vpc_domain_dns")

    @enable_vpc_domain_dns.setter
    def enable_vpc_domain_dns(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_vpc_domain_dns", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[int]]:
        """
        ID of region. Conflict with region_name, can not be set at the same time.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of region. Conflict with region_id, can not be set at the same time.
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)


@pulumi.input_type
class _VpcAttachmentState:
    def __init__(__self__, *,
                 access_ip: Optional[pulumi.Input[str]] = None,
                 enable_public_domain_dns: Optional[pulumi.Input[bool]] = None,
                 enable_vpc_domain_dns: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[int]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcAttachment resources.
        :param pulumi.Input[str] access_ip: IP address of the internal access.
        :param pulumi.Input[bool] enable_public_domain_dns: Whether to enable public domain dns. Default value is `false`.
        :param pulumi.Input[bool] enable_vpc_domain_dns: Whether to enable vpc domain dns. Default value is `false`.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[int] region_id: ID of region. Conflict with region_name, can not be set at the same time.
        :param pulumi.Input[str] region_name: Name of region. Conflict with region_id, can not be set at the same time.
        :param pulumi.Input[str] status: Status of the internal access.
        :param pulumi.Input[str] subnet_id: ID of subnet.
        :param pulumi.Input[str] vpc_id: ID of VPC.
        """
        if access_ip is not None:
            pulumi.set(__self__, "access_ip", access_ip)
        if enable_public_domain_dns is not None:
            pulumi.set(__self__, "enable_public_domain_dns", enable_public_domain_dns)
        if enable_vpc_domain_dns is not None:
            pulumi.set(__self__, "enable_vpc_domain_dns", enable_vpc_domain_dns)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="accessIp")
    def access_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the internal access.
        """
        return pulumi.get(self, "access_ip")

    @access_ip.setter
    def access_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_ip", value)

    @property
    @pulumi.getter(name="enablePublicDomainDns")
    def enable_public_domain_dns(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable public domain dns. Default value is `false`.
        """
        return pulumi.get(self, "enable_public_domain_dns")

    @enable_public_domain_dns.setter
    def enable_public_domain_dns(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_public_domain_dns", value)

    @property
    @pulumi.getter(name="enableVpcDomainDns")
    def enable_vpc_domain_dns(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable vpc domain dns. Default value is `false`.
        """
        return pulumi.get(self, "enable_vpc_domain_dns")

    @enable_vpc_domain_dns.setter
    def enable_vpc_domain_dns(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_vpc_domain_dns", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[int]]:
        """
        ID of region. Conflict with region_name, can not be set at the same time.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of region. Conflict with region_id, can not be set at the same time.
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the internal access.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class VpcAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_public_domain_dns: Optional[pulumi.Input[bool]] = None,
                 enable_vpc_domain_dns: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[int]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VpcAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_public_domain_dns: Whether to enable public domain dns. Default value is `false`.
        :param pulumi.Input[bool] enable_vpc_domain_dns: Whether to enable vpc domain dns. Default value is `false`.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[int] region_id: ID of region. Conflict with region_name, can not be set at the same time.
        :param pulumi.Input[str] region_name: Name of region. Conflict with region_id, can not be set at the same time.
        :param pulumi.Input[str] subnet_id: ID of subnet.
        :param pulumi.Input[str] vpc_id: ID of VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpcAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpcAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_public_domain_dns: Optional[pulumi.Input[bool]] = None,
                 enable_vpc_domain_dns: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[int]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        else:
            opts = copy.copy(opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcAttachmentArgs.__new__(VpcAttachmentArgs)

            __props__.__dict__["enable_public_domain_dns"] = enable_public_domain_dns
            __props__.__dict__["enable_vpc_domain_dns"] = enable_vpc_domain_dns
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region_id"] = region_id
            __props__.__dict__["region_name"] = region_name
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["access_ip"] = None
            __props__.__dict__["status"] = None
        super(VpcAttachment, __self__).__init__(
            'tencentcloud:Tcr/vpcAttachment:VpcAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_ip: Optional[pulumi.Input[str]] = None,
            enable_public_domain_dns: Optional[pulumi.Input[bool]] = None,
            enable_vpc_domain_dns: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            region_id: Optional[pulumi.Input[int]] = None,
            region_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'VpcAttachment':
        """
        Get an existing VpcAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_ip: IP address of the internal access.
        :param pulumi.Input[bool] enable_public_domain_dns: Whether to enable public domain dns. Default value is `false`.
        :param pulumi.Input[bool] enable_vpc_domain_dns: Whether to enable vpc domain dns. Default value is `false`.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[int] region_id: ID of region. Conflict with region_name, can not be set at the same time.
        :param pulumi.Input[str] region_name: Name of region. Conflict with region_id, can not be set at the same time.
        :param pulumi.Input[str] status: Status of the internal access.
        :param pulumi.Input[str] subnet_id: ID of subnet.
        :param pulumi.Input[str] vpc_id: ID of VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcAttachmentState.__new__(_VpcAttachmentState)

        __props__.__dict__["access_ip"] = access_ip
        __props__.__dict__["enable_public_domain_dns"] = enable_public_domain_dns
        __props__.__dict__["enable_vpc_domain_dns"] = enable_vpc_domain_dns
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["region_id"] = region_id
        __props__.__dict__["region_name"] = region_name
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["vpc_id"] = vpc_id
        return VpcAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessIp")
    def access_ip(self) -> pulumi.Output[str]:
        """
        IP address of the internal access.
        """
        return pulumi.get(self, "access_ip")

    @property
    @pulumi.getter(name="enablePublicDomainDns")
    def enable_public_domain_dns(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable public domain dns. Default value is `false`.
        """
        return pulumi.get(self, "enable_public_domain_dns")

    @property
    @pulumi.getter(name="enableVpcDomainDns")
    def enable_vpc_domain_dns(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable vpc domain dns. Default value is `false`.
        """
        return pulumi.get(self, "enable_vpc_domain_dns")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Output[Optional[int]]:
        """
        ID of region. Conflict with region_name, can not be set at the same time.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of region. Conflict with region_id, can not be set at the same time.
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the internal access.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        ID of subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        ID of VPC.
        """
        return pulumi.get(self, "vpc_id")

