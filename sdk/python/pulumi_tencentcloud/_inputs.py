# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ProviderAssumeRoleArgs',
]

@pulumi.input_type
class ProviderAssumeRoleArgs:
    def __init__(__self__, *,
                 role_arn: pulumi.Input[str],
                 session_duration: pulumi.Input[int],
                 session_name: pulumi.Input[str],
                 policy: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "session_duration", session_duration)
        pulumi.set(__self__, "session_name", session_name)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="sessionDuration")
    def session_duration(self) -> pulumi.Input[int]:
        return pulumi.get(self, "session_duration")

    @session_duration.setter
    def session_duration(self, value: pulumi.Input[int]):
        pulumi.set(self, "session_duration", value)

    @property
    @pulumi.getter(name="sessionName")
    def session_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "session_name")

    @session_name.setter
    def session_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "session_name", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


