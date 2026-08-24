import { apiGetAllHosts } from "./api";
import { Host, setBkpHosts, setHostsLoadError, setIfaces } from "./exports";
import { applyHostView } from "./hostView";
import { filterAtStart } from "./filter";
import { sortAtStart } from "./sort";

export function runAtStart() {
  getHosts();

  setInterval(() => {
    getHosts();
  }, 60000); // 60000 ms = 1 minute
}

export async function getHosts() {
  let hosts: Host[];
  try {
    hosts = await apiGetAllHosts();
  } catch {
    setHostsLoadError("Device data could not be refreshed. Showing the last loaded device list.");
    return;
  }

  setHostsLoadError("");

  if (hosts !== null && hosts.length > 0) {
    setBkpHosts(hosts);

    listIfaces(hosts);
    sortAtStart();
    filterAtStart();
    applyHostView();
  }
}

function listIfaces(hosts: Host[]) {

  let ifaces:string[] = [];

  for (let host of hosts) {
    const iface = host.Iface.trim();

    if (iface !== "" && !ifaces.includes(iface)) {
      ifaces.push(iface);
    }
  }

  setIfaces(ifaces);
}
