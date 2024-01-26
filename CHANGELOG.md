# CHANGELOG

## 2024-01-25 v1.23.0

- Based on Comet 23.12.4
- Add TOTP login support

## 2024-01-23 v1.22.0

- Based on Comet 23.12.3

## 2024-01-17 v1.21.0

- Based on Comet 23.12.3

## 2023-12-21 v1.20.0

- Based on Comet 23.12.1

## 2023-12-20 v1.19.0

- Based on Comet 23.11.2

## 2023-11-21 v1.18.0

- Based on Comet 23.9.11

## 2023-11-06 v1.17.0

- Based on Comet 23.9.9
- Add Syncro support

## 2023-11-03 v1.16.0

- Based on Comet 23.9.8

## 2023-10-25 v1.15.0

- Based on Comet 23.9.7
- Add VMware support

## 2023-10-18 v1.14.0

- Based on Comet 23.9.6
- Added Comet Storage Remote Server Type

## 2023-10-12 v1.13.0

- Based on Comet 23.9.5
- Add new field `LogLevel` to control job log verbosity
- Add new `RESTORETYPE_WINDISK_ESXI` for restoring to VMware-compatible
virtual disks

## 2023-09-19 v1.12.0

- Based on Comet 23.9.2
- 'UseObjectLock' for S3 compatible storage settings deprecated. Replaced by 'ObjectLockMode'
- New Streamable event SEVT_DEVICE_LOBBY_CONNECT and SEVT_DEVICE_LOBBY_DISCONNECT
- Added 'TOTPCode' to 'InstallCreds' used for device registration.
- 'GroupedBy' added to 'PSAConfig' for grouping statistics.
- New APIs
	- AdminInstallationDispatchDropConnection
	- AdminInstallationDispatchRegisterDevice
	- AdminInstallationListActive
	- AdminJobAbandon

## 2023-08-29 v1.11.0

- Based on Comet 23.8.0
- Improve documentation of JobStatus constants

## 2023-08-09 v1.10.0

- Based on Comet 23.6.9
- Added WebDAV `DestinationLocation`

## 2023-08-02 v1.9.0

- Based on Comet 23.6.9
- Support new API endpoints for managing external admin authentication sources
- Support Object Lock policy option
- Update docstrings for various types and fields

## 2023-07-11 v1.8.0

- Based on Comet 23.6.5
- Support `DeviceConfig->ClientVersion`
- Support new `OSInfo` and `RestoreJobAdvancedOptions` fields
- Support new `SourceConfig` fields for tracking policy-enforced Protected Items
- Support new optional parameters in `AdminDispatcherRunRestoreCustom` API
- Add many more documentation comments, including behaviour of base64 byte arrays

## 2023-06-01 v1.7.0

- Based on Comet 23.5.0
- Add new `StreamableEventType` constants (`SEVT_*`) used for tracking Comet Server config changes
- Add new `StreamerType` constants (`STREAMER_TYPE_*`)
- Add new `FileOptions` type for configuring the Comet Server to log live events to a file
- Add new field `AuditFileOptions` to the `Organization` type for configuring per-tenant audit log options
- Add `Actor` (authenticated user), `ResourceID`, `Timestamp` and `TypeString` fields to the `StreamableEvent` type
- Deprecate the `UserProfileFragment` type
- Add extra comments to many types

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
