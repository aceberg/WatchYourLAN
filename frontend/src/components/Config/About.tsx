import { For, createSignal, onMount } from "solid-js";
import { apiGetVersion } from "../../functions/api"

function About() {

  const [version, setVersion] = createSignal('');
  const forkHighlights = [
    "Modern responsive interface",
    "Fully local/offline UI assets",
    "Dark and light color modes",
    "Improved Home dashboard and filtering",
    "Presence timeline",
    "Persistent Events system",
    "Unified Events explorer with filtering and grouping",
    "Configurable Presence and Connectivity-event retention",
    "Manual persistent Device Type classification",
    "Host read/edit modes",
    "Responsive Settings experience",
  ];

  onMount(async () => {
    const v = await apiGetVersion();
    setVersion(v);
  });

  return (
    <div class="card wyl-panel config-panel">
      <div class="card-header">
        About
      </div>
      <div class="card-body table-responsive">
        <table class="table config-info-table"><tbody>
          <tr>
            <td class="config-field-label"><b>Project</b></td>
            <td class="config-field-value"><b>LANventory</b></td>
          </tr>
          <tr>
            <td class="config-field-label"><b>Status</b></td>
            <td class="config-field-value">In active development</td>
          </tr>
          <tr>
            <td class="config-field-label config-field-label-top"><b>Development status</b></td>
            <td class="config-field-value">LANventory is an actively developed fork of WatchYourLAN. Features, UI and behavior may continue to change while development is in progress.</td>
          </tr>
          <tr>
            <td class="config-field-label"><b>Repository</b></td>
            <td class="config-field-value"><a href="https://github.com/godlev/WatchYourLAN2" target="_blank" rel="noreferrer">godlev/WatchYourLAN2</a></td>
          </tr>
          <tr>
            <td class="config-field-label"><b>Upstream project</b></td>
            <td class="config-field-value"><a href="https://github.com/aceberg/WatchYourLAN" target="_blank" rel="noreferrer">WatchYourLAN by aceberg</a></td>
          </tr>
          <tr>
            <td class="config-field-label"><b>Base / backend version</b></td>
            <td class="config-field-value">
              {version()}
            </td>
          </tr>
          <tr>
            <td class="config-field-label config-field-label-top"><b>Fork highlights</b></td>
            <td class="config-field-value">
              <ul class="config-highlight-list">
                <For each={forkHighlights}>{highlight =>
                  <li>{highlight}</li>
                }</For>
              </ul>
            </td>
          </tr>
          <tr>
            <td class="config-field-label"><b>Swagger API docs</b></td>
            <td class="config-field-value"><a href="/swagger/index.html" target="_blank" rel="noreferrer">/swagger/index.html</a></td>
          </tr>
          <tr>
            <td class="config-field-label"><b>Network docs</b></td>
            <td class="config-field-value"><a href="https://github.com/aceberg/WatchYourLAN/blob/main/docs/VLAN_ARP_SCAN.md" target="_blank" rel="noreferrer">VLAN and ARP scan guide</a></td>
          </tr>
          <tr>
            <td class="config-field-label config-field-label-top"><b>Upstream credit</b></td>
            <td class="config-field-value">Based on <a href="https://github.com/aceberg/WatchYourLAN" target="_blank" rel="noreferrer">WatchYourLAN by aceberg</a>.</td>
          </tr>
          <tr>
            <td class="config-field-label config-field-label-top"><b>Notifications</b></td>
            <td class="config-field-value">Shoutrrr supports Discord, Email, Gotify, Telegram and other services. <a href="https://shoutrrr.nickfedor.com/services/overview/" target="_blank" rel="noreferrer">Service documentation</a></td>
          </tr>
          <tr>
            <td class="config-field-label config-field-label-top"><b>PostgreSQL URL</b></td>
            <td class="config-field-value">Connection string parameters are documented by <a href="https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters" target="_blank" rel="noreferrer">lib/pq</a>.</td>
          </tr>
        </tbody></table>
      </div>
    </div>
  )
}

export default About
