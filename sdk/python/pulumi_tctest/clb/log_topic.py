# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LogTopicArgs', 'LogTopic']

@pulumi.input_type
class LogTopicArgs:
    def __init__(__self__, *,
                 log_set_id: pulumi.Input[str],
                 topic_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a LogTopic resource.
        :param pulumi.Input[str] log_set_id: Log topic of CLB instance.
        :param pulumi.Input[str] topic_name: Log topic of CLB instance.
        """
        pulumi.set(__self__, "log_set_id", log_set_id)
        pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter(name="logSetId")
    def log_set_id(self) -> pulumi.Input[str]:
        """
        Log topic of CLB instance.
        """
        return pulumi.get(self, "log_set_id")

    @log_set_id.setter
    def log_set_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_set_id", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        Log topic of CLB instance.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)


@pulumi.input_type
class _LogTopicState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 log_set_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[bool]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogTopic resources.
        :param pulumi.Input[str] create_time: Log topic creation time.
        :param pulumi.Input[str] log_set_id: Log topic of CLB instance.
        :param pulumi.Input[bool] status: The status of log topic.
        :param pulumi.Input[str] topic_name: Log topic of CLB instance.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if log_set_id is not None:
            pulumi.set(__self__, "log_set_id", log_set_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Log topic creation time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="logSetId")
    def log_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        Log topic of CLB instance.
        """
        return pulumi.get(self, "log_set_id")

    @log_set_id.setter
    def log_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_set_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[bool]]:
        """
        The status of log topic.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        Log topic of CLB instance.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)


class LogTopic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_set_id: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LogTopic resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] log_set_id: Log topic of CLB instance.
        :param pulumi.Input[str] topic_name: Log topic of CLB instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogTopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LogTopic resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LogTopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogTopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_set_id: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = LogTopicArgs.__new__(LogTopicArgs)

            if log_set_id is None and not opts.urn:
                raise TypeError("Missing required property 'log_set_id'")
            __props__.__dict__["log_set_id"] = log_set_id
            if topic_name is None and not opts.urn:
                raise TypeError("Missing required property 'topic_name'")
            __props__.__dict__["topic_name"] = topic_name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(LogTopic, __self__).__init__(
            'tctest:Clb/logTopic:LogTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            log_set_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[bool]] = None,
            topic_name: Optional[pulumi.Input[str]] = None) -> 'LogTopic':
        """
        Get an existing LogTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Log topic creation time.
        :param pulumi.Input[str] log_set_id: Log topic of CLB instance.
        :param pulumi.Input[bool] status: The status of log topic.
        :param pulumi.Input[str] topic_name: Log topic of CLB instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogTopicState.__new__(_LogTopicState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["log_set_id"] = log_set_id
        __props__.__dict__["status"] = status
        __props__.__dict__["topic_name"] = topic_name
        return LogTopic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Log topic creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="logSetId")
    def log_set_id(self) -> pulumi.Output[str]:
        """
        Log topic of CLB instance.
        """
        return pulumi.get(self, "log_set_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[bool]:
        """
        The status of log topic.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        Log topic of CLB instance.
        """
        return pulumi.get(self, "topic_name")

