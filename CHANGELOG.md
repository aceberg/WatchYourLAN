# Change Log
All notable changes to this project will be documented in this file.

## [v2.0.1] - 2024-09-01
### Added
- Vlans and docker0 support [#47](https://github.com/aceberg/WatchYourLAN/issues/47). Thanks [thehijacker](https://github.com/thehijacker) for help with this issue!
- Remember `sort` field
- InfluxDB error handling

### Fixed
- Bug [#103](https://github.com/aceberg/WatchYourLAN/issues/103)
- Bug [#104](https://github.com/aceberg/WatchYourLAN/issues/104)

## [v2.0.0] - 2024-08-30
### Added
- API
- Arguments for `arp-scan` option
- `InfluxDB` export
- `PostgreSQL` or `SQLite` DB options
- Names from DNS

### Changed
- Better UI with JS
- Switched to `gin` web framework
- Reworked DB schema and config variables

