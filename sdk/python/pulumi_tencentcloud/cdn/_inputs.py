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
    'CdnDomainHttpsConfigArgs',
    'CdnDomainHttpsConfigClientCertificateConfigArgs',
    'CdnDomainHttpsConfigForceRedirectArgs',
    'CdnDomainHttpsConfigServerCertificateConfigArgs',
    'CdnDomainOriginArgs',
    'CdnDomainRequestHeaderArgs',
    'CdnDomainRequestHeaderHeaderRuleArgs',
    'CdnDomainRuleCachArgs',
]

@pulumi.input_type
class CdnDomainHttpsConfigArgs:
    def __init__(__self__, *,
                 https_switch: pulumi.Input[str],
                 client_certificate_config: Optional[pulumi.Input['CdnDomainHttpsConfigClientCertificateConfigArgs']] = None,
                 force_redirect: Optional[pulumi.Input['CdnDomainHttpsConfigForceRedirectArgs']] = None,
                 http2_switch: Optional[pulumi.Input[str]] = None,
                 ocsp_stapling_switch: Optional[pulumi.Input[str]] = None,
                 server_certificate_config: Optional[pulumi.Input['CdnDomainHttpsConfigServerCertificateConfigArgs']] = None,
                 spdy_switch: Optional[pulumi.Input[str]] = None,
                 verify_client: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "https_switch", https_switch)
        if client_certificate_config is not None:
            pulumi.set(__self__, "client_certificate_config", client_certificate_config)
        if force_redirect is not None:
            pulumi.set(__self__, "force_redirect", force_redirect)
        if http2_switch is not None:
            pulumi.set(__self__, "http2_switch", http2_switch)
        if ocsp_stapling_switch is not None:
            pulumi.set(__self__, "ocsp_stapling_switch", ocsp_stapling_switch)
        if server_certificate_config is not None:
            pulumi.set(__self__, "server_certificate_config", server_certificate_config)
        if spdy_switch is not None:
            pulumi.set(__self__, "spdy_switch", spdy_switch)
        if verify_client is not None:
            pulumi.set(__self__, "verify_client", verify_client)

    @property
    @pulumi.getter(name="httpsSwitch")
    def https_switch(self) -> pulumi.Input[str]:
        return pulumi.get(self, "https_switch")

    @https_switch.setter
    def https_switch(self, value: pulumi.Input[str]):
        pulumi.set(self, "https_switch", value)

    @property
    @pulumi.getter(name="clientCertificateConfig")
    def client_certificate_config(self) -> Optional[pulumi.Input['CdnDomainHttpsConfigClientCertificateConfigArgs']]:
        return pulumi.get(self, "client_certificate_config")

    @client_certificate_config.setter
    def client_certificate_config(self, value: Optional[pulumi.Input['CdnDomainHttpsConfigClientCertificateConfigArgs']]):
        pulumi.set(self, "client_certificate_config", value)

    @property
    @pulumi.getter(name="forceRedirect")
    def force_redirect(self) -> Optional[pulumi.Input['CdnDomainHttpsConfigForceRedirectArgs']]:
        return pulumi.get(self, "force_redirect")

    @force_redirect.setter
    def force_redirect(self, value: Optional[pulumi.Input['CdnDomainHttpsConfigForceRedirectArgs']]):
        pulumi.set(self, "force_redirect", value)

    @property
    @pulumi.getter(name="http2Switch")
    def http2_switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "http2_switch")

    @http2_switch.setter
    def http2_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http2_switch", value)

    @property
    @pulumi.getter(name="ocspStaplingSwitch")
    def ocsp_stapling_switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ocsp_stapling_switch")

    @ocsp_stapling_switch.setter
    def ocsp_stapling_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ocsp_stapling_switch", value)

    @property
    @pulumi.getter(name="serverCertificateConfig")
    def server_certificate_config(self) -> Optional[pulumi.Input['CdnDomainHttpsConfigServerCertificateConfigArgs']]:
        return pulumi.get(self, "server_certificate_config")

    @server_certificate_config.setter
    def server_certificate_config(self, value: Optional[pulumi.Input['CdnDomainHttpsConfigServerCertificateConfigArgs']]):
        pulumi.set(self, "server_certificate_config", value)

    @property
    @pulumi.getter(name="spdySwitch")
    def spdy_switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "spdy_switch")

    @spdy_switch.setter
    def spdy_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spdy_switch", value)

    @property
    @pulumi.getter(name="verifyClient")
    def verify_client(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "verify_client")

    @verify_client.setter
    def verify_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verify_client", value)


@pulumi.input_type
class CdnDomainHttpsConfigClientCertificateConfigArgs:
    def __init__(__self__, *,
                 certificate_content: pulumi.Input[str],
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 deploy_time: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "certificate_content", certificate_content)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if deploy_time is not None:
            pulumi.set(__self__, "deploy_time", deploy_time)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)

    @property
    @pulumi.getter(name="certificateContent")
    def certificate_content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "certificate_content")

    @certificate_content.setter
    def certificate_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_content", value)

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_name")

    @certificate_name.setter
    def certificate_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_name", value)

    @property
    @pulumi.getter(name="deployTime")
    def deploy_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "deploy_time")

    @deploy_time.setter
    def deploy_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_time", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)


@pulumi.input_type
class CdnDomainHttpsConfigForceRedirectArgs:
    def __init__(__self__, *,
                 redirect_status_code: Optional[pulumi.Input[int]] = None,
                 redirect_type: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None):
        if redirect_status_code is not None:
            pulumi.set(__self__, "redirect_status_code", redirect_status_code)
        if redirect_type is not None:
            pulumi.set(__self__, "redirect_type", redirect_type)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="redirectStatusCode")
    def redirect_status_code(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "redirect_status_code")

    @redirect_status_code.setter
    def redirect_status_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "redirect_status_code", value)

    @property
    @pulumi.getter(name="redirectType")
    def redirect_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "redirect_type")

    @redirect_type.setter
    def redirect_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redirect_type", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch", value)


@pulumi.input_type
class CdnDomainHttpsConfigServerCertificateConfigArgs:
    def __init__(__self__, *,
                 certificate_content: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 deploy_time: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None):
        if certificate_content is not None:
            pulumi.set(__self__, "certificate_content", certificate_content)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if deploy_time is not None:
            pulumi.set(__self__, "deploy_time", deploy_time)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)

    @property
    @pulumi.getter(name="certificateContent")
    def certificate_content(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_content")

    @certificate_content.setter
    def certificate_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_content", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_name")

    @certificate_name.setter
    def certificate_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_name", value)

    @property
    @pulumi.getter(name="deployTime")
    def deploy_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "deploy_time")

    @deploy_time.setter
    def deploy_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_time", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)


@pulumi.input_type
class CdnDomainOriginArgs:
    def __init__(__self__, *,
                 origin_lists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 origin_type: pulumi.Input[str],
                 backup_origin_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backup_origin_type: Optional[pulumi.Input[str]] = None,
                 backup_server_name: Optional[pulumi.Input[str]] = None,
                 cos_private_access: Optional[pulumi.Input[str]] = None,
                 origin_pull_protocol: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "origin_lists", origin_lists)
        pulumi.set(__self__, "origin_type", origin_type)
        if backup_origin_lists is not None:
            pulumi.set(__self__, "backup_origin_lists", backup_origin_lists)
        if backup_origin_type is not None:
            pulumi.set(__self__, "backup_origin_type", backup_origin_type)
        if backup_server_name is not None:
            pulumi.set(__self__, "backup_server_name", backup_server_name)
        if cos_private_access is not None:
            pulumi.set(__self__, "cos_private_access", cos_private_access)
        if origin_pull_protocol is not None:
            pulumi.set(__self__, "origin_pull_protocol", origin_pull_protocol)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)

    @property
    @pulumi.getter(name="originLists")
    def origin_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "origin_lists")

    @origin_lists.setter
    def origin_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "origin_lists", value)

    @property
    @pulumi.getter(name="originType")
    def origin_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "origin_type")

    @origin_type.setter
    def origin_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "origin_type", value)

    @property
    @pulumi.getter(name="backupOriginLists")
    def backup_origin_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "backup_origin_lists")

    @backup_origin_lists.setter
    def backup_origin_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "backup_origin_lists", value)

    @property
    @pulumi.getter(name="backupOriginType")
    def backup_origin_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "backup_origin_type")

    @backup_origin_type.setter
    def backup_origin_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_origin_type", value)

    @property
    @pulumi.getter(name="backupServerName")
    def backup_server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "backup_server_name")

    @backup_server_name.setter
    def backup_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_server_name", value)

    @property
    @pulumi.getter(name="cosPrivateAccess")
    def cos_private_access(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cos_private_access")

    @cos_private_access.setter
    def cos_private_access(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cos_private_access", value)

    @property
    @pulumi.getter(name="originPullProtocol")
    def origin_pull_protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "origin_pull_protocol")

    @origin_pull_protocol.setter
    def origin_pull_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_pull_protocol", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)


@pulumi.input_type
class CdnDomainRequestHeaderArgs:
    def __init__(__self__, *,
                 header_rules: Optional[pulumi.Input[Sequence[pulumi.Input['CdnDomainRequestHeaderHeaderRuleArgs']]]] = None,
                 switch: Optional[pulumi.Input[str]] = None):
        if header_rules is not None:
            pulumi.set(__self__, "header_rules", header_rules)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="headerRules")
    def header_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CdnDomainRequestHeaderHeaderRuleArgs']]]]:
        return pulumi.get(self, "header_rules")

    @header_rules.setter
    def header_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CdnDomainRequestHeaderHeaderRuleArgs']]]]):
        pulumi.set(self, "header_rules", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch", value)


@pulumi.input_type
class CdnDomainRequestHeaderHeaderRuleArgs:
    def __init__(__self__, *,
                 header_mode: pulumi.Input[str],
                 header_name: pulumi.Input[str],
                 header_value: pulumi.Input[str],
                 rule_paths: pulumi.Input[Sequence[pulumi.Input[str]]],
                 rule_type: pulumi.Input[str]):
        pulumi.set(__self__, "header_mode", header_mode)
        pulumi.set(__self__, "header_name", header_name)
        pulumi.set(__self__, "header_value", header_value)
        pulumi.set(__self__, "rule_paths", rule_paths)
        pulumi.set(__self__, "rule_type", rule_type)

    @property
    @pulumi.getter(name="headerMode")
    def header_mode(self) -> pulumi.Input[str]:
        return pulumi.get(self, "header_mode")

    @header_mode.setter
    def header_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_mode", value)

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "header_name")

    @header_name.setter
    def header_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_name", value)

    @property
    @pulumi.getter(name="headerValue")
    def header_value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "header_value")

    @header_value.setter
    def header_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_value", value)

    @property
    @pulumi.getter(name="rulePaths")
    def rule_paths(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "rule_paths")

    @rule_paths.setter
    def rule_paths(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "rule_paths", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_type", value)


@pulumi.input_type
class CdnDomainRuleCachArgs:
    def __init__(__self__, *,
                 cache_time: pulumi.Input[int],
                 compare_max_age: Optional[pulumi.Input[str]] = None,
                 follow_origin_switch: Optional[pulumi.Input[str]] = None,
                 ignore_cache_control: Optional[pulumi.Input[str]] = None,
                 ignore_set_cookie: Optional[pulumi.Input[str]] = None,
                 no_cache_switch: Optional[pulumi.Input[str]] = None,
                 re_validate: Optional[pulumi.Input[str]] = None,
                 rule_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rule_type: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "cache_time", cache_time)
        if compare_max_age is not None:
            pulumi.set(__self__, "compare_max_age", compare_max_age)
        if follow_origin_switch is not None:
            pulumi.set(__self__, "follow_origin_switch", follow_origin_switch)
        if ignore_cache_control is not None:
            pulumi.set(__self__, "ignore_cache_control", ignore_cache_control)
        if ignore_set_cookie is not None:
            pulumi.set(__self__, "ignore_set_cookie", ignore_set_cookie)
        if no_cache_switch is not None:
            pulumi.set(__self__, "no_cache_switch", no_cache_switch)
        if re_validate is not None:
            pulumi.set(__self__, "re_validate", re_validate)
        if rule_paths is not None:
            pulumi.set(__self__, "rule_paths", rule_paths)
        if rule_type is not None:
            pulumi.set(__self__, "rule_type", rule_type)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="cacheTime")
    def cache_time(self) -> pulumi.Input[int]:
        return pulumi.get(self, "cache_time")

    @cache_time.setter
    def cache_time(self, value: pulumi.Input[int]):
        pulumi.set(self, "cache_time", value)

    @property
    @pulumi.getter(name="compareMaxAge")
    def compare_max_age(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "compare_max_age")

    @compare_max_age.setter
    def compare_max_age(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compare_max_age", value)

    @property
    @pulumi.getter(name="followOriginSwitch")
    def follow_origin_switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "follow_origin_switch")

    @follow_origin_switch.setter
    def follow_origin_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "follow_origin_switch", value)

    @property
    @pulumi.getter(name="ignoreCacheControl")
    def ignore_cache_control(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ignore_cache_control")

    @ignore_cache_control.setter
    def ignore_cache_control(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ignore_cache_control", value)

    @property
    @pulumi.getter(name="ignoreSetCookie")
    def ignore_set_cookie(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ignore_set_cookie")

    @ignore_set_cookie.setter
    def ignore_set_cookie(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ignore_set_cookie", value)

    @property
    @pulumi.getter(name="noCacheSwitch")
    def no_cache_switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "no_cache_switch")

    @no_cache_switch.setter
    def no_cache_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "no_cache_switch", value)

    @property
    @pulumi.getter(name="reValidate")
    def re_validate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "re_validate")

    @re_validate.setter
    def re_validate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "re_validate", value)

    @property
    @pulumi.getter(name="rulePaths")
    def rule_paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "rule_paths")

    @rule_paths.setter
    def rule_paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rule_paths", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_type", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch", value)

