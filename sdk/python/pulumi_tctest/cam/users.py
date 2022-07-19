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
    'UsersResult',
    'AwaitableUsersResult',
    'users',
    'users_output',
]

@pulumi.output_type
class UsersResult:
    """
    A collection of values returned by Users.
    """
    def __init__(__self__, console_login=None, country_code=None, email=None, id=None, name=None, phone_num=None, remark=None, result_output_file=None, uid=None, uin=None, user_lists=None):
        if console_login and not isinstance(console_login, bool):
            raise TypeError("Expected argument 'console_login' to be a bool")
        pulumi.set(__self__, "console_login", console_login)
        if country_code and not isinstance(country_code, str):
            raise TypeError("Expected argument 'country_code' to be a str")
        pulumi.set(__self__, "country_code", country_code)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if phone_num and not isinstance(phone_num, str):
            raise TypeError("Expected argument 'phone_num' to be a str")
        pulumi.set(__self__, "phone_num", phone_num)
        if remark and not isinstance(remark, str):
            raise TypeError("Expected argument 'remark' to be a str")
        pulumi.set(__self__, "remark", remark)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if uid and not isinstance(uid, int):
            raise TypeError("Expected argument 'uid' to be a int")
        pulumi.set(__self__, "uid", uid)
        if uin and not isinstance(uin, int):
            raise TypeError("Expected argument 'uin' to be a int")
        pulumi.set(__self__, "uin", uin)
        if user_lists and not isinstance(user_lists, list):
            raise TypeError("Expected argument 'user_lists' to be a list")
        pulumi.set(__self__, "user_lists", user_lists)

    @property
    @pulumi.getter(name="consoleLogin")
    def console_login(self) -> Optional[bool]:
        return pulumi.get(self, "console_login")

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[str]:
        return pulumi.get(self, "country_code")

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="phoneNum")
    def phone_num(self) -> Optional[str]:
        return pulumi.get(self, "phone_num")

    @property
    @pulumi.getter
    def remark(self) -> Optional[str]:
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def uid(self) -> Optional[int]:
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def uin(self) -> Optional[int]:
        return pulumi.get(self, "uin")

    @property
    @pulumi.getter(name="userLists")
    def user_lists(self) -> Sequence['outputs.UsersUserListResult']:
        return pulumi.get(self, "user_lists")


class AwaitableUsersResult(UsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return UsersResult(
            console_login=self.console_login,
            country_code=self.country_code,
            email=self.email,
            id=self.id,
            name=self.name,
            phone_num=self.phone_num,
            remark=self.remark,
            result_output_file=self.result_output_file,
            uid=self.uid,
            uin=self.uin,
            user_lists=self.user_lists)


def users(console_login: Optional[bool] = None,
          country_code: Optional[str] = None,
          email: Optional[str] = None,
          name: Optional[str] = None,
          phone_num: Optional[str] = None,
          remark: Optional[str] = None,
          result_output_file: Optional[str] = None,
          uid: Optional[int] = None,
          uin: Optional[int] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableUsersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['consoleLogin'] = console_login
    __args__['countryCode'] = country_code
    __args__['email'] = email
    __args__['name'] = name
    __args__['phoneNum'] = phone_num
    __args__['remark'] = remark
    __args__['resultOutputFile'] = result_output_file
    __args__['uid'] = uid
    __args__['uin'] = uin
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Cam/users:Users', __args__, opts=opts, typ=UsersResult).value

    return AwaitableUsersResult(
        console_login=__ret__.console_login,
        country_code=__ret__.country_code,
        email=__ret__.email,
        id=__ret__.id,
        name=__ret__.name,
        phone_num=__ret__.phone_num,
        remark=__ret__.remark,
        result_output_file=__ret__.result_output_file,
        uid=__ret__.uid,
        uin=__ret__.uin,
        user_lists=__ret__.user_lists)


@_utilities.lift_output_func(users)
def users_output(console_login: Optional[pulumi.Input[Optional[bool]]] = None,
                 country_code: Optional[pulumi.Input[Optional[str]]] = None,
                 email: Optional[pulumi.Input[Optional[str]]] = None,
                 name: Optional[pulumi.Input[Optional[str]]] = None,
                 phone_num: Optional[pulumi.Input[Optional[str]]] = None,
                 remark: Optional[pulumi.Input[Optional[str]]] = None,
                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 uid: Optional[pulumi.Input[Optional[int]]] = None,
                 uin: Optional[pulumi.Input[Optional[int]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[UsersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
