# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-10-31

### Added

- `configmaster`: flexible configuration loader supporting YAML, TOML, JSON, XML
  and ENV files, environment variable binding, defaults, merging and
  type-safe getters with dot-notation key access.
- `errstack`: error construction helpers with stack traces built on
  `github.com/pkg/errors`.
- `responswrapper`: generic REST API response wrappers, including paged
  responses and error helpers.
- `validation`: reflection-based validation helpers for non-empty, numeric
  range and format checks.
- `isp`: internet service provider detection for a client IP using the
  ipinfo.io API.
- Initial project scaffolding: MIT license, README, CI workflow and lint
  configuration.

[Unreleased]: https://github.com/MahdiPezeshkian/Onyx.Common/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/MahdiPezeshkian/Onyx.Common/releases/tag/v0.1.0
