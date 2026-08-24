import { For, onMount, Show } from "solid-js";

import { allHosts, filterState, hostsLoadError } from "../functions/exports";

import TableRow from "../components/Body/TableRow";
import TableHead from "../components/Body/TableHead";
import CardHead from "../components/Body/CardHead";
import SummaryCards from "../components/Body/SummaryCards";
import RecentActivityPanel from "../components/Body/RecentActivityPanel";
import { getHosts } from "../functions/atstart";
import { deviceTypeFilterLabel } from "../functions/deviceTypes";

function Body() {

  onMount(() => {
    getHosts();
  });

  const hostLabel = (count: number) => count === 1 ? "host" : "hosts";

  const currentSubtitle = () => {
    const filters = filterState();
    const count = allHosts.length;
    const states = [];

    if (filters.Known === 1) states.push("Known");
    if (filters.Known === 0) states.push("Unknown");
    if (filters.Now === 1) states.push("Online");
    if (filters.Now === 0) states.push("Offline");

    let text = states.length > 0
      ? `${count} ${states.join(" ")} ${hostLabel(count)}`
      : `${count} ${hostLabel(count)}`;

    if (filters.Search !== "") {
      text = states.length > 0
        ? `${text} matching search`
        : `${count} matching ${hostLabel(count)}`;
    }

    if (filters.Iface !== "") {
      text = `${text} on ${filters.Iface}`;
    }

    if (filters.DeviceType !== "") {
      text = `${text} - Type: ${deviceTypeFilterLabel(filters.DeviceType)}`;
    }

    return text;
  };

  return (
    <>
    <Show when={hostsLoadError()}>
      <div class="data-load-warning" role="status">
        <i class="bi bi-exclamation-triangle-fill" aria-hidden="true"></i>
        <span>{hostsLoadError()}</span>
      </div>
    </Show>
    <SummaryCards></SummaryCards>
    <RecentActivityPanel></RecentActivityPanel>
    <div class="card device-panel">
      <div class="card-header device-panel-header">
        <div class="device-panel-title-group">
          <div class="device-panel-title">Devices</div>
          <div class="device-panel-subtitle">{currentSubtitle()}</div>
        </div>
        <CardHead></CardHead>
      </div>
      <div class="card-body table-responsive device-table-wrap">
        <table class="table table-hover device-table">
          <TableHead></TableHead>
          <tbody>
            <For each={allHosts}>{(host, index) =>
            <TableRow host={host} index={index() + 1}></TableRow>
            }</For>
          </tbody> 
        </table>
      </div>
    </div>
    </>
  )
}

export default Body
