# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 console_login: Optional[pulumi.Input[bool]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_reset_password: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 use_api: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[bool] console_login: Indicate whether the CAM user can login to the web console or not.
        :param pulumi.Input[str] country_code: Country code of the phone number, for example: '86'.
        :param pulumi.Input[str] email: Email of the CAM user.
        :param pulumi.Input[bool] force_delete: Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists;
               otherwise the user will be deleted directly. Default is false.
        :param pulumi.Input[str] name: Name of the CAM user.
        :param pulumi.Input[bool] need_reset_password: Indicate whether the CAM user need to reset the password when first logins.
        :param pulumi.Input[str] password: The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes
               uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not
               set, a random password will be automatically generated.
        :param pulumi.Input[str] phone_num: Phone number of the CAM user.
        :param pulumi.Input[str] remark: Remark of the CAM user.
        :param pulumi.Input[bool] use_api: Indicate whether to generate the API secret key or not.
        """
        if console_login is not None:
            pulumi.set(__self__, "console_login", console_login)
        if country_code is not None:
            pulumi.set(__self__, "country_code", country_code)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if need_reset_password is not None:
            pulumi.set(__self__, "need_reset_password", need_reset_password)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if phone_num is not None:
            pulumi.set(__self__, "phone_num", phone_num)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if use_api is not None:
            pulumi.set(__self__, "use_api", use_api)

    @property
    @pulumi.getter(name="consoleLogin")
    def console_login(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether the CAM user can login to the web console or not.
        """
        return pulumi.get(self, "console_login")

    @console_login.setter
    def console_login(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "console_login", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[pulumi.Input[str]]:
        """
        Country code of the phone number, for example: '86'.
        """
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the CAM user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists;
        otherwise the user will be deleted directly. Default is false.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CAM user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="needResetPassword")
    def need_reset_password(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether the CAM user need to reset the password when first logins.
        """
        return pulumi.get(self, "need_reset_password")

    @need_reset_password.setter
    def need_reset_password(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "need_reset_password", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes
        uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not
        set, a random password will be automatically generated.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="phoneNum")
    def phone_num(self) -> Optional[pulumi.Input[str]]:
        """
        Phone number of the CAM user.
        """
        return pulumi.get(self, "phone_num")

    @phone_num.setter
    def phone_num(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_num", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remark of the CAM user.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="useApi")
    def use_api(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether to generate the API secret key or not.
        """
        return pulumi.get(self, "use_api")

    @use_api.setter
    def use_api(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_api", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 console_login: Optional[pulumi.Input[bool]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_reset_password: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[int]] = None,
                 uin: Optional[pulumi.Input[int]] = None,
                 use_api: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[bool] console_login: Indicate whether the CAM user can login to the web console or not.
        :param pulumi.Input[str] country_code: Country code of the phone number, for example: '86'.
        :param pulumi.Input[str] email: Email of the CAM user.
        :param pulumi.Input[bool] force_delete: Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists;
               otherwise the user will be deleted directly. Default is false.
        :param pulumi.Input[str] name: Name of the CAM user.
        :param pulumi.Input[bool] need_reset_password: Indicate whether the CAM user need to reset the password when first logins.
        :param pulumi.Input[str] password: The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes
               uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not
               set, a random password will be automatically generated.
        :param pulumi.Input[str] phone_num: Phone number of the CAM user.
        :param pulumi.Input[str] remark: Remark of the CAM user.
        :param pulumi.Input[str] secret_id: Secret ID of the CAM user.
        :param pulumi.Input[str] secret_key: Secret key of the CAM user.
        :param pulumi.Input[int] uid: ID of the CAM user.
        :param pulumi.Input[int] uin: Uin of the CAM User.
        :param pulumi.Input[bool] use_api: Indicate whether to generate the API secret key or not.
        """
        if console_login is not None:
            pulumi.set(__self__, "console_login", console_login)
        if country_code is not None:
            pulumi.set(__self__, "country_code", country_code)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if need_reset_password is not None:
            pulumi.set(__self__, "need_reset_password", need_reset_password)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if phone_num is not None:
            pulumi.set(__self__, "phone_num", phone_num)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if uin is not None:
            pulumi.set(__self__, "uin", uin)
        if use_api is not None:
            pulumi.set(__self__, "use_api", use_api)

    @property
    @pulumi.getter(name="consoleLogin")
    def console_login(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether the CAM user can login to the web console or not.
        """
        return pulumi.get(self, "console_login")

    @console_login.setter
    def console_login(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "console_login", value)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> Optional[pulumi.Input[str]]:
        """
        Country code of the phone number, for example: '86'.
        """
        return pulumi.get(self, "country_code")

    @country_code.setter
    def country_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_code", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the CAM user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists;
        otherwise the user will be deleted directly. Default is false.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the CAM user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="needResetPassword")
    def need_reset_password(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether the CAM user need to reset the password when first logins.
        """
        return pulumi.get(self, "need_reset_password")

    @need_reset_password.setter
    def need_reset_password(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "need_reset_password", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes
        uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not
        set, a random password will be automatically generated.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="phoneNum")
    def phone_num(self) -> Optional[pulumi.Input[str]]:
        """
        Phone number of the CAM user.
        """
        return pulumi.get(self, "phone_num")

    @phone_num.setter
    def phone_num(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_num", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remark of the CAM user.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        Secret ID of the CAM user.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret key of the CAM user.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[int]]:
        """
        ID of the CAM user.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter
    def uin(self) -> Optional[pulumi.Input[int]]:
        """
        Uin of the CAM User.
        """
        return pulumi.get(self, "uin")

    @uin.setter
    def uin(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uin", value)

    @property
    @pulumi.getter(name="useApi")
    def use_api(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate whether to generate the API secret key or not.
        """
        return pulumi.get(self, "use_api")

    @use_api.setter
    def use_api(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_api", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 console_login: Optional[pulumi.Input[bool]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_reset_password: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 use_api: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a User resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] console_login: Indicate whether the CAM user can login to the web console or not.
        :param pulumi.Input[str] country_code: Country code of the phone number, for example: '86'.
        :param pulumi.Input[str] email: Email of the CAM user.
        :param pulumi.Input[bool] force_delete: Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists;
               otherwise the user will be deleted directly. Default is false.
        :param pulumi.Input[str] name: Name of the CAM user.
        :param pulumi.Input[bool] need_reset_password: Indicate whether the CAM user need to reset the password when first logins.
        :param pulumi.Input[str] password: The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes
               uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not
               set, a random password will be automatically generated.
        :param pulumi.Input[str] phone_num: Phone number of the CAM user.
        :param pulumi.Input[str] remark: Remark of the CAM user.
        :param pulumi.Input[bool] use_api: Indicate whether to generate the API secret key or not.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a User resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 console_login: Optional[pulumi.Input[bool]] = None,
                 country_code: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 need_reset_password: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 phone_num: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 use_api: Optional[pulumi.Input[bool]] = None,
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
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["console_login"] = console_login
            __props__.__dict__["country_code"] = country_code
            __props__.__dict__["email"] = email
            __props__.__dict__["force_delete"] = force_delete
            __props__.__dict__["name"] = name
            __props__.__dict__["need_reset_password"] = need_reset_password
            __props__.__dict__["password"] = password
            __props__.__dict__["phone_num"] = phone_num
            __props__.__dict__["remark"] = remark
            __props__.__dict__["use_api"] = use_api
            __props__.__dict__["secret_id"] = None
            __props__.__dict__["secret_key"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["uin"] = None
        super(User, __self__).__init__(
            'tencentcloud:Cam/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            console_login: Optional[pulumi.Input[bool]] = None,
            country_code: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            force_delete: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            need_reset_password: Optional[pulumi.Input[bool]] = None,
            password: Optional[pulumi.Input[str]] = None,
            phone_num: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            uid: Optional[pulumi.Input[int]] = None,
            uin: Optional[pulumi.Input[int]] = None,
            use_api: Optional[pulumi.Input[bool]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] console_login: Indicate whether the CAM user can login to the web console or not.
        :param pulumi.Input[str] country_code: Country code of the phone number, for example: '86'.
        :param pulumi.Input[str] email: Email of the CAM user.
        :param pulumi.Input[bool] force_delete: Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists;
               otherwise the user will be deleted directly. Default is false.
        :param pulumi.Input[str] name: Name of the CAM user.
        :param pulumi.Input[bool] need_reset_password: Indicate whether the CAM user need to reset the password when first logins.
        :param pulumi.Input[str] password: The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes
               uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not
               set, a random password will be automatically generated.
        :param pulumi.Input[str] phone_num: Phone number of the CAM user.
        :param pulumi.Input[str] remark: Remark of the CAM user.
        :param pulumi.Input[str] secret_id: Secret ID of the CAM user.
        :param pulumi.Input[str] secret_key: Secret key of the CAM user.
        :param pulumi.Input[int] uid: ID of the CAM user.
        :param pulumi.Input[int] uin: Uin of the CAM User.
        :param pulumi.Input[bool] use_api: Indicate whether to generate the API secret key or not.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["console_login"] = console_login
        __props__.__dict__["country_code"] = country_code
        __props__.__dict__["email"] = email
        __props__.__dict__["force_delete"] = force_delete
        __props__.__dict__["name"] = name
        __props__.__dict__["need_reset_password"] = need_reset_password
        __props__.__dict__["password"] = password
        __props__.__dict__["phone_num"] = phone_num
        __props__.__dict__["remark"] = remark
        __props__.__dict__["secret_id"] = secret_id
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["uid"] = uid
        __props__.__dict__["uin"] = uin
        __props__.__dict__["use_api"] = use_api
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consoleLogin")
    def console_login(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate whether the CAM user can login to the web console or not.
        """
        return pulumi.get(self, "console_login")

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> pulumi.Output[str]:
        """
        Country code of the phone number, for example: '86'.
        """
        return pulumi.get(self, "country_code")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[str]]:
        """
        Email of the CAM user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists;
        otherwise the user will be deleted directly. Default is false.
        """
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the CAM user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="needResetPassword")
    def need_reset_password(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate whether the CAM user need to reset the password when first logins.
        """
        return pulumi.get(self, "need_reset_password")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes
        uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not
        set, a random password will be automatically generated.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="phoneNum")
    def phone_num(self) -> pulumi.Output[Optional[str]]:
        """
        Phone number of the CAM user.
        """
        return pulumi.get(self, "phone_num")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Remark of the CAM user.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        """
        Secret ID of the CAM user.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        Secret key of the CAM user.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[int]:
        """
        ID of the CAM user.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def uin(self) -> pulumi.Output[int]:
        """
        Uin of the CAM User.
        """
        return pulumi.get(self, "uin")

    @property
    @pulumi.getter(name="useApi")
    def use_api(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate whether to generate the API secret key or not.
        """
        return pulumi.get(self, "use_api")

