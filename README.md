<h1>
  <img src="frontend/public/fs/public/lanventory-128x128.png" width="48" alt="LANventory icon" />
  LANventory
</h1>

LANventory is an actively developed fork of [WatchYourLAN by aceberg](https://github.com/aceberg/WatchYourLAN), focused on modern LAN inventory, device presence history, event tracking, classification and a self-contained web interface.

> [!IMPORTANT]
> LANventory is still under active development. Features, configuration, packaging and upgrade instructions may continue to change until the first beta release is prepared.

Current repository: [godlev/WatchYourLAN2](https://github.com/godlev/WatchYourLAN2)

The original WatchYourLAN scanning/backend foundation is preserved and credited. LANventory adds a substantially expanded interface, persistent event model, device classification, configurable retention, migration hardening, safer configuration handling and other reliability improvements.

## LANventory dashboard

![LANventory dashboard](assets/image.png)

## Highlights

### Modern local-first interface

- Compact responsive dashboard.
- Dark and light color modes.
- Local Open Sans fonts.
- Local Bootstrap Icons.
- Local Bootswatch themes.
- LANventory favicon and navbar icon assets bundled locally.
- No automatic external UI asset requests after build.

### Home dashboard

- Summary cards for Total, Online, Offline, Known and Unknown devices.
- Shared filtering by interface, device type, recognition state, online state and search.
- Search across device information such as name and IP.
- Device Type filtering integrated with dashboard counts and recent event panels.
- Compact device table with sticky column headings.
- Device names link to read-only Host details.
- Direct edit shortcut from the device table.

### Device management

- Persistent Known / Unknown classification.
- Persistent manual Device Type classification.
- Device Type icon picker.
- Current supported types include router, switch, access point, firewall, server, NAS, desktop, laptop, phone, tablet, TV, printer, camera, IoT, virtual machine, container, game console and other.
- Host page defaults to read mode.
- Explicit Host edit mode for editable properties.
- Wake-on-LAN and port-scan actions remain available where supported.

### Presence

Presence is sampled online/offline visibility over time.

- Search by device name, IP and other shared host fields.
- Filter by interface, device type, recognition and current status.
- Day/night visualization.
- Hour and day timeline boundaries.
- Configurable Presence retention.
- Presence data remains separate from discrete Events.

### Events

Events record meaningful state transitions and device changes instead of storing every scan as an event.

Current event types:

- `discovered`
- `online`
- `offline`
- `known`
- `unknown`
- `device-type-changed`

Events features include:

- Unified Events explorer.
- Summary cards with event filtering.
- Ctrl/Cmd-click multi-select on event summary cards.
- Device filter.
- Event Type presets and custom multiple-event selections.
- Group By: Device, Event, Category, Device Type, IP, Interface and Day.
- Collapse All / Expand All for grouped results.
- Sticky Events table headings.
- Persistent reminder when Events filters/grouping are active.
- Device display preference: `Name + icon`, `Name only`, or `Icon only`.
- `Icon only` is the default for new/invalid preferences.
- Clickable device icons link to Host details when the Host still exists.
- Deleted-host connectivity events remain readable snapshots without stale links.
- Deterministic pagination ordering by event date and ID.

### Home recent events

Home contains two independent recent-event streams:

- Connectivity
- Device Changes

These panels follow the currently filtered device set on Home.

## Retention model

LANventory deliberately separates sampled Presence history from Events.

| Data | Retention behavior |
| --- | --- |
| Presence samples | Controlled by `TRIM_HIST` |
| `online` / `offline` Events | Controlled by `CONNECTIVITY_RETENTION` |
| Device-change Events | Retained while the device record exists |

If `CONNECTIVITY_RETENTION` is absent from an older configuration, LANventory falls back to the existing `TRIM_HIST` value for backward compatibility.

Retention can be configured from **Settings → Data retention** without restarting network scanning.

## Reliability and migration work

Recent release-readiness work includes:

- Non-destructive migration tests against a legacy WatchYourLAN SQLite schema.
- Additive migration for Device Type and Events storage.
- Idempotent database reopen/migration validation.
- Failed scans do not create false offline state changes or false offline Events.
- Repeated scans do not create duplicate transition Events.
- Event retention and Presence retention are tested independently.
- Missing Host IDs are rejected rather than becoming zero-value Hosts.
- Configuration updates are validated atomically before applying changes.
- Configuration access is synchronized for concurrent runtime readers/writers.
- Global database lifecycle/reconnect operations are mutex-guarded.
- Stored sensitive configuration values are no longer returned by `/api/config`.
- Secret settings are write-only in the Settings UI and can be kept, replaced or explicitly cleared.

### Remaining validation work before the first beta

- PostgreSQL runtime integration testing is still required.
- Docker/container packaging and upgrade-path validation are being prepared separately.
- The race detector has not been run in the current Windows development environment because CGO/gcc is unavailable there.

## Security and exposure

LANventory does **not** currently provide built-in authentication.

> [!WARNING]
> Do not expose LANventory directly to the public Internet. Use it on a trusted LAN, through a VPN, or behind an authenticated reverse proxy / SSO solution.

LANventory exposes administrative operations such as device editing/deletion, scan/configuration actions, Wake-on-LAN and port scanning. Protect access accordingly.

Stored PostgreSQL, InfluxDB and Shoutrrr secret values are write-only in the web Settings interface and are redacted from `/api/config`. The local `config_v2.yaml` file still contains the real values and must be protected with normal filesystem permissions.

## Runtime requirements

LANventory inherits the WatchYourLAN ARP discovery model.

- Linux is the intended runtime for real LAN scanning.
- `arp-scan` is required for real ARP discovery.
- Container deployments typically need host networking so `arp-scan` can see the real LAN interfaces.
- The web UI/API normally listens on `0.0.0.0:8840` unless configured otherwise.

Starting the real backend starts network scanning. Use the safe frontend mock mode below when working on UI without touching the LAN.

## Safe development preview

The frontend includes a localhost-only in-memory mock API. It does not run `arp-scan`, scan the LAN, send Wake-on-LAN packets, send notifications, perform real port scans or write production data.

Terminal 1:

```sh
cd frontend
npm install
npm run mock:api
```

Terminal 2:

```sh
cd frontend
npm run dev:local
```

Open:

```text
http://127.0.0.1:5173
```

Vite proxies `/api` and `/fs` to the safe mock API on `127.0.0.1:8840` during development.

## Build from source

Frontend:

```sh
cd frontend
npm install
npm run build
npm audit
npm audit --omit=dev
```

Backend:

```sh
cd backend
go test ./...
go vet ./...
go build ./...
```

To run the real backend from source, provide an intentional data/config directory:

```sh
cd backend
go run ./cmd/WatchYourLAN -d /path/to/dev-data
```

The command path remains inherited for compatibility. Real backend startup begins network scanning, so run it only where LAN scanning is intended and `arp-scan` is installed.

## Installation status

LANventory does not yet advertise a separately published Docker image or packaged release artifact.

Do **not** install `aceberg/watchyourlan` expecting LANventory features; that image belongs to the upstream WatchYourLAN project.

Container packaging, clean-install validation, persistent-volume mapping and the operational WatchYourLAN → LANventory upgrade path are being prepared as part of the next packaging phase.

Until that work is complete, evaluate LANventory by building from this repository checkout.

## Configuration

Configuration can be supplied through `config_v2.yaml`, the web Settings interface or environment variables. Config-file keys use lowercase equivalents of the environment variable names.

### General

| Variable | Description | Default |
| --- | --- | --- |
| `TZ` | Time zone used for timestamps. | |
| `HOST` | Listen address. | `0.0.0.0` |
| `PORT` | Web UI/API port. | `8840` |
| `THEME` | Bundled Bootswatch theme name. | `sand` |
| `COLOR` | Color mode: `dark` or `light`. | `dark` |
| `NODEPATH` | Legacy upstream compatibility setting. LANventory UI assets are bundled locally. | |
| `SHOUTRRR_URL` | Shoutrrr notification URL. See [Shoutrrr documentation](https://shoutrrr.nickfedor.com/services/overview/). | |

### Scanning and database

| Variable | Description | Default |
| --- | --- | --- |
| `IFACES` | Interfaces to scan, separated by spaces. See the upstream [VLAN / ARP scan guide](https://github.com/aceberg/WatchYourLAN/blob/main/docs/VLAN_ARP_SCAN.md). | |
| `TIMEOUT` | Time between scans in seconds. | `120` |
| `ARP_ARGS` | Additional arguments passed to `arp-scan`. | |
| `ARP_STRS`, `ARP_STRS_JOINED` | Optional ARP result strings. See the upstream [VLAN / ARP scan guide](https://github.com/aceberg/WatchYourLAN/blob/main/docs/VLAN_ARP_SCAN.md). | |
| `LOG_LEVEL` | `debug`, `info`, `warn` or `error`. | `info` |
| `USE_DB` | Database backend: `sqlite` or `postgres`. | `sqlite` |
| `PG_CONNECT` | PostgreSQL connection string. Parameters: [lib/pq](https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters). | |

### Data retention

| Variable | Description | Default |
| --- | --- | --- |
| `TRIM_HIST` | Presence sample retention in hours. | `48` |
| `CONNECTIVITY_RETENTION` | `online` / `offline` Event retention in hours. Older configs without this key inherit `TRIM_HIST`. | `TRIM_HIST` |
| `HIST_IN_DB` | Deprecated upstream compatibility setting. History is stored in the database. | |

### InfluxDB2

| Variable | Description | Default |
| --- | --- | --- |
| `INFLUX_ENABLE` | Enable export to InfluxDB2. | `false` |
| `INFLUX_SKIP_TLS` | Skip TLS verification. | `false` |
| `INFLUX_ADDR` | InfluxDB2 server URL. | |
| `INFLUX_BUCKET` | InfluxDB2 bucket. | |
| `INFLUX_ORG` | InfluxDB2 organization. | |
| `INFLUX_TOKEN` | InfluxDB2 token. Stored value is write-only in the UI. | |

### Prometheus

| Variable | Description | Default |
| --- | --- | --- |
| `PROMETHEUS_ENABLE` | Enable the `/metrics` endpoint. | `false` |

## Local / offline UI assets

LANventory's built web UI is self-contained.

The runtime serves local copies of:

- Bootstrap Icons
- Open Sans
- Bootswatch themes
- LANventory favicon/navbar icons
- JavaScript and CSS application bundles

The browser does not need Internet access to render the interface. User-clicked external links to documentation, GitHub, Shoutrrr, package documentation or Revolut remain normal external navigation.

## API and integrations

- API notes: [docs/API.md](docs/API.md)
- Prometheus: enable `PROMETHEUS_ENABLE` and use `/metrics`.
- InfluxDB2: configure the InfluxDB settings above.
- Swagger metadata is branded for LANventory, while inherited API/module compatibility remains intact.

Some integrations and packaging references may still exist only in upstream documentation. Treat them as upstream references unless explicitly updated for LANventory.

## PostgreSQL status

The current code retains PostgreSQL support and fork-added queries have been reviewed for portability, but an isolated real PostgreSQL runtime integration test is still required before LANventory is declared fully validated for PostgreSQL deployments.

SQLite is currently the most extensively tested migration/runtime path in the fork.

## Support development

If you find LANventory useful and would like to support continued development:

[Support via Revolut](https://revolut.me/mirgeo)

## Upstream project and attribution

LANventory is based on [WatchYourLAN by aceberg](https://github.com/aceberg/WatchYourLAN).

The fork intentionally preserves upstream attribution, compatibility-sensitive paths and applicable licensing notices while evolving the user-facing product independently.

Thanks to:

- [WatchYourLAN by aceberg](https://github.com/aceberg/WatchYourLAN)
- [Bootstrap](https://getbootstrap.com/)
- [Bootswatch](https://bootswatch.com/)
- [Bootstrap Icons](https://icons.getbootstrap.com/)
- Open Sans / Fontsource
- Go and JavaScript package authors listed in the dependency manifests

## Repository naming

The product is now branded **LANventory**, while the GitHub repository remains `godlev/WatchYourLAN2` for now. Repository/module renaming is intentionally deferred until compatibility, packaging and beta-release preparation are complete.
