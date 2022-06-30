# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstancesAuditListResult',
]

@pulumi.output_type
class InstancesAuditListResult(dict):
    def __init__(__self__, *,
                 audit_switch: bool,
                 cos_bucket: str,
                 id: str,
                 log_file_prefix: str,
                 name: str):
        pulumi.set(__self__, "audit_switch", audit_switch)
        pulumi.set(__self__, "cos_bucket", cos_bucket)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "log_file_prefix", log_file_prefix)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="auditSwitch")
    def audit_switch(self) -> bool:
        return pulumi.get(self, "audit_switch")

    @property
    @pulumi.getter(name="cosBucket")
    def cos_bucket(self) -> str:
        return pulumi.get(self, "cos_bucket")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logFilePrefix")
    def log_file_prefix(self) -> str:
        return pulumi.get(self, "log_file_prefix")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

