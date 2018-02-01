# caddy-oidc

[![build][travis-image]][travis-url] [![codecov][codecov-image]][codecov-url]

OpenID Connect Relying Party and OAuth 2.0 Resource Server for Caddy HTTP Server

**Table of Contents**

<!-- TOC START min:2 max:2 link:true update:true -->
  - [Status](#status)
  - [Certification](#certification)
  - [Implemented specs & features](#implemented-specs--features)

<!-- TOC END -->

## Status

Still very much in under development.

## Certification

Later.

## Planned specs & features

The following client/RP features from OpenID Connect/OAuth2.0 specifications are implemented by
caddy-oidc.

- [OpenID Connect Core 1.0 incorporating errata set 1][feature-core]
  - Authorization Flows
    - Authorization Code Flow
    - Implicit Flow
    - Hybrid Flow
  - UserInfo Request
  - Offline Access / Refresh Token Grant
  - Client Authentication
    - client_secret_basic
    - client_secret_post
- [OpenID Connect Discovery 1.0 incorporating errata set 1][feature-discovery]
  - Discovery of OpenID Provider (Issuer) Metadata
- [OpenID Connect Dynamic Client Registration 1.0 incorporating errata set 1][feature-registration]
  - Client metadata discovery via Client Configuration Endpoint

The following drafts/experimental specifications are implemented by caddy-oidc.
  - [OpenID Connect Session Management 1.0 - draft 28][feature-session-management]
  - [OpenID Connect Back-Channel Logout 1.0 - draft 04][feature-backchannel-logout]
  - [OpenID Connect Front-Channel Logout 1.0 - draft 02][feature-frontchannel-logout]


[travis-image]: https://img.shields.io/travis/panva/caddy-oidc/master.svg?style=flat-square&maxAge=7200
[travis-url]: https://travis-ci.org/panva/caddy-oidc
[codecov-image]: https://img.shields.io/codecov/c/github/panva/caddy-oidc/master.svg?style=flat-square&maxAge=7200
[codecov-url]: https://codecov.io/gh/panva/caddy-oidc
[feature-core]: https://openid.net/specs/openid-connect-core-1_0.html
[feature-discovery]: https://openid.net/specs/openid-connect-discovery-1_0.html
[feature-registration]: https://openid.net/specs/openid-connect-registration-1_0.html
[feature-backchannel-logout]: https://openid.net/specs/openid-connect-backchannel-1_0-04.html
[feature-frontchannel-logout]: https://openid.net/specs/openid-connect-frontchannel-1_0-02.html
[feature-session-management]: https://openid.net/specs/openid-connect-session-1_0-28.html
