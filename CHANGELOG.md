# Change Log
All notable changes to this project will be documented in this file.

## [v2.0.4] - 2024-
### Fixed
- [#101](https://github.com/aceberg/WatchYourLAN/issues/101) 
- The same problem for Theme, Color mode, Log level

## [v2.0.3] - 2024-09-17
### Fixed
- `ARP_STRS_JOINED` should be empty in config file
- Optimized History Trim

## [v2.0.2] - 2024-09-07
### Added
- Remember Refresh setting in browser [#123](https://github.com/aceberg/WatchYourLAN/issues/123)

### Fixed
- Error when `IFACES` are empty
- Sticky sort bug fix
- Bug [#124](https://github.com/aceberg/WatchYourLAN/issues/124)
- Bug [#128](https://github.com/aceberg/WatchYourLAN/issues/128)


## [v2.0.1] - 2024-09-02
### Added
- `Vlans` and `docker0` support [#47](https://github.com/aceberg/WatchYourLAN/issues/47). Thanks [thehijacker](https://github.com/thehijacker)!
- Remember `sort` field
- `InfluxDB` error handling

### Fixed
- Bug [#103](https://github.com/aceberg/WatchYourLAN/issues/103)
- Bug [#104](https://github.com/aceberg/WatchYourLAN/issues/104). Thanks [Steve Clement](https://github.com/SteveClement)!

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

