<h1><a href="https://github.com/aceberg/WatchYourLAN">
    <img src="https://raw.githubusercontent.com/aceberg/WatchYourLAN/main/assets/logo.png" width="20" />
</a>WatchYourLAN</h1>
<br/>

[![Docker](https://github.com/aceberg/WatchYourLAN/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/WatchYourLAN/actions/workflows/main-docker-all.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/WatchYourLAN)](https://goreportcard.com/report/github.com/aceberg/WatchYourLAN)
[![Maintainability](https://api.codeclimate.com/v1/badges/46b17f99edc1726b5d7d/maintainability)](https://codeclimate.com/github/aceberg/WatchYourLAN/maintainability)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/watchyourlan)
![Docker Pulls](https://img.shields.io/docker/pulls/aceberg/watchyourlan)

Lightweight network IP scanner with web GUI 
- [Quick start](https://github.com/aceberg/WatchYourLAN#quick-start)    
- [Install .deb](https://github.com/aceberg/ppa)    
- [Config](https://github.com/aceberg/WatchYourLAN#config)   
- [Config file](https://github.com/aceberg/WatchYourLAN#config-file)   
- [Options](https://github.com/aceberg/WatchYourLAN#options)  
- [Local network only](https://github.com/aceberg/WatchYourLAN#local-network-only)  
- [Thanks](https://github.com/aceberg/WatchYourLAN#thanks) 

![Screenshot_v0.6](https://raw.githubusercontent.com/aceberg/WatchYourLAN/main/assets/Screenshot_v0.6.png)  

## Quick start

Replace `$YOURTIMEZONE` with correct time zone and `$YOURIFACE` with network interface you want to scan. Network mode must be `host`. Set `$DOCKERDATAPATH` for container to save data:

```sh
docker run --name wyl \
	-e "IFACE=$YOURIFACE" \
	-e "TZ=$YOURTIMEZONE" \
	--network="host" \
	-v $DOCKERDATAPATH/wyl:/data \
    aceberg/watchyourlan
```
Web GUI should be at http://localhost:8840

## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| AUTH | Enable Session-Cookie authentication | false |
| AUTH_EXPIRE | Session expiration time. A number and suffix: **m, h, d** or **M**. | 7d |
| AUTH_USER | Username | "" |
| AUTH_PASSWORD | Encrypted password (bcrypt). [How to encrypt password with bcrypt?](docs/BCRYPT.md) | "" |
| COLOR | Background color: light or dark | light |
| DBPATH    | Path to Database | /data/db.sqlite |
| GUIIP     | Address for web GUI | 0.0.0.0 |
| GUIPORT   | Port for web GUI | 8840 |
| HISTORY_DAYS | Keep devices online/offline history for (days) | 30 |
| IFACE     | Interface to scan. Could be one or more, separated by space. Currently `docker0` is not allowed, as arp-scan wouldn't work with it correctly | enp1s0 |
| IGNOREIP | If you want to detect unknown hosts by MAC only, set this wariable to "yes" | no |
| LOGLEVEL | How much log output you want to see ("short" or "verbose") | verbose |
| SHOUTRRR_URL | Url to any notification service supported by [Shoutrrr](https://github.com/containrrr/shoutrrr) (gotify, email, telegram and others) or [Generic Webhook](https://github.com/containrrr/shoutrrr/blob/main/docs/services/generic.md) | "" |
| THEME | Any theme name from https://bootswatch.com in lowcase | solar |
| TIMEOUT   | Time between scans (seconds) | 60 (1 minute) |


## Config file
> [!WARNING]  
> Config file format has been migrated to YAML in release [v1.0.0](https://github.com/aceberg/WatchYourLAN/releases/tag/1.0.0).   

Config file path is `/data/config.yaml`.
All variables could be set there. Example:
```yaml
color: light
dbpath: /data/db.sqlite
guiip: 192.168.2.1
guiport: "8840"
iface: enp1s0
ignoreip: "no"
loglevel: short
shoutrrr_url: gotify://192.168.2.1:8083/AwQqpAae.rrl5Ob/?title=Unknown host detected&DisableTLS=yes
theme: solar
timeout: 120
```

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -c | Path to config file | /data/config.yaml | 
| -n | Path to node modules (see below) | "" |

## Local network only
By default, this app pulls themes, icons and fonts from the internet. But, in some cases, it may be useful to have an independent from global network setup. I created a separate [image](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap) with all necessary modules and fonts.
Run with Docker:
```sh
docker run --name node-bootstrap          \
    -p 8850:8850                          \
    aceberg/node-bootstrap
```
```sh
docker run --name wyl \
	-e "IFACE=$YOURIFACE" \
	-e "TZ=$YOURTIMEZONE" \
	--network="host" \
	-v $DOCKERDATAPATH/wyl:/data \
    aceberg/watchyourlan -n "http://$YOUR_IP:8850"
```
Or use [docker-compose](docker-compose-local.yml)

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/WatchYourLAN/network/dependencies)
- Favicon and logo: [Access point icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/access-point)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)