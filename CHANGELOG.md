# CHANGELOG

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
