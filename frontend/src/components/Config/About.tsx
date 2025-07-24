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
        <p>● After changing <b>Host</b> or <b>Port</b> the app must be restarted</p>
        <p>● <b>Shoutrrr URL</b> provides notifications to Discord, Email, Gotify, Telegram and other services. <a href="https://nicholas-fedor.github.io/shoutrrr/" target="_blank">Link to documentation</a></p>
        <p>● <b>Interfaces</b> - one or more, space separated</p>
        <p>● <b>Timeout (seconds)</b> - time between scans</p>
        <p>● <b>Args for arp-scan</b> - pass your own arguments to <code>arp-scan</code>. Enable <b>debug</b> log level to see resulting command. (Example: <code>-r 1</code>). See <a href="https://github.com/aceberg/WatchYourLAN/blob/main/docs/VLAN_ARP_SCAN.md" target="_blank">docs</a> for more.</p>
        <p>● <b>Arp Strings</b> - can setup scans for <code>vlans</code>, <code>docker0</code> and etcetera. See <a href="https://github.com/aceberg/WatchYourLAN/blob/main/docs/VLAN_ARP_SCAN.md" target="_blank">docs</a> for more.</p>
        <p>● <b>Trim History</b> - remove history after (hours)</p>
        <p>● <b>Store History in DB</b> - DEPRECATED. Now History is always stored in DB. Use <b>Trim History</b> to reduce DB size</p>
        <p>● <b>PG Connect URL</b> - address to connect to PostgreSQL DB. (Example: <code>postgres://username:password@192.168.0.1:5432/dbname?sslmode=disable</code>). Full list of URL parameters <a href="https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters" target="_blank">here</a></p>
        <p>● If you find this app useful, please, <a href="https://github.com/aceberg#donate" target="_blank">donate</a></p>
      </div>
    </div>
  )
}

export default About