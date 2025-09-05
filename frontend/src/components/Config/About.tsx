import { createSignal, onMount } from "solid-js";
import { apiGetVersion } from "../../functions/api"

function About() {

  const [version, setVersion] = createSignal('');
  const [link, setLink] = createSignal('');

  onMount(async () => {
    const v = await apiGetVersion();
    setVersion(v);
    setLink("https://github.com/aceberg/WatchYourLAN/releases/tag/"+v);
  });

  return (
    <div class="card border-primary">
      <div class="card-header">
        About (<a href={link()} target="_blank">{version()}</a>)
      </div>
      <div class="card-body">
        <table class="table table-striped"><tbody>
          <tr>
            <td><b>Swagger API docs</b></td>
            <td><a href="/swagger/index.html" target="_blank">/swagger/index.html</a></td>
          </tr>
          <tr>
            <td><b>Local node-bootstrap URL</b></td>
            <td>local themes and fonts (optional). If empty, the app will pull everything from <code>cdn</code></td>
          </tr>
          <tr>
            <td><b>Shoutrrr URL</b></td>
            <td>provides notifications to Discord, Email, Gotify, Telegram and other services. <a href="https://shoutrrr.nickfedor.com/services/overview/" target="_blank">Link to documentation</a></td>
          </tr>
          <tr>
            <td><b>Interfaces</b></td>
            <td>one or more, space separated</td>
          </tr>
          <tr>
            <td><b>Timeout (seconds)</b></td>
            <td>time between scans</td>
          </tr>
          <tr>
            <td><b>Args for arp-scan</b></td>
            <td>pass your own arguments to <code>arp-scan</code>. Enable <b>debug</b> log level to see resulting command. (Example: <code>-r 1</code>). See <a href="https://github.com/aceberg/WatchYourLAN/blob/main/docs/VLAN_ARP_SCAN.md" target="_blank">docs</a> for more</td>
          </tr>
          <tr>
            <td><b>Arp Strings</b></td>
            <td>can setup scans for <code>vlans</code>, <code>docker0</code> and etcetera. See <a href="https://github.com/aceberg/WatchYourLAN/blob/main/docs/VLAN_ARP_SCAN.md" target="_blank">docs</a> for more</td>
          </tr>
          <tr>
            <td><b>Trim History</b></td>
            <td>remove history after (hours)</td>
          </tr>
          <tr>
            <td><b>PG Connect URL</b></td>
            <td>address to connect to PostgreSQL DB. (Example: <code>postgres://username:password@192.168.0.1:5432/dbname?sslmode=disable</code>). Full list of URL parameters <a href="https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters" target="_blank">here</a></td>
          </tr>
        </tbody></table>
      </div>
    </div>
  )
}

export default About