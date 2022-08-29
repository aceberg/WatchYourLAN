<div align="center">

<a href="https://github.com/aceberg/WatchYourLAN">
    <img src="https://raw.githubusercontent.com/aceberg/WatchYourLAN/main/assets/logo.png" width="200" />
</a>
</div>
<br/>

# WatchYourLAN

[![Docker](https://github.com/aceberg/WatchYourLAN/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/aceberg/WatchYourLAN/actions/workflows/docker-publish.yml)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/watchyourlan)

Lightweight network IP scanner with web GUI  
https://github.com/aceberg/WatchYourLAN  

![Screenshot_v0.6](https://raw.githubusercontent.com/aceberg/WatchYourLAN/main/assets/Screenshot_v0.6.png)  

## Quick start

Replace `$YOURTIMEZONE` with correct time zone and `$YOURIFACE` with network interface you want to scan. Network mode must be `host`.

```sh
docker run --name wyl \
	-e "IFACE=$YOURIFACE" \
	-e "TZ=$YOURTIMEZONE" \
	--network="host" \
    aceberg/watchyourlan
```

Set `$DOCKERDATAPATH` for container to save data:

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
| IFACE     | Interface to scan. Could be one or more, separated by space. Currently `docker0` is not allowed, as arp-scan wouldn't work with it correctly | enp1s0 |
| DBPATH    | Path to Database | /data/db.sqlite |
| GUIIP     | Address for web GUI | localhost (127.0.0.1) |
| GUIPORT   | Port for web GUI | 8840 |
| TIMEOUT   | Time between scans (seconds) | 60 (1 minute) |
| SHOUTRRR_URL | Url to any notification service supported by [Shoutrrr](https://github.com/containrrr/shoutrrr/tree/main/docs/services) (gotify, email, telegram and others) | "" |
| THEME | Any theme name from https://bootswatch.com in lowcase | solar |

## Config file

Config file path is `/data/config`.
All variables could be set there. Example:
```sh
IFACE="enp2s0 wg0"
DBPATH="/data/hosts.db"
GUIIP="192.168.2.1"     		# To access from LAN
GUIPORT="8840"
TIMEOUT="300"           		# 5 minutes
SHOUTRRR_URL="gotify://192.168.2.1:8083/AwQqpAae.rrl5Ob/?title=Unknown host detected&DisableTLS=yes"	# Url to notify
THEME="darkly"
```

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/WatchYourLAN/network/dependencies)
- Favicon and logo: [Access point icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/access-point)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)