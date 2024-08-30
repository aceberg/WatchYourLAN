## API
```http
GET /api/all
```
Returns all hosts in `json`.


```http
GET /api/history/*mac
```
Returns all History. If `mac` is not empty, returns only history of a device with this `mac`.


```http
GET /api/host/:id
```
Returns host with this `id` in `json`.


```http
GET /api/port/:addr/:port
```
Gets state of one `port` of `addr`. Returns `true` if port is open or `false` otherwise.
<details>
  <summary>Request example</summary>

```bash
curl http://0.0.0.0:8840/api/port/192.168.2.2/8844
```
</details><br>


```http
GET /api/edit/:id/:name/*known
```
Edit host with ID `id`. Can change `name`. `known` is optional, when set to `toggle` will change Known state.


```http
GET /api/host/del/:id
```
Remove host with ID `id`.