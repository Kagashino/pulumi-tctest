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
    'IdlTableInfoArgs',
]

@pulumi.input_type
class IdlTableInfoArgs:
    def __init__(__self__, *,
                 error: Optional[pulumi.Input[str]] = None,
                 index_key_set: Optional[pulumi.Input[str]] = None,
                 key_fields: Optional[pulumi.Input[str]] = None,
                 sum_key_field_size: Optional[pulumi.Input[int]] = None,
                 sum_value_field_size: Optional[pulumi.Input[int]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 value_fields: Optional[pulumi.Input[str]] = None):
        if error is not None:
            pulumi.set(__self__, "error", error)
        if index_key_set is not None:
            pulumi.set(__self__, "index_key_set", index_key_set)
        if key_fields is not None:
            pulumi.set(__self__, "key_fields", key_fields)
        if sum_key_field_size is not None:
            pulumi.set(__self__, "sum_key_field_size", sum_key_field_size)
        if sum_value_field_size is not None:
            pulumi.set(__self__, "sum_value_field_size", sum_value_field_size)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)
        if value_fields is not None:
            pulumi.set(__self__, "value_fields", value_fields)

    @property
    @pulumi.getter
    def error(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "error")

    @error.setter
    def error(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "error", value)

    @property
    @pulumi.getter(name="indexKeySet")
    def index_key_set(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "index_key_set")

    @index_key_set.setter
    def index_key_set(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index_key_set", value)

    @property
    @pulumi.getter(name="keyFields")
    def key_fields(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key_fields")

    @key_fields.setter
    def key_fields(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_fields", value)

    @property
    @pulumi.getter(name="sumKeyFieldSize")
    def sum_key_field_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sum_key_field_size")

    @sum_key_field_size.setter
    def sum_key_field_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sum_key_field_size", value)

    @property
    @pulumi.getter(name="sumValueFieldSize")
    def sum_value_field_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sum_value_field_size")

    @sum_value_field_size.setter
    def sum_value_field_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sum_value_field_size", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="valueFields")
    def value_fields(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value_fields")

    @value_fields.setter
    def value_fields(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value_fields", value)


