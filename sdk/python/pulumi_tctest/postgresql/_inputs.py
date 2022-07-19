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
    'InstanceBackupPlanArgs',
    'InstanceDbNodeSetArgs',
]

@pulumi.input_type
class InstanceBackupPlanArgs:
    def __init__(__self__, *,
                 backup_periods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 base_backup_retention_period: Optional[pulumi.Input[int]] = None,
                 max_backup_start_time: Optional[pulumi.Input[str]] = None,
                 min_backup_start_time: Optional[pulumi.Input[str]] = None):
        if backup_periods is not None:
            pulumi.set(__self__, "backup_periods", backup_periods)
        if base_backup_retention_period is not None:
            pulumi.set(__self__, "base_backup_retention_period", base_backup_retention_period)
        if max_backup_start_time is not None:
            pulumi.set(__self__, "max_backup_start_time", max_backup_start_time)
        if min_backup_start_time is not None:
            pulumi.set(__self__, "min_backup_start_time", min_backup_start_time)

    @property
    @pulumi.getter(name="backupPeriods")
    def backup_periods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "backup_periods")

    @backup_periods.setter
    def backup_periods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "backup_periods", value)

    @property
    @pulumi.getter(name="baseBackupRetentionPeriod")
    def base_backup_retention_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "base_backup_retention_period")

    @base_backup_retention_period.setter
    def base_backup_retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "base_backup_retention_period", value)

    @property
    @pulumi.getter(name="maxBackupStartTime")
    def max_backup_start_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "max_backup_start_time")

    @max_backup_start_time.setter
    def max_backup_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_backup_start_time", value)

    @property
    @pulumi.getter(name="minBackupStartTime")
    def min_backup_start_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "min_backup_start_time")

    @min_backup_start_time.setter
    def min_backup_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "min_backup_start_time", value)


@pulumi.input_type
class InstanceDbNodeSetArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 role: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "zone", zone)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


