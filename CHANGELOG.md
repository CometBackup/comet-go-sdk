# CHANGELOG

## 2025-11-14 v2.6.0

- Based on Comet 25.9.6
- Add new API methods for getting, setting, and deleting Protected Items and
their associated schedules
- Add new API methods for live connection credential/session management
- Add JobID field to API response when starting jobs via the dispatcher on
devices running Comet 25.9.6 or newer
- Add support for custom tags on backup jobs to allow grouping of
snapshots/jobs
- Add support SAS Relic codesigning
- Improve documentation for Proxmox API types

## 2025-08-15 v2.5.0

- Based on Comet 25.6.8
- Add support for Proxmox
- Add profile hash to some API endpoints

## 2025-03-20 v2.4.0

- Based on Comet 25.3.1
- Support yearly retention rules
- Support yearly schedule rules
- Support storage vault key rotation

## 2025-02-17 v2.3.0
Update to 24.12.5

- Based on Comet 24.12.5

## 2025-01-29 v2.2.0

- Based on Comet 24.12.4

## 2025-01-22 v2.1.0

- Based on Comet 24.12.3
- Additional properties for BrandingOptions, BrandingProperties and ServerMetaBrandingProperties

## 2025-01-10 v2.0.0

This version is a semver-MAJOR upgrade. To use this release, update your Go module import path from `"github.com/CometBackup/comet-go-sdk"` to `"github.com/CometBackup/comet-go-sdk/v2"`.

- Based on Comet 24.12.2
- **BREAKING:** Update all client methods to take context parameters
- Add dependency on easyjson to preserve unknown fields from future versions of Comet Server
- Support new Login Protection feature
- Support requiring password changes for admin accounts
- Support new user grouping functionality
- Supoprt creating a first admin account on empty servers
- Support vault-device associations and automatic vault creation

## 2024-10-29 v1.36.0

- Based on Comet 24.9.6
- New API AdminDispatcherTestSmbAuth to instruct a device to test Windows SMB credentials
- New RESTORETYPE_WINDISK_VHDX to restore Disk Image backup as Hyper-V VHDX format
- BackupJobDetail supports ConflictingJobID field to indicate conflicting job if a lock error occurred
- BackupJobProgress can report the total number of items for progress visualization
- DiskDrive adds information about parent disks on Linux

## 2024-09-17 v1.35.0

- Based on Comet 24.9.1
- Added Hyper-V Guest limits per user
- Added VMware Guest limits per user
- Added protected item type (engine) to protected item class
- Added Force Overwrite Restore permissions

## 2024-08-23 v1.34.0

- Based on Comet 24.6.10
- Add TOTP and SessionKey support for multipart/form-data requests

## 2024-08-01 v1.33.0

- Based on Comet 24.6.6
- Add new AdminConvertStorageRole API
- Update data types for new job retry feature (BackupJobDetail, BackupRuleEventTriggers, and new JOB_STATUS_RUNNING_TRYAGAIN)
- Add support for custom Prefix in S3GenericVirtualStorageRole

## 2024-07-16 v1.32.0

- Based on Comet 24.6.4
- Added Server Device and Booster Limits
- Added API to count devices registered on a Server
- Added Software Build Role configuration per tenant

## 2024-06-11 v1.31.0

- Based on Comet 24.6.0
- Added custom HELO/EHLO STMP support
- Added support to set if disabled Office 365 accounts should be backed up
- Added suport for custom headers like User-Agent
- Added support to set default values for User-Agent
- Added EngineType as a property of VaultSnapshots

## 2024-05-31 v1.30.0

- Based on Comet 24.5.0
- Added support for matching Microsoft Office 365 users and groups by Drive ID
- Added support for configuring Object Lock on a Impossible Cloud Partner API Storage Template
- Added support for Custom Body Date when using a Custom Remote Bucket Storage Template

## 2024-05-24 v1.29.0

- Based on Comet 24.3.9

## 2024-05-17 v1.28.0

- Based on Comet 24.3.8
- Added support for configuring Software Build Role builder count
- Added LastStartTime field to job statistics

## 2024-05-14 v1.27.0

- Based on Comet 24.3.7
- Added support for configuring concurrency in Microsoft 365 Protected Items

## 2024-05-02 v1.26.0

- No changes

## 2024-05-02 v1.25.0

- Based on Comet 24.3.6
- Testing new release process

## 2024-04-12 v1.24.0

- Based on Comet 24.3.5

## 2024-01-31 v1.23.0

- Based on Comet 23.12.5-rc
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
- Add new `RESTORETYPE_WINDISK_ESXI` for restoring to VMware-compatible virtual disks

## 2023-09-19 v1.12.0

- Based on Comet 23.9.2
- 'UseObjectLock' for S3 compatible storage settings deprecated. Replaced by 'ObjectLockMode'
- New Streamable event SEVT_DEVICE_LOBBY_CONNECT and SEVT_DEVICE_LOBBY_DISCONNECT
- Added 'TOTPCode' to 'InstallCreds' used for device registration.
- 'GroupedBy' added to 'PSAConfig' for grouping statistics.
- New APIs: AdminInstallationDispatchDropConnection, AdminInstallationDispatchRegisterDevice, AdminInstallationListActiveAdminJobAbandon

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
