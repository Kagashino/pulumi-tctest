# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FileSystemArgs', 'FileSystem']

@pulumi.input_type
class FileSystemArgs:
    def __init__(__self__, *,
                 access_group_id: pulumi.Input[str],
                 availability_zone: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 mount_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a FileSystem resource.
        :param pulumi.Input[str] access_group_id: ID of a access group.
        :param pulumi.Input[str] availability_zone: The available zone that the file system locates at.
        :param pulumi.Input[str] subnet_id: ID of a subnet.
        :param pulumi.Input[str] vpc_id: ID of a VPC network.
        :param pulumi.Input[str] mount_ip: IP of mount point.
        :param pulumi.Input[str] name: Name of a file system.
        :param pulumi.Input[str] protocol: File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
        :param pulumi.Input[str] storage_type: File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        """
        pulumi.set(__self__, "access_group_id", access_group_id)
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if mount_ip is not None:
            pulumi.set(__self__, "mount_ip", mount_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if storage_type is not None:
            pulumi.set(__self__, "storage_type", storage_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> pulumi.Input[str]:
        """
        ID of a access group.
        """
        return pulumi.get(self, "access_group_id")

    @access_group_id.setter
    def access_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_group_id", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        """
        The available zone that the file system locates at.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        ID of a subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        ID of a VPC network.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="mountIp")
    def mount_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP of mount point.
        """
        return pulumi.get(self, "mount_ip")

    @mount_ip.setter
    def mount_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of a file system.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> Optional[pulumi.Input[str]]:
        """
        File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Instance tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _FileSystemState:
    def __init__(__self__, *,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 mount_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FileSystem resources.
        :param pulumi.Input[str] access_group_id: ID of a access group.
        :param pulumi.Input[str] availability_zone: The available zone that the file system locates at.
        :param pulumi.Input[str] create_time: Create time of the file system.
        :param pulumi.Input[str] mount_ip: IP of mount point.
        :param pulumi.Input[str] name: Name of a file system.
        :param pulumi.Input[str] protocol: File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
        :param pulumi.Input[str] storage_type: File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
        :param pulumi.Input[str] subnet_id: ID of a subnet.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        :param pulumi.Input[str] vpc_id: ID of a VPC network.
        """
        if access_group_id is not None:
            pulumi.set(__self__, "access_group_id", access_group_id)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if mount_ip is not None:
            pulumi.set(__self__, "mount_ip", mount_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if storage_type is not None:
            pulumi.set(__self__, "storage_type", storage_type)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a access group.
        """
        return pulumi.get(self, "access_group_id")

    @access_group_id.setter
    def access_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_group_id", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The available zone that the file system locates at.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the file system.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="mountIp")
    def mount_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP of mount point.
        """
        return pulumi.get(self, "mount_ip")

    @mount_ip.setter
    def mount_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of a file system.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> Optional[pulumi.Input[str]]:
        """
        File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Instance tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of a VPC network.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class FileSystem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 mount_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FileSystem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_id: ID of a access group.
        :param pulumi.Input[str] availability_zone: The available zone that the file system locates at.
        :param pulumi.Input[str] mount_ip: IP of mount point.
        :param pulumi.Input[str] name: Name of a file system.
        :param pulumi.Input[str] protocol: File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
        :param pulumi.Input[str] storage_type: File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
        :param pulumi.Input[str] subnet_id: ID of a subnet.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        :param pulumi.Input[str] vpc_id: ID of a VPC network.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileSystemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FileSystem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FileSystemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileSystemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_group_id: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 mount_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
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
            __props__ = FileSystemArgs.__new__(FileSystemArgs)

            if access_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'access_group_id'")
            __props__.__dict__["access_group_id"] = access_group_id
            if availability_zone is None and not opts.urn:
                raise TypeError("Missing required property 'availability_zone'")
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["mount_ip"] = mount_ip
            __props__.__dict__["name"] = name
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["storage_type"] = storage_type
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["create_time"] = None
        super(FileSystem, __self__).__init__(
            'tencentcloud:Cfs/fileSystem:FileSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_group_id: Optional[pulumi.Input[str]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            mount_ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'FileSystem':
        """
        Get an existing FileSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_group_id: ID of a access group.
        :param pulumi.Input[str] availability_zone: The available zone that the file system locates at.
        :param pulumi.Input[str] create_time: Create time of the file system.
        :param pulumi.Input[str] mount_ip: IP of mount point.
        :param pulumi.Input[str] name: Name of a file system.
        :param pulumi.Input[str] protocol: File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
        :param pulumi.Input[str] storage_type: File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
        :param pulumi.Input[str] subnet_id: ID of a subnet.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        :param pulumi.Input[str] vpc_id: ID of a VPC network.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileSystemState.__new__(_FileSystemState)

        __props__.__dict__["access_group_id"] = access_group_id
        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["mount_ip"] = mount_ip
        __props__.__dict__["name"] = name
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["storage_type"] = storage_type
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        return FileSystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> pulumi.Output[str]:
        """
        ID of a access group.
        """
        return pulumi.get(self, "access_group_id")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The available zone that the file system locates at.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the file system.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="mountIp")
    def mount_ip(self) -> pulumi.Output[str]:
        """
        IP of mount point.
        """
        return pulumi.get(self, "mount_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of a file system.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        """
        File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[Optional[str]]:
        """
        File service StorageType. Valid values are `SD` and `HP`. and the default is `SD`.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        ID of a subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Instance tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        ID of a VPC network.
        """
        return pulumi.get(self, "vpc_id")

