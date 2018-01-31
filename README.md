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

## Implemented specs & features

The following client/RP features from OpenID Connect/OAuth2.0 specifications are implemented by
caddy-oidc.

- [OpenID Connect Core 1.0 incorporating errata set 1][feature-core]
  - Authorization Callback
    - Authorization Code Flow
  - UserInfo Request
  - Offline Access / Refresh Token Grant
  - Client Authentication
    - client_secret_basic
    - client_secret_post
- [OpenID Connect Discovery 1.0 incorporating errata set 1][feature-discovery]
  - Discovery of OpenID Provider (Issuer) Metadata

[travis-image]: https://img.shields.io/travis/panva/caddy-oidc/master.svg?style=flat-square&maxAge=7200
[travis-url]: https://travis-ci.org/panva/caddy-oidc
[codecov-image]: https://img.shields.io/codecov/c/github/panva/caddy-oidc/master.svg?style=flat-square&maxAge=7200
[codecov-url]: https://codecov.io/gh/panva/caddy-oidc
[feature-core]: http://openid.net/specs/openid-connect-core-1_0.html
[feature-discovery]: http://openid.net/specs/openid-connect-discovery-1_0.html
