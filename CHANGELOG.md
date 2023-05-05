# CHANGELOG

## 2023-05-05 v1.6.0

- Based on Comet 23.3.7
- No functional changes
- Add significantly many more comments to constants, fields and types

## 2023-04-24 v1.5.0

- Based on Comet 23.3.5
- Support new `AdminDispatcherSearchSnapshots` API to remotely search for files in a Storage Vault
- Support new `AdminMetaRemoteStorageVaultTest` API to test connections for a Storage Template
- Support new `AccentColor` and `BrandingStyleType` branding options for the Comet Server web interface
- New Self-Backup option to include server logs
- Track creation and modification timestamps for `GroupPolicy` objects

## 2023-03-23 v1.4.0

- Based on Comet 23.3.1
- Support filter parameters on `AdminGetJobLogEntries`
- Support S3 Object Lock
- Support Azure Key Vault codesigning

## 2023-02-15 v1.3.0
- Based on 22.12.8
- Add `TimeSpan` option to `EmailReportOptions`
- Add `AlertsDisabled` (default: false) toggle for `PSAConfig` objects
- Add `LastSuspended` for tracking `UserProfileConfig` suspensions
- Improve documentation

## 2023-01-11 v1.2.0
- Based on 22.12.2
- Add new `AdminCountJobsForCustomSearchRequest` API to count total number of jobs for a given search query
- Add new `AdminMetaEmailOptionsGetRequest` / `AdminMetaEmailOptionsSetRequest` / `AdminMetaSendTestReportRequest` APIs for managing tenant email settings
- Add new O365 Protected Accounts quota option in the `UserProfileConfig`
- Add new tenant admin permission `AllowEditEmailOptions`

# 2022-12-09 v1.1.0
- Based on 22.11.1
- New features for PSAs, remote URLs and MS SQL Server restores.
- New features for exporting a self backup for single tenant.
- Fix an issue with detecting errors in API responses.

## 2022-09-08 v1.0.1
- Now based on Comet Server 22.9.0

## 2022-08-24 v1.0.0
- Initial public release

## 2022-08-21 v0.0.2
- Added UserWeb examples
- Moved cometsdk to its own lib module
- Added fix for some marshalling issues

## 2022-07-28 v0.0.1
- Based on 22.6.8
- Initial release
