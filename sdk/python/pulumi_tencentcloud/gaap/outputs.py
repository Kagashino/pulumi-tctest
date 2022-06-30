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
    'CertificatesCertificateResult',
    'DomainErrorPageInfoListErrorPageInfoListResult',
    'HttpDomainsDomainResult',
    'HttpRuleRealserver',
    'HttpRulesRuleResult',
    'HttpRulesRuleRealserverResult',
    'Layer4ListenerRealserverBindSet',
    'Layer4ListenersListenerResult',
    'Layer7ListenersListenerResult',
    'ProxiesProxyResult',
    'RealserversRealserverResult',
    'SecurityRulesRuleResult',
]

@pulumi.output_type
class CertificatesCertificateResult(dict):
    def __init__(__self__, *,
                 begin_time: str,
                 create_time: str,
                 end_time: str,
                 id: str,
                 issuer_cn: str,
                 name: str,
                 subject_cn: str,
                 type: str):
        pulumi.set(__self__, "begin_time", begin_time)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "end_time", end_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "issuer_cn", issuer_cn)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "subject_cn", subject_cn)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="beginTime")
    def begin_time(self) -> str:
        return pulumi.get(self, "begin_time")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="issuerCn")
    def issuer_cn(self) -> str:
        return pulumi.get(self, "issuer_cn")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subjectCn")
    def subject_cn(self) -> str:
        return pulumi.get(self, "subject_cn")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class DomainErrorPageInfoListErrorPageInfoListResult(dict):
    def __init__(__self__, *,
                 body: str,
                 clear_headers: Sequence[str],
                 domain: str,
                 error_codes: Sequence[int],
                 id: str,
                 listener_id: str,
                 new_error_codes: int,
                 set_headers: Mapping[str, Any]):
        pulumi.set(__self__, "body", body)
        pulumi.set(__self__, "clear_headers", clear_headers)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "error_codes", error_codes)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "new_error_codes", new_error_codes)
        pulumi.set(__self__, "set_headers", set_headers)

    @property
    @pulumi.getter
    def body(self) -> str:
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="clearHeaders")
    def clear_headers(self) -> Sequence[str]:
        return pulumi.get(self, "clear_headers")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="errorCodes")
    def error_codes(self) -> Sequence[int]:
        return pulumi.get(self, "error_codes")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="newErrorCodes")
    def new_error_codes(self) -> int:
        return pulumi.get(self, "new_error_codes")

    @property
    @pulumi.getter(name="setHeaders")
    def set_headers(self) -> Mapping[str, Any]:
        return pulumi.get(self, "set_headers")


@pulumi.output_type
class HttpDomainsDomainResult(dict):
    def __init__(__self__, *,
                 basic_auth: bool,
                 basic_auth_id: str,
                 certificate_id: str,
                 client_certificate_id: str,
                 client_certificate_ids: Sequence[str],
                 domain: str,
                 gaap_auth: bool,
                 gaap_auth_id: str,
                 realserver_auth: bool,
                 realserver_certificate_domain: str,
                 realserver_certificate_id: str,
                 realserver_certificate_ids: Sequence[str]):
        pulumi.set(__self__, "basic_auth", basic_auth)
        pulumi.set(__self__, "basic_auth_id", basic_auth_id)
        pulumi.set(__self__, "certificate_id", certificate_id)
        pulumi.set(__self__, "client_certificate_id", client_certificate_id)
        pulumi.set(__self__, "client_certificate_ids", client_certificate_ids)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "gaap_auth", gaap_auth)
        pulumi.set(__self__, "gaap_auth_id", gaap_auth_id)
        pulumi.set(__self__, "realserver_auth", realserver_auth)
        pulumi.set(__self__, "realserver_certificate_domain", realserver_certificate_domain)
        pulumi.set(__self__, "realserver_certificate_id", realserver_certificate_id)
        pulumi.set(__self__, "realserver_certificate_ids", realserver_certificate_ids)

    @property
    @pulumi.getter(name="basicAuth")
    def basic_auth(self) -> bool:
        return pulumi.get(self, "basic_auth")

    @property
    @pulumi.getter(name="basicAuthId")
    def basic_auth_id(self) -> str:
        return pulumi.get(self, "basic_auth_id")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="clientCertificateId")
    def client_certificate_id(self) -> str:
        return pulumi.get(self, "client_certificate_id")

    @property
    @pulumi.getter(name="clientCertificateIds")
    def client_certificate_ids(self) -> Sequence[str]:
        return pulumi.get(self, "client_certificate_ids")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="gaapAuth")
    def gaap_auth(self) -> bool:
        return pulumi.get(self, "gaap_auth")

    @property
    @pulumi.getter(name="gaapAuthId")
    def gaap_auth_id(self) -> str:
        return pulumi.get(self, "gaap_auth_id")

    @property
    @pulumi.getter(name="realserverAuth")
    def realserver_auth(self) -> bool:
        return pulumi.get(self, "realserver_auth")

    @property
    @pulumi.getter(name="realserverCertificateDomain")
    def realserver_certificate_domain(self) -> str:
        return pulumi.get(self, "realserver_certificate_domain")

    @property
    @pulumi.getter(name="realserverCertificateId")
    def realserver_certificate_id(self) -> str:
        return pulumi.get(self, "realserver_certificate_id")

    @property
    @pulumi.getter(name="realserverCertificateIds")
    def realserver_certificate_ids(self) -> Sequence[str]:
        return pulumi.get(self, "realserver_certificate_ids")


@pulumi.output_type
class HttpRuleRealserver(dict):
    def __init__(__self__, *,
                 id: str,
                 ip: str,
                 port: int,
                 weight: Optional[int] = None):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def weight(self) -> Optional[int]:
        return pulumi.get(self, "weight")


@pulumi.output_type
class HttpRulesRuleResult(dict):
    def __init__(__self__, *,
                 connect_timeout: int,
                 domain: str,
                 forward_host: str,
                 health_check: bool,
                 health_check_method: str,
                 health_check_path: str,
                 health_check_status_codes: Sequence[int],
                 id: str,
                 interval: int,
                 listener_id: str,
                 path: str,
                 realserver_type: str,
                 realservers: Sequence['outputs.HttpRulesRuleRealserverResult'],
                 scheduler: str,
                 sni: str,
                 sni_switch: str):
        pulumi.set(__self__, "connect_timeout", connect_timeout)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "forward_host", forward_host)
        pulumi.set(__self__, "health_check", health_check)
        pulumi.set(__self__, "health_check_method", health_check_method)
        pulumi.set(__self__, "health_check_path", health_check_path)
        pulumi.set(__self__, "health_check_status_codes", health_check_status_codes)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "realserver_type", realserver_type)
        pulumi.set(__self__, "realservers", realservers)
        pulumi.set(__self__, "scheduler", scheduler)
        pulumi.set(__self__, "sni", sni)
        pulumi.set(__self__, "sni_switch", sni_switch)

    @property
    @pulumi.getter(name="connectTimeout")
    def connect_timeout(self) -> int:
        return pulumi.get(self, "connect_timeout")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="forwardHost")
    def forward_host(self) -> str:
        return pulumi.get(self, "forward_host")

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> bool:
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter(name="healthCheckMethod")
    def health_check_method(self) -> str:
        return pulumi.get(self, "health_check_method")

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> str:
        return pulumi.get(self, "health_check_path")

    @property
    @pulumi.getter(name="healthCheckStatusCodes")
    def health_check_status_codes(self) -> Sequence[int]:
        return pulumi.get(self, "health_check_status_codes")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interval(self) -> int:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="realserverType")
    def realserver_type(self) -> str:
        return pulumi.get(self, "realserver_type")

    @property
    @pulumi.getter
    def realservers(self) -> Sequence['outputs.HttpRulesRuleRealserverResult']:
        return pulumi.get(self, "realservers")

    @property
    @pulumi.getter
    def scheduler(self) -> str:
        return pulumi.get(self, "scheduler")

    @property
    @pulumi.getter
    def sni(self) -> str:
        return pulumi.get(self, "sni")

    @property
    @pulumi.getter(name="sniSwitch")
    def sni_switch(self) -> str:
        return pulumi.get(self, "sni_switch")


@pulumi.output_type
class HttpRulesRuleRealserverResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 id: str,
                 ip: str,
                 port: int,
                 status: int,
                 weight: int):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def status(self) -> int:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def weight(self) -> int:
        return pulumi.get(self, "weight")


@pulumi.output_type
class Layer4ListenerRealserverBindSet(dict):
    def __init__(__self__, *,
                 id: str,
                 ip: str,
                 port: int,
                 weight: Optional[int] = None):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def weight(self) -> Optional[int]:
        return pulumi.get(self, "weight")


@pulumi.output_type
class Layer4ListenersListenerResult(dict):
    def __init__(__self__, *,
                 connect_timeout: int,
                 create_time: str,
                 health_check: bool,
                 id: str,
                 interval: int,
                 name: str,
                 port: int,
                 protocol: str,
                 realserver_type: str,
                 scheduler: str,
                 status: int):
        pulumi.set(__self__, "connect_timeout", connect_timeout)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "health_check", health_check)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "realserver_type", realserver_type)
        pulumi.set(__self__, "scheduler", scheduler)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="connectTimeout")
    def connect_timeout(self) -> int:
        return pulumi.get(self, "connect_timeout")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> bool:
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interval(self) -> int:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="realserverType")
    def realserver_type(self) -> str:
        return pulumi.get(self, "realserver_type")

    @property
    @pulumi.getter
    def scheduler(self) -> str:
        return pulumi.get(self, "scheduler")

    @property
    @pulumi.getter
    def status(self) -> int:
        return pulumi.get(self, "status")


@pulumi.output_type
class Layer7ListenersListenerResult(dict):
    def __init__(__self__, *,
                 auth_type: int,
                 certificate_id: str,
                 client_certificate_id: str,
                 client_certificate_ids: Sequence[str],
                 create_time: str,
                 forward_protocol: str,
                 id: str,
                 name: str,
                 port: int,
                 protocol: str,
                 status: int):
        pulumi.set(__self__, "auth_type", auth_type)
        pulumi.set(__self__, "certificate_id", certificate_id)
        pulumi.set(__self__, "client_certificate_id", client_certificate_id)
        pulumi.set(__self__, "client_certificate_ids", client_certificate_ids)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "forward_protocol", forward_protocol)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> int:
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="clientCertificateId")
    def client_certificate_id(self) -> str:
        return pulumi.get(self, "client_certificate_id")

    @property
    @pulumi.getter(name="clientCertificateIds")
    def client_certificate_ids(self) -> Sequence[str]:
        return pulumi.get(self, "client_certificate_ids")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="forwardProtocol")
    def forward_protocol(self) -> str:
        return pulumi.get(self, "forward_protocol")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def status(self) -> int:
        return pulumi.get(self, "status")


@pulumi.output_type
class ProxiesProxyResult(dict):
    def __init__(__self__, *,
                 access_region: str,
                 bandwidth: int,
                 concurrent: int,
                 create_time: str,
                 domain: str,
                 forward_ip: str,
                 id: str,
                 ip: str,
                 name: str,
                 policy_id: str,
                 project_id: int,
                 realserver_region: str,
                 scalable: bool,
                 status: str,
                 support_protocols: Sequence[str],
                 tags: Mapping[str, Any],
                 version: str):
        pulumi.set(__self__, "access_region", access_region)
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "concurrent", concurrent)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "forward_ip", forward_ip)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "realserver_region", realserver_region)
        pulumi.set(__self__, "scalable", scalable)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "support_protocols", support_protocols)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="accessRegion")
    def access_region(self) -> str:
        return pulumi.get(self, "access_region")

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def concurrent(self) -> int:
        return pulumi.get(self, "concurrent")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="forwardIp")
    def forward_ip(self) -> str:
        return pulumi.get(self, "forward_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> str:
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="realserverRegion")
    def realserver_region(self) -> str:
        return pulumi.get(self, "realserver_region")

    @property
    @pulumi.getter
    def scalable(self) -> bool:
        return pulumi.get(self, "scalable")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="supportProtocols")
    def support_protocols(self) -> Sequence[str]:
        return pulumi.get(self, "support_protocols")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> str:
        return pulumi.get(self, "version")


@pulumi.output_type
class RealserversRealserverResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 id: str,
                 ip: str,
                 name: str,
                 project_id: int,
                 tags: Mapping[str, Any]):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        return pulumi.get(self, "tags")


@pulumi.output_type
class SecurityRulesRuleResult(dict):
    def __init__(__self__, *,
                 action: str,
                 cidr_ip: str,
                 id: str,
                 name: str,
                 port: str,
                 protocol: str):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "cidr_ip", cidr_ip)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> str:
        return pulumi.get(self, "cidr_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> str:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")


