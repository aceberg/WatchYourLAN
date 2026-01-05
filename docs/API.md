## API
```http
GET {BASE_PATH}/api/all
```
Returns all hosts in `json`.


```http
GET {BASE_PATH}/api/history
```
Returns all History. Not recommended, the output can be a lot.

```http
GET {BASE_PATH}/api/history/:mac/:date
```
Returns only history of a device with this `mac` filtered by `date`. `date` format can be anything from `2` to `2025-07` to `2025-07-26`.

```http
GET {BASE_PATH}/api/history/:mac?num=20
```
Returns only last 20 lines of history of a device with this `mac`.


```http
GET {BASE_PATH}/api/host/:id
```
Returns host with this `id` in `json`.


```http
GET {BASE_PATH}/api/port/:addr/:port
```
Gets state of one `port` of `addr`. Returns `true` if port is open or `false` otherwise.
<details>
  <summary>Request example</summary>

```bash
# No base path
curl "http://0.0.0.0:8840/api/port/192.168.2.2/8844"

# Generic form
curl "http://0.0.0.0:8840/{BASE_PATH}/api/port/192.168.2.2/8844"
```
</details><br>


```http
GET {BASE_PATH}/api/edit/:id/:name/*known
```
Edit host with ID `id`. Can change `name`. `known` is optional, when set to `toggle` will change Known state.


```http
GET {BASE_PATH}/api/host/del/:id
```
Remove host with ID `id`.


```http
GET {BASE_PATH}/api/notify_test
```
Send test notification.


```http
GET {BASE_PATH}/api/status/*iface
```
Show status (Total number of hosts, online/offline, known/unknown). The `iface` parameter is optional and shows status for one interface only. For all interfaces just call `{BASE_PATH}/api/status/`.